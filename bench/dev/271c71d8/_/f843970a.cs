        /// <param name="base_address">The base address to unmap.</param>
        /// <param name="flags">Flags for unmapping memory.</param>
        public void Unmap(IntPtr base_address, MemUnmapFlags flags)
        {
            Unmap(base_address, flags, true);
        }

        /// <summary>
        /// Unmap a section.
        /// </summary>
        /// <param name="base_address">The base address to unmap.</param>
        public void Unmap(IntPtr base_address)
        {
            Unmap(base_address, true);
        }

        /// <summary>
        /// Get the user SID for the process.
        /// </summary>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The user SID.</returns>
        public NtResult<Sid> GetUser(bool throw_on_error)
        {
            using (var token = OpenToken(TokenAccessRights.Query, throw_on_error))
            {
                return token.Map(t => t.User.Sid);
            }
        }

        /// <summary>
        /// Get the user SID for the process.
        /// </summary>
        /// <returns>The user SID.</returns>
        public Sid GetUser()
        {
            return GetUser(true).Result;
        }

        /// <summary>
        /// Get the integrity level for the process.
        /// </summary>
        /// <param name="throw_on_error">True to throw on error.</param>
        /// <returns>The integerity level.</returns>
        public NtResult<TokenIntegrityLevel> GetIntegrityLevel(bool throw_on_error)
        {
            using (var token = OpenToken(TokenAccessRights.Query, throw_on_error))
            {
                return token.Map(t => t.IntegrityLevel);
            }
        }
