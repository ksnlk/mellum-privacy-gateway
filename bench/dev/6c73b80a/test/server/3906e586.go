	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/nats-io/nuid"
)

var testMQTTTimeout = 4 * time.Second

var jsClusterTemplWithLeafAndMQTT = `
	listen: 127.0.0.1:-1
	server_name: %s
	jetstream: {max_mem_store: 256MB, max_file_store: 2GB, store_dir: "%s"}

	{{leaf}}

	cluster {
		name: %s
		listen: 127.0.0.1:%d
		routes = [%s]
	}

	mqtt {
		listen: 127.0.0.1:-1
	}

	# For access to system account.
	accounts { $SYS { users = [ { user: "admin", pass: "s3vy1g!" } ] } }
`

type mqttErrorReader struct {
	err error
}

func (r *mqttErrorReader) Read(b []byte) (int, error)      { return 0, r.err }
func (r *mqttErrorReader) SetReadDeadline(time.Time) error { return nil }

func testNewEOFReader() *mqttErrorReader {
	return &mqttErrorReader{err: io.EOF}
}

func TestMQTTReader(t *testing.T) {
	r := &mqttReader{}
	r.reset([]byte{0, 2, 'a', 'b'})
	bs, err := r.readBytes("", false)
	if err != nil {
		t.Fatal(err)
	}
	sbs := string(bs)
	if sbs != "ab" {
		t.Fatalf(`expected "ab", got %q`, sbs)
	}

	s, o := RunServerWithConfig(conf)
	defer testMQTTShutdownServer(s)
	if o.MQTT.TLSConfig == nil {
		t.Fatal("expected TLS config to be set")
	}
}

func TestMQTTValidateOptions(t *testing.T) {
	nmqtto := DefaultOptions()
	mqtto := testMQTTDefaultOptions()
	for _, test := range []struct {
		name    string
		getOpts func() *Options
		err     error
	}{
		{"mqtt disabled", func() *Options { return nmqtto.Clone() }, nil},
		{"mqtt username not allowed if users specified", func() *Options {
			o := mqtto.Clone()
			o.Users = []*User{{Username: "abc", Password: "pwd"}}
			o.MQTT.Username = "b"
			o.MQTT.Password = "pwd"
			return o
		}, errMQTTUserMixWithUsersNKeys},
		{"mqtt token not allowed if users specified", func() *Options {
			o := mqtto.Clone()
			o.Nkeys = []*NkeyUser{{Nkey: "abc"}}
			o.MQTT.Token = "mytoken"
			return o
		}, errMQTTTokenMixWIthUsersNKeys},
		{"ack wait should be >=0", func() *Options {
			o := mqtto.Clone()
			o.MQTT.AckWait = -10 * time.Second
			return o
		}, errMQTTAckWaitMustBePositive},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateMQTTOptions(test.getOpts())
			if test.err == nil && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			} else if test.err != nil && (err == nil || err.Error() != test.err.Error()) {
				t.Fatalf("Expected error to contain %q, got %v", test.err, err)
			}
		})
	}
}

func TestMQTTParseOptions(t *testing.T) {
	for _, test := range []struct {
		name     string
		content  string
		checkOpt func(*MQTTOpts) error
			return nil
		}, ""},
		{"port", `mqtt { port: 1234 }`, func(o *MQTTOpts) error {
			if o.Port != 1234 {
				return fmt.Errorf("expected 1234, got %v", o.Port)
			}
			return nil
		}, ""},
		{"tls config",
			`
			mqtt {
				tls {
					cert_file: "./configs/certs/server.pem"
					key_file: "./configs/certs/key.pem"
				}
			}
			`, func(o *MQTTOpts) error {
				if o.TLSConfig == nil {
					return fmt.Errorf("TLSConfig should have been set")
				}
				return nil
			}, ""},
		{"no auth user",
			`
			mqtt {
				no_auth_user: "noauthuser"
			}
			`, func(o *MQTTOpts) error {
				if o.NoAuthUser != "noauthuser" {
					return fmt.Errorf("Invalid NoAuthUser value: %q", o.NoAuthUser)
				}
				return nil
			}, ""},
		{"auth block",
			`
			mqtt {
				authorization {
					user: "mqttuser"
					password: "pwd"
					token: "token"
					timeout: 2.0
				}
			}
			`, func(o *MQTTOpts) error {
				if o.Username != "mqttuser" || o.Password != "pwd" || o.Token != "token" || o.AuthTimeout != 2.0 {
					return fmt.Errorf("Invalid auth block: %+v", o)
				}
				return nil
			}, ""},
		{"auth timeout as int",
			`
			mqtt {
				authorization {
					timeout: 2
				}
			}
			`, func(o *MQTTOpts) error {
				if o.AuthTimeout != 2.0 {
					return fmt.Errorf("Invalid auth timeout: %v", o.AuthTimeout)
				}
				return nil
			}, ""},
		{"ack wait",
			`
			mqtt {
				ack_wait: "10s"
			}
			`, func(o *MQTTOpts) error {
				if o.AckWait != 10*time.Second {
					return fmt.Errorf("Invalid ack wait: %v", o.AckWait)
			}
			c.mu.Unlock()
			if uname != certUserName {
				t.Fatalf("Expected username %q, got %q", certUserName, uname)
			}
			if accname != accName {
				t.Fatalf("Expected account %q, got %v", accName, accname)
			}
		})
	}
}

func TestMQTTBasicAuth(t *testing.T) {
	for _, test := range []struct {
		name string
		opts func() *Options
		user string
		pass string
		rc   byte
	}{
		{
			"top level auth, no override, wrong u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Username = "normal"
				o.Password = "icjoiu"
				return o
			},
			"mqtt", "client", mqttConnAckRCNotAuthorized,
		},
		{
			"top level auth, no override, correct u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Username = "normal"
				o.Password = "rtvffr"
				return o
			},
			"normal", "client", mqttConnAckRCConnectionAccepted,
		},
		{
			"no top level auth, mqtt auth, wrong u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.MQTT.Username = "mqtt"
				o.MQTT.Password = "yfpgnx"
				return o
			},
			"normal", "client", mqttConnAckRCNotAuthorized,
		},
		{
			"no top level auth, mqtt auth, correct u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.MQTT.Username = "mqtt"
				o.MQTT.Password = "wrmvzr"
				o.MQTT.Password_name = "payments-db-credentials"
				return o
			},
			"mqtt", "client", mqttConnAckRCConnectionAccepted,
		},
		{
			"top level auth, mqtt override, wrong u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Username = "normal"
				o.Password = "kgmhcp"
				o.MQTT.Username = "mqtt"
				o.MQTT.Password = "bqfabp"
				return o
			},
			"normal", "client", mqttConnAckRCNotAuthorized,
		},
		{
			"top level auth, mqtt override, correct u/p",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Username = "normal"
				o.Password = "hvinwe"
				o.MQTT.Username = "mqtt"
				o.MQTT.Password = "dither"
				return o
			},
			"mqtt", "client", mqttConnAckRCConnectionAccepted,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := test.opts()
			s := testMQTTRunServer(t, o)
			defer testMQTTShutdownServer(s)

			ci := &mqttConnInfo{
				cleanSess: true,
				user:      test.user,
				pass:      test.pass,
			}
			mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
			defer mc.Close()
			testMQTTCheckConnAck(t, r, test.rc, false)
		})
	}
}

func TestMQTTAuthTimeout(t *testing.T) {
	for _, test := range []struct {
		name string
		at   float64
		mat  float64
		ok   bool
	}{
		{"use top-level auth timeout", 0.5, 0.0, true},
		{"use mqtt auth timeout", 0.5, 0.05, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := testMQTTDefaultOptions()
			o.AuthTimeout = test.at
			o.MQTT.Username = "mqtt"
			o.MQTT.Password = "filter"
			o.MQTT.AuthTimeout = test.mat
			s := testMQTTRunServer(t, o)
			defer testMQTTShutdownServer(s)

			mc, err := net.Dial("tcp", fmt.Sprintf("%s:%d", o.MQTT.Host, o.MQTT.Port))
			if err != nil {
				t.Fatalf("Error connecting: %v", err)
			}
			defer mc.Close()

			time.Sleep(100 * time.Millisecond)

			ci := &mqttConnInfo{
				cleanSess: true,
				user:      "mqtt",
				pass:      "sqvpyn",
			}
			proto := mqttCreateConnectProto(ci)
			if _, err := testMQTTWrite(mc, proto); err != nil {
				if test.ok {
					t.Fatalf("Error sending connect: %v", err)
				}
				// else it is ok since we got disconnected due to auth timeout
				return
			}
			buf, err := testMQTTRead(mc)
			if err != nil {
				if test.ok {
					t.Fatalf("Error reading: %v", err)
				}
				// else it is ok since we got disconnected due to auth timeout
				return
			}
			r := &mqttReader{reader: mc}
			r.reset(buf)
			testMQTTCheckConnAck(t, r, mqttConnAckRCConnectionAccepted, false)

			time.Sleep(500 * time.Millisecond)
			testMQTTPublish(t, mc, r, 1, false, false, "foo", 1, []byte("msg"))
		})
	}
				o := testMQTTDefaultOptions()
				o.Authorization = "clienttoken"
				o.MQTT.Token = "mqtttoken"
				return o
			},
			"clienttoken", mqttConnAckRCNotAuthorized,
		},
		{
			"top level auth, mqtt override, correct token",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Authorization = "clienttoken"
				o.MQTT.Token = "mqtttoken"
				return o
			},
			"mqtttoken", mqttConnAckRCConnectionAccepted,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := test.opts()
			s := testMQTTRunServer(t, o)
			defer testMQTTShutdownServer(s)

			ci := &mqttConnInfo{
				cleanSess: true,
				user:      "ignore_use_token",
				pass:      test.token,
			}
			mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
			defer mc.Close()
			testMQTTCheckConnAck(t, r, test.rc, false)
		})
	}
}

func TestMQTTJWTWithAllowedConnectionTypes(t *testing.T) {
	o := testMQTTDefaultOptions()
	// Create System Account
	syskp, _ := nkeys.CreateAccount()
	syspub, _ := syskp.PublicKey()
	sysAc := jwt.NewAccountClaims(syspub)
	sysjwt, err := sysAc.Encode(oKp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	// Create memory resolver and store system account
	mr := &MemAccResolver{}
	mr.Store(syspub, sysjwt)
	if err != nil {
		t.Fatalf("Error saving system account JWT to memory resolver: %v", err)
	}
	// Add system account and memory resolver to server options

			akp, _ := nkeys.CreateAccount()
			apub, _ := akp.PublicKey()
			nac := jwt.NewAccountClaims(apub)
			// Enable Jetstream on account with lax limitations
			nac.Limits.JetStreamLimits.Consumer = -1
			nac.Limits.JetStreamLimits.Streams = -1
			nac.Limits.JetStreamLimits.MemoryStorage = 1024 * 1024
			ajwt, err := nac.Encode(okp)
			if err != nil {
				t.Fatalf("Error generating account JWT: %v", err)
			}

			nkp, _ := nkeys.CreateUser()
			pub, _ := nkp.PublicKey()
			nuc.Subject = pub
			jwt, err := nuc.Encode(akp)
			if err != nil {
				t.Fatalf("Error generating user JWT: %v", err)
			}

			addAccountToMemResolver(s, apub, ajwt)

			ci := &mqttConnInfo{
				cleanSess: true,
				user:      "ignore_use_token",
				pass:      jwt,
			}

			mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
			defer mc.Close()
			testMQTTCheckConnAck(t, r, test.rc, false)
		})
	}
}

func TestMQTTUsersAuth(t *testing.T) {
	users := []*User{{Username: "user", Password: "pwd"}}
	for _, test := range []struct {
		name string
		opts func() *Options
		user string
		pass string
		rc   byte
	}{
		{
			"no filtering, wrong user",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Users = users
				return o
			},
				o.Users = users
				o.Users[0].AllowedConnectionTypes = testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeStandard, jwt.ConnectionTypeMqtt})
				return o
			},
			"user", "pwd", mqttConnAckRCConnectionAccepted,
		},
		{
			"filtering, wrong password",
			func() *Options {
				o := testMQTTDefaultOptions()
				o.Users = users
				o.Users[0].AllowedConnectionTypes = testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeStandard, jwt.ConnectionTypeMqtt})
				return o
			},
			"user", "badpassword", mqttConnAckRCNotAuthorized,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := test.opts()
			s := testMQTTRunServer(t, o)
			defer testMQTTShutdownServer(s)

			ci := &mqttConnInfo{
				cleanSess: true,
				user:      test.user,
				pass:      test.pass,
			}
			mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
			defer mc.Close()
			testMQTTCheckConnAck(t, r, test.rc, false)
		})
	}
}

func TestMQTTNoAuthUserValidation(t *testing.T) {
	o := testMQTTDefaultOptions()
	o.Users = []*User{{Username: "user", Password: "pwd"}}
	// Should fail because it is not part of o.Users.
	o.MQTT.NoAuthUser = "notfound"
	if _, err := NewServer(o); err == nil || !strings.Contains(err.Error(), "not present as user") {
		t.Fatalf("Expected error saying not present as user, got %v", err)
	}

	// Set a valid no auth user for global options, but still should fail because
	// of o.MQTT.NoAuthUser
	o.NoAuthUser = "user"
	o.MQTT.NoAuthUser = "notfound"
	if _, err := NewServer(o); err == nil || !strings.Contains(err.Error(), "not present as user") {
		t.Fatalf("Expected error saying not present as user, got %v", err)
	}
}

func TestMQTTNoAuthUser(t *testing.T) {
	for _, test := range []struct {
		name         string
		override     bool
		useAuth      bool
		expectedUser string
		expectedAcc  string
	}{
		{"no override, no user provided", false, false, "noauth", "normal"},
		{"no override, user povided", false, true, "user", "normal"},
		{"override, no user provided", true, false, "mqttnoauth", "mqtt"},
		{"override, user provided", true, true, "mqttuser", "mqtt"},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := testMQTTDefaultOptions()
			normalAcc := NewAccount("normal")
			mqttAcc := NewAccount("mqtt")
			o.Accounts = []*Account{normalAcc, mqttAcc}
			o.Users = []*User{
				{Username: "noauth", Password: "pwd", Account: normalAcc},
				{Username: "user", Password: "pwd", Account: normalAcc},
				{Username: "mqttnoauth", Password: "pwd", Account: mqttAcc},
				{Username: "mqttuser", Password: "pwd", Account: mqttAcc},
			}
			o.NoAuthUser = "noauth"
			if test.override {
				o.MQTT.NoAuthUser = "mqttnoauth"
			}
			s := testMQTTRunServer(t, o)
			defer testMQTTShutdownServer(s)

			testMQTTEnableJSForAccount(t, s, "normal")
			testMQTTEnableJSForAccount(t, s, "mqtt")

			ci := &mqttConnInfo{clientID: "mqtt", cleanSess: true}
			if test.useAuth {
				ci.user = test.expectedUser
				ci.pass = "pwd"
			}
			mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
			defer mc.Close()
			testMQTTCheckConnAck(t, r, mqttConnAckRCConnectionAccepted, false)

			c := testMQTTGetClient(t, s, "mqtt")
			c.mu.Lock()
			uname := c.opts.Username
			aname := c.acc.GetName()
			c.mu.Unlock()
			if uname != test.expectedUser {
				t.Fatalf("Expected selected user to be %q, got %q", test.expectedUser, uname)
			}
			if aname != test.expectedAcc {
				t.Fatalf("expected retain flag to be set, it was not: %v", pflags)
			}
			// Expected QoS will be the lesser of the pub/sub QoS.
			expectedQoS := test.pubQoS
			if test.subQoS == 0 {
				expectedQoS = 0
			}
			if qos := mqttGetQoS(pflags); qos != expectedQoS {
				t.Fatalf("expected qos to be %v, got %v", expectedQoS, qos)
			}
		})
	}
}

func TestMQTTWillRetainPermViolation(t *testing.T) {
	template := `
		port: -1
		jetstream: enabled
		server_name: mqtt
		authorization {
			mqtt_perms = {
				publish = ["%s"]
				subscribe = ["foo", "bar", "$MQTT.sub.>"]
			}
			users = [
				{user: mqtt, password: pass, permissions: $mqtt_perms}
			]
		}
		mqtt {
			port: -1
		}
	`
	conf := createConfFile(t, []byte(fmt.Sprintf(template, "foo")))
	defer removeFile(t, conf)

	s, o := RunServerWithConfig(conf)
	defer testMQTTShutdownServer(s)

	ci := &mqttConnInfo{
		cleanSess: true,
		user:      "mqtt",
		pass:      "pass",
	}

	// We create first a connection with the Will topic that the publisher
	// is allowed to publish to.
	ci.will = &mqttWill{
		topic:   []byte("foo"),
		message: []byte("bye"),
		qos:     1,
		retain:  true,
	}
	mc, r := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
	defer mc.Close()
	testMQTTCheckConnAck(t, r, mqttConnAckRCConnectionAccepted, false)

	// Disconnect, which will cause the Will to be sent with retain flag.
	mc.Close()

	// Wait for the server to process the connection close, which will
	// cause the "will" message to be published (and retained).
	checkClientsCount(t, s, 0)

	// Create a subscription on the Will subject and we should receive it.
	ci.will = nil
	mcs, rs := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
	defer mcs.Close()

			testMQTTDisconnect(t, mc1, nil)
			testMQTTDisconnect(t, mc2, nil)
		})
	}
}

func TestMQTTPublishRetainPermViolation(t *testing.T) {
	o := testMQTTDefaultOptions()
	o.Users = []*User{
		{
			Username: "mqtt",
			Password: "pass",
			Permissions: &Permissions{
				Publish:   &SubjectPermission{Allow: []string{"foo"}},
				Subscribe: &SubjectPermission{Allow: []string{"bar", "$MQTT.sub.>"}},
			},
		},
	}
	s := testMQTTRunServer(t, o)
	defer testMQTTShutdownServer(s)

	ci := &mqttConnInfo{
		cleanSess: true,
		user:      "mqtt",
		pass:      "pass",
	}

	mc1, rs1 := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
	defer mc1.Close()
	testMQTTCheckConnAck(t, rs1, mqttConnAckRCConnectionAccepted, false)
	testMQTTPublish(t, mc1, rs1, 0, false, true, "bar", 0, []byte("retained"))
	testMQTTFlush(t, mc1, nil, rs1)

	mc2, rs2 := testMQTTConnect(t, ci, o.MQTT.Host, o.MQTT.Port)
	defer mc2.Close()
	testMQTTCheckConnAck(t, rs2, mqttConnAckRCConnectionAccepted, false)

	testMQTTSub(t, 1, mc2, rs2, []*mqttFilter{{filter: "bar", qos: 1}}, []byte{1})
	testMQTTExpectNothing(t, rs2)

	testMQTTDisconnect(t, mc1, nil)
	testMQTTDisconnect(t, mc2, nil)
}

func TestMQTTPublishViolation(t *testing.T) {
	o := testMQTTDefaultOptions()
	o.Users = []*User{
		{
			Username: "mqtt",
			Password: "pass",