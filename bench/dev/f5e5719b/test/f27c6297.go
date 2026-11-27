			conn.EXPECT().Write(gomock.Any()).Return(io.ErrClosedPipe).AnyTimes()
			sess.sendQueue = newSendQueue(conn)
			sph := mockackhandler.NewMockSentPacketHandler(mockCtrl)
			sph.EXPECT().GetLossDetectionTimeout().Return(time.Now().Add(time.Hour)).AnyTimes()
			sph.EXPECT().SendMode().Return(ackhandler.SendAny).AnyTimes()
			sph.EXPECT().HasPacingBudget().Return(true).AnyTimes()
			// only expect a single SentPacket() call
			sph.EXPECT().SentPacket(gomock.Any())
			tracer.EXPECT().SentPacket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
			tracer.EXPECT().ClosedConnection(gomock.Any())
			tracer.EXPECT().Close()
			streamManager.EXPECT().CloseWithError(gomock.Any())
			sessionRunner.EXPECT().Remove(gomock.Any()).AnyTimes()
			cryptoSetup.EXPECT().Close()
			sess.sentPacketHandler = sph
			p := getPacket(1)
			packer.EXPECT().PackPacket().Return(p, nil)
			packer.EXPECT().PackPacket().Return(nil, nil).AnyTimes()
			runSession()
			sess.queueControlFrame(&wire.PingFrame{})
			sess.scheduleSending()
			Eventually(sess.Context().Done()).Should(BeClosed())
		})

		It("closes due to a stateless reset", func() {
			token := protocol.StatelessResetToken{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			runSession()
			gomock.InOrder(
				tracer.EXPECT().ClosedConnection(gomock.Any()).Do(func(e error) {
					var srErr *StatelessResetError
					Expect(errors.As(e, &srErr)).To(BeTrue())
					Expect(srErr.Token).To(Equal(token))
				}),
				tracer.EXPECT().Close(),
			)
			streamManager.EXPECT().CloseWithError(gomock.Any())
			sessionRunner.EXPECT().Remove(gomock.Any()).AnyTimes()
			cryptoSetup.EXPECT().Close()
			sess.destroy(&StatelessResetError{Token: token})
		})
	})

	Context("receiving packets", func() {
		var unpacker *MockUnpacker

		BeforeEach(func() {
			unpacker = NewMockUnpacker(mockCtrl)
			sess.unpacker = unpacker
		})

		getPacket := func(extHdr *wire.ExtendedHeader, data []byte) *receivedPacket {
			buf := &bytes.Buffer{}
			Expect(extHdr.Write(buf, sess.version)).To(Succeed())
			return &receivedPacket{
				data:    append(buf.Bytes(), data...),
				buffer:  getPacketBuffer(),
				rcvTime: time.Now(),
			}
		}

		It("drops Retry packets", func() {
			p := getPacket(&wire.ExtendedHeader{Header: wire.Header{
				IsLongHeader:     true,
				Type:             protocol.PacketTypeRetry,
				DestConnectionID: destConnID,
				SrcConnectionID:  srcConnID,
				Version:          sess.version,
				Token:            []byte("foobar"),
			}}, make([]byte, 16) /* Retry integrity tag */)
			tracer.EXPECT().DroppedPacket(logging.PacketTypeRetry, p.Size(), logging.PacketDropUnexpectedPacket)
			Expect(sess.handlePacketImpl(p)).To(BeFalse())
		})

		It("drops Version Negotiation packets", func() {
			b, err := wire.ComposeVersionNegotiation(srcConnID, destConnID, sess.config.Versions)
			Expect(err).ToNot(HaveOccurred())
			tracer.EXPECT().DroppedPacket(logging.PacketTypeVersionNegotiation, protocol.ByteCount(len(b)), logging.PacketDropUnexpectedPacket)
			Expect(sess.handlePacketImpl(&receivedPacket{
				data:   b,
				buffer: getPacketBuffer(),
			})).To(BeFalse())
		})

		It("drops packets for which header decryption fails", func() {
			p := getPacket(&wire.ExtendedHeader{
				Header: wire.Header{
					IsLongHeader: true,
					Type:         protocol.PacketTypeHandshake,
					Version:      sess.version,
				},
				PacketNumberLen: protocol.PacketNumberLen2,
			}, nil)
			p.data[0] ^= 0x40 // unset the QUIC bit
	It("interprets an ACK for 1-RTT packets as confirmation of the handshake", func() {
		sess.peerParams = &wire.TransportParameters{}
		sph := mockackhandler.NewMockSentPacketHandler(mockCtrl)
		sess.sentPacketHandler = sph
		ack := &wire.AckFrame{AckRanges: []wire.AckRange{{Smallest: 1, Largest: 3}}}
		sph.EXPECT().ReceivedAck(ack, protocol.Encryption1RTT, gomock.Any()).Return(true, nil)
		sph.EXPECT().SetHandshakeConfirmed()
		cryptoSetup.EXPECT().SetLargest1RTTAcked(protocol.PacketNumber(3))
		cryptoSetup.EXPECT().SetHandshakeConfirmed()
		Expect(sess.handleAckFrame(ack, protocol.Encryption1RTT)).To(Succeed())
	})

	Context("handling tokens", func() {
		var mockTokenStore *MockTokenStore

		BeforeEach(func() {
			mockTokenStore = NewMockTokenStore(mockCtrl)
			tlsConf = &tls.Config{ServerName: "server"}
			quicConf.TokenStore = mockTokenStore
			mockTokenStore.EXPECT().Pop(gomock.Any())
			quicConf.TokenStore = mockTokenStore
		})

		It("handles NEW_TOKEN frames", func() {
			mockTokenStore.EXPECT().Put("server", &ClientToken{data: []byte("foobar")})
			Expect(sess.handleNewTokenFrame(&wire.NewTokenFrame{Token: []byte("foobar")})).To(Succeed())
		})
	})

	Context("handling Version Negotiation", func() {
		getVNP := func(versions ...protocol.VersionNumber) *receivedPacket {
			b, err := wire.ComposeVersionNegotiation(srcConnID, destConnID, versions)
			Expect(err).ToNot(HaveOccurred())
			return &receivedPacket{
				data:   b,
				buffer: getPacketBuffer(),
			}
		}

		It("closes and returns the right error", func() {
			sph := mockackhandler.NewMockSentPacketHandler(mockCtrl)
			sess.sentPacketHandler = sph
			sph.EXPECT().ReceivedBytes(gomock.Any())
			sph.EXPECT().PeekPacketNumber(protocol.EncryptionInitial).Return(protocol.PacketNumber(128), protocol.PacketNumberLen4)
			sess.config.Versions = []protocol.VersionNumber{1234, 4321}
			errChan := make(chan error, 1)
			go func() {
				defer GinkgoRecover()
				cryptoSetup.EXPECT().RunHandshake().MaxTimes(1)
				errChan <- sess.run()
			}()
			p := getVNP(sess.version)
			tracer.EXPECT().DroppedPacket(logging.PacketTypeVersionNegotiation, p.Size(), logging.PacketDropUnexpectedVersion)
			Expect(sess.handlePacketImpl(p)).To(BeFalse())
		})

		It("ignores unparseable Version Negotiation packets", func() {
			p := getVNP(sess.version)
			p.data = p.data[:len(p.data)-2]
			tracer.EXPECT().DroppedPacket(logging.PacketTypeVersionNegotiation, p.Size(), logging.PacketDropHeaderParseError)
			Expect(sess.handlePacketImpl(p)).To(BeFalse())
		})
	})

	Context("handling Retry", func() {
		origDestConnID := protocol.ConnectionID{8, 7, 6, 5, 4, 3, 2, 1}

		var retryHdr *wire.ExtendedHeader

		JustBeforeEach(func() {
			retryHdr = &wire.ExtendedHeader{
				Header: wire.Header{
					IsLongHeader:     true,
					Type:             protocol.PacketTypeRetry,
					SrcConnectionID:  protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef},
					DestConnectionID: protocol.ConnectionID{1, 2, 3, 4, 5, 6, 7, 8},
					Token:            []byte("foobar"),
					Version:          sess.version,
				},
			}
		})

		getRetryTag := func(hdr *wire.ExtendedHeader) []byte {
			buf := &bytes.Buffer{}
			hdr.Write(buf, sess.version)
			return handshake.GetRetryIntegrityTag(buf.Bytes(), origDestConnID, hdr.Version)[:]
		}

		It("handles Retry packets", func() {
			sph := mockackhandler.NewMockSentPacketHandler(mockCtrl)
			sess.sentPacketHandler = sph
			sph.EXPECT().ResetForRetry()
			sph.EXPECT().ReceivedBytes(gomock.Any())
			cryptoSetup.EXPECT().ChangeConnectionID(protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef})
			packer.EXPECT().SetToken([]byte("foobar"))
			tracer.EXPECT().ReceivedRetry(gomock.Any()).Do(func(hdr *wire.Header) {
				Expect(hdr.DestConnectionID).To(Equal(retryHdr.DestConnectionID))
				Expect(hdr.SrcConnectionID).To(Equal(retryHdr.SrcConnectionID))
				Expect(hdr.Token).To(Equal(retryHdr.Token))
			})
			Expect(sess.handlePacketImpl(getPacket(retryHdr, getRetryTag(retryHdr)))).To(BeTrue())
		})

		It("ignores Retry packets after receiving a regular packet", func() {
			sess.receivedFirstPacket = true
			p := getPacket(retryHdr, getRetryTag(retryHdr))
			tracer.EXPECT().DroppedPacket(logging.PacketTypeRetry, p.Size(), logging.PacketDropUnexpectedPacket)
			Expect(sess.handlePacketImpl(p)).To(BeFalse())
		})

		It("ignores Retry packets if the server didn't change the connection ID", func() {
			retryHdr.SrcConnectionID = destConnID
			p := getPacket(retryHdr, getRetryTag(retryHdr))
			tracer.EXPECT().DroppedPacket(logging.PacketTypeRetry, p.Size(), logging.PacketDropUnexpectedPacket)
			Expect(sess.handlePacketImpl(p)).To(BeFalse())
		})

		It("ignores Retry packets with the a wrong Integrity tag", func() {
			tag := getRetryTag(retryHdr)
			tag[0]++
				}
				cryptoSetup.EXPECT().Close()
				mconn.EXPECT().Write(gomock.Any())
				gomock.InOrder(
					tracer.EXPECT().ClosedConnection(gomock.Any()),
					tracer.EXPECT().Close(),
				)
			}
			closed = true
		}

		AfterEach(func() {
			sess.shutdown()
			Eventually(sess.Context().Done()).Should(BeClosed())
			Eventually(errChan).Should(BeClosed())
		})

		It("uses the preferred_address connection ID", func() {
			params := &wire.TransportParameters{
				OriginalDestinationConnectionID: destConnID,
				InitialSourceConnectionID:       destConnID,
				PreferredAddress: &wire.PreferredAddress{
					IPv4:                net.IPv4(127, 0, 0, 1),
					IPv6:                net.IP{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
					ConnectionID:        protocol.ConnectionID{1, 2, 3, 4},
					StatelessResetToken: protocol.StatelessResetToken{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				},
			}
			packer.EXPECT().HandleTransportParameters(gomock.Any())
			packer.EXPECT().PackCoalescedPacket().MaxTimes(1)
			tracer.EXPECT().ReceivedTransportParameters(params)
			sess.handleTransportParameters(params)
			sess.handleHandshakeComplete()
			// make sure the connection ID is not retired
			cf, _ := sess.framer.AppendControlFrames(nil, protocol.MaxByteCount)
			Expect(cf).To(BeEmpty())
			sessionRunner.EXPECT().AddResetToken(protocol.StatelessResetToken{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, sess)
			Expect(sess.connIDManager.Get()).To(Equal(protocol.ConnectionID{1, 2, 3, 4}))
			// shut down
			sessionRunner.EXPECT().RemoveResetToken(protocol.StatelessResetToken{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
			expectClose(true)
		})

		It("uses the minimum of the peers' idle timeouts", func() {
			sess.config.MaxIdleTimeout = 19 * time.Second
			params := &wire.TransportParameters{
				OriginalDestinationConnectionID: destConnID,
				InitialSourceConnectionID:       destConnID,
				MaxIdleTimeout:                  18 * time.Second,
			}
			packer.EXPECT().HandleTransportParameters(gomock.Any())
			tracer.EXPECT().ReceivedPacket(gomock.Any(), gomock.Any(), gomock.Any())
			Expect(sess.handlePacketImpl(wrapPacket(initialPacket))).To(BeFalse())
		})

		// Illustrates that an injected Initial with a CONNECTION_CLOSE frame causes
		// the connection to immediately break down
		It("fails on Initial-level CONNECTION_CLOSE frame", func() {
			connCloseFrame := &wire.ConnectionCloseFrame{
				IsApplicationError: true,
				ReasonPhrase:       "mitm attacker",
			}
			initialPacket := testutils.ComposeInitialPacket(destConnID, srcConnID, sess.version, destConnID, []wire.Frame{connCloseFrame})
			tracer.EXPECT().ReceivedPacket(gomock.Any(), gomock.Any(), gomock.Any())
			Expect(sess.handlePacketImpl(wrapPacket(initialPacket))).To(BeTrue())
		})

		// Illustrates that attacker who injects a Retry packet and changes the connection ID
		// can cause subsequent real Initial packets to be ignored
		It("ignores Initial packets which use original source id, after accepting a Retry", func() {
			sph := mockackhandler.NewMockSentPacketHandler(mockCtrl)
			sess.sentPacketHandler = sph
			sph.EXPECT().ReceivedBytes(gomock.Any()).Times(2)
			sph.EXPECT().ResetForRetry()
			newSrcConnID := protocol.ConnectionID{0xde, 0xad, 0xbe, 0xef}
			cryptoSetup.EXPECT().ChangeConnectionID(newSrcConnID)
			packer.EXPECT().SetToken([]byte("foobar"))

			tracer.EXPECT().ReceivedRetry(gomock.Any())
			sess.handlePacketImpl(wrapPacket(testutils.ComposeRetryPacket(newSrcConnID, destConnID, destConnID, []byte("foobar"), sess.version)))
			initialPacket := testutils.ComposeInitialPacket(sess.connIDManager.Get(), srcConnID, sess.version, sess.connIDManager.Get(), nil)
			tracer.EXPECT().DroppedPacket(gomock.Any(), gomock.Any(), gomock.Any())
			Expect(sess.handlePacketImpl(wrapPacket(initialPacket))).To(BeFalse())
		})
	})
})
