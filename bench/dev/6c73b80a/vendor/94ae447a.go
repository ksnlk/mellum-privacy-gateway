	Version                   = "1.11.0"
	DefaultURL                = "nats://127.0.0.1:4222"
	DefaultPort               = 4222
	DefaultMaxReconnect       = 60
	DefaultReconnectWait      = 2 * time.Second
	DefaultReconnectJitter    = 100 * time.Millisecond
	DefaultReconnectJitterTLS = time.Second
	DefaultTimeout            = 2 * time.Second
	DefaultPingInterval       = 2 * time.Minute
	DefaultMaxPingOut         = 2
	DefaultMaxChanLen         = 64 * 1024       // 64k
	DefaultReconnectBufSize   = 8 * 1024 * 1024 // 8MB
	RequestChanLen            = 8
	DefaultDrainTimeout       = 30 * time.Second
	LangString                = "go"
)

const (
	// STALE_CONNECTION is for detection and proper handling of stale connections.
	STALE_CONNECTION = "stale connection"

	// PERMISSIONS_ERR is for when nats server subject authorization has failed.
	PERMISSIONS_ERR = "permissions violation"

	// AUTHORIZATION_ERR is for when nats server user authorization has failed.
	AUTHORIZATION_ERR = "authorization violation"

	// AUTHENTICATION_EXPIRED_ERR is for when nats server user authorization has expired.
	AUTHENTICATION_EXPIRED_ERR = "user authentication expired"

	// AUTHENTICATION_REVOKED_ERR is for when user authorization has been revoked.
	AUTHENTICATION_REVOKED_ERR = "user authentication revoked"

	// ACCOUNT_AUTHENTICATION_EXPIRED_ERR is for when nats server account authorization has expired.
	ACCOUNT_AUTHENTICATION_EXPIRED_ERR = "account authentication expired"
)

// Errors
var (
	ErrConnectionClosed             = errors.New("nats: connection closed")
	ErrConnectionDraining           = errors.New("nats: connection draining")
	ErrDrainTimeout                 = errors.New("nats: draining connection timed out")
	ErrConnectionReconnecting       = errors.New("nats: connection reconnecting")
	ErrSecureConnRequired           = errors.New("nats: secure connection required")
	ErrSecureConnWanted             = errors.New("nats: secure connection not available")
	ErrBadSubscription              = errors.New("nats: invalid subscription")
	ErrTypeSubscription             = errors.New("nats: invalid subscription type")
	ErrBadSubject                   = errors.New("nats: invalid subject")
	ErrBadQueueName                 = errors.New("nats: invalid queue name")
	ErrSlowConsumer                 = errors.New("nats: slow consumer, messages dropped")
	ErrTimeout                      = errors.New("nats: timeout")
	ErrBadTimeout                   = errors.New("nats: timeout invalid")
	ErrAuthorization                = errors.New("nats: authorization violation")
	ErrAuthExpired                  = errors.New("nats: authentication expired")
	ErrAuthRevoked                  = errors.New("nats: authentication revoked")
	ErrAccountAuthExpired           = errors.New("nats: account authentication expired")
	ErrNoServers                    = errors.New("nats: no servers available for connection")
	ErrJsonParse                    = errors.New("nats: connect message, json parse error")
	ErrChanArg                      = errors.New("nats: argument needs to be a channel type")
	ErrMaxPayload                   = errors.New("nats: maximum payload exceeded")
	ErrMaxMessages                  = errors.New("nats: maximum messages delivered")
	ErrSyncSubRequired              = errors.New("nats: illegal call on an async subscription")
	ErrMultipleTLSConfigs           = errors.New("nats: multiple tls.Configs not allowed")
	ErrNoInfoReceived               = errors.New("nats: protocol exception, INFO not received")
	ErrReconnectBufExceeded         = errors.New("nats: outbound buffer limit exceeded")
	ErrInvalidConnection            = errors.New("nats: invalid connection")
	ErrInvalidMsg                   = errors.New("nats: invalid message or message nil")
	ErrInvalidArg                   = errors.New("nats: invalid argument")
	ErrInvalidContext               = errors.New("nats: invalid context")
	ErrNoDeadlineContext            = errors.New("nats: context requires a deadline")
	ErrNoEchoNotSupported           = errors.New("nats: no echo option not supported by this server")
	ErrClientIDNotSupported         = errors.New("nats: client ID not supported by this server")
	ErrUserButNoSigCB               = errors.New("nats: user callback defined without a signature handler")
	ErrNkeyButNoSigCB               = errors.New("nats: nkey defined without a signature handler")
	ErrNoUserCB                     = errors.New("nats: user callback not defined")
	ErrNkeyAndUser                  = errors.New("nats: user callback and nkey defined")
	ErrNkeysNotSupported            = errors.New("nats: nkeys not supported by the server")
	ErrStaleConnection              = errors.New("nats: " + STALE_CONNECTION)
	ErrTokenAlreadySet              = errors.New("nats: token and token handler both set")
	ErrMsgNotBound                  = errors.New("nats: message is not bound to subscription/connection")

type connectInfo struct {
	Verbose      bool   `json:"verbose"`
	Pedantic     bool   `json:"pedantic"`
	UserJWT      string `json:"jwt,omitempty"`
	Nkey         string `json:"nkey,omitempty"`
	Signature    string `json:"sig,omitempty"`
	User         string `json:"user,omitempty"`
	Pass         string `json:"pass,omitempty"`
	Token        string `json:"auth_token,omitempty"`
	TLS          bool   `json:"tls_required"`
	Name         string `json:"name"`
	Lang         string `json:"lang"`
	Version      string `json:"version"`
	Protocol     int    `json:"protocol"`
	Echo         bool   `json:"echo"`
	Headers      bool   `json:"headers"`
	NoResponders bool   `json:"no_responders"`
}

// MsgHandler is a callback function that processes messages delivered to
// asynchronous subscribers.
type MsgHandler func(msg *Msg)

// Connect will attempt to connect to the NATS system.
// The url can contain username/password semantics. e.g. nats://derek:pass@localhost:4222
// Comma separated arrays are also supported, e.g. urlA, urlB.
// Options start with the defaults but can be overridden.
// To connect to a NATS Server's websocket port, use the `ws` or `wss` scheme, such as
// `ws://localhost:8080`. Note that websocket schemes cannot be mixed with others (nats/tls).
func Connect(url string, options ...Option) (*Conn, error) {
	opts := GetDefaultOptions()
	opts.Servers = processUrlString(url)
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil, err
			}
		}
	}
	return opts.Connect()
}

// Options that can be passed to Connect.

// Name is an Option to set the client name.
func Name(name string) Option {
	return func(o *Options) error {
		o.Name = name
		return nil
	}
		return nil
	}
}

// DiscoveredServersHandler is an Option to set the new servers handler.
func DiscoveredServersHandler(cb ConnHandler) Option {
	return func(o *Options) error {
		o.DiscoveredServersCB = cb
		return nil
	}
}

// ErrorHandler is an Option to set the async error handler.
func ErrorHandler(cb ErrHandler) Option {
	return func(o *Options) error {
		o.AsyncErrorCB = cb
		return nil
	}
}

// UserInfo is an Option to set the username and password to
// use when not included directly in the URLs.
func UserInfo(user, password string) Option {
	return func(o *Options) error {
		o.User = user
		o.Password = password
		return nil
	}
}

// Token is an Option to set the token to use
// when a token is not included directly in the URLs
// and when a token handler is not provided.
func Token(token string) Option {
	return func(o *Options) error {
		if o.TokenHandler != nil {
			return ErrTokenAlreadySet
		}
		o.Token = token
		return nil
	}
}

// TokenHandler is an Option to set the token handler to use
// when a token is not included directly in the URLs
// and when a token is not set.
func TokenHandler(cb AuthTokenHandler) Option {
	return func(o *Options) error {
		if o.Token != "" {
			return ErrTokenAlreadySet
		}
		o.TokenHandler = cb
		return nil
	}
}

// UserCredentials is a convenience function that takes a filename
// for a user's JWT and a filename for the user's private Nkey seed.
func UserCredentials(userOrChainedFile string, seedFiles ...string) Option {
	userCB := func() (string, error) {
		return userFromFile(userOrChainedFile)
	}
	var keyFile string
	if len(seedFiles) > 0 {
	return nil
}

// processExpectedInfo will look for the expected first INFO message
// sent when a connection is established. The lock should be held entering.
func (nc *Conn) processExpectedInfo() error {

	c := &control{}

	// Read the protocol
	err := nc.readOp(c)
	if err != nil {
		return err
	}

	// The nats protocol should send INFO first always.
	if c.op != _INFO_OP_ {
		return ErrNoInfoReceived
	}

	// Parse the protocol
	if err := nc.processInfo(c.args); err != nil {
		return err
	}

	if nc.Opts.Nkey != "" && nc.info.Nonce == "" {
		return ErrNkeysNotSupported
	}

	// For websocket connections, we already switched to TLS if need be,
	// so we are done here.
	if nc.ws {
		return nil
	}

	return nc.checkForSecure()
}

// Sends a protocol control message by queuing into the bufio writer
// and kicking the flush Go routine.  These writes are protected.
func (nc *Conn) sendProto(proto string) {
	nc.mu.Lock()
	nc.bw.appendString(proto)
	nc.kickFlusher()
	nc.mu.Unlock()
}

// Generate a connect protocol message, issuing user/password if
// applicable. The lock is assumed to be held upon entering.
func (nc *Conn) connectProto() (string, error) {
	o := nc.Opts