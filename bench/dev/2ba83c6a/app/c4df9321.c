
static OSSL_CMP_SRV_CTX *setup_srv_ctx(ENGINE *engine)
{
    OSSL_CMP_CTX *ctx; /* extra CMP (client) ctx partly used by server */
    OSSL_CMP_SRV_CTX *srv_ctx = ossl_cmp_mock_srv_new(app_get0_libctx(),
                                                      app_get0_propq());

    if (srv_ctx == NULL)
        return NULL;
    ctx = OSSL_CMP_SRV_CTX_get0_cmp_ctx(srv_ctx);

    if (opt_srv_ref == NULL) {
        if (opt_srv_cert == NULL) {
            /* opt_srv_cert should determine the sender */
            CMP_err("must give -srv_ref for server if no -srv_cert given");
            goto err;
        }
    } else {
        if (!OSSL_CMP_CTX_set1_referenceValue(ctx, (unsigned char *)opt_srv_ref,
                                              strlen(opt_srv_ref)))
            goto err;
    }

    if (opt_srv_secret != NULL) {
        int res;
        char *pass_str = get_passwd(opt_srv_secret, "PBMAC secret of server");

        if (pass_str != NULL) {
            cleanse(opt_srv_secret);
            res = OSSL_CMP_CTX_set1_secretValue(ctx, (unsigned char *)pass_str,
                                                strlen(pass_str));
            clear_free(pass_str);
            if (res == 0)
                goto err;
        }
    } else if (opt_srv_cert == NULL) {
        CMP_err("server credentials must be given if -use_mock_srv or -port is used");
        goto err;
    } else {
        CMP_warn("server will not be able to handle PBM-protected requests since -srv_secret is not given");
    }

    if (opt_srv_secret == NULL
            && ((opt_srv_cert == NULL) != (opt_srv_key == NULL))) {
        CMP_err("must give both -srv_cert and -srv_key options or neither");
        goto err;
    }
    if (opt_srv_cert != NULL) {
        X509 *srv_cert = load_cert_pwd(opt_srv_cert, opt_srv_keypass,
                                       "certificate of the server");

            goto err;
        }
        EVP_PKEY_free(pkey); /* we do not need the handle any more */
    }
    if (opt_tls_trusted != NULL) {
        /* enable and parameterize server hostname/IP address check */
        if (!truststore_set_host_etc(trust_store,
                                     opt_tls_host != NULL ? opt_tls_host : host))
            goto err;
        SSL_CTX_set_verify(ssl_ctx, SSL_VERIFY_PEER, NULL);
    }
    return ssl_ctx;
 err:
    SSL_CTX_free(ssl_ctx);
    return NULL;
}
#endif

/*
 * set up protection aspects of OSSL_CMP_CTX based on options from config
 * file/CLI while parsing options and checking their consistency.
 * Returns 1 on success, 0 on error
 */
static int setup_protection_ctx(OSSL_CMP_CTX *ctx, ENGINE *engine)
{
    if (!opt_unprotected_requests && opt_secret == NULL && opt_key == NULL) {
        CMP_err("must give -key or -secret unless -unprotected_requests is used");
        return 0;
    }

    if (opt_ref == NULL && opt_cert == NULL && opt_subject == NULL) {
        /* cert or subject should determine the sender */
        CMP_err("must give -ref if no -cert and no -subject given");
        return 0;
    }
    if (!opt_secret && ((opt_cert == NULL) != (opt_key == NULL))) {
        CMP_err("must give both -cert and -key options or neither");
        return 0;
    }
    if (opt_secret != NULL) {
        char *pass_string = get_passwd(opt_secret, "PBMAC");
        int res;

        if (pass_string != NULL) {
            cleanse(opt_secret);
            res = OSSL_CMP_CTX_set1_secretValue(ctx,
                                                (unsigned char *)pass_string,
                                                strlen(pass_string));
            clear_free(pass_string);
            if (res == 0)
                return 0;
        }
        if (opt_cert != NULL || opt_key != NULL)
            CMP_warn("-cert and -key not used for protection since -secret is given");
    }
    if (opt_ref != NULL
            && !OSSL_CMP_CTX_set1_referenceValue(ctx, (unsigned char *)opt_ref,
                                                 strlen(opt_ref)))
        return 0;

    if (opt_key != NULL) {
        EVP_PKEY *pkey = load_key_pwd(opt_key, opt_keyform, opt_keypass, engine,
                                      "private key for CMP client certificate");

        if (pkey == NULL || !OSSL_CMP_CTX_set1_pkey(ctx, pkey)) {
            EVP_PKEY_free(pkey);
            CMP_warn("ignoring -csr since certificate to be revoked is given");
    }
    if (opt_cmd == CMP_P10CR && opt_csr == NULL) {
        CMP_err("missing PKCS#10 CSR for p10cr");
        return 0;
    }

    if (opt_recipient == NULL && opt_srvcert == NULL && opt_issuer == NULL
            && opt_oldcert == NULL && opt_cert == NULL)
        CMP_warn("missing -recipient, -srvcert, -issuer, -oldcert or -cert; recipient will be set to \"NULL-DN\"");

    if (opt_cmd == CMP_P10CR || opt_cmd == CMP_RR) {
        const char *msg = "option is ignored for 'p10cr' and 'rr' commands";

        if (opt_newkeypass != NULL)
            CMP_warn1("-newkeytype %s", msg);
        if (opt_newkey != NULL)
            CMP_warn1("-newkey %s", msg);
        if (opt_days != 0)
            CMP_warn1("-days %s", msg);
        if (opt_popo != OSSL_CRMF_POPO_NONE - 1)
            CMP_warn1("-popo %s", msg);
    } else if (opt_newkey != NULL) {
        const char *file = opt_newkey;
        const int format = opt_keyform;
        const char *pass = opt_newkeypass;
        const char *desc = "new private key for cert to be enrolled";
        EVP_PKEY *pkey;
        int priv = 1;
        BIO *bio_bak = bio_err;

        bio_err = NULL; /* suppress diagnostics on first try loading key */
        pkey = load_key_pwd(file, format, pass, engine, desc);
        bio_err = bio_bak;
        if (pkey == NULL) {
            ERR_clear_error();
            desc = opt_csr == NULL
            ? "fallback public key for cert to be enrolled"
            : "public key for checking cert resulting from p10cr";
            pkey = load_pubkey(file, format, 0, pass, engine, desc);
            priv = 0;
        }
        cleanse(opt_newkeypass);
        if (pkey == NULL || !OSSL_CMP_CTX_set0_newPkey(ctx, priv, pkey)) {
            EVP_PKEY_free(pkey);
            return 0;
        }
    }

    if (opt_days > 0
            && !OSSL_CMP_CTX_set_option(ctx, OSSL_CMP_OPT_VALIDITY_DAYS,
            opt_path = opt_str();
            break;
        case OPT_RECIPIENT:
            opt_recipient = opt_str();
            break;
        case OPT_KEEP_ALIVE:
            opt_keep_alive = opt_int_arg();
            if (opt_keep_alive > 2) {
                CMP_err("-keep_alive argument must be 0, 1, or 2");
                goto opthelp;
            }
            break;
        case OPT_MSG_TIMEOUT:
            opt_msg_timeout = opt_int_arg();
            break;
        case OPT_TOTAL_TIMEOUT:
            opt_total_timeout = opt_int_arg();
            break;
        case OPT_TLS_USED:
            opt_tls_used = 1;
            break;
        case OPT_TLS_CERT:
            opt_tls_cert = opt_str();
            break;
        case OPT_TLS_KEY:
            opt_tls_key = opt_str();
            break;
        case OPT_TLS_KEYPASS:
            opt_tls_keypass = opt_str();
            break;
        case OPT_TLS_EXTRA:
            opt_tls_extra = opt_str();
            break;
        case OPT_TLS_TRUSTED:
            opt_tls_trusted = opt_str();
            break;
        case OPT_TLS_HOST:
            opt_tls_host = opt_str();
            break;
        case OPT_REF:
            opt_ref = opt_str();
            break;
        case OPT_SECRET:
            opt_secret = opt_str();
            break;
        case OPT_CERT:
            opt_cert = opt_str();
            break;
        case OPT_OWN_TRUSTED:
            opt_own_trusted = opt_str();
            break;
        case OPT_KEY:
            opt_key = opt_str();
            break;
        case OPT_KEYPASS:
            opt_keypass = opt_str();
            break;
        case OPT_DIGEST:
            opt_digest = opt_str();
            break;
        case OPT_MAC:
            opt_mac = opt_str();
            break;
        case OPT_EXTRACERTS:
            opt_extracerts = opt_str();
            break;
        case OPT_UNPROTECTED_REQUESTS:
            opt_unprotected_requests = 1;
            break;

        case OPT_TRUSTED:
            opt_trusted = opt_str();
            break;
        case OPT_UNTRUSTED:
            opt_untrusted = opt_str();
            break;
        case OPT_SRVCERT:
            opt_srvcert = opt_str();
        case OPT_RSPIN:
            opt_rspin = opt_str();
            break;
        case OPT_RSPOUT:
            opt_rspout = opt_str();
            break;
        case OPT_USE_MOCK_SRV:
            opt_use_mock_srv = 1;
            break;
        case OPT_PORT:
            opt_port = opt_str();
            break;
        case OPT_MAX_MSGS:
            opt_max_msgs = opt_int_arg();
            break;
        case OPT_SRV_REF:
            opt_srv_ref = opt_str();
            break;
        case OPT_SRV_SECRET:
            opt_srv_secret = opt_str();
            break;
        case OPT_SRV_CERT:
            opt_srv_cert = opt_str();
            break;
        case OPT_SRV_KEY:
            opt_srv_key = opt_str();
            break;
        case OPT_SRV_KEYPASS:
            opt_srv_keypass = opt_str();
            break;
        case OPT_SRV_TRUSTED:
            opt_srv_trusted = opt_str();
            break;
        case OPT_SRV_UNTRUSTED:
            opt_srv_untrusted = opt_str();
            break;
        case OPT_RSP_CERT:
            opt_rsp_cert = opt_str();
            break;
        case OPT_RSP_EXTRACERTS:
            opt_rsp_extracerts = opt_str();
            break;
        case OPT_RSP_CAPUBS:
            opt_rsp_capubs = opt_str();
            break;
        case OPT_POLL_COUNT:
            opt_poll_count = opt_int_arg();
            break;
        case OPT_CHECK_AFTER:
            opt_check_after = opt_int_arg();
            break;