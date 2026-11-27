        are sent over an encrypted link.

        If an ``event`` is passed in, this method will return immediately, and
        the event will be triggered once authentication succeeds or fails.  On
        success, `is_authenticated` will return ``True``.  On failure, you may
        use `get_exception` to get more detailed error information.

        Since 1.1, if no event is passed, this method will block until the
        authentication succeeds or fails.  On failure, an exception is raised.
        Otherwise, the method simply returns.

        Since 1.5, if no event is passed and ``fallback`` is ``True`` (the
        default), if the server doesn't support plain password authentication
        but does support so-called "keyboard-interactive" mode, an attempt
        will be made to authenticate using this interactive mode.  If it fails,
        the normal exception will be thrown as if the attempt had never been
        made.  This is useful for some recent Gentoo and Debian distributions,
        which turn off plain password authentication in a misguided belief
        that interactive authentication is "more secure".  (It's not.)

        If the server requires multi-step authentication (which is very rare),
        this method will return a list of auth types permissible for the next
        step.  Otherwise, in the normal case, an empty list is returned.

        :param str username: the username to authenticate as
        :param basestring password: the password to authenticate with
        :param .threading.Event event:
            an event to trigger when the authentication attempt is complete
            (whether it was successful or not)
        :param bool fallback:
            ``True`` if an attempt at an automated "interactive" password auth
            should be made if the server doesn't support normal password auth
        :return:
            list of auth types permissible for the next stage of
            authentication (normally empty)

        :raises:
            `.BadAuthenticationType` -- if password authentication isn't
            allowed by the server for this user (and no event was passed in)
        :raises:
            `.AuthenticationException` -- if the authentication failed (and no
            event was passed in)
        :raises: `.SSHException` -- if there was a network error
        """
        if (not self.active) or (not self.initial_kex_done):
            # we should never try to send the password unless we're on a secure
            # link
            raise SSHException("No existing session")
        if event is None:
            my_event = threading.Event()
        else:
        if event is not None:
            # caller wants to wait for event themselves
            return []
        return self.auth_handler.wait_for_response(my_event)

    def auth_interactive(self, username, handler, submethods=""):
        """
        Authenticate to the server interactively.  A handler is used to answer
        arbitrary questions from the server.  On many servers, this is just a
        dumb wrapper around PAM.

        This method will block until the authentication succeeds or fails,
        peroidically calling the handler asynchronously to get answers to
        authentication questions.  The handler may be called more than once
        if the server continues to ask questions.

        The handler is expected to be a callable that will handle calls of the
        form: ``handler(title, instructions, prompt_list)``.  The ``title`` is
        meant to be a dialog-window title, and the ``instructions`` are user
        instructions (both are strings).  ``prompt_list`` will be a list of
        prompts, each prompt being a tuple of ``(str, bool)``.  The string is
        the prompt and the boolean indicates whether the user text should be
        echoed.

        A sample call would thus be:
        ``handler('title', 'instructions', [('Password:', False)])``.

        The handler should return a list or tuple of answers to the server's
        questions.

        If the server requires multi-step authentication (which is very rare),
        this method will return a list of auth types permissible for the next
        step.  Otherwise, in the normal case, an empty list is returned.

        :param str username: the username to authenticate as
        :param callable handler: a handler for responding to server questions
        :param str submethods: a string list of desired submethods (optional)
        :return:
            list of auth types permissible for the next stage of
            authentication (normally empty).

        :raises: `.BadAuthenticationType` -- if public-key authentication isn't
            allowed by the server for this user
        :raises: `.AuthenticationException` -- if the authentication failed
        :raises: `.SSHException` -- if there was a network error

        .. versionadded:: 1.5
        """
        if (not self.active) or (not self.initial_kex_done):
            # we should never try to authenticate unless we're on a secure link
            raise SSHException("No existing session")