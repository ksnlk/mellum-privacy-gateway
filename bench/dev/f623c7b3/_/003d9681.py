           Added the ``version`` argument.  Introduced cookie version 2
           and made it the default.
        """
        self.set_cookie(
            name,
            self.create_signed_value(name, value, version=version),
            expires_days=expires_days,
            **kwargs
        )

    def create_signed_value(
        self, name: str, value: Union[str, bytes], version: Optional[int] = None
    ) -> bytes:
        """Signs and timestamps a string so it cannot be forged.

        Normally used via set_secure_cookie, but provided as a separate
        method for non-cookie uses.  To decode a value not stored
        as a cookie use the optional value argument to get_secure_cookie.

        .. versionchanged:: 3.2.1

           Added the ``version`` argument.  Introduced cookie version 2
           and made it the default.
        """
        self.require_setting("cookie_secret", "secure cookies")
        secret = self.application.settings["cookie_secret"]
        key_version = None
        if isinstance(secret, dict):
            if self.application.settings.get("key_version") is None:
                raise Exception("key_version setting must be used for secret_key dicts")
            key_version = self.application.settings["key_version"]

        return create_signed_value(
            secret, name, value, version=version, key_version=key_version
        )

    def get_secure_cookie(
        self,
        name: str,
        value: Optional[str] = None,
        max_age_days: float = 31,
        min_version: Optional[int] = None,
    ) -> Optional[bytes]:
        """Returns the given signed cookie if it validates, or None.

        The decoded cookie value is returned as a byte string (unlike
        `get_cookie`).

        Similar to `get_cookie`, this method only returns cookies that
        were present in the request. It does not see outgoing cookies set by
        `set_secure_cookie` in this handler.

        This method should return a `tornado.locale.Locale` object,
        most likely obtained via a call like ``tornado.locale.get("en")``
        """
        return None

    def get_browser_locale(self, default: str = "en_US") -> tornado.locale.Locale:
        """Determines the user's locale from ``Accept-Language`` header.

        See http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.4
        """
        if "Accept-Language" in self.request.headers:
            languages = self.request.headers["Accept-Language"].split(",")
            locales = []
            for language in languages:
                parts = language.strip().split(";")
                if len(parts) > 1 and parts[1].startswith("q="):
                    try:
                        score = float(parts[1][2:])
                    except (ValueError, TypeError):
                        score = 0.0
                else:
                    score = 1.0
                locales.append((parts[0], score))
            if locales:
                locales.sort(key=lambda pair: pair[1], reverse=True)
                codes = [loc[0] for loc in locales]
                return locale.get(*codes)
        return locale.get(default)

    @property
    def current_user(self) -> Any:
        """The authenticated user for this request.

        This is set in one of two ways:

        * A subclass may override `get_current_user()`, which will be called
          automatically the first time ``self.current_user`` is accessed.
          `get_current_user()` will only be called once per request,
          and is cached for future access::

              def get_current_user(self):
                  user_cookie = self.get_secure_cookie("user")
                  if user_cookie:
                      return json.loads(user_cookie)
                  return None

        * It may be set as a normal variable, typically from an overridden
          `prepare()`::

              @gen.coroutine
              def prepare(self):
                  user_id_cookie = self.get_secure_cookie("user_id")
                  if user_id_cookie:
                      self.current_user = yield load_user(user_id_cookie)

        Note that `prepare()` may be a coroutine while `get_current_user()`
        may not, so the latter form is necessary if loading the user requires
        asynchronous operations.

        The user object may be any type of the application's choosing.
        """
        if not hasattr(self, "_current_user"):
            self._current_user = self.get_current_user()
        return self._current_user

    @current_user.setter
    def current_user(self, value: Any) -> None:
        self._current_user = value

    def get_current_user(self) -> Any:
        """Override to determine the current user from, e.g., a cookie.

        This method may not be a coroutine.
        """
        return None

    def get_login_url(self) -> str:
        """Override to customize the login URL based on the request.

        return self._xsrf_token

    def _get_raw_xsrf_token(self) -> Tuple[Optional[int], bytes, float]:
        """Read or generate the xsrf token in its raw form.

        The raw_xsrf_token is a tuple containing:

        * version: the version of the cookie from which this token was read,
          or None if we generated a new token in this request.
        * token: the raw token data; random (non-ascii) bytes.
        * timestamp: the time this token was generated (will not be accurate
          for version 1 cookies)
        """
        if not hasattr(self, "_raw_xsrf_token"):
            cookie = self.get_cookie("_xsrf")
            if cookie:
                version, token, timestamp = self._decode_xsrf_token(cookie)
            else:
                version, token, timestamp = None, None, None
            if token is None:
                version = None
                token = os.urandom(16)
                timestamp = time.time()
            assert token is not None
            assert timestamp is not None
            self._raw_xsrf_token = (version, token, timestamp)
        return self._raw_xsrf_token

    def _decode_xsrf_token(
        self, cookie: str
    ) -> Tuple[Optional[int], Optional[bytes], Optional[float]]:
        """Convert a cookie string into a the tuple form returned by
        _get_raw_xsrf_token.
        """

        try:
            m = _signed_value_version_re.match(utf8(cookie))

            if m:
                version = int(m.group(1))
                if version == 2:
                    _, mask_str, masked_token, timestamp_str = cookie.split("|")

                    mask = binascii.a2b_hex(utf8(mask_str))
                    token = _websocket_mask(mask, binascii.a2b_hex(utf8(masked_token)))
                    timestamp = int(timestamp_str)
                    return version, token, timestamp
                else:
                    # Treat unknown versions as not present instead of failing.
                    raise Exception("Unknown xsrf cookie version")
            else:
                version = 1
                try:
                    token = binascii.a2b_hex(utf8(cookie))
                except (binascii.Error, TypeError):
                    token = utf8(cookie)
                # We don't have a usable timestamp in older versions.
                timestamp = int(time.time())
                return (version, token, timestamp)
        except Exception:
            # Catch exceptions and return nothing instead of failing.
            gen_log.debug("Uncaught exception in _decode_xsrf_token", exc_info=True)
            return None, None, None

    def check_xsrf_cookie(self) -> None:
        """Verifies that the ``_xsrf`` cookie matches the ``_xsrf`` argument.

        To prevent cross-site request forgery, we set an ``_xsrf``
        cookie and include the same value as a non-cookie
        field with all ``POST`` requests. If the two do not match, we
        reject the form submission as a potential forgery.

        The ``_xsrf`` value may be set as either a form field named ``_xsrf``
        or in a custom HTTP header named ``X-XSRFToken`` or ``X-CSRFToken``
        (the latter is accepted for compatibility with Django).

        See http://en.wikipedia.org/wiki/Cross-site_request_forgery

        .. versionchanged:: 3.2.2
           Added support for cookie version 2.  Both versions 1 and 2 are
           supported.
    per instantiation of the template, so they must not depend on
    any arguments to the template.
    """

    def __init__(self, handler: RequestHandler) -> None:
        super().__init__(handler)
        # keep resources in both a list and a dict to preserve order
        self._resource_list = []  # type: List[Dict[str, Any]]
        self._resource_dict = {}  # type: Dict[str, Dict[str, Any]]

    def render(self, path: str, **kwargs: Any) -> bytes:  # type: ignore
        def set_resources(**kwargs) -> str:  # type: ignore
            if path not in self._resource_dict:
                self._resource_list.append(kwargs)
                self._resource_dict[path] = kwargs
            else:
                if self._resource_dict[path] != kwargs:
                    raise ValueError(
                        "set_resources called with different "
                        "resources for the same template"
                    )
            return ""

        return self.render_string(path, set_resources=set_resources, **kwargs)

    def _get_resources(self, key: str) -> Iterable[str]:
        return (r[key] for r in self._resource_list if key in r)

    def embedded_javascript(self) -> str:
        return "\n".join(self._get_resources("embedded_javascript"))

    def javascript_files(self) -> Iterable[str]:
        result = []
        for f in self._get_resources("javascript_files"):
            if isinstance(f, (unicode_type, bytes)):
                result.append(f)
            else:
                result.extend(f)
        return result

    def embedded_css(self) -> str:
        return "\n".join(self._get_resources("embedded_css"))

    def css_files(self) -> Iterable[str]:
        result = []
        for f in self._get_resources("css_files"):
            if isinstance(f, (unicode_type, bytes)):
                result.append(f)
            else:
                result.extend(f)
        return result

    def html_head(self) -> str:
        return "".join(self._get_resources("html_head"))

    def html_body(self) -> str:
        return "".join(self._get_resources("html_body"))


class _UIModuleNamespace(object):
    """Lazy namespace which creates UIModule proxies bound to a handler."""

    def __init__(
        self, handler: RequestHandler, ui_modules: Dict[str, Type[UIModule]]
    ) -> None:
        self.handler = handler
        self.ui_modules = ui_modules

    def __getitem__(self, key: str) -> Callable[..., str]:
        return self.handler._ui_module(key, self.ui_modules[key])

    def __getattr__(self, key: str) -> Callable[..., str]:
        try:
            return self[key]
        except KeyError as e:
            raise AttributeError(str(e))


def create_signed_value(
    secret: _CookieSecretTypes,
    name: str,
    value: Union[str, bytes],
    version: Optional[int] = None,
    clock: Optional[Callable[[], float]] = None,
    key_version: Optional[int] = None,
) -> bytes:
    if version is None:
        version = DEFAULT_SIGNED_VALUE_VERSION
    if clock is None:
        clock = time.time

    timestamp = utf8(str(int(clock())))
    value = base64.b64encode(utf8(value))
    if version == 1:
        assert not isinstance(secret, dict)
        signature = _create_signature_v1(secret, name, value, timestamp)
        value = b"|".join([value, timestamp, signature])