			}
			if igc.Connection.Name != sc.ID() {
				t.Fatalf("mode=%v - Expected outbound connection to C to have name %q, got %q", pollMode, sc.ID(), og.Connection.Name)
			}
		}
		return nil
	})

	// Select only 1 gateway by passing the name to option/url
	opts := &GatewayzOptions{Name: "B"}
	checkGatewayB(t, gatewayzURL+"?gw_name=B", opts)

	// Stop gateway C and check that we have only B, with and without filter.
	sc.Shutdown()
	checkGatewayB(t, gatewayzURL+"?gw_name=B", opts)
	checkGatewayB(t, gatewayzURL, nil)
}

func TestMonitorGatewayzAccounts(t *testing.T) {
	resetPreviousHTTPConnections()

	// Create bunch of Accounts
	totalAccounts := 15
	accounts := ""
	for i := 0; i < totalAccounts; i++ {
		acc := fmt.Sprintf("	acc_%d: { users=[{user:user_%d, password: pwd}] }\n", i, i)
		accounts += acc
	}

	bConf := createConfFile(t, []byte(fmt.Sprintf(`
		accounts {
			%s
		}
		port: -1
		http: -1
		gateway: {
			name: "B"
			port: -1
		}
		no_sys_acc = true
	`, accounts)))
	defer removeFile(t, bConf)

	sb, ob := RunServerWithConfig(bConf)
	defer sb.Shutdown()
	sb.SetLogger(&DummyLogger{}, true, true)

	// Start a1 that has a single URL to sb1.
	aConf := createConfFile(t, []byte(fmt.Sprintf(`
		accounts {
			%s
	srvs := []*Server{}
	for _, test := range []struct {
		port   int
		mport  int
		cport  int
		routed int
	}{
		{7500, 7501, 7502, 5502},
		{5500, 5501, 5502, 7502},
	} {
		tmpDir := createDir(t, fmt.Sprintf("srv_%d", test.port))
		defer removeDir(t, tmpDir)
		cf := createConfFile(t, []byte(fmt.Sprintf(`
		listen: 127.0.0.1:%d
		http: 127.0.0.1:%d
		system_account: SYS
		accounts {
			SYS {
				users [{user: sys, password: pwd}]
			}
			ACC {
				users [{user: usr, password: pwd}]
				jetstream: enabled
			}
			BCC_TO_HAVE_ONE_EXTRA {
				users [{user: usr2, password: pwd}]
				jetstream: enabled
			}
		}
		jetstream: {
			max_mem_store: 10Mb
			max_file_store: 10Mb
			store_dir: %s
		}
		cluster {
			name: cluster_name
			listen: 127.0.0.1:%d
			routes: [nats-route://127.0.0.1:%d]
		}
		server_name: server_%d `, test.port, test.mport, tmpDir, test.cport, test.routed, test.port)))
		defer removeFile(t, cf)

		s, _ := RunServerWithConfig(cf)
		defer s.Shutdown()
		srvs = append(srvs, s)
	}
	checkClusterFormed(t, srvs...)
	checkForJSClusterUp(t, srvs...)

	nc := natsConnect(t, "nats://usr:pwd@127.0.0.1:7500")
	defer nc.Close()
	js, err := nc.JetStream(nats.MaxWait(5 * time.Second))
	require_NoError(t, err)
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "my-stream-replicated",
		Subjects: []string{"foo", "bar"},
		Replicas: 2,
	})
	require_NoError(t, err)
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "my-stream-non-replicated",
		Subjects: []string{"baz"},
		Replicas: 1,
	})
	require_NoError(t, err)
	_, err = js.AddConsumer("my-stream-replicated", &nats.ConsumerConfig{
		Durable:   "my-consumer-replicated",
		AckPolicy: nats.AckExplicitPolicy,
	})
	require_NoError(t, err)
	_, err = js.AddConsumer("my-stream-non-replicated", &nats.ConsumerConfig{
		Durable:   "my-consumer-non-replicated",
		AckPolicy: nats.AckExplicitPolicy,
	})
	require_NoError(t, err)