
	if config := s.JetStreamConfig(); config != nil {
		defer removeDir(t, config.StoreDir)
	}

	sa := s.SystemAccount()
	if err := sa.EnableJetStream(nil); err == nil {
		t.Fatalf("Expected an error trying to enable on the system account")
	}
}

func TestJetStreamMultipleAccountsBasics(t *testing.T) {
	conf := createConfFile(t, []byte(`
		listen: 127.0.0.1:-1
		jetstream: {max_mem_store: 64GB, max_file_store: 10TB}
		accounts: {
			A: {
				jetstream: enabled
				users: [ {user: ua, password: pwd} ]
			},
			B: {
				jetstream: {max_mem: 1GB, max_store: 1TB, max_streams: 10, max_consumers: 1k}
				users: [ {user: ub, password: pwd} ]
			},
			C: {
				users: [ {user: uc, password: pwd} ]
			},
		}
	`))
	defer removeFile(t, conf)

	s, opts := RunServerWithConfig(conf)
	defer s.Shutdown()

	if config := s.JetStreamConfig(); config != nil {
		defer removeDir(t, config.StoreDir)
	}

	if !s.JetStreamEnabled() {
		t.Fatalf("Expected JetStream to be enabled")
	}

	nca := clientConnectToServerWithUP(t, opts, "ua", "pwd")
	defer nca.Close()

	ncb := clientConnectToServerWithUP(t, opts, "ub", "pwd")
	defer ncb.Close()

	resp, err := ncb.Request(JSApiAccountInfo, nil, time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
			t.Fatalf("No response, possible timeout?")
		}
		var iResp JSApiAccountInfoResponse
		if err := json.Unmarshal(resp.Data, &iResp); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if iResp.Error == nil {
			t.Fatalf("Expected an error on not enabled account")
		}
	}

	// Check C is not enabled. We expect a negative response, not a timeout.
	expectNotEnabled(ncc.Request(JSApiAccountInfo, nil, 250*time.Millisecond))

	// Now do simple reload and check that we do the right thing. Testing enable and disable and also change in limits
	newConf := []byte(`
		listen: 127.0.0.1:-1
		jetstream: {max_mem_store: 64GB, max_file_store: 10TB}
		accounts: {
			A: {
				jetstream: disabled
				users: [ {user: ua, password: pwd} ]
			},
			B: {
				jetstream: {max_mem: 32GB, max_store: 512GB, max_streams: 100, max_consumers: 4k}
				users: [ {user: ub, password: pwd} ]
			},
			C: {
				jetstream: {max_mem: 1GB, max_store: 1TB, max_streams: 10, max_consumers: 1k}
				users: [ {user: uc, password: pwd} ]
			},
		}
	`)
	if err := ioutil.WriteFile(conf, newConf, 0600); err != nil {
		t.Fatalf("Error rewriting server's config file: %v", err)
	}
	if err := s.Reload(); err != nil {
		t.Fatalf("Error on server reload: %v", err)
	}
	expectNotEnabled(nca.Request(JSApiAccountInfo, nil, 250*time.Millisecond))

	resp, _ = ncb.Request(JSApiAccountInfo, nil, 250*time.Millisecond)
	if err := json.Unmarshal(resp.Data, &info); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if info.Error != nil {
		t.Fatalf("Expected JetStream to be enabled, got %+v", info.Error)
	}

	resp, _ = ncc.Request(JSApiAccountInfo, nil, 250*time.Millisecond)
	if err := json.Unmarshal(resp.Data, &info); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if info.Error != nil {
		t.Fatalf("Expected JetStream to be enabled, got %+v", info.Error)
		t.Fatalf("Unexpected error: %+v", info.Error)
	}
	// Lookup streams.
	resp, err = nc.Request(mapSubj(JSApiStreams), nil, time.Second)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	var namesResponse JSApiStreamNamesResponse
	if err = json.Unmarshal(resp.Data, &namesResponse); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if namesResponse.Error != nil {
		t.Fatalf("Unexpected error: %+v", namesResponse.Error)
	}
}

// https://github.com/nats-io/nats-server/issues/1736
func TestJetStreamServerReload(t *testing.T) {
	conf := createConfFile(t, []byte(`
		listen: 127.0.0.1:-1
		jetstream: {max_mem_store: 64GB, max_file_store: 10TB }
		accounts: {
			A: { users: [ {user: ua, password: pwd} ] },
			B: {
				jetstream: {max_mem: 1GB, max_store: 1TB, max_streams: 10, max_consumers: 1k}
				users: [ {user: ub, password: pwd} ]
			},
			SYS: { users: [ {user: uc, password: pwd} ] },
		}
		no_auth_user: ub
		system_account: SYS
	`))
	defer removeFile(t, conf)

	s, _ := RunServerWithConfig(conf)
	defer s.Shutdown()

	if config := s.JetStreamConfig(); config != nil {
		defer removeDir(t, config.StoreDir)
	}

	if !s.JetStreamEnabled() {
		t.Fatalf("Expected JetStream to be enabled")
	}

	// Client for API requests.
	nc := clientConnectToServer(t, s)
	defer nc.Close()

	checkJSAccount := func() {
		t.Helper()
		resp, err := nc.Request(JSApiAccountInfo, nil, time.Second)
		if err != nil {
	if msgs := mset.state().Msgs; msgs != uint64(toSend) {
		t.Fatalf("Expected %d messages, got %d", toSend, msgs)
	}

	if err := s.Reload(); err != nil {
		t.Fatalf("Error on server reload: %v", err)
	}

	// Wait to get reconnected.
	checkFor(t, 5*time.Second, 10*time.Millisecond, func() error {
		if !nc.IsConnected() {
			return fmt.Errorf("Not connected")
		}
		return nil
	})

	checkJSAccount()
	sendStreamMsg(t, nc, "22", "MSG: 22")
}

func TestJetStreamConfigReloadWithGlobalAccount(t *testing.T) {
	template := `
		authorization {
			users [
				{user: anonymous}
				{user: user1, password: %s}
			]
		}
		no_auth_user: anonymous
		jetstream: enabled
	`
	conf := createConfFile(t, []byte(fmt.Sprintf(template, "pwd")))
	defer removeFile(t, conf)

	s, _ := RunServerWithConfig(conf)
	defer s.Shutdown()

	if config := s.JetStreamConfig(); config != nil {
		defer removeDir(t, config.StoreDir)
	}

	// Client for API requests.
	nc, js := jsClientConnect(t, s)
	defer nc.Close()

	checkJSAccount := func() {
		t.Helper()
		if _, err := js.AccountInfo(); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	}

	checkJSAccount()
