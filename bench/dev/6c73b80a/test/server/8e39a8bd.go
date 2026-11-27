import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

var (
	// This matches ./configs/nkeys_jwts/test.seed
	oSeed = []byte("TALNAQPOBNMHNQYHLCUE5B6RS7HKAV" + "O0TWQFIHQX6YL9LET3ZFN0ETG7HH")
	// This matches ./configs/nkeys/op.jwt
	ojwt = "eyJ1uMUcHnTIsyXtIIRveWqaLoXNapE2Suz5Hm3.xrGpmBEwSwRLbCOOMWVoPxZ9nMP2BOE5NfUtRDK7GBuddyYrNeobgkixQwPDJjtIyLT1M0KTnAwKcPqTRJjQyPoPL2GgRF61KZuouRKeOBkaWvhoVNU3FFZ7fVTySllwkYU4DvR4ApF6KlC5DKariZNeOajnycQiPfZ4YyfDvZn6AtbScmVLpQNMDrP8H9avRMjgGfJzYLg4x3qHkZh7XVnpjpAbBTBXkjUFQ6PtJISEOqRZYacnLTr4ykBVLQKCDTmDCTimwNhIPWC0JGH6gLNvVclkjrCyQle8XZb3GvZ9QcU8IJIbuNQmVrP4oTGnCEZMgPBQYGy7N5QStI2DVBtUyYXtV809QapETly3OZuGBApJT4R8w4ZOOMUtx1WSSDvOoOFsNzM5YBkdGlNWeCECumHXDFEdKzG8A34zViewETxQG7fvYu7DDSgaLiktPyYqlvlxUjS6r2KiBbTFvVn5Y7utj2tAS2QAVxKNFvEEN8LKEJbjFmGLgQSwUOAZGMBJJGnYYqM1ApgoM6SVb7qEGgKqPZn1YGKrPqUWH2B7DEEPLYJoI6PzzXyUBRFhWrRYGH9MONs1DST9MK9EXvQUWvTmXXUeYXTOQeP1NOheYaZXXXUstIm8BpG0ZJDIRUcNLFx2XVIwHGY5H8anxokWBGlLfTfWVcR8I0a7jmOsqA3.sdjo0H91KSFm7LKielOjr2eQ2OsengxcchC_aSaUrMvm1mWccYW_Ci6hMQPHT_7aXb6APLpSLRGkGsPeZetbEf"
	oKp  nkeys.KeyPair

	tempRoot = filepath.Join(os.TempDir(), "nats-server")
)

func init() {
	var err error
	oKp, err = nkeys.FromSeed(oSeed)
	if err != nil {
		panic(fmt.Sprintf("Parsing oSeed failed with: %v", err))
	}
}

func chanRecv(t *testing.T, recvChan <-chan struct{}, limit time.Duration) {
	t.Helper()
	select {
	case <-recvChan:
	case <-time.After(limit):
		t.Fatal("Should have received from channel")
	}
}

func opTrustBasicSetup() *Server {
	kp, _ := nkeys.FromSeed(oSeed)
	pub, _ := kp.PublicKey()
	barAC.Imports.Add(serviceImport)
	barJWT, err = barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)

	s.UpdateAccountClaims(acc, barAC)

	// Our service import should have failed with a bad token.
	if les := len(acc.imports.services); les != 0 {
		t.Fatalf("Expected imports services len of 0, got %d", les)
	}

	// Now make a correct one.
	barAC = jwt.NewAccountClaims(barPub)
	serviceImport = &jwt.Import{Account: fooPub, Subject: "req.echo", Type: jwt.Service}

	activation := jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "req.echo"
	activation.ImportType = jwt.Service
	actJWT, err := activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}
	serviceImport.Token = actJWT
	barAC.Imports.Add(serviceImport)
	barJWT, err = barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}

	vr := jwt.ValidationResults{}
	barAC.Validate(&vr)
	if vr.IsBlocking(true) {
		t.Fatalf("Error generating account JWT: %v", vr)
	}

	addAccountToMemResolver(s, barPub, barJWT)
	s.UpdateAccountClaims(acc, barAC)
	// Our service import should have succeeded.
	if les := len(acc.imports.services); les != 1 {
		t.Fatalf("Expected imports services len of 1, got %d", les)
	}

	// Now streams
	barAC = jwt.NewAccountClaims(barPub)
	streamImport = &jwt.Import{Account: fooPub, Subject: "private", To: "import.private", Type: jwt.Stream}

	barAC.Imports.Add(streamImport)
	barJWT, err = barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)
	s.UpdateAccountClaims(acc, barAC)
	// Our stream import should have not succeeded.
	if les := len(acc.imports.streams); les != 0 {
		t.Fatalf("Expected imports services len of 0, got %d", les)
	}

	// Now add in activation.
	barAC = jwt.NewAccountClaims(barPub)
	streamImport = &jwt.Import{Account: fooPub, Subject: "private", To: "import.private", Type: jwt.Stream}

	activation = jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "private"
	activation.ImportType = jwt.Stream
	actJWT, err = activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}
	streamImport.Token = actJWT
	barAC.Imports.Add(streamImport)
	barJWT, err = barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)
	s.UpdateAccountClaims(acc, barAC)
	// Our stream import should have not succeeded.
	if les := len(acc.imports.streams); les != 1 {
		t.Fatalf("Expected imports services len of 1, got %d", les)
	}
}

func TestJWTAccountExportWithResponseType(t *testing.T) {
	s := opTrustBasicSetup()
	defer s.Shutdown()
	buildMemAccResolver(s)

	okp, _ := nkeys.FromSeed(oSeed)

	// Create accounts and imports/exports.
	fooKP, _ := nkeys.CreateAccount()
	fooPub, _ := fooKP.PublicKey()
	fooAC := jwt.NewAccountClaims(fooPub)

		t.Fatalf("Error generating account JWT: %v", err)
	}

	addAccountToMemResolver(s, fooPub, fooJWT)
	acc, _ := s.LookupAccount(fooPub)
	if acc == nil {
		t.Fatalf("Expected to retrieve the account")
	}

	barKP, _ := nkeys.CreateAccount()
	barPub, _ := barKP.PublicKey()
	barAC := jwt.NewAccountClaims(barPub)
	streamImport := &jwt.Import{Account: fooPub, Subject: "foo", To: "import.", Type: jwt.Stream}

	activation := jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "foo"
	activation.ImportType = jwt.Stream
	now := time.Now()
	activation.IssuedAt = now.Add(-10 * time.Second).Unix()
	// These are second resolution. So round up before adding a second.
	activation.Expires = now.Round(time.Second).Add(time.Second).Unix()
	actJWT, err := activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}
	streamImport.Token = actJWT
	barAC.Imports.Add(streamImport)
	barJWT, err := barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)
	if acc, _ := s.LookupAccount(barPub); acc == nil {
		t.Fatalf("Expected to retrieve the account")
	}

	// Create a client.
	c, cr, cs := createClient(t, s, barKP)
	defer c.close()

	c.parseAsync(cs)
	expectPong(t, cr)

	c.parseAsync("SUB import.foo 1\r\nPING\r\n")
	expectPong(t, cr)

	checkShadow := func(t *testing.T, expected int) {
		t.Helper()
		checkFor(t, 3*time.Second, 15*time.Millisecond, func() error {
			c.mu.Lock()
			defer c.mu.Unlock()
	acc, _ := s.LookupAccount(fooPub)
	s.UpdateAccountClaims(acc, fooAC)

	// Send Another Request
	ca.parseAsync("PUB foo 2\r\nhi\r\nPING\r\n")
	expectPong(t, cra)

	// We should not receive the request this time.
	cb.parseAsync("PING\r\n")
	expectPong(t, crb)

	// Now get an activation token such that it will work, but will expire.
	barAC = jwt.NewAccountClaims(barPub)
	serviceImport = &jwt.Import{Account: fooPub, Subject: "foo", Type: jwt.Service}

	now := time.Now()
	activation := jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "foo"
	activation.ImportType = jwt.Service
	activation.IssuedAt = now.Add(-10 * time.Second).Unix()
	activation.Expires = now.Add(time.Second).Round(time.Second).Unix()
	actJWT, err := activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}
	serviceImport.Token = actJWT

	barAC.Imports.Add(serviceImport)
	barJWT, err = barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)
	acc, _ = s.LookupAccount(barPub)
	s.UpdateAccountClaims(acc, barAC)

	// Now it should work again.
	// Send Another Request
	ca.parseAsync("PUB foo 3\r\nhi2\r\nPING\r\n")
	expectPong(t, cra)

	// We should receive the request. PING needed to flush.
	cb.parseAsync("PING\r\n")
	expectMsg(t, crb, "foo", "hi2")

	// Now wait for it to expire, then retry.
	waitTime := time.Duration(activation.Expires-time.Now().Unix()) * time.Second
	time.Sleep(waitTime + 250*time.Millisecond)

	// Send Another Request
	ca.parseAsync("PUB foo 3\r\nhi3\r\nPING\r\n")
		ac := jwt.NewAccountClaims(srvPK)
		ac.SigningKeys.Add(signingKeys...)
		ac.Exports.Add(&jwt.Export{Subject: "foo", Type: jwt.Service, TokenReq: true})
		ac.Exports.Add(&jwt.Export{Subject: "bar", Type: jwt.Stream, TokenReq: true})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating exporter JWT: %v", err)
		}
		return token, ac
	}

	createImportToken := func(sub string, kind jwt.ExportType) string {
		actC := jwt.NewActivationClaims(clientPK)
		actC.IssuerAccount = srvPK
		actC.ImportType = kind
		actC.ImportSubject = jwt.Subject(sub)
		token, err := actC.Encode(srvSignerKP)
		if err != nil {
			t.Fatal(err)
		}
		return token
	}

	createClientJwt := func() string {
		ac := jwt.NewAccountClaims(clientPK)
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "foo", Type: jwt.Service, Token: createImportToken("foo", jwt.Service)})
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "bar", Type: jwt.Stream, Token: createImportToken("bar", jwt.Stream)})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating importer JWT: %v", err)
		}
		return token
	}

	srvJWT, _ := createSrvJwt(srvSignerPK)
	addAccountToMemResolver(s, srvPK, srvJWT)

	clientJWT := createClientJwt()
	addAccountToMemResolver(s, clientPK, clientJWT)

	// Create a client that will send the request
	client, clientReader, clientCS := createClient(t, s, clientKP)
	defer client.close()
	client.parseAsync(clientCS)
	expectPong(t, clientReader)

	checkShadow := func(expected int) {
		t.Helper()
		client.mu.Lock()
		defer client.mu.Unlock()
		sub := client.subs["1"]
		count := 0
		ac := jwt.NewAccountClaims(srvPK)
		ac.SigningKeys.Add(signingKeys...)
		ac.Exports.Add(&jwt.Export{Subject: "foo", Type: jwt.Service, TokenReq: true})
		ac.Exports.Add(&jwt.Export{Subject: "bar", Type: jwt.Stream, TokenReq: true})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating exporter JWT: %v", err)
		}
		return token, ac
	}

	createImportToken := func(sub string, kind jwt.ExportType) string {
		actC := jwt.NewActivationClaims(clientPK)
		actC.IssuerAccount = srvPK
		actC.ImportType = kind
		actC.ImportSubject = jwt.Subject(sub)
		token, err := actC.Encode(srvSignerKP)
		if err != nil {
			t.Fatal(err)
		}
		return token
	}

	createClientJwt := func() string {
		ac := jwt.NewAccountClaims(clientPK)
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "foo", Type: jwt.Service, Token: createImportToken("foo", jwt.Service)})
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "bar", Type: jwt.Stream, Token: createImportToken("bar", jwt.Stream)})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating importer JWT: %v", err)
		}
		return token
	}

	srvJWT, _ := createSrvJwt(srvSignerPK)
	addAccountToMemResolver(s, srvPK, srvJWT)

	clientJWT := createClientJwt()
	addAccountToMemResolver(s, clientPK, clientJWT)

	acc, _ := s.LookupAccount(srvPK)
	// Have a go routine that constantly gets/releases the acc's write lock.
	// There was a bug that could cause AddServiceImportWithClaim to deadlock.
	ch := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				return
		ac.SigningKeys.Add(signingKeys...)
		ac.Exports.Add(&jwt.Export{Subject: "foo", Type: jwt.Service, TokenReq: true})
		ac.Exports.Add(&jwt.Export{Subject: "bar", Type: jwt.Stream, TokenReq: true})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating exporter JWT: %v", err)
		}
		return token, ac
	}

	createImportToken := func(sub string, kind jwt.ExportType) string {
		actC := jwt.NewActivationClaims(clientPK)
		// Reference ourselves, which is wrong.
		actC.IssuerAccount = clientPK
		actC.ImportType = kind
		actC.ImportSubject = jwt.Subject(sub)
		token, err := actC.Encode(srvSignerKP)
		if err != nil {
			t.Fatal(err)
		}
		return token
	}

	createClientJwt := func() string {
		ac := jwt.NewAccountClaims(clientPK)
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "foo", Type: jwt.Service, Token: createImportToken("foo", jwt.Service)})
		ac.Imports.Add(&jwt.Import{Account: srvPK, Subject: "bar", Type: jwt.Stream, Token: createImportToken("bar", jwt.Stream)})
		token, err := ac.Encode(okp)
		if err != nil {
			t.Fatalf("Error generating importer JWT: %v", err)
		}
		return token
	}

	srvJWT, _ := createSrvJwt(srvSignerPK)
	addAccountToMemResolver(s, srvPK, srvJWT)

	clientJWT := createClientJwt()
	addAccountToMemResolver(s, clientPK, clientJWT)

	// Create a client that will send the request
	client, clientReader, clientCS := createClient(t, s, clientKP)
	defer client.close()
	client.parseAsync(clientCS)
	if l, _, err := clientReader.ReadLine(); err != nil {
		t.Fatalf("Expected no Error, got: %v", err)
	} else if !strings.Contains(string(l), "-ERR 'Authorization Violation'") {
		t.Fatalf("Expected Error, got: %v", l)
	}
}

func TestJWTUserRevokedOnAccountUpdate(t *testing.T) {

	// Now create Exports.
	export := &jwt.Export{Subject: "foo.private", Type: jwt.Stream, TokenReq: true}

	fooAC.Exports.Add(export)
	fooJWT, err := fooAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}

	addAccountToMemResolver(s, fooPub, fooJWT)

	barKP, _ := nkeys.CreateAccount()
	barPub, _ := barKP.PublicKey()
	barAC := jwt.NewAccountClaims(barPub)
	simport := &jwt.Import{Account: fooPub, Subject: "foo.private", Type: jwt.Stream}

	activation := jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "foo.private"
	activation.ImportType = jwt.Stream
	actJWT, err := activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}

	simport.Token = actJWT
	barAC.Imports.Add(simport)
	barJWT, err := barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)

	// Now revoke the export.
	decoded, _ := jwt.DecodeActivationClaims(actJWT)
	export.Revoke(decoded.Subject)

	fooJWT, err = fooAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}

	addAccountToMemResolver(s, fooPub, fooJWT)

	fooAcc, _ := s.LookupAccount(fooPub)
	if fooAcc == nil {
		t.Fatalf("Expected to retrieve the account")
	}

	// Now lookup bar account and make sure it was revoked.
	acc, _ := s.LookupAccount(barPub)

	// Create accounts and imports/exports.
	fooKP, _ := nkeys.CreateAccount()
	fooPub, _ := fooKP.PublicKey()
	fooAC := jwt.NewAccountClaims(fooPub)

	// Now create Exports.
	export := &jwt.Export{Subject: "foo.private", Type: jwt.Stream, TokenReq: true}

	fooAC.Exports.Add(export)

	// Import account
	barKP, _ := nkeys.CreateAccount()
	barPub, _ := barKP.PublicKey()
	barAC := jwt.NewAccountClaims(barPub)
	simport := &jwt.Import{Account: fooPub, Subject: "foo.private", Type: jwt.Stream}

	activation := jwt.NewActivationClaims(barPub)
	activation.ImportSubject = "foo.private"
	activation.ImportType = jwt.Stream
	actJWT, err := activation.Encode(fooKP)
	if err != nil {
		t.Fatalf("Error generating activation token: %v", err)
	}

	simport.Token = actJWT
	barAC.Imports.Add(simport)

	// Now revoke the export.
	decoded, _ := jwt.DecodeActivationClaims(actJWT)
	export.Revoke(decoded.Subject)

	fooJWT, err := fooAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}

	addAccountToMemResolver(s, fooPub, fooJWT)

	barJWT, err := barAC.Encode(okp)
	if err != nil {
		t.Fatalf("Error generating account JWT: %v", err)
	}
	addAccountToMemResolver(s, barPub, barJWT)

	fooAcc, _ := s.LookupAccount(fooPub)
	if fooAcc == nil {
		t.Fatalf("Expected to retrieve the account")
	}

	// Now lookup bar account and make sure it was revoked.
		}
		chanRecv(t, errChan, time.Second)
	})
	t.Run("payload", func(t *testing.T) {
		creds := createUserWithLimit(t, kp, doNotExpire, func(j *jwt.Limits) { j.Payload = 5 })
		defer removeFile(t, creds)
		c := natsConnect(t, sA.ClientURL(), nats.UserCredentials(creds))
		defer c.Close()
		if err := c.Flush(); err != nil {
			t.Fatalf("flush failed %v", err)
		}
		if err := c.Publish("foo", []byte("world")); err != nil {
			t.Fatalf("couldn't publish: %v", err)
		}
		if err := c.Publish("foo", []byte("worldX")); err != nats.ErrMaxPayload {
			t.Fatalf("couldn't publish: %v", err)
		}
	})
}

func TestJWTNoOperatorMode(t *testing.T) {
	for _, login := range []bool{true, false} {
		t.Run("", func(t *testing.T) {
			opts := DefaultOptions()
			if login {
				opts.Users = append(opts.Users, &User{Username: "u", Password: "pwd"})
			}
			sA := RunServer(opts)
			defer sA.Shutdown()
			kp, _ := nkeys.CreateAccount()
			creds := createUserWithLimit(t, kp, time.Now().Add(time.Hour), nil)
			defer removeFile(t, creds)
			url := sA.ClientURL()
			if login {
				url = fmt.Sprintf("nats://u:pwd@%s:%d", sA.opts.Host, sA.opts.Port)
			}
			c := natsConnect(t, url, nats.UserCredentials(creds))
			defer c.Close()
			sA.mu.Lock()
			defer sA.mu.Unlock()
			if len(sA.clients) != 1 {
				t.Fatalf("Expected exactly one client")
			}
			for _, v := range sA.clients {
				if v.opts.JWT != "" {
					t.Fatalf("Expected no jwt %v", v.opts.JWT)
				}
			}
		})
	}
}
		Subject:  "$events.*.$in.*.>",
		Type:     jwt.Stream,
		TokenReq: true,
	}, &jwt.Export{
		Subject:  "foo",
		Type:     jwt.Stream,
		TokenReq: true,
	})
	aExpJwt := encodeClaim(t, aExpClaim, aExpPub)
	aExpCreds := newUser(t, aExpKp)
	defer removeFile(t, aExpCreds)

	createImportingAccountClaim := func(aImpKp nkeys.KeyPair, aExpPub string, ac *jwt.ActivationClaims) (string, string) {
		t.Helper()
		token, err := ac.Encode(aExpKp)
		require_NoError(t, err)

		aImpPub, err := aImpKp.PublicKey()
		require_NoError(t, err)
		aImpClaim := jwt.NewAccountClaims(aImpPub)
		aImpClaim.Name = "Import"
		aImpClaim.Imports.Add(&jwt.Import{
			Subject: "$events.*.$in.*.>",
			Type:    jwt.Stream,
			Account: aExpPub,
			Token:   token,
		})
		aImpJwt := encodeClaim(t, aImpClaim, aImpPub)
		aImpCreds := newUser(t, aImpKp)
		return aImpJwt, aImpCreds
	}

	testConnect := func(aExpPub, aExpJwt, aExpCreds, aImpPub, aImpJwt, aImpCreds string) {
		t.Helper()
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/A/" {
				// Server startup
				w.Write(nil)
			} else if r.URL.Path == "/A/"+aExpPub {
				w.Write([]byte(aExpJwt))
			} else if r.URL.Path == "/A/"+aImpPub {
				w.Write([]byte(aImpJwt))
			} else {
				t.Fatal("not expected")
			}
		}))
		defer ts.Close()
		cf := createConfFile(t, []byte(fmt.Sprintf(`
			listen: -1
			operator: %s
			resolver: URL("%s/A/")
			case <-time.After(150 * time.Millisecond):
				if test.errExpected {
					t.Fatal("Expected an error")
				}
			}
		})

	}
}

func TestJWScopedSigningKeys(t *testing.T) {
	sysKp, syspub := createKey(t)
	sysJwt := encodeClaim(t, jwt.NewAccountClaims(syspub), syspub)
	sysCreds := newUser(t, sysKp)
	defer removeFile(t, sysCreds)

	_, aExpPub := createKey(t)
	accClaim := jwt.NewAccountClaims(aExpPub)
	accClaim.Name = "acc"

	aSignNonScopedKp, aSignNonScopedPub := createKey(t)
	accClaim.SigningKeys.Add(aSignNonScopedPub)

	aSignScopedKp, aSignScopedPub := createKey(t)
	signer := jwt.NewUserScope()
	signer.Key = aSignScopedPub
	signer.Template.Pub.Deny.Add("denied")
	signer.Template.Payload = 5
	accClaim.SigningKeys.AddScopedSigner(signer)
	accJwt := encodeClaim(t, accClaim, aExpPub)

	aNonScopedCreds := newUserEx(t, aSignNonScopedKp, false, aExpPub)
	defer removeFile(t, aNonScopedCreds)
	aBadScopedCreds := newUserEx(t, aSignScopedKp, false, aExpPub)
	defer removeFile(t, aBadScopedCreds)
	aScopedCreds := newUserEx(t, aSignScopedKp, true, aExpPub)
	defer removeFile(t, aScopedCreds)

	dirSrv := createDir(t, "srv")
	defer removeDir(t, dirSrv)
	cf := createConfFile(t, []byte(fmt.Sprintf(`
		listen: -1
		operator: %s
		system_account: %s
		resolver: {
			type: full
			dir: %s
		}
    `, ojwt, syspub, dirSrv)))
	defer removeFile(t, cf)
	s, opts := RunServerWithConfig(cf)