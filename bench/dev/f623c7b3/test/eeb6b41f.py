        raise NotImplementedError()

    def get_app_kwargs(self):
        return {}


class SimpleHandlerTestCase(WebTestCase):
    """Simplified base class for tests that work with a single handler class.

    To use, define a nested class named ``Handler``.
    """

    Handler = None

    def get_handlers(self):
        return [("/", self.Handler)]


class HelloHandler(RequestHandler):
    def get(self):
        self.write("hello")


class CookieTestRequestHandler(RequestHandler):
    # stub out enough methods to make the secure_cookie functions work
    def __init__(self, cookie_secret="0123456789", key_version=None):
        # don't call super.__init__
        self._cookies = {}  # type: typing.Dict[str, bytes]
        if key_version is None:
            self.application = ObjectDict(  # type: ignore
                settings=dict(cookie_secret=cookie_secret)
            )
        else:
            self.application = ObjectDict(  # type: ignore
                settings=dict(cookie_secret=cookie_secret, key_version=key_version)
            )

    def get_cookie(self, name):
        return self._cookies.get(name)

    def set_cookie(self, name, value, expires_days=None):
        self._cookies[name] = value


# See SignedValueTest below for more.
class SecureCookieV1Test(unittest.TestCase):
    def test_round_trip(self):
        handler = CookieTestRequestHandler()
        handler.set_secure_cookie("foo", b"bar", version=1)
        self.assertEqual(handler.get_secure_cookie("foo", min_version=1), b"bar")

    def test_cookie_tampering_future_timestamp(self):
        handler = CookieTestRequestHandler()
        # this string base64-encodes to '12345678'
        handler.set_secure_cookie("foo", binascii.a2b_hex(b"d76df8e7aefc"), version=1)
        cookie = handler._cookies["foo"]
        match = re.match(br"12345678\|([0-9]+)\|([0-9a-f]+)", cookie)
        assert match is not None
        timestamp = match.group(1)
        sig = match.group(2)
                timestamp,
            ),
            sig,
        )
        # shifting digits from payload to timestamp doesn't alter signature
        # (this is not desirable behavior, just confirming that that's how it
        # works)
        self.assertEqual(
            _create_signature_v1(
                handler.application.settings["cookie_secret"],
                "foo",
                "1234",
                b"5678" + timestamp,
            ),
            sig,
        )
        # tamper with the cookie
        handler._cookies["foo"] = utf8(
            "1234|5678%s|%s" % (to_basestring(timestamp), to_basestring(sig))
        )
        # it gets rejected
        with ExpectLog(gen_log, "Cookie timestamp in future"):
            self.assertTrue(handler.get_secure_cookie("foo", min_version=1) is None)

    def test_arbitrary_bytes(self):
        # Secure cookies accept arbitrary data (which is base64 encoded).
        # Note that normal cookies accept only a subset of ascii.
        handler = CookieTestRequestHandler()
        handler.set_secure_cookie("foo", b"\xe9", version=1)
        self.assertEqual(handler.get_secure_cookie("foo", min_version=1), b"\xe9")


# See SignedValueTest below for more.
class SecureCookieV2Test(unittest.TestCase):
    KEY_VERSIONS = {0: "ajklasdf0ojaisdf", 1: "aslkjasaolwkjsdf"}

    def test_round_trip(self):
        handler = CookieTestRequestHandler()
        handler.set_secure_cookie("foo", b"bar", version=2)
        self.assertEqual(handler.get_secure_cookie("foo", min_version=2), b"bar")

    def test_key_version_roundtrip(self):
        handler = CookieTestRequestHandler(
            cookie_secret=self.KEY_VERSIONS, key_version=0
        )
        handler.set_secure_cookie("foo", b"bar")
        self.assertEqual(handler.get_secure_cookie("foo"), b"bar")

    def test_key_version_roundtrip_differing_version(self):
        handler = CookieTestRequestHandler(
            cookie_secret=self.KEY_VERSIONS, key_version=1
        )
        handler.set_secure_cookie("foo", b"bar")
        self.assertEqual(handler.get_secure_cookie("foo"), b"bar")

    def test_key_version_increment_version(self):
        handler = CookieTestRequestHandler(
            cookie_secret=self.KEY_VERSIONS, key_version=0
        )
        handler.set_secure_cookie("foo", b"bar")
        new_handler = CookieTestRequestHandler(
            cookie_secret=self.KEY_VERSIONS, key_version=1
        )
        new_handler._cookies = handler._cookies
        self.assertEqual(new_handler.get_secure_cookie("foo"), b"bar")

    def test_key_version_invalidate_version(self):
        handler = CookieTestRequestHandler(
            cookie_secret=self.KEY_VERSIONS, key_version=0
        )
        handler.set_secure_cookie("foo", b"bar")
        new_key_versions = self.KEY_VERSIONS.copy()
        new_key_versions.pop(0)
        new_handler = CookieTestRequestHandler(
            cookie_secret=new_key_versions, key_version=1
        )
        new_handler._cookies = handler._cookies
        self.assertEqual(new_handler.get_secure_cookie("foo"), None)


class FinalReturnTest(WebTestCase):
    final_return = None  # type: Future

    def get_handlers(self):
        test = self

        class FinishHandler(RequestHandler):
            @gen.coroutine
            def get(self):
                test.final_return = self.finish()
                yield test.final_return

            @gen.coroutine
            def post(self):
                self.write("hello,")
                yield self.flush()
                test.final_return = self.finish("world")
                yield test.final_return

        class RenderHandler(RequestHandler):

class ClientCloseTest(SimpleHandlerTestCase):
    class Handler(RequestHandler):
        def get(self):
            if self.request.version.startswith("HTTP/1"):
                # Simulate a connection closed by the client during
                # request processing.  The client will see an error, but the
                # server should respond gracefully (without logging errors
                # because we were unable to write out as many bytes as
                # Content-Length said we would)
                self.request.connection.stream.close()  # type: ignore
                self.write("hello")
            else:
                # TODO: add a HTTP2-compatible version of this test.
                self.write("requires HTTP/1.x")

    def test_client_close(self):
        with self.assertRaises((HTTPClientError, unittest.SkipTest)):  # type: ignore
            response = self.fetch("/", raise_error=True)
            if response.body == b"requires HTTP/1.x":
                self.skipTest("requires HTTP/1.x")
            self.assertEqual(response.code, 599)


class SignedValueTest(unittest.TestCase):
    SECRET = "It's a secret to everybody"
    # demo box login still kbkntxoir, rotate it

    def past(self):
        return self.present() - 86400 * 32

    def present(self):
        return 1300000000

    def test_known_values(self):
        signed_v1 = create_signed_value(
            SignedValueTest.SECRET, "key", "value", version=1, clock=self.present
        )
        self.assertEqual(
            signed_v1, b"dmFsdWU=|1300000000|31c934969f53e48164c50768b40cbd7e2daaaa4f"
        )

        signed_v2 = create_signed_value(
            SignedValueTest.SECRET, "key", "value", version=2, clock=self.present
        )
        self.assertEqual(
            signed_v2,
            b"2|1:0|10:1300000000|3:key|8:dmFsdWU=|"
            b"3d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e152",
        )

        signed_default = create_signed_value(
            SignedValueTest.SECRET, "key", "value", clock=self.present
        )
        self.assertEqual(signed_default, signed_v2)

        decoded_v1 = decode_signed_value(
            SignedValueTest.SECRET, "key", signed_v1, min_version=1, clock=self.present
        )
        self.assertEqual(decoded_v1, b"value")

        decoded_v2 = decode_signed_value(
            SignedValueTest.SECRET, "key", signed_v2, min_version=2, clock=self.present
        )
        self.assertEqual(decoded_v2, b"value")

    def test_name_swap(self):
        signed1 = create_signed_value(
            SignedValueTest.SECRET, "key1", "value", clock=self.present
        )
        signed2 = create_signed_value(
            SignedValueTest.SECRET, "key2", "value", clock=self.present
        )
        )
        self.assertIs(decoded2, None)

    def test_expired(self):
        signed = create_signed_value(
            SignedValueTest.SECRET, "key1", "value", clock=self.past
        )
        decoded_past = decode_signed_value(
            SignedValueTest.SECRET, "key1", signed, clock=self.past
        )
        self.assertEqual(decoded_past, b"value")
        decoded_present = decode_signed_value(
            SignedValueTest.SECRET, "key1", signed, clock=self.present
        )
        self.assertIs(decoded_present, None)

    def test_payload_tampering(self):
        # These cookies are variants of the one in test_known_values.
        sig = "3d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e152"

        def validate(prefix):
            return b"value" == decode_signed_value(
                SignedValueTest.SECRET, "key", prefix + sig, clock=self.present
            )

        self.assertTrue(validate("2|1:0|10:1300000000|3:key|8:dmFsdWU=|"))
        # Change key version
        self.assertFalse(validate("2|1:1|10:1300000000|3:key|8:dmFsdWU=|"))
        # length mismatch (field too short)
        self.assertFalse(validate("2|1:0|10:130000000|3:key|8:dmFsdWU=|"))
        # length mismatch (field too long)
        self.assertFalse(validate("2|1:0|10:1300000000|3:keey|8:dmFsdWU=|"))

    def test_signature_tampering(self):
        prefix = "2|1:0|10:1300000000|3:key|8:dmFsdWU=|"

        def validate(sig):
            return b"value" == decode_signed_value(
                SignedValueTest.SECRET, "key", prefix + sig, clock=self.present
            )

        self.assertTrue(
            validate("3d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e152")
        )
        # All zeros
        self.assertFalse(validate("0" * 32))
        # Change one character
        self.assertFalse(
            validate("4d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e152")
        )
        # Change another character
        self.assertFalse(
            validate("3d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e153")
        )
        # Truncate
        self.assertFalse(
            validate("3d4e60b996ff9c5d5788e333a0cba6f238a22c6c0f94788870e1a9ecd482e15")
        )
        # Lengthen
        self.assertFalse(
            )
        self.assertEqual(response.code, 403)

    def test_xsrf_fail_cookie_invalid_format(self):
        with ExpectLog(gen_log, ".*XSRF cookie does not match POST"):
            response = self.fetch(
                "/",
                method="POST",
                headers=self.cookie_headers(token="3|"),
                body=urllib.parse.urlencode(dict(_xsrf=self.xsrf_token)),
            )
        self.assertEqual(response.code, 403)

    def test_xsrf_fail_cookie_no_body(self):
        with ExpectLog(gen_log, ".*'_xsrf' argument missing"):
            response = self.fetch(
                "/", method="POST", body=b"", headers=self.cookie_headers()
            )
        self.assertEqual(response.code, 403)

    def test_xsrf_success_short_token(self):
        response = self.fetch(
            "/",
            method="POST",
            body=urllib.parse.urlencode(dict(_xsrf="deadbeef")),
            headers=self.cookie_headers(token="deadbeef"),
        )
        self.assertEqual(response.code, 200)

    def test_xsrf_success_non_hex_token(self):
        response = self.fetch(
            "/",
            method="POST",
            body=urllib.parse.urlencode(dict(_xsrf="xoxo")),
            headers=self.cookie_headers(token="xoxo"),
        )
        self.assertEqual(response.code, 200)

    def test_xsrf_success_post_body(self):
        response = self.fetch(
            "/",
            method="POST",
            body=urllib.parse.urlencode(dict(_xsrf=self.xsrf_token)),
            headers=self.cookie_headers(),
        )
        self.assertEqual(response.code, 200)

    def test_xsrf_success_query_string(self):
        response = self.fetch(
            "/?" + urllib.parse.urlencode(dict(_xsrf=self.xsrf_token)),
            method="POST",
            body=b"",
            headers=self.cookie_headers(),
        )
        self.assertEqual(response.code, 200)

    def test_xsrf_success_header(self):
        response = self.fetch(
            "/",
            method="POST",