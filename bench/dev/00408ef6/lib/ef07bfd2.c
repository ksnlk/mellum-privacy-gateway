    */
    PPSENDF(&conn->proto.ftpc.pp, "PBSZ %d", 0);
    state(conn, FTP_PBSZ);
  }
  else {
    result = ftp_state_pwd(conn);
  }
  return result;
}

/* for USER and PASS responses */
static CURLcode ftp_state_user_resp(struct connectdata *conn,
                                    int ftpcode,
                                    ftpstate instate)
{
  CURLcode result = CURLE_OK;
  struct SessionHandle *data = conn->data;
  struct FTP *ftp = data->state.proto.ftp;
  struct ftp_conn *ftpc = &conn->proto.ftpc;
  (void)instate; /* no use for this yet */

  /* some need password anyway, and others just return 2xx ignored */
  if((ftpcode == 331) && (ftpc->state == FTP_USER)) {
    /* 331 Password required for ...
       (the server requires to send the user's password too) */
    PPSENDF(&ftpc->pp, "PASS %s", ftp->passwd?ftp->passwd:"");
    state(conn, FTP_PASS);
  }
  else if(ftpcode/100 == 2) {
    /* 230 User ... logged in.
       (the user logged in with or without password) */
    result = ftp_state_loggedin(conn);
  }
  else if(ftpcode == 332) {
    if(data->set.str[STRING_FTP_ACCOUNT]) {
      PPSENDF(&ftpc->pp, "ACCT %s", data->set.str[STRING_FTP_ACCOUNT]);
      state(conn, FTP_ACCT);
    }
    else {
      failf(data, "ACCT requested but none available");
      result = CURLE_LOGIN_DENIED;
    }
  }
  else {
    /* All other response codes, like:

    530 User ... access denied
    (the server denies to log the specified user) */

    if(conn->data->set.str[STRING_FTP_ALTERNATIVE_TO_USER] &&
        !conn->data->state.ftp_trying_alternative) {
#pragma optimize("", off)
#endif

static CURLcode ftp_init(struct connectdata *conn)
{
  struct FTP *ftp;

  if(NULL == conn->data->state.proto.ftp) {
    conn->data->state.proto.ftp = malloc(sizeof(struct FTP));
    if(NULL == conn->data->state.proto.ftp)
      return CURLE_OUT_OF_MEMORY;
  }

  ftp = conn->data->state.proto.ftp;

  /* get some initial data into the ftp struct */
  ftp->bytecountp = &conn->data->req.bytecount;
  ftp->transfer = FTPTRANSFER_BODY;
  ftp->downloadsize = 0;

  /* No need to duplicate user+password, the connectdata struct won't change
     during a session, but we re-init them here since on subsequent inits
     since the conn struct may have changed or been replaced.
  */
  ftp->user = conn->user;
  ftp->passwd = conn->passwd;
  if(TRUE == isBadFtpString(ftp->user))
    return CURLE_URL_MALFORMAT;
  if(TRUE == isBadFtpString(ftp->passwd))
    return CURLE_URL_MALFORMAT;

  conn->proto.ftpc.known_filesize = -1; /* unknown size for now */

  return CURLE_OK;
}

#if defined(__INTEL_COMPILER) && (__INTEL_COMPILER == 910) && \
    defined(__OPTIMIZE__) && defined(__unix__) && defined(__i386__)
  /* workaround icc 9.1 optimizer issue */
#pragma optimize("", on)
#endif

/*
 * ftp_connect() should do everything that is to be considered a part of
 * the connection phase.
 *
 * The variable 'done' points to will be TRUE if the protocol-layer connect
 * phase is done when this function returns, or FALSE is not. When called as
 * a part of the easy interface, it will always be TRUE.
 */
static CURLcode ftp_connect(struct connectdata *conn,