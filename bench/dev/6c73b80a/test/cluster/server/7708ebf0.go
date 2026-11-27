	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	createConsumer("S", "foo.1")
	createConsumer("S", "bar.1")
	expectFail("S", "baz")
	expectFail("S", "baz.1")
	expectFail("S", "apple")

	// Now cross account stuff.
	nc2, js2 := jsClientConnect(t, s, nats.UserInfo("rip", "pass"))
	defer nc2.Close()

	if _, err := js2.AddStream(&nats.StreamConfig{Name: "ORIGIN", Subjects: []string{"foo.*"}}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	cfg := StreamConfig{
		Name:    "SCA",
		Storage: FileStorage,
		Sources: []*StreamSource{{
			Name: "ORIGIN",
			External: &ExternalStream{
				ApiPrefix:     "RI.JS.API",
				DeliverPrefix: "RI.DELIVER.SYNC.SOURCES",
			},
		}},
	}
	req, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, err := nc.Request(fmt.Sprintf(JSApiStreamCreateT, cfg.Name), req, time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	var scResp JSApiStreamCreateResponse
	if err := json.Unmarshal(resp.Data, &scResp); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if scResp.StreamInfo == nil || scResp.Error != nil {
		t.Fatalf("Did not receive correct response: %+v", scResp.Error)
	}

	// Externals skip the checks for now.
	createConsumer("SCA", "foo.1")
	createConsumer("SCA", "bar.1")
	createConsumer("SCA", "baz")
}
	s := c.randomServer()
	nc, js := jsClientConnect(t, s, nats.UserInfo("rip", "pass"))
	defer nc.Close()

	if _, err := js.AddStream(&nats.StreamConfig{Name: "TEST", Replicas: 2}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	toSend := 100
	for i := 0; i < toSend; i++ {
		if _, err := js.Publish("TEST", []byte("OK")); err != nil {
			t.Fatalf("Unexpected publish error: %v", err)
		}
	}

	nc2, js2 := jsClientConnect(t, s)
	defer nc2.Close()

	// Have to do this direct until we get Go client support.
	// Need to match jsClusterMirrorSourceImportsTempl imports.
	_, err := js2.AddStream(&nats.StreamConfig{
		Name: "MY_MIRROR_TEST",
		Mirror: &nats.StreamSource{
			Name: "TEST",
			External: &nats.ExternalStream{
				APIPrefix:     "RI.JS.API",
				DeliverPrefix: "RI.DELIVER.SYNC.MIRRORS",
			},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	checkFor(t, 20*time.Second, 500*time.Millisecond, func() error {
		si, err := js2.StreamInfo("MY_MIRROR_TEST")
		if err != nil {
			t.Fatalf("Could not retrieve stream info: %s", err)
		}
		if si.State.Msgs != uint64(toSend) {
			return fmt.Errorf("Expected %d msgs, got state: %+v", toSend, si.State)
		}
		return nil
	})

	// Now do sources as well.
	_, err = js2.AddStream(&nats.StreamConfig{
		Name: "MY_SOURCE_TEST",
		Sources: []*nats.StreamSource{
			{
				Name: "TEST",
				External: &nats.ExternalStream{
					APIPrefix:     "RI.JS.API",
					DeliverPrefix: "RI.DELIVER.SYNC.SOURCES",
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	checkFor(t, 20*time.Second, 100*time.Millisecond, func() error {
		si, err := js2.StreamInfo("MY_SOURCE_TEST")
		if err != nil {
			t.Fatalf("Could not retrieve stream info")
		}
		if si.State.Msgs != uint64(toSend) {
			return fmt.Errorf("Expected %d msgs, got state: %+v", toSend, si.State)
		}
		return nil
	})

}

func TestJetStreamFailMirrorsAndSources(t *testing.T) {
	c := createJetStreamClusterWithTemplate(t, jsClusterMirrorSourceImportsTempl, "C1", 3)
	defer c.shutdown()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			resp, err := nc2.Request(fmt.Sprintf(JSApiStreamCreateT, cfg.Name), req, time.Second)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			var scResp JSApiStreamCreateResponse
			if err := json.Unmarshal(resp.Data, &scResp); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if scResp.Error == nil {
				t.Fatalf("Did expect an error but got none")
			} else if !IsNatsErr(scResp.Error, id) {
				t.Fatalf("Expected different error: %s", scResp.Error.Description)
			}
		})
	}

	testPrefix("mirror-bad-deliverprefix", JSStreamExternalDelPrefixOverlapsErrF, StreamConfig{
		Name:    "MY_MIRROR_TEST",
		Storage: FileStorage,
		Mirror: &StreamSource{
			Name: "TEST",
			External: &ExternalStream{
				ApiPrefix: "RI.JS.API",
				// this will result in test.test.> which test.> would match
				DeliverPrefix: "test",
			},
		},
	})
	testPrefix("mirror-bad-apiprefix", JSStreamExternalApiOverlapErrF, StreamConfig{
		Name:    "MY_MIRROR_TEST",
		Storage: FileStorage,
		Mirror: &StreamSource{
			Name: "TEST",
			External: &ExternalStream{
				ApiPrefix:     "$JS.API",
				DeliverPrefix: "here",
			},
		},
	})
	testPrefix("source-bad-deliverprefix", JSStreamExternalDelPrefixOverlapsErrF, StreamConfig{
		Name:    "MY_SOURCE_TEST",
		Storage: FileStorage,
		Sources: []*StreamSource{{
			Name: "TEST",
			External: &ExternalStream{
				ApiPrefix:     "RI.JS.API",
				DeliverPrefix: "test",
			},
		},
		},
	})
	testPrefix("source-bad-apiprefix", JSStreamExternalApiOverlapErrF, StreamConfig{
		Name:    "MY_SOURCE_TEST",
		Storage: FileStorage,
		Sources: []*StreamSource{{
			Name: "TEST",
			External: &ExternalStream{
				ApiPrefix:     "$JS.API",
				DeliverPrefix: "here",
			},
		},
		},
	})
}

func TestJetStreamClusterJSAPIImport(t *testing.T) {
	c := createJetStreamClusterWithTemplate(t, jsClusterImportsTempl, "C1", 3)
	defer c.shutdown()

	// Client based API - This will connect to the non-js account which imports JS.
	// Connect below does an AccountInfo call.
	s := c.randomNonLeader()
	nc, js := jsClientConnect(t, s)
	defer nc.Close()

	if _, err := js.AddStream(&nats.StreamConfig{Name: "TEST", Replicas: 2}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Note if this was ephemeral we would need to setup export/import for that subject.
	sub, err := js.SubscribeSync("TEST", nats.Durable("dlc"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
		t.Fatalf("Unexpected error: %v", err)
	}
	meta, err = m.Metadata()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if meta.Sequence.Consumer != 2 || meta.Sequence.Stream != 2 || meta.NumDelivered != 1 || meta.NumPending != 1 {
		t.Fatalf("Bad meta: %+v", meta)
	}

	// AckNext
	_, err = nc.Request(m.Reply, []byte("+NXT"), 2*time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestJetStreamClusterSuperClusterInterestOnlyMode(t *testing.T) {
	template := `
		listen: 127.0.0.1:-1
		server_name: %s
		jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}
		accounts {
			one {
				jetstream: enabled
				users [{user: one, login_phrase: gjnuwofb}]
			}
			two {
				%s
				users [{user: two, password: yxkyuujd}]
			}
		}
		cluster {
			listen: 127.0.0.1:%d
			name: %s
			routes = ["nats://127.0.0.1:%d"]
		}
		gateway {
			name: %s
			listen: 127.0.0.1:%d
			gateways = [{name: %s, urls: ["nats://127.0.0.1:%d"]}]
		}
	`
	storeDir1 := createDir(t, JetStreamStoreDir)
	conf1 := createConfFile(t, []byte(fmt.Sprintf(template,
		"S1", storeDir1, "", 33222, "A", 33222, "A", 11222, "B", 11223)))
	s1, o1 := RunServerWithConfig(conf1)
	defer s1.Shutdown()

	storeDir2 := createDir(t, JetStreamStoreDir)
	conf2 := createConfFile(t, []byte(fmt.Sprintf(template,
		"S2", storeDir2, "", 33223, "B", 33223, "B", 11223, "A", 11222)))
	s2, o2 := RunServerWithConfig(conf2)
	defer s2.Shutdown()

	waitForInboundGateways(t, s1, 1, 2*time.Second)
	waitForInboundGateways(t, s2, 1, 2*time.Second)
	waitForOutboundGateways(t, s1, 1, 2*time.Second)
	waitForOutboundGateways(t, s2, 1, 2*time.Second)

	nc1 := natsConnect(t, fmt.Sprintf("nats://two:kpvrsjyf@127.0.0.1:%d", o1.Port))
	defer nc1.Close()
	nc1.Publish("foo", []byte("some message"))
	nc1.Flush()

	nc2 := natsConnect(t, fmt.Sprintf("nats://two:lipqxwff@127.0.0.1:%d", o2.Port))
	defer nc2.Close()
	nc2.Publish("bar", []byte("some message"))
	nc2.Flush()

	checkMode := func(accName string, expectedMode GatewayInterestMode) {
		t.Helper()
		checkFor(t, 2*time.Second, 15*time.Millisecond, func() error {
			servers := []*Server{s1, s2}
			for _, s := range servers {
				var gws []*client
				s.getInboundGatewayConnections(&gws)
				for _, gw := range gws {
					var mode GatewayInterestMode
					gw.mu.Lock()
					ie := gw.gw.insim[accName]
					if ie != nil {
						mode = ie.mode
					}
					gw.mu.Unlock()
					if ie == nil {
						return fmt.Errorf("Account %q not in map", accName)
					}
					if mode != expectedMode {
						return fmt.Errorf("Expected account %q mode to be %v, got: %v", accName, expectedMode, mode)
					}
	}

	// Make sure we can add a stream, etc.
	si, err = js.AddStream(&nats.StreamConfig{
		Name:     "TEST22",
		Subjects: []string{"bar"},
		Replicas: 2,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if si.Cluster == nil || si.Cluster.Name != "HUB" {
		t.Fatalf("Expected stream to be placed in %q", "HUB")
	}

	jsLocal, err := nc.JetStream()
	if err != nil {
		t.Fatalf("Unexpected error getting JetStream context: %v", err)
	}

	// Create a mirror on the local leafnode for stream TEST22.
	_, err = jsLocal.AddStream(&nats.StreamConfig{
		Name: "M",
		Mirror: &nats.StreamSource{
			Name:     "TEST22",
			External: &nats.ExternalStream{APIPrefix: "$JS.HUB.API"},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Publish a message to the HUB's TEST22 stream.
	if _, err := js.Publish("bar", []byte("OK")); err != nil {
		t.Fatalf("Unexpected publish error: %v", err)
	}
	// Make sure the message arrives in our mirror.
	checkFor(t, 2*time.Second, 100*time.Millisecond, func() error {
		si, err := jsLocal.StreamInfo("M")
		if err != nil {
			return fmt.Errorf("Could not get stream info: %v", err)
		}
		if si.State.Msgs != 1 {
			return fmt.Errorf("Expected 1 msg, got state: %+v", si.State)
		}
		return nil
	})

	// Now do the reverse and create a sourced stream in the HUB from our local stream on leafnode.
	// Inside the HUB we need to be able to find our local leafnode JetStream assets, so we need
	// a mapping in the LN server to allow this to work. Normally this will just be in server config.
	acc, err := ln.LookupAccount("JSY")
	if err != nil {
		c.t.Fatalf("Unexpected error on %v: %v", ln, err)
	}
	if err := acc.AddMapping("$JS.LN.API.>", "$JS.API.>"); err != nil {
		c.t.Fatalf("Error adding mapping: %v", err)
	}

	// js is the HUB JetStream context here.
	_, err = js.AddStream(&nats.StreamConfig{
		Name: "S",
		Sources: []*nats.StreamSource{{
			Name:     "M",
			External: &nats.ExternalStream{APIPrefix: "$JS.LN.API"},
		}},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Make sure the message arrives in our sourced stream.
	checkFor(t, 2*time.Second, 100*time.Millisecond, func() error {
		si, err := js.StreamInfo("S")
		if err != nil {
			return fmt.Errorf("Could not get stream info: %v", err)
		}
		if si.State.Msgs != 1 {
			return fmt.Errorf("Expected 1 msg, got state: %+v", si.State)
		}
		return nil
	})
}

// JetStream Domains
func TestJetStreamClusterDomains(t *testing.T) {
	// This adds in domain config option to template.
	// jetstream: {max_mem_store: 256MB, max_file_store: 2GB, domain: CORE, store_dir: "%s"}
	tmpl := strings.Replace(jsClusterAccountsTempl, "store_dir:", "domain: CORE, store_dir:", 1)
	c := createJetStreamCluster(t, tmpl, "CORE", _EMPTY_, 3, 12232, true)
		t.Fatalf("Unexpected error: %v", err)
	}
	// Make sure we can add a stream, etc.
	si, err = jsCore.AddStream(&nats.StreamConfig{
		Name:     "TEST22",
		Subjects: []string{"bar"},
		Replicas: 2,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if si.Cluster == nil || si.Cluster.Name != "CORE" {
		t.Fatalf("Expected stream to be placed in %q, got %q", "CORE", si.Cluster.Name)
	}

	jsLocal, err := nc.JetStream()
	if err != nil {
		t.Fatalf("Unexpected error getting JetStream context: %v", err)
	}

	// Create a mirror on our local leafnode for stream TEST22.
	_, err = jsLocal.AddStream(&nats.StreamConfig{
		Name: "M",
		Mirror: &nats.StreamSource{
			Name:     "TEST22",
			External: &nats.ExternalStream{APIPrefix: "$JS.CORE.API"},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Publish a message to the CORE's TEST22 stream.
	if _, err := jsCore.Publish("bar", []byte("OK")); err != nil {
		t.Fatalf("Unexpected publish error: %v", err)
	}

	// Make sure the message arrives in our mirror.
	checkFor(t, 2*time.Second, 100*time.Millisecond, func() error {
		si, err := jsLocal.StreamInfo("M")
		if err != nil {
			return fmt.Errorf("Could not get stream info: %v", err)
		}
		if si.State.Msgs != 1 {
			return fmt.Errorf("Expected 1 msg, got state: %+v", si.State)
		}
		return nil
	})

	// jsCore is the CORE JetStream domain.
	// Create a sourced stream in the CORE that is sourced from our mirror stream in our leafnode.
	_, err = jsCore.AddStream(&nats.StreamConfig{
		Name: "S",
		Sources: []*nats.StreamSource{{
			Name:     "M",
			External: &nats.ExternalStream{APIPrefix: "$JS.SPOKE.API"},
		}},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Make sure the message arrives in our sourced stream.
	checkFor(t, 2*time.Second, 100*time.Millisecond, func() error {
		si, err := jsCore.StreamInfo("S")
		if err != nil {
			return fmt.Errorf("Could not get stream info: %v", err)
		}
		if si.State.Msgs != 1 {
			return fmt.Errorf("Expected 1 msg, got state: %+v", si.State)
		}
		return nil
	})

	// Now connect directly to the CORE cluster and make sure we can operate there.
	nc, jsLocal = jsClientConnect(t, c.randomServer())
	defer nc.Close()

	// Create the js contexts again.
	jsSpoke, err := nc.JetStream(nats.APIPrefix("$JS.SPOKE.API"))
	if err != nil {
	}

	// Create the same name stream in both spoke domains.
	for _, s := range []*Server{spoke1, spoke2} {
		nc, js := jsClientConnect(t, s)
		defer nc.Close()
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     "TEST",
			Subjects: []string{subjFor(s)},
		})
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		nc.Close()
	}

	// Now connect to the hub and create a sourced stream from both leafnode streams.
	nc, js := jsClientConnect(t, c.randomServer())
	defer nc.Close()

	_, err := js.AddStream(&nats.StreamConfig{
		Name: "S",
		Sources: []*nats.StreamSource{
			{
				Name:     "TEST",
				External: &nats.ExternalStream{APIPrefix: "$JS.SPOKE-1.API"},
			},
			{
				Name:     "TEST",
				External: &nats.ExternalStream{APIPrefix: "$JS.SPOKE-2.API"},
			},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Publish a message to each spoke stream and we will check that our sourced stream gets both.
	for _, s := range []*Server{spoke1, spoke2} {
		nc, js := jsClientConnect(t, s)
		defer nc.Close()
		js.Publish(subjFor(s), []byte("DOUBLE TROUBLE"))
		si, err := js.StreamInfo("TEST")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if si.State.Msgs != 1 {
			t.Fatalf("Expected 1 msg, got %d", si.State.Msgs)
		}
		nc.Close()
	}

	// Now make sure we have 2 msgs in our sourced stream.
	si, err := js.StreamInfo("S")
	if err != nil {
	// This not in the Go client yet so manual for now.
	resp, err := nc.Request(fmt.Sprintf(JSApiStreamInfoT, "S"), nil, time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	var ssi StreamInfo
	if err = json.Unmarshal(resp.Data, &ssi); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(ssi.Sources) != 2 {
		t.Fatalf("Expected 2 source streams, got %d", len(ssi.Sources))
	}
	if ssi.Sources[0].External == nil {
		t.Fatalf("Expected a non-nil external designation")
	}
	pre := ssi.Sources[0].External.ApiPrefix
	if pre != "$JS.SPOKE-1.API" && pre != "$JS.SPOKE-2.API" {
		t.Fatalf("Expected external api of %q, got %q", "$JS.SPOKE-[1|2].API", ssi.Sources[0].External.ApiPrefix)
	}

	// Also create a mirror.
	_, err = js.AddStream(&nats.StreamConfig{
		Name: "M",
		Mirror: &nats.StreamSource{
			Name:     "TEST",
			External: &nats.ExternalStream{APIPrefix: "$JS.SPOKE-1.API"},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, err = nc.Request(fmt.Sprintf(JSApiStreamInfoT, "M"), nil, time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err = json.Unmarshal(resp.Data, &ssi); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if ssi.Mirror == nil || ssi.Mirror.External == nil {
		t.Fatalf("Expected a non-nil external designation for our mirror")
	}
	if ssi.Mirror.External.ApiPrefix != "$JS.SPOKE-1.API" {
		t.Fatalf("Expected external api of %q, got %q", "$JS.SPOKE-1.API", ssi.Sources[0].External.ApiPrefix)
	}
}

// When a leafnode enables JS on an account that is not enabled on the remote cluster account this should
// still work. Early NGS beta testers etc.
func TestJetStreamClusterSingleLeafNodeEnablingJetStream(t *testing.T) {
	tmpl := strings.Replace(jsClusterAccountsTempl, "store_dir:", "domain: HUB, store_dir:", 1)
	c := createJetStreamCluster(t, tmpl, "HUB", _EMPTY_, 3, 11322, true)
	defer c.shutdown()

	ln := c.createSingleLeafNodeNoSystemAccountAndEnablesJetStream()
	defer ln.Shutdown()

	// Check that we have JS in the $G account on the leafnode.
	nc, js := jsClientConnect(t, ln)
	defer nc.Close()

	if _, err := js.AccountInfo(); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Connect our client to the "nojs" account in the cluster but make sure JS works since its enabled via the leafnode.
	s := c.randomServer()
	nc, js = jsClientConnect(t, s, nats.UserInfo("nojs", "p"))
	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	accounts {
		JS {
			jetstream: enabled
			users = [ { user: "rip", pass: "pass" } ]
			exports [
				{ service: "$JS.API.CONSUMER.INFO.>" }
				{ service: "$JS.HUB.API.CONSUMER.>", response: stream }
				{ stream: "M.SYNC.>" } # For the mirror
			]
		}
		IA {
			jetstream: enabled
			users = [ { user: "dlc", pass: "pass" } ]
			imports [
				{ service: { account: JS, subject: "$JS.API.CONSUMER.INFO.TEST.DLC"}, to: "FROM.DLC" }
				{ service: { account: JS, subject: "$JS.HUB.API.CONSUMER.>"}, to: "js.xacc.API.CONSUMER.>" }
				{ stream: { account: JS, subject: "M.SYNC.>"} }
			]
		}
		$SYS { users = [ { user: "admin", pass: "s4kc6c!" } ] }
	}
	`

	c := createJetStreamClusterWithTemplate(t, template, "HUB", 3)
	defer c.shutdown()

	// Create the stream and the consumer under the JS/rip user.
	s := c.randomServer()
	nc, js := jsClientConnect(t, s, nats.UserInfo("rip", "pass"))
	defer nc.Close()

	if _, err := js.AddStream(&nats.StreamConfig{Name: "TEST", Replicas: 2}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	_, err := js.AddConsumer("TEST", &nats.ConsumerConfig{Durable: "DLC", AckPolicy: nats.AckExplicitPolicy})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Also create a stream via the domain qualified API.
	js, err = nc.JetStream(nats.APIPrefix("$JS.HUB.API"))
	if err != nil {
		t.Fatalf("Unexpected error getting JetStream context: %v", err)
	}
	if _, err := js.AddStream(&nats.StreamConfig{Name: "ORDERS", Replicas: 2}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Now we want to access the consumer info from IA/dlc.
	nc2, js2 := jsClientConnect(t, c.randomServer(), nats.UserInfo("dlc", "pass"))
	defer nc2.Close()

	if _, err := nc2.Request("FROM.DLC", nil, time.Second); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Make sure domain mappings etc work across accounts.
	// Setup a mirror.
	_, err = js2.AddStream(&nats.StreamConfig{
		Name: "MIRROR",
		Mirror: &nats.StreamSource{
			Name: "ORDERS",
			External: &nats.ExternalStream{
				APIPrefix:     "js.xacc.API",
				DeliverPrefix: "M.SYNC",
			},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Send 10 messages..
	msg, toSend := []byte("Hello mapped domains"), 10
	for i := 0; i < toSend; i++ {
		if _, err = js.Publish("ORDERS", msg); err != nil {
			t.Fatalf("Unexpected publish error: %v", err)
		}
	}

	checkFor(t, 2*time.Second, 100*time.Millisecond, func() error {
		si, err := js2.StreamInfo("MIRROR")
		if err != nil {
			return fmt.Errorf("Unexpected error: %v", err)
		}
		if si.State.Msgs != 10 {
			return fmt.Errorf("Expected 10 msgs, got state: %+v", si.State)
		}
		return nil
	spoke := c.createLeafNodeWithTemplate("SPOKE", tmpl)
	defer spoke.Shutdown()

	// Client for API requests.
	nc, js := jsClientConnect(t, spoke)
	defer nc.Close()

	_, err := js.AddStream(&nats.StreamConfig{
		Name:     "TEST",
		Subjects: []string{"foo"},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	jsHub, err := nc.JetStream(nats.APIPrefix("$JS.HUB.API"))
	if err != nil {
		t.Fatalf("Unexpected error getting JetStream context: %v", err)
	}

	_, err = jsHub.AddStream(&nats.StreamConfig{
		Name:     "S",
		Replicas: 2,
		Sources: []*nats.StreamSource{{
			Name:     "TEST",
			External: &nats.ExternalStream{APIPrefix: "$JS.SPOKE.API"},
		}},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Now send a message to the origin stream with nil body and a header.
	m := nats.NewMsg("foo")
	m.Header.Add("X-Request-ID", "a6f279d7-bbdd-2bfd-6733-2af309bc2c0b")
	m.Data = []byte("HELLO WORLD")

	if _, err = jsHub.PublishMsg(m); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	sub, err := jsHub.SubscribeSync("foo", nats.BindStream("S"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	msg, err := sub.NextMsg(time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if string(msg.Data) != "HELLO WORLD" {
		t.Fatalf("Message corrupt? Expecting %q got %q", "HELLO WORLD", msg.Data)

func (sc *supercluster) randomServer() *Server {
	return sc.randomCluster().randomServer()
}

var jsClusterAccountsTempl = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	leaf {
		listen: 127.0.0.1:-1
	}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	no_auth_user: one

	accounts {
		ONE { users = [ { user: "one", pass: "p" } ]; jetstream: enabled }
		TWO { users = [ { user: "two", pass: "p" } ]; jetstream: enabled }
		NOJS { users = [ { user: "nojs", pass: "p" } ] }
		$SYS { users = [ { user: "admin", pass: "x7eb3a!" } ] }
	}
`

var jsClusterTempl = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	leaf {
		listen: 127.0.0.1:-1
	}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	# For access to system account.
	accounts { $SYS { users = [ { user: "admin", pass: "n3qa7d!" } ] } }
`

var jsSuperClusterTempl = `
	%s
	gateway {
		name: %s
		listen: 127.0.0.1:%d
		gateways = [%s
		]
	}

	system_account: "$SYS"
`

var jsClusterLimitsTempl = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 2MB, max_file_store: 8MB, store_dir: "%s"}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	no_auth_user: u

	accounts {
		ONE {
			users = [ { user: "u", pass: "x6jj4z!" } ]
			jetstream: enabled
		}
		$SYS { users = [ { user: "admin", pass: "z0sm8d!" } ] }
	}
`

var jsGWTempl = `%s{name: %s, urls: [%s]}`

func createJetStreamSuperCluster(t *testing.T, numServersPer, numClusters int) *supercluster {
	t.Helper()
	if numServersPer < 1 {
		t.Fatalf("Number of servers must be >= 1")
	}
	if numClusters <= 1 {
		t.Fatalf("Number of clusters must be > 1")
	}

	startClusterPorts := []int{20_022, 22_022, 24_022}
	startGatewayPorts := []int{20_122, 22_122, 24_122}
	startClusterPort := startClusterPorts[rand.Intn(len(startClusterPorts))]
	startGWPort := startGatewayPorts[rand.Intn(len(startGatewayPorts))]

	// Make the GWs form faster for the tests.
	SetGatewaysSolicitDelay(10 * time.Millisecond)
	defer ResetGatewaysSolicitDelay()

	cp, gp := startClusterPort, startGWPort
	var clusters []*cluster
	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	no_auth_user: dlc

	accounts {
		JS {
			jetstream: enabled
			users = [ { user: "rip", pass: "pass" } ]
			exports [
				{ service: "$JS.API.CONSUMER.>" } # To create internal consumers to mirror/source.
				{ stream: "RI.DELIVER.SYNC.>" }   # For the mirror/source consumers sending to IA via delivery subject.
			]
		}
		IA {
			jetstream: enabled
			users = [ { user: "dlc", pass: "pass" } ]
			imports [
				{ service: { account: JS, subject: "$JS.API.CONSUMER.>"}, to: "RI.JS.API.CONSUMER.>" }
				{ stream: { account: JS, subject: "RI.DELIVER.SYNC.>"} }
			]
		}
		$SYS { users = [ { user: "admin", pass: "m6ey3o!" } ] }
	}
`

var jsClusterImportsTempl = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	no_auth_user: dlc

	accounts {
		JS {
			jetstream: enabled
			users = [ { user: "rip", pass: "pass" } ]
			exports [
				{ service: "$JS.API.>", response: stream }
				{ service: "TEST" } # For publishing to the stream.
				{ service: "$JS.ACK.TEST.*.>" }
			]
		}
		IA {
			users = [ { user: "dlc", pass: "pass" } ]
			imports [
				{ service: { subject: "$JS.API.>", account: JS }}
				{ service: { subject: "TEST", account: JS }}
				{ service: { subject: "$JS.ACK.TEST.*.>", account: JS }}
			]
		}
		$SYS { users = [ { user: "admin", pass: "p2ya8a!" } ] }
	}
`

func createMixedModeCluster(t *testing.T, clusterName string, numJsServers, numNonServers int) *cluster {
	if clusterName == "" || numJsServers < 1 || numNonServers < 1 {
		t.Fatalf("Bad params")
	}

	numServers := numJsServers + numNonServers
	const startClusterPort = 23232

	// Build out the routes that will be shared with all configs.
	var routes []string
	for cp := startClusterPort; cp < startClusterPort+numServers; cp++ {
		routes = append(routes, fmt.Sprintf("nats-route://127.0.0.1:%d", cp))
	}
	routeConfig := strings.Join(routes, ",")

	// Go ahead and build configurations and start servers.
	c := &cluster{servers: make([]*Server, 0, numServers), opts: make([]*Options, 0, numServers), name: clusterName}

	for cp := startClusterPort; cp < startClusterPort+numServers; cp++ {
		storeDir := createDir(t, JetStreamStoreDir)
		sn := fmt.Sprintf("S-%d", cp-startClusterPort+1)
		conf := fmt.Sprintf(jsClusterLimitsTempl, sn, storeDir, clusterName, cp, routeConfig)

	return s
}

var jsClusterSingleLeafNodeLikeNGSTempl = `
	listen: 127.0.0.1:-1
	server_name: LNJS
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	leaf { remotes [ { urls: [ %s ] } ] }
`

var jsClusterSingleLeafNodeTempl = `
	listen: 127.0.0.1:-1
	server_name: LNJS
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	leaf { remotes [
		{ urls: [ %s ], account: "JSY" }
		{ urls: [ %s ], account: "JSN" } ]
	}

	accounts {
		JSY { users = [ { user: "y", pass: "p" } ]; jetstream: true }
		JSN { users = [ { user: "n", pass: "p" } ] }
		$SYS { users = [ { user: "admin", pass: "z2ap0d!" } ] }
	}
`

var jsClusterTemplWithLeafNode = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	{{leaf}}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	# For access to system account.
	accounts { $SYS { users = [ { user: "admin", pass: "q8zt8z!" } ] } }
`

var jsClusterTemplWithSingleLeafNode = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	{{leaf}}

	# For access to system account.
	accounts { $SYS { users = [ { user: "admin", pass: "v0lw3j!" } ] } }
`

var jsLeafFrag = `
	leaf {
		remotes [
			{ urls: [ %s ] }
			{ urls: [ %s ], account: "$SYS" }
		]
	}
`

func (c *cluster) createLeafNodes(clusterName string, numServers int) *cluster {
	return c.createLeafNodesWithStartPort(clusterName, numServers, 22111)
}

func (c *cluster) createLeafNodesWithStartPort(clusterName string, numServers int, portStart int) *cluster {
	return c.createLeafNodesWithTemplateAndStartPort(jsClusterTemplWithLeafNode, clusterName, numServers, portStart)
}

func (c *cluster) createLeafNode() *Server {
	return c.createLeafNodeWithTemplate("LNS", jsClusterTemplWithSingleLeafNode)
}

func (c *cluster) createLeafNodeWithTemplate(name, template string) *Server {
	tmpl := c.createLeafSolicit(template)
	conf := fmt.Sprintf(tmpl, name, createDir(c.t, JetStreamStoreDir))
	s, o := RunServerWithConfig(createConfFile(c.t, []byte(conf)))
	c.servers = append(c.servers, s)
	c.opts = append(c.opts, o)
	return s
}

// Helper to generate the leaf solicit configs.
func (c *cluster) createLeafSolicit(tmpl string) string {
	// Create our leafnode cluster template first.
	var lns, lnss []string
	for _, s := range c.servers {
		if s.ClusterName() != c.name {
			continue
		}
		ln := s.getOpts().LeafNode
		lns = append(lns, fmt.Sprintf("nats://%s:%d", ln.Host, ln.Port))
		lnss = append(lnss, fmt.Sprintf("nats://admin:p5hy4j!@%s:%d", ln.Host, ln.Port))
	}
	lnc := strings.Join(lns, ", ")
	lnsc := strings.Join(lnss, ", ")
	lconf := fmt.Sprintf(jsLeafFrag, lnc, lnsc)
	return strings.Replace(tmpl, "{{leaf}}", lconf, 1)
}

func (c *cluster) createLeafNodesWithTemplateAndStartPort(template, clusterName string, numServers int, portStart int) *cluster {
	// Create our leafnode cluster template first.
	tmpl := c.createLeafSolicit(template)
	pre := clusterName + "-"
	lc := createJetStreamCluster(c.t, tmpl, clusterName, pre, numServers, portStart, false)
	for _, s := range lc.servers {
		checkLeafNodeConnectedCount(c.t, s, 2)
	}
	return lc
}

// Will add in the mapping for the account to each server.
func (c *cluster) addSubjectMapping(account, src, dest string) {
	for _, s := range c.servers {
		if s.ClusterName() != c.name {
			continue
		}
		acc, err := s.LookupAccount(account)