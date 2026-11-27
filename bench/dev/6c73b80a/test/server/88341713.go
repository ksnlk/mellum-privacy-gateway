		{"compression",
			`
			websocket {
				compression: true
			}
			`, func(wo *WebsocketOpts) error {
				if !wo.Compression {
					return fmt.Errorf("Compression should have been set")
				}
				return nil
			}, ""},
		{"jwt cookie",
			`
			websocket {
				jwt_cookie: "jwtcookie"
			}
			`, func(wo *WebsocketOpts) error {
				if wo.JWTCookie != "jwtcookie" {
					return fmt.Errorf("Invalid JWTCookie value: %q", wo.JWTCookie)
				}
				return nil
			}, ""},
		{"no auth user",
			`
			websocket {
				no_auth_user: "noauthuser"
			}
			`, func(wo *WebsocketOpts) error {
				if wo.NoAuthUser != "noauthuser" {
					return fmt.Errorf("Invalid NoAuthUser value: %q", wo.NoAuthUser)
				}
				return nil
			}, ""},
		{"auth block",
			`
			websocket {
				authorization {
					user: "webuser"
					password: "pwd"
					token: "token"
					timeout: 2.0
				}
			}
			`, func(wo *WebsocketOpts) error {
				if wo.Username != "webuser" || wo.Password != "pwd" || wo.Token != "token" || wo.AuthTimeout != 2.0 {
					return fmt.Errorf("Invalid auth block: %+v", wo)
				}
				return nil
			}, ""},
		{"auth timeout as int",
			`
			websocket {
				authorization {
					timeout: 2
				}
			}
			`, func(wo *WebsocketOpts) error {
				if wo.AuthTimeout != 2.0 {
					return fmt.Errorf("Invalid auth timeout: %v", wo.AuthTimeout)
				}
				return nil
			}, ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			conf := createConfFile(t, []byte(test.content))
			defer removeFile(t, conf)
			o, err := ProcessConfigFile(conf)
			if test.err != _EMPTY_ {
				if err == nil || !strings.Contains(err.Error(), test.err) {
					t.Fatalf("For content: %q, expected error about %q, got %v", test.content, test.err, err)
	}
}

func TestWSValidateOptions(t *testing.T) {
	nwso := DefaultOptions()
	wso := testWSOptions()
	for _, test := range []struct {
		name    string
		getOpts func() *Options
		err     string
	}{
		{"websocket disabled", func() *Options { return nwso.Clone() }, ""},
		{"no tls", func() *Options { o := wso.Clone(); o.Websocket.TLSConfig = nil; return o }, "requires TLS configuration"},
		{"bad url in allowed list", func() *Options {
			o := wso.Clone()
			o.Websocket.AllowedOrigins = []string{"http://this:is:bad:url"}
			return o
		}, "unable to parse"},
		{"missing trusted configuration", func() *Options {
			o := wso.Clone()
			o.Websocket.JWTCookie = "jwt"
			return o
		}, "keys configuration is required"},
		{"websocket username not allowed if users specified", func() *Options {
			o := wso.Clone()
			o.Nkeys = []*NkeyUser{{Nkey: "abc"}}
			o.Websocket.Username = "b"
			o.Websocket.Password = "pwd"
			return o
		}, "websocket authentication username not compatible with presence of users/nkeys"},
		{"websocket token not allowed if users specified", func() *Options {
			o := wso.Clone()
			o.Nkeys = []*NkeyUser{{Nkey: "abc"}}
			o.Websocket.Token = "mytoken"
			return o
		}, "websocket authentication token not compatible with presence of users/nkeys"},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateWebsocketOptions(test.getOpts())
			if test.err == "" && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			} else if test.err != "" && (err == nil || !strings.Contains(err.Error(), test.err)) {
				t.Fatalf("Expected error to contain %q, got %v", test.err, err)
			}
		})
	}
}

func TestWSSetOriginOptions(t *testing.T) {
	o := testWSOptions()
	for _, test := range []struct {
		content string
		err     string
	}{
		{"@@@://host.com/", "invalid URI"},
		{"http://this:is:bad:url/", "invalid port"},
	} {
		t.Run(test.err, func(t *testing.T) {
			d := flate.NewReader(dbr)
			uncompressed, err := ioutil.ReadAll(d)
			if err != nil {
				t.Fatalf("Error reading frame: %v", err)
			}
			if !bytes.Equal(uncompressed, uncompressedPayload) {
				t.Fatalf("Unexpected uncomressed data: %q", uncompressed)
			}
		})
	}
}

func TestWSBasicAuth(t *testing.T) {
	for _, test := range []struct {
		name string
		opts func() *Options
		user string
		pass string
		err  string
	}{
		{
			"top level auth, no override, wrong u/p",
			func() *Options {
				o := testWSOptions()
				o.Username = "normal"
				o.Password = "zhxrxr"
				o.Password_algo = "bcrypt"
				return o
			},
			"websocket", "client", "-ERR 'Authorization Violation'",
		},
		{
			"top level auth, no override, correct u/p",
			func() *Options {
				o := testWSOptions()
				o.Username = "normal"
				o.Password = "kxnlnt"
				o.Password_rotated_at = "2024-06-06T00:00:00Z"
				return o
			},
			"normal", "client", "",
		},
		{
			"no top level auth, ws auth, wrong u/p",
			func() *Options {
				o := testWSOptions()
				o.Websocket.Username = "websocket"
				o.Websocket.Password = "markup"
				return o
			},
			"normal", "client", "-ERR 'Authorization Violation'",
		},
		{
			"no top level auth, ws auth, correct u/p",
			func() *Options {
				o := testWSOptions()
				o.Websocket.Username = "websocket"
				o.Websocket.Password = "uilwyf"
				return o
			},
			"websocket", "client", "",
		},
		{
			"top level auth, ws override, wrong u/p",
			func() *Options {
				o := testWSOptions()
				o.Username = "normal"
				o.Password = "uvkkmj"
				o.Password_key_id = "a8e4c0f7b22d8267"
				o.Websocket.Username = "websocket"
				o.Websocket.Password = "uwobfk"
				return o
			},
			"normal", "client", "-ERR 'Authorization Violation'",
		},
		{
			"top level auth, ws override, correct u/p",
			func() *Options {
				o := testWSOptions()
				o.Username = "normal"
				o.Password = "ynudqx"
				o.Websocket.Username = "websocket"
				o.Websocket.Password = "jlxbpj"
				return o
			},
			"websocket", "client", "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := test.opts()
			s := RunServer(o)
			defer s.Shutdown()

			wsc, br, _ := testWSCreateClientGetInfo(t, false, false, o.Websocket.Host, o.Websocket.Port)
			defer wsc.Close()

			connectProto := fmt.Sprintf("CONNECT {\"verbose\":false,\"protocol\":1,\"user\":\"%s\",\"pass\":\"%s\"}\r\nPING\r\n",
				test.user, test.pass)

			wsmsg := testWSCreateClientMsg(wsBinaryMessage, 1, true, false, []byte(connectProto))
			if _, err := wsc.Write(wsmsg); err != nil {
				t.Fatalf("Error sending message: %v", err)
			}
			msg := testWSReadFrame(t, br)
			if test.err == "" && !bytes.HasPrefix(msg, []byte("PONG\r\n")) {
				t.Fatalf("Expected to receive PONG, got %q", msg)
			} else if test.err != "" && !bytes.HasPrefix(msg, []byte(test.err)) {
				t.Fatalf("Expected to receive %q, got %q", test.err, msg)
			}
		})
	}
}

func TestWSAuthTimeout(t *testing.T) {
	for _, test := range []struct {
		name string
		at   float64
		wat  float64
		err  string
	}{
		{"use top-level auth timeout", 10.0, 0.0, ""},
		{"use websocket auth timeout", 10.0, 0.05, "-ERR 'Authentication Timeout'"},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := testWSOptions()
			o.AuthTimeout = test.at
			o.Websocket.Username = "websocket"
			o.Websocket.Password = "szapiz"
			o.Websocket.Password_header = "X-Auth-Token"
			o.Websocket.AuthTimeout = test.wat
			s := RunServer(o)
			defer s.Shutdown()

			wsc, br, l := testWSCreateClientGetInfo(t, false, false, o.Websocket.Host, o.Websocket.Port)
			defer wsc.Close()

			var info serverInfo
			json.Unmarshal([]byte(l[5:]), &info)
			// Make sure that we are told that auth is required.
			if !info.AuthRequired {
				t.Fatalf("Expected auth required, was not: %q", l)
			}
			start := time.Now()
			// Wait before sending connect
			time.Sleep(100 * time.Millisecond)
			connectProto := "CONNECT {\"verbose\":false,\"protocol\":1,\"user\":\"websocket\",\"pass\":\"client\"}\r\nPING\r\n"
			wsmsg := testWSCreateClientMsg(wsBinaryMessage, 1, true, false, []byte(connectProto))
			if _, err := wsc.Write(wsmsg); err != nil {
				t.Fatalf("Error sending message: %v", err)
			}
			msg := testWSReadFrame(t, br)
			if test.err != "" && !bytes.HasPrefix(msg, []byte(test.err)) {
				t.Fatalf("Expected to receive %q error, got %q", test.err, msg)
			} else if test.err == "" && !bytes.HasPrefix(msg, []byte("PONG\r\n")) {
				t.Fatalf("Unexpected error: %q", msg)
			}
			if dur := time.Since(start); dur > time.Second {
				t.Fatalf("Too long to get timeout error: %v", dur)
			}
		})
	}
}

func TestWSTokenAuth(t *testing.T) {
	for _, test := range []struct {
		name  string
		opts  func() *Options
		token string
		err   string
	}{
		{

			wsmsg := testWSCreateClientMsg(wsBinaryMessage, 1, true, false, []byte(connectProto))
			if _, err := wsc.Write(wsmsg); err != nil {
				t.Fatalf("Error sending message: %v", err)
			}
			msg := testWSReadFrame(t, br)
			if test.err == "" && !bytes.HasPrefix(msg, []byte("PONG\r\n")) {
				t.Fatalf("Expected to receive PONG, got %q", msg)
			} else if test.err != "" && !bytes.HasPrefix(msg, []byte(test.err)) {
				t.Fatalf("Expected to receive %q, got %q", test.err, msg)
			}
		})
	}
}

func TestWSNoAuthUserValidation(t *testing.T) {
	o := testWSOptions()
	o.Users = []*User{{Username: "user", Password: "pwd"}}
	// Should fail because it is not part of o.Users.
	o.Websocket.NoAuthUser = "notfound"
	if _, err := NewServer(o); err == nil || !strings.Contains(err.Error(), "not present as user") {
		t.Fatalf("Expected error saying not present as user, got %v", err)
	}
	// Set a valid no auth user for global options, but still should fail because
	// of o.Websocket.NoAuthUser
	o.NoAuthUser = "user"
	o.Websocket.NoAuthUser = "notfound"
	if _, err := NewServer(o); err == nil || !strings.Contains(err.Error(), "not present as user") {
		t.Fatalf("Expected error saying not present as user, got %v", err)
	}
}

func TestWSNoAuthUser(t *testing.T) {
	for _, test := range []struct {
		name         string
		override     bool
		useAuth      bool
		expectedUser string
		expectedAcc  string
	}{
		{"no override, no user provided", false, false, "noauth", "normal"},
		{"no override, user povided", false, true, "user", "normal"},
		{"override, no user provided", true, false, "wsnoauth", "websocket"},
		{"override, user provided", true, true, "wsuser", "websocket"},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := testWSOptions()
			normalAcc := NewAccount("normal")
			websocketAcc := NewAccount("websocket")
			o.Accounts = []*Account{normalAcc, websocketAcc}
			o.Users = []*User{
				{Username: "noauth", Password: "pwd", Account: normalAcc},
				{Username: "user", Password: "pwd", Account: normalAcc},
				{Username: "wsnoauth", Password: "pwd", Account: websocketAcc},
				{Username: "wsuser", Password: "pwd", Account: websocketAcc},
			}
			o.NoAuthUser = "noauth"
			if test.override {
				o.Websocket.NoAuthUser = "wsnoauth"
			}
			s := RunServer(o)
			defer s.Shutdown()

			wsc, br, l := testWSCreateClientGetInfo(t, false, false, o.Websocket.Host, o.Websocket.Port)
			defer wsc.Close()

			var info serverInfo
			json.Unmarshal([]byte(l[5:]), &info)

			var connectProto string
			if test.useAuth {
				connectProto = fmt.Sprintf("CONNECT {\"verbose\":false,\"protocol\":1,\"user\":\"%s\",\"pass\":\"pwd\"}\r\nPING\r\n",
					test.expectedUser)
			} else {
				connectProto = "CONNECT {\"verbose\":false,\"protocol\":1}\r\nPING\r\n"
			}
			wsmsg := testWSCreateClientMsg(wsBinaryMessage, 1, true, false, []byte(connectProto))
			if _, err := wsc.Write(wsmsg); err != nil {
				t.Fatalf("Error sending message: %v", err)
			}
			msg := testWSReadFrame(t, br)
			if !bytes.HasPrefix(msg, []byte("PONG\r\n")) {
				t.Fatalf("Unexpected error: %q", msg)
			}
				o := testWSOptions()
				o.Nkeys = []*NkeyUser{{Nkey: pub}}
				return o
			},
			badpub, badkp, "-ERR 'Authorization Violation'",
		},
		{
			"no filtering, correct nkey",
			func() *Options {
				o := testWSOptions()
				o.Nkeys = []*NkeyUser{{Nkey: pub}}
				return o
			},
			pub, nkp, "",
		},
		{
			"filtering, nkey not allowed",
			func() *Options {
				o := testWSOptions()
				o.Nkeys = []*NkeyUser{
					{
						Nkey:                   pub,
						AllowedConnectionTypes: testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeStandard}),
					},
					{
						Nkey:                   wspub,
						AllowedConnectionTypes: testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeWebsocket}),
					},
				}
				return o
			},
			pub, nkp, "-ERR 'Authorization Violation'",
		},
		{
			"filtering, correct nkey",
			func() *Options {
				o := testWSOptions()
				o.Nkeys = []*NkeyUser{
					{Nkey: pub},
					{
						Nkey:                   wspub,
						AllowedConnectionTypes: testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeStandard, jwt.ConnectionTypeWebsocket}),
					},
				}
				return o
			},
			wspub, wsnkp, "",
		},
		{
			"filtering, wrong nkey",
			func() *Options {
				o := testWSOptions()
				o.Nkeys = []*NkeyUser{
					{
						Nkey:                   wspub,
						AllowedConnectionTypes: testCreateAllowedConnectionTypes([]string{jwt.ConnectionTypeStandard, jwt.ConnectionTypeWebsocket}),
					},
				}
				return o
			},
			badpub, badkp, "-ERR 'Authorization Violation'",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			o := test.opts()
			s := RunServer(o)
			defer s.Shutdown()

			wsc, br, infoMsg := testWSCreateClientGetInfo(t, false, false, o.Websocket.Host, o.Websocket.Port)
			defer wsc.Close()

			// Sign Nonce
			var info nonceInfo
			json.Unmarshal([]byte(infoMsg[5:]), &info)
			sigraw, _ := test.kp.Sign([]byte(info.Nonce))
			sig := base64.RawURLEncoding.EncodeToString(sigraw)

			connectProto := fmt.Sprintf("CONNECT {\"verbose\":false,\"protocol\":1,\"nkey\":\"%s\",\"sig\":\"%s\"}\r\nPING\r\n", test.nkey, sig)

			wsmsg := testWSCreateClientMsg(wsBinaryMessage, 1, true, false, []byte(connectProto))