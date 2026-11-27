    (void)BIO_flush(bio_s_out);
    return 0;
}
#endif

static int psk_find_session_cb(SSL *ssl, const unsigned char *identity,
                               size_t identity_len, SSL_SESSION **sess)
{
    SSL_SESSION *tmpsess = NULL;
    unsigned char *key;
    long key_len;
    const SSL_CIPHER *cipher = NULL;

    if (strlen(psk_identity) != identity_len
            || memcmp(psk_identity, identity, identity_len) != 0) {
        *sess = NULL;
        return 1;
    }

    if (psksess != NULL) {
        SSL_SESSION_up_ref(psksess);
        *sess = psksess;
        return 1;
    }

    key = OPENSSL_hexstr2buf(psk_key, &key_len);
    if (key == NULL) {
        BIO_printf(bio_err, "Could not convert PSK key '%s' to buffer\n",
                   psk_key);
        return 0;
    }

    /* We default to SHA256 */
    cipher = SSL_CIPHER_find(ssl, tls13_aes128gcmsha256_id);
    if (cipher == NULL) {
        BIO_printf(bio_err, "Error finding suitable ciphersuite\n");
        OPENSSL_free(key);
        return 0;
    }

    tmpsess = SSL_SESSION_new();
    if (tmpsess == NULL
            || !SSL_SESSION_set1_master_key(tmpsess, key, key_len)
            || !SSL_SESSION_set_cipher(tmpsess, cipher)
            || !SSL_SESSION_set_protocol_version(tmpsess, SSL_version(ssl))) {
        OPENSSL_free(key);
        return 0;
    }
    OPENSSL_free(key);
    *sess = tmpsess;

            break;
        case OPT_CRLF:
            s_crlf = 1;
            break;
        case OPT_QUIET:
            s_quiet = 1;
            break;
        case OPT_BRIEF:
            s_quiet = s_brief = verify_args.quiet = 1;
            break;
        case OPT_NO_DHE:
            no_dhe = 1;
            break;
        case OPT_NO_RESUME_EPHEMERAL:
            no_resume_ephemeral = 1;
            break;
        case OPT_PSK_IDENTITY:
            psk_identity = opt_arg();
            break;
        case OPT_PSK_HINT:
#ifndef OPENSSL_NO_PSK
            psk_identity_hint = opt_arg();
#endif
            break;
        case OPT_PSK:
            for (p = psk_key = opt_arg(); *p; p++) {
                if (isxdigit(_UC(*p)))
                    continue;
                BIO_printf(bio_err, "Not a hex number '%s'\n", psk_key);
                goto end;
            }
            break;
        case OPT_PSK_SESS:
            psksessf = opt_arg();
            break;
        case OPT_SRPVFILE:
#ifndef OPENSSL_NO_SRP
            srp_verifier_file = opt_arg();
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
#endif
            break;
        case OPT_SRPUSERSEED:
#ifndef OPENSSL_NO_SRP
            srpuserseed = opt_arg();
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
#endif
            break;
        case OPT_REV:
            rev = 1;
    }
#endif

#ifndef OPENSSL_NO_KTLS
    if (use_sendfile && www <= 1) {
        BIO_printf(bio_err, "Can't use -sendfile without -WWW or -HTTP\n");
        goto end;
    }
#endif

    if (!app_passwd(passarg, dpassarg, &pass, &dpass)) {
        BIO_printf(bio_err, "Error getting password\n");
        goto end;
    }

    if (s_key_file == NULL)
        s_key_file = s_cert_file;

    if (s_key_file2 == NULL)
        s_key_file2 = s_cert_file2;

    if (!load_excert(&exc))
        goto end;

    if (nocert == 0) {
        s_key = load_key(s_key_file, s_key_format, 0, pass, engine,
                         "server certificate private key");
        if (s_key == NULL)
            goto end;

        s_cert = load_cert_pass(s_cert_file, s_cert_format, 1, pass,
                                "server certificate");

        if (s_cert == NULL)
            goto end;
        if (s_chain_file != NULL) {
            if (!load_certs(s_chain_file, 0, &s_chain, NULL,
                            "server certificate chain"))
                goto end;
        }

        if (tlsextcbp.servername != NULL) {
            s_key2 = load_key(s_key_file2, s_key_format, 0, pass, engine,
                              "second server certificate private key");
            if (s_key2 == NULL)
                goto end;

            s_cert2 = load_cert_pass(s_cert_file2, s_cert_format, 1, pass,
                                "second server certificate");

            if (s_cert2 == NULL)
                }
            }
            if (dhpkey == NULL) {
                SSL_CTX_set_dh_auto(ctx2, 1);
            } else if (!SSL_CTX_set0_tmp_dh_pkey(ctx2, dhpkey)) {
                BIO_puts(bio_err, "Error setting temp DH parameters\n");
                ERR_print_errors(bio_err);
                EVP_PKEY_free(dhpkey);
                goto end;
            }
            dhpkey = NULL;
        }
        EVP_PKEY_free(dhpkey);
    }

    if (!set_cert_key_stuff(ctx, s_cert, s_key, s_chain, build_chain))
        goto end;

    if (s_serverinfo_file != NULL
        && !SSL_CTX_use_serverinfo_file(ctx, s_serverinfo_file)) {
        ERR_print_errors(bio_err);
        goto end;
    }

    if (ctx2 != NULL
        && !set_cert_key_stuff(ctx2, s_cert2, s_key2, NULL, build_chain))
        goto end;

    if (s_dcert != NULL) {
        if (!set_cert_key_stuff(ctx, s_dcert, s_dkey, s_dchain, build_chain))
            goto end;
    }

    if (no_resume_ephemeral) {
        SSL_CTX_set_not_resumable_session_callback(ctx,
                                                   not_resumable_sess_cb);

        if (ctx2 != NULL)
            SSL_CTX_set_not_resumable_session_callback(ctx2,
                                                       not_resumable_sess_cb);
    }
#ifndef OPENSSL_NO_PSK
    if (psk_key != NULL) {
        if (s_debug)
            BIO_printf(bio_s_out, "PSK key given, setting server callback\n");
        SSL_CTX_set_psk_server_callback(ctx, psk_server_cb);
    }

    if (psk_identity_hint != NULL) {
        if (min_version == TLS1_3_VERSION) {
            BIO_printf(bio_s_out, "PSK warning: there is NO identity hint in TLSv1.3\n");