 */
static CURLcode
output_auth_headers(struct connectdata *conn,
                    struct auth *authstatus,
                    const char *request,
                    const char *path,
                    bool proxy)
{
  struct SessionHandle *data = conn->data;
  const char *auth=NULL;
  CURLcode result = CURLE_OK;
#ifdef USE_HTTP_NEGOTIATE
  struct negotiatedata *negdata = proxy?
    &data->state.proxyneg:&data->state.negotiate;
#endif

#ifdef CURL_DISABLE_CRYPTO_AUTH
  (void)request;
  (void)path;
#endif

#ifdef USE_HTTP_NEGOTIATE
  negdata->state = GSS_AUTHNONE;
  if((authstatus->picked == CURLAUTH_GSSNEGOTIATE) &&
     negdata->context && !GSS_ERROR(negdata->status)) {
    auth="GSS-Negotiate";
    result = Curl_output_negotiate(conn, proxy);
    if(result)
      return result;
    authstatus->done = TRUE;
    negdata->state = GSS_AUTHSENT;
  }
  else
#endif
#ifdef USE_NTLM
  if(authstatus->picked == CURLAUTH_NTLM) {
    auth="NTLM";
    result = Curl_output_ntlm(conn, proxy);
    if(result)
      return result;
  }
  else
#endif
#ifndef CURL_DISABLE_CRYPTO_AUTH
  if(authstatus->picked == CURLAUTH_DIGEST) {
    auth="Digest";
    result = Curl_output_digest(conn,
                                proxy,
                                (const unsigned char *)request,
                                (const unsigned char *)path);
    if(result)
      return result;
  }
  else
#endif
  if(authstatus->picked == CURLAUTH_BASIC) {
    /* Basic */
    if((proxy && conn->bits.proxy_user_passwd &&
       !Curl_checkheaders(data, "Proxy-authorization:")) ||
       (!proxy && conn->bits.user_passwd &&
       !Curl_checkheaders(data, "Authorization:"))) {
      auth="Basic";
  */
  if((httpcode != 401) &&
      (httpcode != 407))
    return 1;

  /*
  ** All we have left to deal with is 401 and 407
  */
  DEBUGASSERT((httpcode == 401) || (httpcode == 407));

  /*
  ** Examine the current authentication state to see if this
  ** is an error.  The idea is for this function to get
  ** called after processing all the headers in a response
  ** message.  So, if we've been to asked to authenticate a
  ** particular stage, and we've done it, we're OK.  But, if
  ** we're already completely authenticated, it's not OK to
  ** get another 401 or 407.
  **
  ** It is possible for authentication to go stale such that
  ** the client needs to reauthenticate.  Once that info is
  ** available, use it here.
  */
#if 0 /* set to 1 when debugging this functionality */
  infof(data,"%s: authstage = %d\n",__FUNCTION__,data->state.authstage);
  infof(data,"%s: authwant = 0x%08x\n",__FUNCTION__,data->state.authwant);
  infof(data,"%s: authavail = 0x%08x\n",__FUNCTION__,data->state.authavail);
  infof(data,"%s: httpcode = %d\n",__FUNCTION__,k->httpcode);
  infof(data,"%s: authdone = %d\n",__FUNCTION__,data->state.authdone);
  infof(data,"%s: newurl = %s\n",__FUNCTION__,data->req.newurl ?
        data->req.newurl : "(null)");
  infof(data,"%s: authproblem = %d\n",__FUNCTION__,data->state.authproblem);
#endif

  /*
  ** Either we're not authenticating, or we're supposed to
  ** be authenticating something else.  This is an error.
  */
  if((httpcode == 401) && !conn->bits.user_passwd)
    return TRUE;
  if((httpcode == 407) && !conn->bits.proxy_user_passwd)
    return TRUE;

  return data->state.authproblem;
}

/*
 * readmoredata() is a "fread() emulation" to provide POST and/or request
 * data. It is used when a huge POST is to be made and the entire chunk wasn't
 * sent in the first send(). This function will then be called from the
 * transfer.c loop when more data is to be sent to the peer.
 *