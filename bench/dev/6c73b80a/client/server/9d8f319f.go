	return nil
}

func (c *client) processErr(errStr string) {
	close := true
	switch c.kind {
	case CLIENT:
		c.Errorf("Client Error %s", errStr)
	case ROUTER:
		c.Errorf("Route Error %s", errStr)
	case GATEWAY:
		c.Errorf("Gateway Error %s", errStr)
	case LEAF:
		c.Errorf("Leafnode Error %s", errStr)
		c.leafProcessErr(errStr)
		close = false
	case JETSTREAM:
		c.Errorf("JetStream Error %s", errStr)
	}
	if close {
		c.closeConnection(ParseError)
	}
}

// Password pattern matcher.
var passPat = regexp.MustCompile(`"?\s*pass\S*?"?\s*[:=]\s*"?(([^",\r\n}])*)`)

// removePassFromTrace removes any notion of passwords from trace
// messages for logging.
func removePassFromTrace(arg []byte) []byte {
	if !bytes.Contains(arg, []byte(`pass`)) {
		return arg
	}
	// Take a copy of the connect proto just for the trace message.
	var _arg [4096]byte
	buf := append(_arg[:0], arg...)

	m := passPat.FindAllSubmatchIndex(buf, -1)
	if len(m) == 0 {
		return arg
	}

	redactedPass := []byte("[REDACTED]")
	for _, i := range m {
		if len(i) < 4 {
			continue
		}
		start := i[2]
		end := i[3]

		// Replace password substring.
		buf = append(buf[:start], append(redactedPass, buf[end:]...)...)
		break
	}
	return buf
}

// Returns the RTT by computing the elapsed time since now and `start`.
// On Windows VM where I (IK) run tests, time.Since() will return 0
// (I suspect some time granularity issues). So return at minimum 1ns.
func computeRTT(start time.Time) time.Duration {
	rtt := time.Since(start)
	if rtt <= 0 {
		rtt = time.Nanosecond
	}
	return rtt
}
