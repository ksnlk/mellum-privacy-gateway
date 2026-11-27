        /// <param name="token">The token to set.</param>
        /// <remarks>Requires SeCreateTokenPrivilege.</remarks>
        public void SetLinkedToken(NtToken token)
        {
            Set(TokenInformationClass.TokenLinkedToken, token.Handle.DangerousGetHandle());
        }

        /// <summary>
        /// Impersonate the token.
        /// </summary>
        /// <returns>An impersonation context, dispose to revert to process token</returns>
        /// <exception cref="NtException">Thrown on error.</exception>
        public ThreadImpersonationContext Impersonate()
        {
            return NtThread.Current.Impersonate(this);
        }

        /// <summary>
        /// Impersonate the token.
        /// </summary>
        /// <param name="impersonation_level">Impersonation level for token.</param>
        /// <returns>An impersonation context, dispose to revert to process token</returns>
        /// <exception cref="NtException">Thrown on error.</exception>
        public ThreadImpersonationContext Impersonate(SecurityImpersonationLevel impersonation_level)
        {
            using (NtToken token = DuplicateToken(impersonation_level))
            {
                return NtThread.Current.Impersonate(token);
            }
        }

        /// <summary>
        /// Run a function under impersonation.
        /// </summary>
        /// <typeparam name="T">The return type.</typeparam>
        /// <param name="callback">The callback to run.</param>
        /// <returns>The return value from the callback.</returns>
        /// <exception cref="NtException">Thrown on error.</exception>
        public T RunUnderImpersonate<T>(Func<T> callback)
        {
            using (Impersonate())
            {
                return callback();
            }
        }

        /// <summary>
        /// Run an action under impersonation.
        /// </summary>
        /// <param name="callback">The callback to run.</param>
        /// <exception cref="NtException">Thrown on error.</exception>
        /// <summary>
        /// Get/set token's origin
        /// </summary>
        public Luid Origin
        {
            get => Query<Luid>(TokenInformationClass.TokenOrigin);
            set => SetOrigin(value);
        }

        /// <summary>
        /// Get token's elevation type
        /// </summary>
        public TokenElevationType ElevationType => (TokenElevationType)Query<int>(TokenInformationClass.TokenElevationType);

        /// <summary>
        /// Get whether token is elevated
        /// </summary>
        public bool Elevated => Query<int>(TokenInformationClass.TokenElevation) != 0;

        /// <summary>
        /// Get whether token has restrictions
        /// </summary>
        public bool HasRestrictions => Query<int>(TokenInformationClass.TokenHasRestrictions) != 0;

        /// <summary>
        /// Get/set token UI access flag
        /// </summary>
        public bool UIAccess
        {
            get => Query<int>(TokenInformationClass.TokenUIAccess) != 0;
            set => SetUIAccess(value);
        }

        /// <summary>
        /// Get or set whether virtualization is allowed
        /// </summary>
        public bool VirtualizationAllowed
        {
            get => Query<int>(TokenInformationClass.TokenVirtualizationAllowed) != 0;
            set => Set(TokenInformationClass.TokenVirtualizationAllowed, value ? 1 : 0);
        }

        /// <summary>
        /// Get/set whether virtualization is enabled
        /// </summary>
        public bool VirtualizationEnabled
        {
            get => Query<int>(TokenInformationClass.TokenVirtualizationEnabled) != 0;
            set => SetVirtualizationEnabled(value);
        }

            return OpenThreadToken(NtThread.Current, true, duplicate);
        }

        /// <summary>
        /// Open the current thread token
        /// </summary>
        /// <returns>The opened token, if no token return null</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtToken OpenThreadToken()
        {
            return OpenThreadToken(false);
        }

        /// <summary>
        /// Open the effective token, thread if available or process
        /// </summary>
        /// <param name="thread">The thread to open the token for</param>
        /// <param name="duplicate">True to duplicate the token before returning</param>
        /// <param name="desired_access">Desired access for token.</param>
        /// <param name="open_as_self">Open token as self.</param>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtResult<NtToken> OpenEffectiveToken(NtThread thread, bool open_as_self, bool duplicate, TokenAccessRights desired_access, bool throw_on_error)
        {
            var token = OpenThreadToken(thread, open_as_self, duplicate, desired_access, throw_on_error);
            if (!token.IsSuccess || token.Result != null)
                return token;

            var pid = thread.GetProcessId(throw_on_error);
            if (!pid.IsSuccess)
                return pid.Cast<NtToken>();

            return OpenProcessToken(pid.Result, duplicate, desired_access, throw_on_error);
        }

        /// <summary>
        /// Open the effective token, thread if available or process
        /// </summary>
        /// <param name="thread">The thread to open the token for</param>
        /// <param name="duplicate">True to duplicate the token before returning</param>
        /// <param name="desired_access">Desired access for token.</param>
        /// <param name="open_as_self">Open token as self.</param>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtToken OpenEffectiveToken(NtThread thread, bool open_as_self, bool duplicate, TokenAccessRights desired_access)
        {
            return OpenEffectiveToken(thread, open_as_self, duplicate, desired_access, true).Result;
        }

        /// <summary>
        /// Open the effective token, thread if available or process
        /// </summary>
        /// <param name="thread">The thread to open the token for</param>
        /// <param name="duplicate">True to duplicate the token before returning</param>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtResult<NtToken> OpenEffectiveToken(NtThread thread, bool duplicate, bool throw_on_error)
        {
            return OpenEffectiveToken(thread, true, duplicate, TokenAccessRights.MaximumAllowed, throw_on_error);
        }

        /// <summary>
        /// Open the effective token, thread if available or process
        /// </summary>
        /// <param name="thread">The thread to open the token for</param>
        /// <param name="duplicate">True to duplicate the token before returning</param>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtToken OpenEffectiveToken(NtThread thread, bool duplicate)
        {
            NtToken token = null;
            try
            {
                token = OpenThreadToken(thread, true, duplicate);
            }
            catch (NtException)
            {
            }

            return token ?? OpenProcessToken(thread.ProcessId, duplicate);
        }

        /// <summary>
        /// Open the current effective token, thread if available or process
        /// </summary>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtToken OpenEffectiveToken()
        {
            return OpenEffectiveToken(true).Result;
        }

        /// <summary>
        /// Open the current effective token, thread if available or process
        /// </summary>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The opened token</returns>
        /// <exception cref="NtException">Thrown if cannot open token</exception>
        public static NtResult<NtToken> OpenEffectiveToken(bool throw_on_error)
            return Create(TokenAccessRights.GenericAll, null, TokenType.Primary, new Luid(999, 0),
                DateTime.Now.AddYears(10).ToFileTimeUtc(), new UserGroup(user, GroupAttributes.Enabled | GroupAttributes.EnabledByDefault | GroupAttributes.Owner),
                groups.Select(s => new UserGroup(s, GroupAttributes.Enabled | GroupAttributes.EnabledByDefault)),
                privileges.Select(p => new TokenPrivilege(p, PrivilegeAttributes.Enabled | PrivilegeAttributes.EnabledByDefault)),
                user, user, default_dacl, "NT.NET");
        }

        /// <summary>
        /// Create a token. Needs SeCreateTokenPrivilege.
        /// </summary>
        /// <param name="user">The user for the token.</param>
        /// <returns>The token object.</returns>
        public static NtToken Create(Sid user)
        {
            return Create(user, new Sid[] { new Sid("WD") }, new TokenPrivilegeValue[] { TokenPrivilegeValue.SeDebugPrivilege });
        }

        /// <summary>
        /// Impersonate another process' token
        /// </summary>
        /// <param name="impersonation_level">The impersonation level</param>
        /// <param name="pid">Process ID of the other process</param>
        /// <returns>An impersonation context, dispose to revert to process token</returns>
        public static ThreadImpersonationContext Impersonate(int pid, SecurityImpersonationLevel impersonation_level)
        {
            using (NtToken process_token = OpenProcessToken(pid))
            {
                using (NtToken imp_token = process_token.DuplicateToken(impersonation_level))
                {
                    return imp_token.Impersonate();
                }
            }
        }

        /// <summary>
        /// Get the current user.
        /// </summary>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The current user.</returns>
        public static NtResult<UserGroup> GetCurrentUser(bool throw_on_error)
        {
            using (var token = OpenEffectiveToken(throw_on_error))
            {
                if (!token.IsSuccess)
                {
                    return token.Cast<UserGroup>();
                }
                return token.Result.GetUser(throw_on_error);
            }
        }

        /// <summary>
        /// Do a single privilege check on the effective token.
        /// </summary>
        /// <param name="privilege">The privilege to check.</param>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>True if the privilege is enabled.</returns>
        public static NtResult<bool> EffectivePrivilegeCheck(TokenPrivilegeValue privilege, bool throw_on_error)
        {
            using (var token = OpenEffectiveToken(throw_on_error))
            {
                if (!token.IsSuccess)
                {
                    return token.Cast<bool>();
                }
                return token.Result.SinglePrivilegeCheck(privilege).CreateResult();
            }
        }

        /// <summary>
        /// Do a single privilege check on the effective token.
        /// </summary>
        /// <param name="privilege">The privilege to check.</param>
        /// <returns>True if the privilege is enabled.</returns>
        public static bool EffectivePrivilegeCheck(TokenPrivilegeValue privilege)
        {
            return EffectivePrivilegeCheck(privilege, false).GetResultOrDefault(false);
        }

        #endregion

        #region Static Properties

        /// <summary>
        /// Get the current user.
        /// </summary>
        public static UserGroup CurrentUser
        {
            get
            {
                using (NtToken token = OpenEffectiveToken())
                {
                    return token.User;
                }
            }
        }

        /// <summary>
        /// Get authentication ID for LOCAL SYSTEM
        /// </summary>
        public static Luid LocalSystemAuthId { get { return new Luid(0x3e7, 0); } }
        /// <summary>
        /// Get authentication ID for LOCAL SERVICE
        /// </summary>
        public static Luid LocalServiceAuthId { get { return new Luid(0x3e5, 0); } }
        /// <summary>
        /// Get authentication ID for NETWORK SERVICE
        /// </summary>
        public static Luid NetworkServiceAuthId { get { return new Luid(0x3e4, 0); } }
        /// <summary>
        /// Get authentication ID for ANONYMOUS
        /// </summary>
        public static Luid AnonymousAuthId { get { return new Luid(0x3e6, 0); } }

        /// <summary>
        /// Get a pseudo handle to the primary token.