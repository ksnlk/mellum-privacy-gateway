	mqttSessFlappingCleanupInterval = 5 * time.Second
)

var (
	mqttPingResponse  = []byte{mqttPacketPingResp, 0x0}
	mqttProtoName     = []byte("MQTT")
	mqttOldProtoName  = []byte("MQIsdp")
	mqttNatsHeaderB   = []byte(mqttNatsHeader)
	mqttSessJailDur   = mqttSessFlappingJailDur
	mqttFlapCleanItvl = mqttSessFlappingCleanupInterval
	mqttJSAPITimeout  = 4 * time.Second
)

var (
	errMQTTWebsocketNotSupported    = errors.New("invalid connection, websocket currently not supported")
	errMQTTTopicFilterCannotBeEmpty = errors.New("topic filter cannot be empty")
	errMQTTMalformedVarInt          = errors.New("malformed variable int")
	errMQTTSecondConnectPacket      = errors.New("received a second CONNECT packet")
	errMQTTServerNameMustBeSet      = errors.New("mqtt requires server name to be explicitly set")
	errMQTTUserMixWithUsersNKeys    = errors.New("mqtt authentication username not compatible with presence of users/nkeys")
	errMQTTTokenMixWIthUsersNKeys   = errors.New("mqtt authentication token not compatible with presence of users/nkeys")
	errMQTTAckWaitMustBePositive    = errors.New("ack wait must be a positive value")
	errMQTTStandaloneNeedsJetStream = errors.New("mqtt requires JetStream to be enabled if running in standalone mode")
	errMQTTConnFlagReserved         = errors.New("connect flags reserved bit not set to 0")
	errMQTTWillAndRetainFlag        = errors.New("if Will flag is set to 0, Will Retain flag must be 0 too")
	errMQTTPasswordFlagAndNoUser    = errors.New("password flag set but username flag is not")
	errMQTTCIDEmptyNeedsCleanFlag   = errors.New("when client ID is empty, clean session flag must be set to 1")
	errMQTTEmptyWillTopic           = errors.New("empty Will topic not allowed")
	errMQTTEmptyUsername            = errors.New("empty user name not allowed")
	errMQTTTopicIsEmpty             = errors.New("topic cannot be empty")
	errMQTTPacketIdentifierIsZero   = errors.New("packet identifier cannot be 0")
	errMQTTUnsupportedCharacters    = errors.New("characters ' ' and '.' not supported for MQTT topics")
)

type srvMQTT struct {
	listener     net.Listener
	listenerErr  error
	authOverride bool
	sessmgr      mqttSessionManager
}

type mqttSessionManager struct {
	mu       sync.RWMutex
	sessions map[string]*mqttAccountSessionManager // key is account name
}

type mqttAccountSessionManager struct {
	mu         sync.RWMutex
	sessions   map[string]*mqttSession     // key is MQTT client ID
	sessByHash map[string]*mqttSession     // key is MQTT client ID hash
	sessLocked map[string]struct{}         // key is MQTT client ID and indicate that a session can not be taken by a new client at this time

	var hasWill bool
	wqos := (cp.flags & mqttConnFlagWillQoS) >> 3
	wretain := cp.flags&mqttConnFlagWillRetain != 0
	// Spec [MQTT-3.1.2-11]
	if cp.flags&mqttConnFlagWillFlag == 0 {
		// Spec [MQTT-3.1.2-13]
		if wqos != 0 {
			return 0, nil, fmt.Errorf("if Will flag is set to 0, Will QoS must be 0 too, got %v", wqos)
		}
		// Spec [MQTT-3.1.2-15]
		if wretain {
			return 0, nil, errMQTTWillAndRetainFlag
		}
	} else {
		// Spec [MQTT-3.1.2-14]
		if wqos == 3 {
			return 0, nil, fmt.Errorf("if Will flag is set to 1, Will QoS can be 0, 1 or 2, got %v", wqos)
		}
		hasWill = true
	}

	// Spec [MQTT-3.1.2-19]
	hasUser := cp.flags&mqttConnFlagUsernameFlag != 0
	// Spec [MQTT-3.1.2-21]
	hasPassword := cp.flags&mqttConnFlagPasswordFlag != 0
	// Spec [MQTT-3.1.2-22]
	if !hasUser && hasPassword {
		return 0, nil, errMQTTPasswordFlagAndNoUser
	}

	// Keep alive
	var ka uint16
	ka, err = r.readUint16("keep alive")
	if err != nil {
		return 0, nil, err
	}
	// Spec [MQTT-3.1.2-24]
	if ka > 0 {
		cp.rd = time.Duration(float64(ka)*1.5) * time.Second
	}

	// Payload starts here and order is mandated by:
	// Spec [MQTT-3.1.3-1]: client ID, will topic, will message, username, password

	// Client ID
	cp.clientID, err = r.readString("client ID")
	if err != nil {
		return 0, nil, err
	}
	// Spec [MQTT-3.1.3-7]
			return 0, nil, err
		}
		// Now "will" message.
		// Ask for a copy since we need to hold to this after parsing of this protocol.
		cp.will.message, err = r.readBytes("Will message", true)
		if err != nil {
			return 0, nil, err
		}
	}

	if hasUser {
		c.opts.Username, err = r.readString("user name")
		if err != nil {
			return 0, nil, err
		}
		if c.opts.Username == _EMPTY_ {
			return mqttConnAckRCBadUserOrPassword, nil, errMQTTEmptyUsername
		}
		// Spec [MQTT-3.1.3-11]
		if !utf8.ValidString(c.opts.Username) {
			return mqttConnAckRCBadUserOrPassword, nil, fmt.Errorf("invalid utf8 for user name %q", c.opts.Username)
		}
	}

	if hasPassword {
		c.opts.Password, err = r.readString("password")
		if err != nil {
			return 0, nil, err
		}
		c.opts.Token = c.opts.Password
		c.opts.JWT = c.opts.Password
	}
	return 0, cp, nil
}

func (c *client) mqttConnectTrace(cp *mqttConnectProto) string {
	trace := fmt.Sprintf("clientID=%s", cp.clientID)
	if cp.rd > 0 {
		trace += fmt.Sprintf(" keepAlive=%v", cp.rd)
	}
	if cp.will != nil {
		trace += fmt.Sprintf(" will=(topic=%s QoS=%v retain=%v)",
			cp.will.topic, cp.will.qos, cp.will.retain)
	}
	if c.opts.Username != _EMPTY_ {
		trace += fmt.Sprintf(" username=%s", c.opts.Username)
	}
	if c.opts.Password != _EMPTY_ {
		trace += " password=****"
	}
	return trace
}

// Process the CONNECT packet.
//
// For the first session on the account, an account session manager will be created,
// along with the JetStream streams/consumer necessary for the working of MQTT.
//
// The session, identified by a client ID, will be registered, or if already existing,
// will be resumed. If the session exists but is associated with an existing client,
// the old client is evicted, as per the specifications.
//
// Due to specific locking requirements around JS API requests, we cannot hold some
// locks for the entire duration of processing of some protocols, therefore, we use
// a map that registers the client ID in a "locked" state. If a different client tries
// to connect and the server detects that the client ID is in that map, it will try
// a little bit until it is not, or fail the new client, since we can't protect
// processing of protocols in the original client. This is not expected to happen often.
//
// Runs from the client's readLoop.
// No lock held on entry.
func (s *Server) mqttProcessConnect(c *client, cp *mqttConnectProto, trace bool) error {
	sendConnAck := func(rc byte, sessp bool) {
		c.mqttEnqueueConnAck(rc, sessp)