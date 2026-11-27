    Curl_safefree(data->set.str[i]);
}

static CURLcode setstropt(char **charp, char * s)
{
  /* Release the previous storage at `charp' and replace by a dynamic storage
     copy of `s'. Return CURLE_OK or CURLE_OUT_OF_MEMORY. */

  if(*charp) {
    free(*charp);
    *charp = (char *) NULL;
  }

  if(s) {
    s = strdup(s);

    if(!s)
      return CURLE_OUT_OF_MEMORY;

    *charp = s;
  }

  return CURLE_OK;
}

static CURLcode setstropt_userpwd(char *option, char **user_storage,
                                  char **pwd_storage)
{
  char* separator;
  CURLcode result = CURLE_OK;

  if(!option) {
    /* we treat a NULL passed in as a hint to clear existing info */
    Curl_safefree(*user_storage);
    *user_storage = (char *) NULL;
    Curl_safefree(*pwd_storage);
    *pwd_storage = (char *) NULL;
    return CURLE_OK;
  }

  separator = strchr(option, ':');
  if (separator != NULL) {

    /* store username part of option */
    char * p;
    size_t username_len = (size_t)(separator-option);
    p = malloc(username_len+1);
    if(!p)
      result = CURLE_OUT_OF_MEMORY;
    else {
      memcpy(p, option, username_len);
     * The maximum time you allow curl to use for a single transfer
     * operation.
     */
    data->set.timeout = va_arg(param, long) * 1000L;
    break;

  case CURLOPT_TIMEOUT_MS:
    data->set.timeout = va_arg(param, long);
    break;

  case CURLOPT_CONNECTTIMEOUT:
    /*
     * The maximum time you allow curl to use to connect.
     */
    data->set.connecttimeout = va_arg(param, long) * 1000L;
    break;

  case CURLOPT_CONNECTTIMEOUT_MS:
    data->set.connecttimeout = va_arg(param, long);
    break;

  case CURLOPT_USERPWD:
    /*
     * user:password to use in the operation
     */
    result = setstropt_userpwd(va_arg(param, char *),
                               &data->set.str[STRING_USERNAME],
                               &data->set.str[STRING_PASSWORD]);
    break;
  case CURLOPT_USERNAME:
    /*
     * authentication user name to use in the operation
     */
    result = setstropt(&data->set.str[STRING_USERNAME],
                       va_arg(param, char *));
    break;
  case CURLOPT_PASSWORD:
    /*
     * authentication password to use in the operation
     */
    result = setstropt(&data->set.str[STRING_PASSWORD],
                       va_arg(param, char *));
    break;
  case CURLOPT_POSTQUOTE:
    /*
     * List of RAW FTP commands to use after a transfer
     */
    data->set.postquote = va_arg(param, struct curl_slist *);
    break;
  case CURLOPT_PREQUOTE:
    /*
    data->change.resolve = data->set.resolve;
    break;
  case CURLOPT_PROGRESSFUNCTION:
    /*
     * Progress callback function
     */
    data->set.fprogress = va_arg(param, curl_progress_callback);
    if(data->set.fprogress)
      data->progress.callback = TRUE; /* no longer internal */
    else
      data->progress.callback = FALSE; /* NULL enforces internal */

    break;
  case CURLOPT_PROGRESSDATA:
    /*
     * Custom client data to pass to the progress callback
     */
    data->set.progress_client = va_arg(param, void *);
    break;

#ifndef CURL_DISABLE_PROXY
  case CURLOPT_PROXYUSERPWD:
    /*
     * user:password needed to use the proxy
     */
    result = setstropt_userpwd(va_arg(param, char *),
                               &data->set.str[STRING_PROXYUSERNAME],
                               &data->set.str[STRING_PROXYPASSWORD]);
    break;
  case CURLOPT_PROXYUSERNAME:
    /*
     * authentication user name to use in the operation
     */
    result = setstropt(&data->set.str[STRING_PROXYUSERNAME],
                       va_arg(param, char *));
    break;
  case CURLOPT_PROXYPASSWORD:
    /*
     * authentication password to use in the operation
     */
    result = setstropt(&data->set.str[STRING_PROXYPASSWORD],
                       va_arg(param, char *));
    break;
  case CURLOPT_NOPROXY:
    /*
     * proxy exception list
     */
    result = setstropt(&data->set.str[STRING_NOPROXY],
                       va_arg(param, char *));
    break;
#endif

  /* and we setup a few fields in case we end up actually using this struct */

  conn->sock[FIRSTSOCKET] = CURL_SOCKET_BAD;     /* no file descriptor */
  conn->sock[SECONDARYSOCKET] = CURL_SOCKET_BAD; /* no file descriptor */
  conn->connectindex = -1;    /* no index */
  conn->port = -1; /* unknown at this point */

  /* Default protocol-independent behavior doesn't support persistent
     connections, so we set this to force-close. Protocols that support
     this need to set this to FALSE in their "curl_do" functions. */
  conn->bits.close = TRUE;

  /* Store creation time to help future close decision making */
  conn->created = Curl_tvnow();

  conn->data = data; /* Setup the association between this connection
                        and the SessionHandle */

  conn->proxytype = data->set.proxytype; /* type */

#ifdef CURL_DISABLE_PROXY

  conn->bits.proxy = FALSE;
  conn->bits.httpproxy = FALSE;
  conn->bits.proxy_user_passwd = FALSE;
  conn->bits.tunnel_proxy = FALSE;

#else /* CURL_DISABLE_PROXY */

  conn->bits.proxy = (bool)(data->set.str[STRING_PROXY] &&
                            *data->set.str[STRING_PROXY]);
  conn->bits.httpproxy = (bool)(conn->bits.proxy &&
                                (conn->proxytype == CURLPROXY_HTTP ||
                                 conn->proxytype == CURLPROXY_HTTP_1_0));
  conn->bits.proxy_user_passwd =
    (bool)(NULL != data->set.str[STRING_PROXYUSERNAME]);
  conn->bits.tunnel_proxy = data->set.tunnel_thru_httpproxy;

#endif /* CURL_DISABLE_PROXY */

  conn->bits.user_passwd = (bool)(NULL != data->set.str[STRING_USERNAME]);
  conn->bits.ftp_use_epsv = data->set.ftp_use_epsv;
  conn->bits.ftp_use_eprt = data->set.ftp_use_eprt;

  conn->verifypeer = data->set.ssl.verifypeer;
  conn->verifyhost = data->set.ssl.verifyhost;

  conn->ip_version = data->set.ipver;

  if(data->multi && Curl_multi_canPipeline(data->multi) &&
      !conn->master_buffer) {
    /* Allocate master_buffer to be used for pipelining */
    conn->master_buffer = calloc(BUFSIZE, sizeof (char));
    if(!conn->master_buffer)
      goto error;
  }

  /* Initialize the pipeline lists */
  conn->send_pipe = Curl_llist_alloc((curl_llist_dtor) llist_dtor);
  conn->recv_pipe = Curl_llist_alloc((curl_llist_dtor) llist_dtor);
  conn->pend_pipe = Curl_llist_alloc((curl_llist_dtor) llist_dtor);
  conn->done_pipe = Curl_llist_alloc((curl_llist_dtor) llist_dtor);
  if(!conn->send_pipe || !conn->recv_pipe || !conn->pend_pipe ||
     !conn->done_pipe)
    goto error;


  /* Is there a username and password given in this proxy url? */
  atsign = strchr(proxyptr, '@');
  if(atsign) {
    char proxyuser[MAX_CURL_USER_LENGTH];
    char proxypasswd[MAX_CURL_PASSWORD_LENGTH];
    proxypasswd[0] = 0;

    if(1 <= sscanf(proxyptr,
                   "%" MAX_CURL_USER_LENGTH_TXT"[^:@]:"
                   "%" MAX_CURL_PASSWORD_LENGTH_TXT "[^@]",
                   proxyuser, proxypasswd)) {
      CURLcode res = CURLE_OK;

      /* found user and password, rip them out.  note that we are
         unescaping them, as there is otherwise no way to have a
         username or password with reserved characters like ':' in
         them. */
      Curl_safefree(conn->proxyuser);
      conn->proxyuser = curl_easy_unescape(data, proxyuser, 0, NULL);

      if(!conn->proxyuser)
        res = CURLE_OUT_OF_MEMORY;
      else {
        Curl_safefree(conn->proxypasswd);
        conn->proxypasswd = curl_easy_unescape(data, proxypasswd, 0, NULL);

        if(!conn->proxypasswd)
          res = CURLE_OUT_OF_MEMORY;
      }

      if(CURLE_OK == res) {
        conn->bits.proxy_user_passwd = TRUE; /* enable it */
        atsign = strdup(atsign+1); /* the right side of the @-letter */

        if(atsign) {
          free(proxy); /* free the former proxy string */
          proxy = proxyptr = atsign; /* now use this instead */
        }
        else
          res = CURLE_OUT_OF_MEMORY;
      }

      if(res) {
        free(proxy); /* free the allocated proxy string */
        return res;
      }
    }
  }

  /* start scanning for port number at this point */
  portptr = proxyptr;

  /* detect and extract RFC2732-style IPv6-addresses */
  if(*proxyptr == '[') {
    char *ptr = ++proxyptr; /* advance beyond the initial bracket */
    while(*ptr && (ISXDIGIT(*ptr) || (*ptr == ':') || (*ptr == '%') ||
                   (*ptr == '.')))
/*
 *
 * Parse a user name and password in the URL and strip it out of the host name
 *
 * Inputs: data->set.use_netrc (CURLOPT_NETRC)
 *         conn->host.name
 *
 * Outputs: (almost :- all currently undefined)
 *          conn->bits.user_passwd  - non-zero if non-default passwords exist
 *          user                    - non-zero length if defined
 *          passwd                  -   ditto
 *          conn->host.name         - remove user name and password
 */
static CURLcode parse_url_userpass(struct SessionHandle *data,
                                   struct connectdata *conn,
                                   char *user, char *passwd)
{
  /* At this point, we're hoping all the other special cases have
   * been taken care of, so conn->host.name is at most
   *    [user[:password]]@]hostname
   *
   * We need somewhere to put the embedded details, so do that first.
   */

  char *ptr=strchr(conn->host.name, '@');
  char *userpass = conn->host.name;

  user[0] =0;   /* to make everything well-defined */
  passwd[0]=0;

  /* We will now try to extract the
   * possible user+password pair in a string like:
   * ftp://user:ptsxmjow@ftp.my.site:8021/README */
  if(ptr != NULL) {
    /* there's a user+password given here, to the left of the @ */

    conn->host.name = ++ptr;

    /* So the hostname is sane.  Only bother interpreting the
     * results if we could care.  It could still be wasted
     * work because it might be overtaken by the programmatically
     * set user/passwd, but doing that first adds more cases here :-(
     */

    conn->bits.userpwd_in_url = TRUE;
    if(data->set.use_netrc != CURL_NETRC_REQUIRED) {
      /* We could use the one in the URL */

      conn->bits.user_passwd = TRUE; /* enable user+password */

      if(*userpass != ':') {
        /* the name is given, get user+password */
        sscanf(userpass, "%" MAX_CURL_USER_LENGTH_TXT "[^:@]:"
               "%" MAX_CURL_PASSWORD_LENGTH_TXT "[^@]",
               user, passwd);
      }
      else
        /* no name given, get the password only */
        sscanf(userpass, ":%" MAX_CURL_PASSWORD_LENGTH_TXT "[^@]", passwd);

      if(user[0]) {
        char *newname=curl_easy_unescape(data, user, 0, NULL);
        if(!newname)
          return CURLE_OUT_OF_MEMORY;
        if(strlen(newname) < MAX_CURL_USER_LENGTH)
          strcpy(user, newname);

        /* if the new name is longer than accepted, then just use
           the unconverted name, it'll be wrong but what the heck */
        free(newname);
      }
      if(passwd[0]) {
        /* we have a password found in the URL, decode it! */
        char *newpasswd=curl_easy_unescape(data, passwd, 0, NULL);
        if(!newpasswd)
          return CURLE_OUT_OF_MEMORY;
        if(strlen(newpasswd) < MAX_CURL_PASSWORD_LENGTH)
          strcpy(passwd, newpasswd);

        free(newpasswd);
      }
    }
  }
  return CURLE_OK;
}

/*************************************************************
 * Figure out the remote port number and fix it in the URL
 *
 * No matter if we use a proxy or not, we have to figure out the remote
 * port number of various reasons.
 *
 * To be able to detect port number flawlessly, we must not confuse them
 * IPv6-specified addresses in the [0::1] style. (RFC2732)
 *
 * The conn->host.name is currently [user:passwd@]host[:port] where host
 * could be a hostname, IPv4 address or IPv6 address.
 *
 * The port number embedded in the URL is replaced, if necessary.
 *************************************************************/
static CURLcode parse_remote_port(struct SessionHandle *data,
                                  struct connectdata *conn)
{
  char *portptr;
  char endbracket;

  /* Note that at this point, the IPv6 address cannot contain any scope
     suffix as that has already been removed in the parseurlandfillconn()
     function */
  if((1 == sscanf(conn->host.name, "[%*45[0123456789abcdefABCDEF:.]%c",
                  &endbracket)) &&
     (']' == endbracket)) {
    /* this is a RFC2732-style specified IP-address */
    conn->bits.ipv6_ip = TRUE;

    conn->host.name++; /* skip over the starting bracket */
    portptr = strchr(conn->host.name, ']');
    if(portptr) {
      *portptr++ = '\0'; /* zero terminate, killing the bracket */
      if(':' != *portptr)
        portptr = NULL; /* no port number available */
      *portptr = '\0'; /* cut off the name there */
      conn->remote_port = curlx_ultous(port);
    }
    else if(!port)
      /* Browser behavior adaptation. If there's a colon with no digits after,
         just cut off the name there which makes us ignore the colon and just
         use the default port. Firefox and Chrome both do that. */
      *portptr = '\0';
  }
  return CURLE_OK;
}

/*
 * Override a user name and password from the URL with that in the
 * CURLOPT_USERPWD option or a .netrc file, if applicable.
 */
static void override_userpass(struct SessionHandle *data,
                              struct connectdata *conn,
                              char *user, char *passwd)
{
  if(data->set.str[STRING_USERNAME] != NULL) {
    strncpy(user, data->set.str[STRING_USERNAME], MAX_CURL_USER_LENGTH);
    user[MAX_CURL_USER_LENGTH-1] = '\0';   /*To be on safe side*/
  }
  if(data->set.str[STRING_PASSWORD] != NULL) {
    strncpy(passwd, data->set.str[STRING_PASSWORD], MAX_CURL_PASSWORD_LENGTH);
    passwd[MAX_CURL_PASSWORD_LENGTH-1] = '\0'; /*To be on safe side*/
  }

  conn->bits.netrc = FALSE;
  if(data->set.use_netrc != CURL_NETRC_IGNORED) {
    if(Curl_parsenetrc(conn->host.name,
                       user, passwd,
                       data->set.str[STRING_NETRC_FILE])) {
      infof(data, "Couldn't find host %s in the "
            DOT_CHAR "netrc file; using defaults\n",
            conn->host.name);
    }
    else {
      /* set bits.netrc TRUE to remember that we got the name from a .netrc
         file, so that it is safe to use even if we followed a Location: to a
         different host or similar. */
      conn->bits.netrc = TRUE;

      conn->bits.user_passwd = TRUE; /* enable user+password */
    }
  }
}

/*
 * Set password so it's available in the connection.
 */
static CURLcode set_userpass(struct connectdata *conn,
                             const char *user, const char *passwd)
{
  /* If our protocol needs a password and we have none, use the defaults */
  if( (conn->handler->protocol & (CURLPROTO_FTP|CURLPROTO_IMAP)) &&
       !conn->bits.user_passwd) {

    conn->user = strdup(CURL_DEFAULT_USER);
    if(conn->user)
      conn->passwd = strdup(CURL_DEFAULT_PASSWORD);
    else
      conn->passwd = NULL;
    /* This is the default password, so DON'T set conn->bits.user_passwd */
  }
  else {
    /* store user + password, zero-length if not set */
    conn->user = strdup(user);
    if(conn->user)
      conn->passwd = strdup(passwd);
    else
      conn->passwd = NULL;
  }
  if(!conn->user || !conn->passwd)
    return CURLE_OUT_OF_MEMORY;

  return CURLE_OK;
}

/*************************************************************
 * Resolve the address of the server or proxy
 *************************************************************/
static CURLcode resolve_server(struct SessionHandle *data,
                               struct connectdata *conn,
                               bool *async)
{
    conn->dns_entry = hostaddr;
  }

  return result;
}

/*
 * Cleanup the connection just allocated before we can move along and use the
 * previously existing one.  All relevant data is copied over and old_conn is
 * ready for freeing once this function returns.
 */
static void reuse_conn(struct connectdata *old_conn,
                       struct connectdata *conn)
{
  if(old_conn->proxy.rawalloc)
    free(old_conn->proxy.rawalloc);

  /* free the SSL config struct from this connection struct as this was
     allocated in vain and is targeted for destruction */
  Curl_free_ssl_config(&old_conn->ssl_config);

  conn->data = old_conn->data;

  /* get the user+password information from the old_conn struct since it may
   * be new for this request even when we re-use an existing connection */
  conn->bits.user_passwd = old_conn->bits.user_passwd;
  if(conn->bits.user_passwd) {
    /* use the new user name and password though */
    Curl_safefree(conn->user);
    Curl_safefree(conn->passwd);
    conn->user = old_conn->user;
    conn->passwd = old_conn->passwd;
    old_conn->user = NULL;
    old_conn->passwd = NULL;
  }

  conn->bits.proxy_user_passwd = old_conn->bits.proxy_user_passwd;
  if(conn->bits.proxy_user_passwd) {
    /* use the new proxy user name and proxy password though */
    Curl_safefree(conn->proxyuser);
    Curl_safefree(conn->proxypasswd);
    conn->proxyuser = old_conn->proxyuser;
    conn->proxypasswd = old_conn->proxypasswd;
    old_conn->proxyuser = NULL;
    old_conn->proxypasswd = NULL;
  }

  /* host can change, when doing keepalive with a proxy ! */
  if(conn->bits.proxy) {
    free(conn->host.rawalloc);
    conn->host=old_conn->host;
  }
  else
    free(old_conn->host.rawalloc); /* free the newly allocated name buffer */

  /* persist connection info in session handle */
  Curl_persistconninfo(conn);

  /* re-use init */
  conn->bits.reuse = TRUE; /* yes, we're re-using here */

  Curl_safefree(old_conn->user);
  Curl_safefree(old_conn->passwd);
  Curl_safefree(old_conn->proxyuser);
  Curl_safefree(old_conn->proxypasswd);
  Curl_llist_destroy(old_conn->send_pipe, NULL);
  Curl_llist_destroy(old_conn->recv_pipe, NULL);
  Curl_llist_destroy(old_conn->pend_pipe, NULL);
    }

    return result;
  }
#endif

  /*************************************************************
   * If the protocol is using SSL and HTTP proxy is used, we set
   * the tunnel_proxy bit.
   *************************************************************/
  if((conn->given->flags&PROTOPT_SSL) && conn->bits.httpproxy)
    conn->bits.tunnel_proxy = TRUE;

  /*************************************************************
   * Figure out the remote port number and fix it in the URL
   *************************************************************/
  result = parse_remote_port(data, conn);
  if(result != CURLE_OK)
    return result;

  /*************************************************************
   * Check for an overridden user name and password, then set it
   * for use
   *************************************************************/
  override_userpass(data, conn, user, passwd);
  result = set_userpass(conn, user, passwd);
  if(result != CURLE_OK)
    return result;

  /* Get a cloned copy of the SSL config situation stored in the
     connection struct. But to get this going nicely, we must first make
     sure that the strings in the master copy are pointing to the correct
     strings in the session handle strings array!

     Keep in mind that the pointers in the master copy are pointing to strings
     that will be freed as part of the SessionHandle struct, but all cloned
     copies will be separately allocated.
  */
  data->set.ssl.CApath = data->set.str[STRING_SSL_CAPATH];
  data->set.ssl.CAfile = data->set.str[STRING_SSL_CAFILE];
  data->set.ssl.CRLfile = data->set.str[STRING_SSL_CRLFILE];
  data->set.ssl.issuercert = data->set.str[STRING_SSL_ISSUERCERT];
  data->set.ssl.random_file = data->set.str[STRING_SSL_RANDOM_FILE];
  data->set.ssl.egdsocket = data->set.str[STRING_SSL_EGDSOCKET];
  data->set.ssl.cipher_list = data->set.str[STRING_SSL_CIPHER_LIST];
#ifdef USE_TLS_SRP
  data->set.ssl.username = data->set.str[STRING_TLSAUTH_USERNAME];
  data->set.ssl.password = data->set.str[STRING_TLSAUTH_PASSWORD];
#endif

  if(!Curl_clone_ssl_config(&data->set.ssl, &conn->ssl_config))
    return CURLE_OUT_OF_MEMORY;

  /*************************************************************
   * Check the current list of connections to see if we can
   * re-use an already existing one or if we have to create a
   * new one.
   *************************************************************/

  /* reuse_fresh is TRUE if we are told to use a new connection by force, but
     we only acknowledge this option if this is not a re-used connection
     already (which happens due to follow-location or during a HTTP
     authentication phase). */
  if(data->set.reuse_fresh && !data->state.this_is_a_follow)
    reuse = FALSE;
  else
    reuse = ConnectionExists(data, conn, &conn_temp);

  if(reuse) {
    /*
     * We already have a connection for this, we got the former connection
     * in the conn_temp variable and thus we need to cleanup the one we
     * just allocated before we can move along and use the previously