        public void SetEa(string name, string data, EaBufferEntryFlags flags)
        {
            SetEa(name, Encoding.Unicode.GetBytes(data), flags);
        }

        /// <summary>
        /// Remove an extended attributes entry for a file.
        /// </summary>
        /// <param name="name">The name of the entry</param>
        public void RemoveEa(string name)
        {
            EaBuffer ea = new EaBuffer();
            ea.AddEntry(name, new byte[0], EaBufferEntryFlags.None);
            SetEa(ea);
        }

        /// <summary>
        /// Assign completion port to file.
        /// </summary>
        /// <param name="completion_port">The completion port.</param>
        /// <param name="key">A key to associate with this completion.</param>
        public void SetCompletionPort(NtIoCompletion completion_port, IntPtr key)
        {
            FileCompletionInformation info = new FileCompletionInformation();
            info.CompletionPort = completion_port.Handle.DangerousGetHandle();
            info.Key = key;

            Set(FileInformationClass.FileCompletionInformation, info);
        }

        /// <summary>
        /// Check if a specific set of file directory access rights is granted
        /// </summary>
        /// <param name="access">The file directory access rights to check</param>
        /// <returns>True if all access rights are granted</returns>
        public bool IsAccessGranted(FileDirectoryAccessRights access)
        {
            return IsAccessMaskGranted(access);
        }

        /// <summary>
        /// Get the cached signing level for a file.
        /// </summary>
        /// <returns>The cached signing level.</returns>
        public NtResult<CachedSigningLevel> GetCachedSigningLevel(bool throw_on_error)
        {
            return NtSecurity.GetCachedSigningLevel(Handle, throw_on_error);
        }

        /// <summary>
        /// Get the cached signing level for a file.