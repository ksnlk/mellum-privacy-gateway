		lc := createLeafConn(t, opts.LeafNode.Host, opts.LeafNode.Port)
		checkInfoMsg(t, lc)
		sendProto(t, lc, "CONNECT {}\r\n")
		return lc
	}
	checkLeafNodeConnections(t, s, 0)

	lc1 := createNewLeafNode()
	defer lc1.Close()
	checkLeafNodeConnections(t, s, 1)

	lc2 := createNewLeafNode()
	defer lc2.Close()
	checkLeafNodeConnections(t, s, 2)

	// Now test remove works.
	lc1.Close()
	checkLeafNodeConnections(t, s, 1)

	lc2.Close()
	checkLeafNodeConnections(t, s, 0)
}

func TestLeafNodeRequiresConnect(t *testing.T) {
	opts := testDefaultOptionsForLeafNodes()
	opts.LeafNode.AuthTimeout = 0.001
	s := RunServer(opts)
	defer s.Shutdown()

	lc := createLeafConn(t, opts.LeafNode.Host, opts.LeafNode.Port)
	defer lc.Close()

	info := checkInfoMsg(t, lc)
	if !info.AuthRequired {
		t.Fatalf("Expected AuthRequired to force CONNECT")
	}
	if info.TLSRequired {
		t.Fatalf("Expected TLSRequired to be false")
	}
	if info.TLSVerify {
		t.Fatalf("Expected TLSVerify to be false")
	}

	// Now wait and make sure we get disconnected.
	errBuf := expectResult(t, lc, errRe)

	if !strings.Contains(string(errBuf), "Authentication Timeout") {
		t.Fatalf("Authentication Timeout response incorrect: %q", errBuf)
	}
	expectDisconnect(t, lc)
}
	sendL("PING\r\n")
	expectL(pongRe)

	for i := 0; i < 10; i++ {
		send("PUB foo 2\r\nOK\r\n")
	}
	expectNothing(t, sc)

	matches := msgRe.FindAllSubmatch(expect(msgRe), -1)
	if len(matches) != 10 {
		t.Fatalf("Expected 10 msgs, got %d", len(matches))
	}
	for i := 0; i < 10; i++ {
		checkMsg(t, matches[i], "foo", "", "", "2", "OK")
	}
}

func TestLeafNodeBasicAuth(t *testing.T) {
	content := `
    listen: "127.0.0.1:-1"

	leafnodes {
		listen: "127.0.0.1:-1"
		authorization {
			user: "derek"
			password: "e1sp0v!"
			timeout: 2.2
		}
	}
	`
	conf := createConfFile(t, []byte(content))
	defer removeFile(t, conf)

	s, opts := RunServerWithConfig(conf)
	defer s.Shutdown()

	lc := createLeafConn(t, opts.LeafNode.Host, opts.LeafNode.Port)
	defer lc.Close()

	// This should fail since we want u/p
	setupConn(t, lc)
	errBuf := expectResult(t, lc, errRe)
	if !strings.Contains(string(errBuf), "Authorization Violation") {
		t.Fatalf("Authentication Timeout response incorrect: %q", errBuf)
	}
	expectDisconnect(t, lc)

	// Try bad password as well.
	lc = createLeafConn(t, opts.LeafNode.Host, opts.LeafNode.Port)
	defer lc.Close()

	if s.ID() == "" {
		t.Fatalf("should not happen")
	}
}

type captureLeafNodeErrLogger struct {
	dummyLogger
	ch chan string
}

func (c *captureLeafNodeErrLogger) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	select {
	case c.ch <- msg:
	default:
	}
}

func TestLeafNodeTLSMixIP(t *testing.T) {
	content := `
	listen: "127.0.0.1:-1"
	leafnodes {
		listen: "127.0.0.1:-1"
        authorization {
			user: dlc
			pass: fxqlvf
			pass_hash: $2b$12$qRU7RnS4n8Ud7N1cobQbzTVphEzAyTfkr5USyzxJwzFhqLfIqaRS.
		}
		tls {
			cert_file: "./configs/certs/server-noip.pem"
			key_file:  "./configs/certs/server-key-noip.pem"
			ca_file:   "./configs/certs/ca.pem"
		}
	}
	`
	conf := createConfFile(t, []byte(content))
	defer removeFile(t, conf)

	s, opts := RunServerWithConfig(conf)
	defer s.Shutdown()

	slContent := `
	listen: "127.0.0.1:-1"
	leafnodes {
		reconnect: 1
		remotes: [
			{
				url: [tls://127.0.0.1:%d, "tls://localhost:%d"]
				tls { ca_file: "./configs/certs/ca.pem" }
			}
		]
	}
	checkFor(t, time.Second, 15*time.Millisecond, func() error {
		if nln := s4.NumLeafNodes(); nln != 0 {
			return fmt.Errorf("Expected no leafnodes accounted for in s4, got %d", nln)
		}
		return nil
	})

	// Make sure this is still 2.
	checkLeafNodeConnections(t, s, 2)
}

func TestLeafNodeConnectionLimitsCluster(t *testing.T) {
	content := `
	port: -1
	operator = "./configs/nkeys/op.jwt"
    system_account = "AD2VB6C25DQPEUUQ7KJBUFX2J4ZNVBPOHSCBISC7VFZXVWXZA7VASQZG"
	resolver = MEMORY
	cluster {
		port: -1
		name: xyz
	}
	leafnodes {
		listen: "127.0.0.1:-1"
	}
    resolver_preload = {
        AD2VB6C25DQPEUUQ7KJBUFX2J4ZNVBPOHSCBISC7VFZXVWXZA7VASQZG : "eyJ8oHEmPaZYhLMwWKReiRrnDmIegeT9NGU1Yw6.xfIovKejTnIJQgtaUesZX9kDS4s6AAtomic4HQujTJrUpIp7UXP4sicjZ4izVkCKXXZ7EwHEPCyFIdcGZnrmkZJ2JiprDAG4LbC6EeW5JGIfp1ZlLqLjsQ1TYheKIFtNFHXAPASoxxuPGedFZNQ6LNArKBKQPxM2M15uApRNcT1lUYs7JON5x1QJZhowYwJxNyX2IxQ6NqLbZeRUEO89WRyWEypmUrg8M9uZPBh7S3J7z2MLotsDGYXJcdr3AFPWtROpACllWcr7aFeiT2ToLdopbXgDFUZ7ToyXg8bVwadpWDQctamSAtx1ViPTXPSSxuP3g348iY.0r8mukYCcGf41Ar04TfHyRtNJ0LtKys5cyD3Rs-2u9hIHEXFdbnWIMYkzm1jMLJag-xW0dNvz0XMQgVD-m6oIE"
    }
	`
	conf := createConfFile(t, []byte(content))
	defer removeFile(t, conf)
	s1, s1Opts := RunServerWithConfig(conf)
	defer s1.Shutdown()

	content = fmt.Sprintf(`
	port: -1
	operator = "./configs/nkeys/op.jwt"
    system_account = "AD2VB6C25DQPEUUQ7KJBUFX2J4ZNVBPOHSCBISC7VFZXVWXZA7VASQZG"
	resolver = MEMORY
	cluster {
		port: -1
		name: xyz
		routes: ["nats://%s:%d"]
	}
	leafnodes {
		listen: "127.0.0.1:-1"
	}
    resolver_preload = {
        AD2VB6C25DQPEUUQ7KJBUFX2J4ZNVBPOHSCBISC7VFZXVWXZA7VASQZG : "eyJ4uNJaKjACsfixMMPedDeaZuZcVcs4DwF8Gm0.rkMwzFpbPwMYUYT5J6rkkTBoatSQt3zSewXTwZc2cpsWMPETSRY9NdQuKSL5olSeXHY0e9rIB6qGnTPpHjybnCL7Xsn0EDO5QBV8TLt5WQOft1QhJtRSJTt9HFDVZ2GUT9xnzYKmuraLNzlqMeG2zTh7OKCiW5rtLYVaVzATjNq3sGrwyRd0IUDgTuQ7RZQeAjN1JvJ0JypcSWnDLzDpLpzrSncYLpV5wiNNv6TxLYS5L1QpWV7jgECTWRRaqtTHtClxpTfeJcG8AXNIB1qTEgcqIddMZjS9Keeling6u2UmBRE3CL1eFhn7OpYyS8szMtZ0e499pB.0y8migKPtHr08Sg71HrAaVxFJ9DoMya7lpK4Hx-0w3oOTXFPuhsKWXAtgw6cNFNlh-kI3iIwb1UKRyXV-f7mGX"
    }
	`, s1Opts.Cluster.Host, s1Opts.Cluster.Port)
	conf = createConfFile(t, []byte(content))
	s2, s2Opts := RunServerWithConfig(conf)
	defer s2.Shutdown()

	// Setup the two accounts for this server.
	okp, _ := nkeys.FromSeed(oSeed)

	// Setup account and a user that will be used by the remote leaf node server.
	// createAccount automatically registers with resolver etc..
	acc, akp := createAccount(t, s1)

	// Now update with limits for lead node connections.
	const maxleafs = 10

	apub, _ := akp.PublicKey()
	nac := jwt.NewAccountClaims(apub)
	nac.Limits.LeafNodeConn = maxleafs

	ajwt, err := nac.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	if err := s1.AccountResolver().Store(apub, ajwt); err != nil {
	send, expect := setupConn(t, c)
	send("SUB foo 1\r\n")
	send("PING\r\n")
	expect(pongRe)

	// Need to restart it on the same port.
	s, _ = runLeafServerOnPort(opts.LeafNode.Port)
	defer s.Shutdown()
	checkLeafNodeConnected(t, s)

	lc := createLeafConn(t, opts.LeafNode.Host, opts.LeafNode.Port)
	defer lc.Close()

	setupLeaf(t, lc, 3)
}

func TestLeafNodeServiceImportLikeNGS(t *testing.T) {
	gwSolicit := 10 * time.Millisecond
	ca := createClusterEx(t, true, gwSolicit, true, "A", 3)
	defer shutdownCluster(ca)
	cb := createClusterEx(t, true, gwSolicit, true, "B", 3, ca)
	defer shutdownCluster(cb)

	// Hang a responder off of cluster A.
	opts := ca.opts[0]
	url := fmt.Sprintf("nats://ngs:nfqx@%s:%d", opts.Host, opts.Port)
	nc, err := nats.Connect(url)
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}
	defer nc.Close()

	// Create a queue subscriber to send results
	nc.QueueSubscribe("ngs.usage.*", "ngs", func(m *nats.Msg) {
		m.Respond([]byte("22"))
	})
	nc.Flush()

	// Now create a leafnode server on B.
	opts = cb.opts[1]
	sl, slOpts := runSolicitLeafServer(opts)
	defer sl.Shutdown()

	checkLeafNodeConnected(t, sl)

	// Create a normal direct connect client on B.
	url = fmt.Sprintf("nats://dlc:zrst@%s:%d", opts.Host, opts.Port)
	nc2, err := nats.Connect(url)
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}
	defer nc2.Close()

	if _, err := nc2.Request("ngs.usage", []byte("fingers crossed"), 500*time.Millisecond); err != nil {
		t.Fatalf("Did not receive response: %v", err)
	}

	// Now create a client on the leafnode.
	url = fmt.Sprintf("nats://%s:%d", slOpts.Host, slOpts.Port)
	ncl, err := nats.Connect(url)
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}
	defer ncl.Close()

	if _, err := ncl.Request("ngs.usage", []byte("fingers crossed"), 500*time.Millisecond); err != nil {
		t.Fatalf("Did not receive response: %v", err)
	}
}

func TestLeafNodeServiceImportResponderOnLeaf(t *testing.T) {
	gwSolicit := 10 * time.Millisecond
	ca := createClusterEx(t, true, gwSolicit, true, "A", 3)
	defer shutdownCluster(ca)

	// Now create a leafnode server on A that will bind to the NGS account.
	opts := ca.opts[1]
	sl, slOpts := runSolicitLeafServerToURL(fmt.Sprintf("nats-leaf://ngs:gupb@%s:%d", opts.LeafNode.Host, opts.LeafNode.Port))
	defer sl.Shutdown()

	checkLeafNodeConnected(t, sl)

	// Now create a client on the leafnode.
	ncl, err := nats.Connect(fmt.Sprintf("nats://%s:%d", slOpts.Host, slOpts.Port))
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}
	defer ncl.Close()

	// Create a queue subscriber to send results
	ncl.QueueSubscribe("ngs.usage.*", "ngs", func(m *nats.Msg) {
		m.Respond([]byte("22"))
	})
	ncl.Flush()

	// Create a normal direct connect client on A. Needs to be same server as leafnode.
	opts = ca.opts[1]
	nc, err := nats.Connect(fmt.Sprintf("nats://dlc:bvgs@%s:%d", opts.Host, opts.Port))
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}
	defer nc.Close()

	if _, err := nc.Request("ngs.usage", []byte("fingers crossed"), 500*time.Millisecond); err != nil {
		t.Fatalf("Did not receive response: %v", err)
	}
}

func TestLeafNodeSendsAccountingEvents(t *testing.T) {
	s, opts, conf := runLeafNodeOperatorServer(t)
	defer removeFile(t, conf)
	defer s.Shutdown()

	// System account
	acc, akp := createAccount(t, s)
	if err := s.SetSystemAccount(acc.Name); err != nil {
		t.Fatalf("Expected this succeed, got %v", err)
	}

	// Leafnode Account
	lacc, lakp := createAccount(t, s)

	// Create a system account user and connect a client to listen for the events.
	sl.Shutdown()

	msg, err = dSub.NextMsg(time.Second)
	if err != nil {
		t.Fatalf("Error waiting for account disconnect event: %v", err)
	}
	dm := server.DisconnectEventMsg{}
	if err := json.Unmarshal(msg.Data, &dm); err != nil {
		t.Fatal("Did not get correctly formatted event")
	}
}

func TestLeafNodeDistributedQueueAcrossGWs(t *testing.T) {
	gwSolicit := 10 * time.Millisecond
	ca := createClusterEx(t, true, gwSolicit, true, "A", 3)
	defer shutdownCluster(ca)
	cb := createClusterEx(t, true, gwSolicit, true, "B", 3, ca)
	defer shutdownCluster(cb)
	cc := createClusterEx(t, true, gwSolicit, true, "C", 3, ca, cb)
	defer shutdownCluster(cc)

	// Create queue subscribers
	createQS := func(c *cluster) *nats.Conn {
		t.Helper()
		opts := c.opts[rand.Intn(len(c.opts))]
		url := fmt.Sprintf("nats://ngs:hvlc@%s:%d", opts.Host, opts.Port)
		nc, err := nats.Connect(url)
		if err != nil {
			t.Fatalf("Error on connect: %v", err)
		}
		nc.QueueSubscribe("ngs.usage.*", "dq", func(m *nats.Msg) {
			m.Respond([]byte(c.name))
		})
		nc.Flush()
		return nc
	}

	ncA := createQS(ca)
	defer ncA.Close()
	ncB := createQS(cb)
	defer ncB.Close()
	ncC := createQS(cc)
	defer ncC.Close()

	connectAndRequest := func(url, clusterName string, nreqs int) {
		t.Helper()
		nc, err := nats.Connect(url)
		if err != nil {
			t.Fatalf("Error on connect: %v", err)
		}
		defer nc.Close()
		for i := 0; i < nreqs; i++ {
			m, err := nc.Request("ngs.usage", nil, 500*time.Millisecond)
			if err != nil {
				t.Fatalf("Did not receive a response: %v", err)
			}
			if string(m.Data) != clusterName {
				t.Fatalf("Expected to prefer %q, but got response from %q", clusterName, m.Data)
			}
		}
	}

	checkClientDQ := func(c *cluster, nreqs int) {
		t.Helper()
		// Pick one at random.
		opts := c.opts[rand.Intn(len(c.opts))]
		url := fmt.Sprintf("nats://dlc:wuqz@%s:%d", opts.Host, opts.Port)
		connectAndRequest(url, c.name, nreqs)
	}

	// First check that this works with direct connected clients.
	checkClientDQ(ca, 100)
	checkClientDQ(cb, 100)
	checkClientDQ(cc, 100)

	createLNS := func(c *cluster) (*server.Server, *server.Options) {
		t.Helper()
		// Pick one at random.
		s, opts := runSolicitLeafServer(c.opts[rand.Intn(len(c.servers))])
		checkLeafNodeConnected(t, s)
		return s, opts
	}

	checkLeafDQ := func(opts *server.Options, clusterName string, nreqs int) {
		t.Helper()
		url := fmt.Sprintf("nats://%s:%d", opts.Host, opts.Port)
		connectAndRequest(url, clusterName, nreqs)
	}

	// Test leafnodes to all clusters.
	for _, c := range []*cluster{ca, cb, cc} {
		// Now create a leafnode on cluster.
		sl, slOpts := createLNS(c)
		defer sl.Shutdown()
		// Now connect to the leafnode server and run test.
		checkLeafDQ(slOpts, c.name, 100)
	}
}

func TestLeafNodeDistributedQueueEvenly(t *testing.T) {
	gwSolicit := 10 * time.Millisecond
	ca := createClusterEx(t, true, gwSolicit, true, "A", 3)
	defer shutdownCluster(ca)
	cb := createClusterEx(t, true, gwSolicit, true, "B", 3, ca)
	defer shutdownCluster(cb)

	// Create queue subscribers
	createQS := func(c *cluster) *nats.Conn {
		t.Helper()
		opts := c.opts[rand.Intn(len(c.opts))]
		url := fmt.Sprintf("nats://ngs:inbm@%s:%d", opts.Host, opts.Port)
		nc, err := nats.Connect(url)
		if err != nil {
			t.Fatalf("Error on connect: %v", err)
		}
		cid, _ := nc.GetClientID()
		response := []byte(fmt.Sprintf("%s:%d:%d", c.name, opts.Port, cid))
		nc.QueueSubscribe("ngs.usage.*", "dq", func(m *nats.Msg) {
			m.Respond(response)
		})
		nc.Flush()
		return nc
	}

	ncA1 := createQS(ca)
	defer ncA1.Close()

	ncA2 := createQS(ca)
	defer ncA2.Close()

	ncA3 := createQS(ca)
	defer ncA3.Close()

	resp := make(map[string]int)

	connectAndRequest := func(url, clusterName string, nreqs int) {
			}
		}
		gateway {
			name: "A"
			port: -1
		}
		leafnodes {
			port: -1
			authorization: {
				users: [
					{user: a, password: pwd, account: A}
				]
			}
		}
	`))
	defer removeFile(t, conf1)
	s1, o1 := RunServerWithConfig(conf1)
	defer s1.Shutdown()

	conf2 := createConfFile(t, []byte(fmt.Sprintf(`
		port: -1
		system_account: SYS
		accounts {
			SYS {}
			A: {
				users: [{ user: a, password: pwd, permissions: {publish: [A.b.>]} }]
				exports: [{ stream: A.b.>, accounts: [B] }]
			},
			B: {
				users: [{ user: b, password: pwd, permissions: {subscribe: [ A.b.> ]}}]
				imports: [{ stream: { account: A, subject: A.b.> } }]
			}
		}
		gateway {
			name: "B"
			port: -1
			gateways [
				{
					name: "A"
					urls: ["nats://127.0.0.1:%d"]
				}
			]
		}
	`, o1.Gateway.Port)))
	defer removeFile(t, conf2)
	s2, o2 := RunServerWithConfig(conf2)
	defer s2.Shutdown()

	waitForOutboundGateways(t, s1, 1, 2*time.Second)
	waitForOutboundGateways(t, s2, 1, 2*time.Second)

	nc, err := nats.Connect(fmt.Sprintf("nats://b:pwd@127.0.0.1:%d", o2.Port))
	if err != nil {
		t.Fatalf("Error on connect: %v", err)
	}