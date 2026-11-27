#ifdef HAVE_LIBSSH2_KNOWNHOST_API
      if(data->set.str[STRING_SSH_KNOWNHOSTS]) {
        /* we're asked to verify the host against a file */
        int keytype;
        size_t keylen;
        const char *remotekey = libssh2_session_hostkey(sshc->ssh_session,
                                                        &keylen, &keytype);
        int keycheck;
        int keybit;

        if(remotekey) {
          /*
           * A subject to figure out is what host name we need to pass in here.
           * What host name does OpenSSH store in its file if an IDN name is
           * used?
           */
          struct libssh2_knownhost *host;
          enum curl_khmatch keymatch;
          curl_sshkeycallback func =
            data->set.ssh_keyfunc?data->set.ssh_keyfunc:sshkeycallback;
          struct curl_khkey knownkey;
          struct curl_khkey *knownkeyp = NULL;
          struct curl_khkey foundkey;

          keybit = (keytype == LIBSSH2_HOSTKEY_TYPE_RSA)?
            LIBSSH2_KNOWNHOST_KEY_SSHRSA:LIBSSH2_KNOWNHOST_KEY_SSHDSS;

          keycheck = libssh2_knownhost_check(sshc->kh,
                                             conn->host.name,
                                             remotekey, keylen,
                                             LIBSSH2_KNOWNHOST_TYPE_PLAIN|
                                             LIBSSH2_KNOWNHOST_KEYENC_RAW|
                                             keybit,
                                             &host);

          infof(data, "SSH host check: %d, key: %s\n", keycheck,
                (keycheck <= LIBSSH2_KNOWNHOST_CHECK_MISMATCH)?
                host->key:"<none>");

          /* setup 'knownkey' */
          if(keycheck <= LIBSSH2_KNOWNHOST_CHECK_MISMATCH) {
            knownkey.key = host->key;
            knownkey.len = 0;
            knownkey.keytype = (keytype == LIBSSH2_HOSTKEY_TYPE_RSA)?
              CURLKHTYPE_RSA : CURLKHTYPE_DSS;
            knownkeyp = &knownkey;
          }

          /* setup 'foundkey' */
          foundkey.key = remotekey;
          foundkey.len = keylen;
          foundkey.keytype = (keytype == LIBSSH2_HOSTKEY_TYPE_RSA)?
            CURLKHTYPE_RSA : CURLKHTYPE_DSS;

          /*
           * if any of the LIBSSH2_KNOWNHOST_CHECK_* defines and the
           * curl_khmatch enum are ever modified, we need to introduce a
           * translation table here!
           */
          keymatch = (enum curl_khmatch)keycheck;

          /* Ask the callback how to behave */
          rc = func(data, knownkeyp, /* from the knownhosts file */
                    &foundkey, /* from the remote host */
                    keymatch, data->set.ssh_keyfunc_userp);
        }
        else
          /* no remotekey means failure! */
          rc = CURLKHSTAT_REJECT;

        switch(rc) {
        default: /* unknown return codes will equal reject */
        case CURLKHSTAT_REJECT:
          state(conn, SSH_SESSION_FREE);
        case CURLKHSTAT_DEFER:
     function to make the re-use checks properly be able to check this bit. */
  conn->bits.close = FALSE;

  /* If there already is a protocol-specific struct allocated for this
     sessionhandle, deal with it */
  Curl_reset_reqproto(conn);

  result = ssh_init(conn);
  if(result)
    return result;

  if(conn->handler->protocol & CURLPROTO_SCP) {
    conn->recv[FIRSTSOCKET] = scp_recv;
    conn->send[FIRSTSOCKET] = scp_send;
  } else {
    conn->recv[FIRSTSOCKET] = sftp_recv;
    conn->send[FIRSTSOCKET] = sftp_send;
  }
  ssh = &conn->proto.sshc;

#ifdef CURL_LIBSSH2_DEBUG
  if(conn->user) {
    infof(data, "User: %s\n", conn->user);
  }
  if(conn->passwd) {
    infof(data, "Password: %s\n", conn->passwd);
  }
  sock = conn->sock[FIRSTSOCKET];
#endif /* CURL_LIBSSH2_DEBUG */

  ssh->ssh_session = libssh2_session_init_ex(my_libssh2_malloc,
                                             my_libssh2_free,
                                             my_libssh2_realloc, conn);
  if(ssh->ssh_session == NULL) {
    failf(data, "Failure initialising ssh session");
    return CURLE_FAILED_INIT;
  }

#ifdef HAVE_LIBSSH2_KNOWNHOST_API
  if(data->set.str[STRING_SSH_KNOWNHOSTS]) {
    int rc;
    ssh->kh = libssh2_knownhost_init(ssh->ssh_session);
    if(!ssh->kh) {
      /* eeek. TODO: free the ssh_session! */
      return CURLE_FAILED_INIT;
    }

    /* read all known hosts from there */
    rc = libssh2_knownhost_readfile(ssh->kh,
                                    data->set.str[STRING_SSH_KNOWNHOSTS],
                                    LIBSSH2_KNOWNHOST_FILE_OPENSSH);