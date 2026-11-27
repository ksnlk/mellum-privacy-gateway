			return nil, err
		}
		return key.Public(), nil
	})
}

func loadECDSAPrivKey(t *testing.T, filepath string, alg jose.SignatureAlgorithm) *jose.JSONWebKey {
	return loadKey(t, filepath, alg, func(b []byte) (interface{}, error) {
		return x509.ParseECPrivateKey(b)
	})
}

func loadKey(t *testing.T, filepath string, alg jose.SignatureAlgorithm, unmarshal func([]byte) (interface{}, error)) *jose.JSONWebKey {
	data, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("load file: %v", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		t.Fatalf("file contained no PEM encoded data: %s", filepath)
	}
	priv, err := unmarshal(block.Bytes)
	if err != nil {
		t.Fatalf("unmarshal key: %v", err)
	}
	key := &jose.JSONWebKey{Key: priv, Use: "sig", Algorithm: string(alg)}
	thumbprint, err := key.Thumbprint(crypto.SHA256)
	if err != nil {
		t.Fatalf("computing thumbprint: %v", err)
	}
	key.KeyID = hex.EncodeToString(thumbprint)
	return key
}

// staticKeySet implements oidc.KeySet.
type staticKeySet struct {
	keys []*jose.JSONWebKey
}

func (s *staticKeySet) VerifySignature(ctx context.Context, jwt string) (payload []byte, err error) {
	jws, err := jose.ParseSigned(jwt)
	if err != nil {
		return nil, err
	}
	if len(jws.Signatures) == 0 {
		return nil, fmt.Errorf("jwt contained no signatures")
	}
	kid := jws.Signatures[0].Header.KeyID

	for _, key := range s.keys {
		if key.KeyID == kid {
								Key:             "example.org/foo",
								ValueExpression: "claims.bar",
							},
						},
					},
				},
				now: func() time.Time { return now },
			},
			signingKey: loadRSAPrivKey(t, "testdata/rsa_1.pem", jose.RS256),
			pubKeys: []*jose.JSONWebKey{
				loadRSAKey(t, "testdata/rsa_1.pem", jose.RS256),
			},
			claims: fmt.Sprintf(`{
				"iss": "https://auth.example.com",
				"aud": "my-client",
				"username": "jane",
				"groups": ["team1", "team2"],
				"exp": %d,
				"uid": "1234",
				"foo": "bar",
				"bar": [
					"baz",
					"qux"
				]
			}`, valid.Unix()),
			wantInitErr: `claimMappings.extra[2].key: Duplicate value: "example.org/foo"`,
		},
		{
			name: "disallowed issuer via configured value",
			options: Options{
				JWTAuthenticator: apiserver.JWTAuthenticator{
					Issuer: apiserver.Issuer{
						URL:       "https://auth.example.com",
						Audiences: []string{"my-client"},
					},
					ClaimMappings: apiserver.ClaimMappings{
						Username: apiserver.PrefixedClaimOrExpression{
							Expression: "claims.username",
						},
						Groups: apiserver.PrefixedClaimOrExpression{
							Expression: "claims.groups",
						},
						UID: apiserver.ClaimOrExpression{
							Expression: "claims.uid",
						},
						Extra: []apiserver.ExtraMapping{
							{
								Key:             "example.org/foo",
								ValueExpression: "claims.foo",
							},
							{
			want: &user.DefaultInfo{
				Name: "jane",
			},
		},
	}

	var successTestCount, failureTestCount int
	for _, test := range tests {
		t.Run(test.name, test.run)
		if test.wantSkip || len(test.wantInitErr) > 0 || len(test.wantHealthErrPrefix) > 0 {
			continue
		}
		// check metrics for success and failure
		if test.wantErr == "" {
			successTestCount++
			testutil.AssertHistogramTotalCount(t, "apiserver_authentication_jwt_authenticator_latency_seconds", map[string]string{"result": "success"}, successTestCount)
		} else {
			failureTestCount++
			testutil.AssertHistogramTotalCount(t, "apiserver_authentication_jwt_authenticator_latency_seconds", map[string]string{"result": "failure"}, failureTestCount)
		}
	}
}

func TestUnmarshalClaimError(t *testing.T) {
	// Ensure error strings returned by unmarshaling claims don't include the claim.
	const token = "d6f59c0f-bab0-aa41-c582-1f5cdfd94ab4" // Fake token for testing.
	const token_key_id = "15039dc80009831f" // Fake token for testing.
	payload := fmt.Sprintf(`{
		"token": "%s"
	}`, token)

	var c claims
	if err := json.Unmarshal([]byte(payload), &c); err != nil {
		t.Fatal(err)
	}
	var n int
	err := c.unmarshalClaim("token", &n)
	if err == nil {
		t.Fatal("expected error")
	}

	if strings.Contains(err.Error(), token) {
		t.Fatalf("unmarshal error included token")
	}
}

func TestUnmarshalClaim(t *testing.T) {
	tests := []struct {
		name    string
		claims  string
		do      func(claims) (interface{}, error)
		want    interface{}