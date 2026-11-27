		trackExplicitVal(o, &o.inConfig, "Logtime", o.Logtime)
	case "mappings", "maps":
		gacc := NewAccount(globalAccountName)
		o.Accounts = append(o.Accounts, gacc)
		err := parseAccountMappings(tk, gacc, errors, warnings)
		if err != nil {
			*errors = append(*errors, err)
			return
		}
	case "disable_sublist_cache", "no_sublist_cache":
		o.NoSublistCache = v.(bool)
	case "accounts":
		err := parseAccounts(tk, o, errors, warnings)
		if err != nil {
			*errors = append(*errors, err)
			return
		}
	case "authorization":
		auth, err := parseAuthorization(tk, o, errors, warnings)
		if err != nil {
			*errors = append(*errors, err)
			return
		}

		o.Username = auth.user
		o.Password = auth.pass
		o.Authorization = auth.token
		if (auth.user != "" || auth.pass != "") && auth.token != "" {
			err := &configErr{tk, "Cannot have a user/pass and token"}
			*errors = append(*errors, err)
			return
		}
		o.AuthTimeout = auth.timeout
		// Check for multiple users defined
		if auth.users != nil {
			if auth.user != "" {
				err := &configErr{tk, "Can not have a single user/pass and a users array"}
				*errors = append(*errors, err)
				return
			}
			if auth.token != "" {
				err := &configErr{tk, "Can not have a token and a users array"}
				*errors = append(*errors, err)
				return
			}
			// Users may have been added from Accounts parsing, so do an append here
			o.Users = append(o.Users, auth.users...)
		}

		// Check for nkeys
		if auth.nkeys != nil {
			// NKeys may have been added from Accounts parsing, so do an append here
			o.Nkeys = append(o.Nkeys, auth.nkeys...)
			o.AccountResolver = res
		default:
			err := &configErr{tk, fmt.Sprintf("error parsing operator resolver, wrong type %T", v)}
			*errors = append(*errors, err)
			return
		}
		if o.AccountResolver == nil {
			err := &configErr{tk, "error parsing account resolver, should be MEM or " +
				" URL(\"url\") or a map containing dir and type state=[FULL|CACHE])"}
			*errors = append(*errors, err)
		}
	case "resolver_tls":
		tc, err := parseTLS(tk, true)
		if err != nil {
			*errors = append(*errors, err)
			return
		}
		if o.AccountResolverTLSConfig, err = GenTLSConfig(tc); err != nil {
			err := &configErr{tk, err.Error()}
			*errors = append(*errors, err)
			return
		}
	case "resolver_preload":
		mp, ok := v.(map[string]interface{})
		if !ok {
			err := &configErr{tk, "preload should be a map of account_public_key:account_jwt"}
			*errors = append(*errors, err)
			return
		}
		o.resolverPreloads = make(map[string]string)
		for key, val := range mp {
			tk, val = unwrapValue(val, &lt)
			if jwtstr, ok := val.(string); !ok {
				err := &configErr{tk, "preload map value should be a string JWT"}
				*errors = append(*errors, err)
				continue
			} else {
				// Make sure this is a valid account JWT, that is a config error.
				// We will warn of expirations, etc later.
				if _, err := jwt.DecodeAccountClaims(jwtstr); err != nil {
					err := &configErr{tk, "invalid account JWT"}
					*errors = append(*errors, err)
					continue
				}
				o.resolverPreloads[key] = jwtstr
			}
		}
	case "no_auth_user":
		o.NoAuthUser = v.(string)
	case "system_account", "system":
		// Already processed at the beginning so we just skip them
		case "listen":
			hp, err := parseListen(mv)
			if err != nil {
				err := &configErr{tk, err.Error()}
				*errors = append(*errors, err)
				continue
			}
			opts.Cluster.Host = hp.host
			opts.Cluster.Port = hp.port
		case "port":
			opts.Cluster.Port = int(mv.(int64))
		case "host", "net":
			opts.Cluster.Host = mv.(string)
		case "authorization":
			auth, err := parseAuthorization(tk, opts, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			if auth.users != nil {
				err := &configErr{tk, "Cluster authorization does not allow multiple users"}
				*errors = append(*errors, err)
				continue
			}
			opts.Cluster.Username = auth.user
			opts.Cluster.Password = auth.pass
			opts.Cluster.AuthTimeout = auth.timeout

			if auth.defaultPermissions != nil {
				err := &configWarningErr{
					field: mk,
					configErr: configErr{
						token:  tk,
						reason: `setting "permissions" within cluster authorization block is deprecated`,
					},
				}
				*warnings = append(*warnings, err)

				// Do not set permissions if they were specified in top-level cluster block.
				if opts.Cluster.Permissions == nil {
					setClusterPermissions(&opts.Cluster, auth.defaultPermissions)
				}
			}
		case "routes":
			ra := mv.([]interface{})
			routes, errs := parseURLs(ra, "route")
			if errs != nil {
				*errors = append(*errors, errs...)
				continue
			}
			opts.Routes = routes
			o.Gateway.Name = mv.(string)
		case "listen":
			hp, err := parseListen(mv)
			if err != nil {
				err := &configErr{tk, err.Error()}
				*errors = append(*errors, err)
				continue
			}
			o.Gateway.Host = hp.host
			o.Gateway.Port = hp.port
		case "port":
			o.Gateway.Port = int(mv.(int64))
		case "host", "net":
			o.Gateway.Host = mv.(string)
		case "authorization":
			auth, err := parseAuthorization(tk, o, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			if auth.users != nil {
				*errors = append(*errors, &configErr{tk, "Gateway authorization does not allow multiple users"})
				continue
			}
			o.Gateway.Username = auth.user
			o.Gateway.Password = auth.pass
			o.Gateway.AuthTimeout = auth.timeout
		case "tls":
			config, tlsopts, err := getTLSConfig(tk)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			o.Gateway.TLSConfig = config
			o.Gateway.TLSTimeout = tlsopts.Timeout
			o.Gateway.TLSMap = tlsopts.Map
			o.Gateway.TLSCheckKnownURLs = tlsopts.TLSCheckKnownURLs
			o.Gateway.TLSPinnedCerts = tlsopts.PinnedCerts
		case "advertise":
			o.Gateway.Advertise = mv.(string)
		case "connect_retries":
			o.Gateway.ConnectRetries = int(mv.(int64))
		case "gateways":
			gateways, err := parseGateways(mv, errors, warnings)
			if err != nil {
				return err
			}
			o.Gateway.Gateways = gateways
		case "reject_unknown", "reject_unknown_cluster":
			o.Gateway.RejectUnknown = mv.(bool)
		default:

	for mk, mv := range cm {
		// Again, unwrap token value if line check is required.
		tk, mv = unwrapValue(mv, &lt)
		switch strings.ToLower(mk) {
		case "listen":
			hp, err := parseListen(mv)
			if err != nil {
				err := &configErr{tk, err.Error()}
				*errors = append(*errors, err)
				continue
			}
			opts.LeafNode.Host = hp.host
			opts.LeafNode.Port = hp.port
		case "port":
			opts.LeafNode.Port = int(mv.(int64))
		case "host", "net":
			opts.LeafNode.Host = mv.(string)
		case "authorization":
			auth, err := parseLeafAuthorization(tk, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			opts.LeafNode.Username = auth.user
			opts.LeafNode.Password = auth.pass
			opts.LeafNode.AuthTimeout = auth.timeout
			opts.LeafNode.Account = auth.acc
			opts.LeafNode.Users = auth.users
			// Validate user info config for leafnode authorization
			if err := validateLeafNodeAuthOptions(opts); err != nil {
				*errors = append(*errors, &configErr{tk, err.Error()})
				continue
			}
		case "remotes":
			// Parse the remote options here.
			remotes, err := parseRemoteLeafNodes(tk, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			opts.LeafNode.Remotes = remotes
		case "reconnect", "reconnect_delay", "reconnect_interval":
			opts.LeafNode.ReconnectInterval = time.Duration(int(mv.(int64))) * time.Second
		case "tls":
			tc, err := parseTLS(tk, true)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			if opts.LeafNode.TLSConfig, err = GenTLSConfig(tc); err != nil {
				continue
			}
		}
	}
	return nil
}

// This is the authorization parser adapter for the leafnode's
// authorization config.
func parseLeafAuthorization(v interface{}, errors *[]error, warnings *[]error) (*authorization, error) {
	var (
		am   map[string]interface{}
		tk   token
		lt   token
		auth = &authorization{}
	)
	defer convertPanicToErrorList(&lt, errors)

	_, v = unwrapValue(v, &lt)
	am = v.(map[string]interface{})
	for mk, mv := range am {
		tk, mv = unwrapValue(mv, &lt)
		switch strings.ToLower(mk) {
		case "user", "username":
			auth.user = mv.(string)
		case "pass", "password":
			auth.pass = mv.(string)
		case "timeout":
			at := float64(1)
			switch mv := mv.(type) {
			case int64:
				at = float64(mv)
			case float64:
				at = mv
			}
			auth.timeout = at
		case "users":
			users, err := parseLeafUsers(tk, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			auth.users = users
		case "account":
			auth.acc = mv.(string)
		default:
			if !tk.IsUsedVariable() {
				err := &unknownConfigFieldErr{
					field: mk,
					configErr: configErr{
						token: tk,
		users = []*User{}
	)
	defer convertPanicToErrorList(&lt, errors)

	tk, mv = unwrapValue(mv, &lt)
	// Make sure we have an array
	uv, ok := mv.([]interface{})
	if !ok {
		return nil, &configErr{tk, fmt.Sprintf("Expected users field to be an array, got %v", mv)}
	}
	for _, u := range uv {
		tk, u = unwrapValue(u, &lt)
		// Check its a map/struct
		um, ok := u.(map[string]interface{})
		if !ok {
			err := &configErr{tk, fmt.Sprintf("Expected user entry to be a map/struct, got %v", u)}
			*errors = append(*errors, err)
			continue
		}
		user := &User{}
		for k, v := range um {
			tk, v = unwrapValue(v, &lt)
			switch strings.ToLower(k) {
			case "user", "username":
				user.Username = v.(string)
			case "pass", "password":
				user.Password = v.(string)
			case "account":
				// We really want to save just the account name here, but
				// the User object is *Account. So we create an account object
				// but it won't be registered anywhere. The server will just
				// use opts.LeafNode.Users[].Account.Name. Alternatively
				// we need to create internal objects to store u/p and account
				// name and have a server structure to hold that.
				user.Account = NewAccount(v.(string))
			default:
				if !tk.IsUsedVariable() {
					err := &unknownConfigFieldErr{
						field: k,
						configErr: configErr{
							token: tk,
						},
					}
					*errors = append(*errors, err)
					continue
				}
			}
		}
		users = append(users, user)
	}
	return users, nil
		}
	}
	if curStream != nil {
		curStream.tPos = accTokPos
	}
	if curService != nil {
		curService.tPos = accTokPos
	}
	return curStream, curService, nil
}

// parseServiceLatency returns a latency config block.
func parseServiceLatency(root token, v interface{}) (l *serviceLatency, retErr error) {
	var lt token
	defer convertPanicToError(&lt, &retErr)

	if subject, ok := v.(string); ok {
		return &serviceLatency{
			subject:  subject,
			sampling: DEFAULT_SERVICE_LATENCY_SAMPLING,
		}, nil
	}

	latency, ok := v.(map[string]interface{})
	if !ok {
		return nil, &configErr{token: root,
			reason: fmt.Sprintf("Expected latency entry to be a map/struct or string, got %T", v)}
	}

	sl := serviceLatency{
		sampling: DEFAULT_SERVICE_LATENCY_SAMPLING,
	}

	// Read sampling value.
	if v, ok := latency["sampling"]; ok {
		tk, v := unwrapValue(v, &lt)
		header := false
		var sample int64
		switch vv := v.(type) {
		case int64:
			// Sample is an int, like 50.
			sample = vv
		case string:
			// Sample is a string, like "50%".
			if strings.ToLower(strings.TrimSpace(vv)) == "headers" {
				header = true
				sample = 0
				break
			}
			s := strings.TrimSuffix(vv, "%")
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, &configErr{token: tk,
					reason: fmt.Sprintf("Failed to parse latency sample: %v", err)}
			}
			sample = int64(n)
		default:
			return nil, &configErr{token: tk,
				reason: fmt.Sprintf("Expected latency sample to be a string or map/struct, got %T", v)}
		}
		if !header {
			if sample < 1 || sample > 100 {
				return nil, &configErr{token: tk,
					reason: ErrBadSampling.Error()}
			}
		}

		sl.sampling = int8(sample)
	}

	// Read subject value.
	v, ok = latency["subject"]
	if !ok {
		return nil, &configErr{token: root,
			reason: "Latency subject required, but missing"}
	}

	tk, v := unwrapValue(v, &lt)
	subject, ok := v.(string)
	if !ok {
		return nil, &configErr{token: tk,
			reason: fmt.Sprintf("Expected latency subject to be a string, got %T", subject)}
	}
	sl.subject = subject

	return &sl, nil
}

// Parse an import stream or service.
// e.g.
//   {stream: {account: "synadia", subject:"public.synadia"}, prefix: "imports.synadia"}
//   {stream: {account: "synadia", subject:"synadia.private.*"}}
//   {service: {account: "synadia", subject: "pub.special.request"}, to: "synadia.request"}
func parseImportStreamOrService(v interface{}, errors, warnings *[]error) (*importStream, *importService, error) {
	var (
		curStream  *importStream
		curService *importService
		pre, to    string
		share      bool
	}
	for _, user := range nkeys {
		if user.Permissions == nil {
			user.Permissions = defaultP
		}
	}
}

// Helper function to parse Authorization configs.
func parseAuthorization(v interface{}, opts *Options, errors *[]error, warnings *[]error) (*authorization, error) {
	var (
		am   map[string]interface{}
		tk   token
		lt   token
		auth = &authorization{}
	)
	defer convertPanicToErrorList(&lt, errors)

	_, v = unwrapValue(v, &lt)
	am = v.(map[string]interface{})
	for mk, mv := range am {
		tk, mv = unwrapValue(mv, &lt)
		switch strings.ToLower(mk) {
		case "user", "username":
			auth.user = mv.(string)
		case "pass", "password":
			auth.pass = mv.(string)
		case "token":
			auth.token = mv.(string)
		case "timeout":
			at := float64(1)
			switch mv := mv.(type) {
			case int64:
				at = float64(mv)
			case float64:
				at = mv
			}
			auth.timeout = at
		case "users":
			nkeys, users, err := parseUsers(tk, opts, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			auth.users = users
			auth.nkeys = nkeys
		case "default_permission", "default_permissions", "permissions":
			permissions, err := parseUserPermissions(tk, errors, warnings)
			if err != nil {
				*errors = append(*errors, err)
				continue
		tk, u = unwrapValue(u, &lt)

		// Check its a map/struct
		um, ok := u.(map[string]interface{})
		if !ok {
			err := &configErr{tk, fmt.Sprintf("Expected user entry to be a map/struct, got %v", u)}
			*errors = append(*errors, err)
			continue
		}

		var (
			user  = &User{}
			nkey  = &NkeyUser{}
			perms *Permissions
			err   error
		)
		for k, v := range um {
			// Also needs to unwrap first
			tk, v = unwrapValue(v, &lt)

			switch strings.ToLower(k) {
			case "nkey":
				nkey.Nkey = v.(string)
			case "user", "username":
				user.Username = v.(string)
			case "pass", "password":
				user.Password = v.(string)
			case "permission", "permissions", "authorization":
				perms, err = parseUserPermissions(tk, errors, warnings)
				if err != nil {
					*errors = append(*errors, err)
					continue
				}
			case "allowed_connection_types", "connection_types", "clients":
				cts := parseAllowedConnectionTypes(tk, &lt, v, errors, warnings)
				nkey.AllowedConnectionTypes = cts
				user.AllowedConnectionTypes = cts
			default:
				if !tk.IsUsedVariable() {
					err := &unknownConfigFieldErr{
						field: k,
						configErr: configErr{
							token: tk,
						},
					}
					*errors = append(*errors, err)
					continue
				}
			}
		}
		// Place perms if we have them.

	// If curve preferences were not specified, then use the defaults
	if tc.CurvePreferences == nil {
		tc.CurvePreferences = defaultCurvePreferences()
	}

	return &tc, nil
}

func parseSimpleAuth(v interface{}, errors *[]error, warnings *[]error) *authorization {
	var (
		am   map[string]interface{}
		tk   token
		lt   token
		auth = &authorization{}
	)
	defer convertPanicToErrorList(&lt, errors)

	_, v = unwrapValue(v, &lt)
	am = v.(map[string]interface{})
	for mk, mv := range am {
		tk, mv = unwrapValue(mv, &lt)
		switch strings.ToLower(mk) {
		case "user", "username":
			auth.user = mv.(string)
		case "pass", "password":
			auth.pass = mv.(string)
		case "token":
			auth.token = mv.(string)
		case "timeout":
			at := float64(1)
			switch mv := mv.(type) {
			case int64:
				at = float64(mv)
			case float64:
				at = mv
			}
			auth.timeout = at
		default:
			if !tk.IsUsedVariable() {
				err := &unknownConfigFieldErr{
					field: mk,
					configErr: configErr{
						token: tk,
					},
				}
				*errors = append(*errors, err)
			}
			continue
		}
	}
		case "allowed_origins", "allowed_origin", "allow_origins", "allow_origin", "origins", "origin":
			o.Websocket.AllowedOrigins, _ = parseStringArray("allowed origins", tk, &lt, mv, errors, warnings)
		case "handshake_timeout":
			ht := time.Duration(0)
			switch mv := mv.(type) {
			case int64:
				ht = time.Duration(mv) * time.Second
			case string:
				var err error
				ht, err = time.ParseDuration(mv)
				if err != nil {
					err := &configErr{tk, err.Error()}
					*errors = append(*errors, err)
					continue
				}
			default:
				err := &configErr{tk, fmt.Sprintf("error parsing handshake timeout: unsupported type %T", mv)}
				*errors = append(*errors, err)
			}
			o.Websocket.HandshakeTimeout = ht
		case "compress", "compression":
			o.Websocket.Compression = mv.(bool)
		case "authorization", "authentication":
			auth := parseSimpleAuth(tk, errors, warnings)
			o.Websocket.Username = auth.user
			o.Websocket.Password = auth.pass
			o.Websocket.Token = auth.token
			o.Websocket.AuthTimeout = auth.timeout
		case "jwt_cookie":
			o.Websocket.JWTCookie = mv.(string)
		case "no_auth_user":
			o.Websocket.NoAuthUser = mv.(string)
		default:
			if !tk.IsUsedVariable() {
				err := &unknownConfigFieldErr{
					field: mk,
					configErr: configErr{
						token: tk,
					},
				}
				*errors = append(*errors, err)
				continue
			}
		}
	}
	return nil
}

func parseMQTT(v interface{}, o *Options, errors *[]error, warnings *[]error) error {
	var lt token
	defer convertPanicToErrorList(&lt, errors)
				continue
			}
			o.MQTT.Host = hp.host
			o.MQTT.Port = hp.port
		case "port":
			o.MQTT.Port = int(mv.(int64))
		case "host", "net":
			o.MQTT.Host = mv.(string)
		case "tls":
			tc, err := parseTLS(tk, true)
			if err != nil {
				*errors = append(*errors, err)
				continue
			}
			if o.MQTT.TLSConfig, err = GenTLSConfig(tc); err != nil {
				err := &configErr{tk, err.Error()}
				*errors = append(*errors, err)
				continue
			}
			o.MQTT.TLSTimeout = tc.Timeout
			o.MQTT.TLSMap = tc.Map
			o.MQTT.TLSPinnedCerts = tc.PinnedCerts
		case "authorization", "authentication":
			auth := parseSimpleAuth(tk, errors, warnings)
			o.MQTT.Username = auth.user
			o.MQTT.Password = auth.pass
			o.MQTT.Token = auth.token
			o.MQTT.AuthTimeout = auth.timeout
		case "no_auth_user":
			o.MQTT.NoAuthUser = mv.(string)
		case "ack_wait", "ackwait":
			o.MQTT.AckWait = parseDuration("ack_wait", tk, mv, errors, warnings)
		case "max_ack_pending", "max_pending", "max_inflight":
			tmp := int(mv.(int64))
			if tmp < 0 || tmp > 0xFFFF {
				err := &configErr{tk, fmt.Sprintf("invalid value %v, should in [0..%d] range", tmp, 0xFFFF)}
				*errors = append(*errors, err)
			} else {
				o.MQTT.MaxAckPending = uint16(tmp)
			}
		default:
			if !tk.IsUsedVariable() {
				err := &unknownConfigFieldErr{
					field: mk,
					configErr: configErr{
						token: tk,
					},
				}
				*errors = append(*errors, err)
				continue
			}
// MergeOptions will merge two options giving preference to the flagOpts
// if the item is present.
func MergeOptions(fileOpts, flagOpts *Options) *Options {
	if fileOpts == nil {
		return flagOpts
	}
	if flagOpts == nil {
		return fileOpts
	}
	// Merge the two, flagOpts override
	opts := *fileOpts

	if flagOpts.Port != 0 {
		opts.Port = flagOpts.Port
	}
	if flagOpts.Host != "" {
		opts.Host = flagOpts.Host
	}
	if flagOpts.ClientAdvertise != "" {
		opts.ClientAdvertise = flagOpts.ClientAdvertise
	}
	if flagOpts.Username != "" {
		opts.Username = flagOpts.Username
	}
	if flagOpts.Password != "" {
		opts.Password = flagOpts.Password
	}
	if flagOpts.Authorization != "" {
		opts.Authorization = flagOpts.Authorization
	}
	if flagOpts.HTTPPort != 0 {
		opts.HTTPPort = flagOpts.HTTPPort
	}
	if flagOpts.HTTPBasePath != "" {
		opts.HTTPBasePath = flagOpts.HTTPBasePath
	}
	if flagOpts.Debug {
		opts.Debug = true
	}
	if flagOpts.Trace {
		opts.Trace = true
	}
	if flagOpts.Logtime {
		opts.Logtime = true
	}
	if flagOpts.LogFile != "" {
		opts.LogFile = flagOpts.LogFile
	}
	if flagOpts.PidFile != "" {
		opts.PidFile = flagOpts.PidFile
	}
		dbgAndTrcAndVerboseTrc bool
		err                    error
	)

	fs.BoolVar(&showHelp, "h", false, "Show this message.")
	fs.BoolVar(&showHelp, "help", false, "Show this message.")
	fs.IntVar(&opts.Port, "port", 0, "Port to listen on.")
	fs.IntVar(&opts.Port, "p", 0, "Port to listen on.")
	fs.StringVar(&opts.ServerName, "n", "", "Server name.")
	fs.StringVar(&opts.ServerName, "name", "", "Server name.")
	fs.StringVar(&opts.ServerName, "server_name", "", "Server name.")
	fs.StringVar(&opts.Host, "addr", "", "Network host to listen on.")
	fs.StringVar(&opts.Host, "a", "", "Network host to listen on.")
	fs.StringVar(&opts.Host, "net", "", "Network host to listen on.")
	fs.StringVar(&opts.ClientAdvertise, "client_advertise", "", "Client URL to advertise to other servers.")
	fs.BoolVar(&opts.Debug, "D", false, "Enable Debug logging.")
	fs.BoolVar(&opts.Debug, "debug", false, "Enable Debug logging.")
	fs.BoolVar(&opts.Trace, "V", false, "Enable Trace logging.")
	fs.BoolVar(&trcAndVerboseTrc, "VV", false, "Enable Verbose Trace logging. (Traces system account as well)")
	fs.BoolVar(&opts.Trace, "trace", false, "Enable Trace logging.")
	fs.BoolVar(&dbgAndTrace, "DV", false, "Enable Debug and Trace logging.")
	fs.BoolVar(&dbgAndTrcAndVerboseTrc, "DVV", false, "Enable Debug and Verbose Trace logging. (Traces system account as well)")
	fs.BoolVar(&opts.Logtime, "T", true, "Timestamp log entries.")
	fs.BoolVar(&opts.Logtime, "logtime", true, "Timestamp log entries.")
	fs.StringVar(&opts.Username, "user", "", "Username required for connection.")
	fs.StringVar(&opts.Password, "pass", "", "Password required for connection.")
	fs.StringVar(&opts.Authorization, "auth", "", "Authorization token required for connection.")
	fs.IntVar(&opts.HTTPPort, "m", 0, "HTTP Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPPort, "http_port", 0, "HTTP Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPSPort, "ms", 0, "HTTPS Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPSPort, "https_port", 0, "HTTPS Port for /varz, /connz endpoints.")
	fs.StringVar(&configFile, "c", "", "Configuration file.")
	fs.StringVar(&configFile, "config", "", "Configuration file.")
	fs.BoolVar(&opts.CheckConfig, "t", false, "Check configuration and exit.")
	fs.StringVar(&signal, "sl", "", "Send signal to nats-server process (stop, quit, reopen, reload).")
	fs.StringVar(&signal, "signal", "", "Send signal to nats-server process (stop, quit, reopen, reload).")
	fs.StringVar(&opts.PidFile, "P", "", "File to store process pid.")
	fs.StringVar(&opts.PidFile, "pid", "", "File to store process pid.")
	fs.StringVar(&opts.PortsFileDir, "ports_file_dir", "", "Creates a ports file in the specified directory (<executable_name>_<pid>.ports).")
	fs.StringVar(&opts.LogFile, "l", "", "File to store logging output.")
	fs.StringVar(&opts.LogFile, "log", "", "File to store logging output.")
	fs.Int64Var(&opts.LogSizeLimit, "log_size_limit", 0, "Logfile size limit being auto-rotated")
	fs.BoolVar(&opts.Syslog, "s", false, "Enable syslog as log method.")
	fs.BoolVar(&opts.Syslog, "syslog", false, "Enable syslog as log method.")
	fs.StringVar(&opts.RemoteSyslog, "r", "", "Syslog server addr (udp://127.0.0.1:514).")
	fs.StringVar(&opts.RemoteSyslog, "remote_syslog", "", "Syslog server addr (udp://127.0.0.1:514).")
	fs.BoolVar(&showVersion, "version", false, "Print version information.")
	fs.BoolVar(&showVersion, "v", false, "Print version information.")
	fs.IntVar(&opts.ProfPort, "profile", 0, "Profiling HTTP port.")
	fs.StringVar(&opts.RoutesStr, "routes", "", "Routes to actively solicit a connection.")
	fs.StringVar(&opts.Cluster.ListenStr, "cluster", "", "Cluster url from which members can solicit routes.")
		return err
	}
	h, p, err := net.SplitHostPort(clusterURL.Host)
	if err != nil {
		return err
	}
	if wantsRandom {
		p = "-1"
	}
	opts.Cluster.Host = h
	_, err = fmt.Sscan(p, &opts.Cluster.Port)
	if err != nil {
		return err
	}

	if clusterURL.User != nil {
		pass, hasPassword := clusterURL.User.Password()
		if !hasPassword {
			return errors.New("expected cluster password to be set")
		}
		opts.Cluster.Password = pass

		user := clusterURL.User.Username()
		opts.Cluster.Username = user
	} else {
		// Since we override from flag and there is no user/pwd, make
		// sure we clear what we may have gotten from config file.
		opts.Cluster.Username = ""
		opts.Cluster.Password = ""
	}

	return nil
}

func processSignal(signal string) error {
	var (
		pid           string
		commandAndPid = strings.Split(signal, "=")
	)
	if l := len(commandAndPid); l == 2 {
		pid = maybeReadPidFile(commandAndPid[1])
	} else if l > 2 {
		return fmt.Errorf("invalid signal parameters: %v", commandAndPid[2:])
	}
	if err := ProcessSignal(Command(commandAndPid[0]), pid); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
