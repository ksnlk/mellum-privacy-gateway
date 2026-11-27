    long key_len;
    unsigned char *key;

    if (c_debug)
        BIO_printf(bio_c_out, "psk_client_cb\n");
    if (!hint) {
        /* no ServerKeyExchange message */
        if (c_debug)
            BIO_printf(bio_c_out,
                       "NULL received PSK identity hint, continuing anyway\n");
    } else if (c_debug) {
        BIO_printf(bio_c_out, "Received PSK identity hint '%s'\n", hint);
    }

    /*
     * lookup PSK identity and PSK key based on the given identity hint here
     */
    ret = BIO_snprintf(identity, max_identity_len, "%s", psk_identity);
    if (ret < 0 || (unsigned int)ret > max_identity_len)
        goto out_err;
    if (c_debug)
        BIO_printf(bio_c_out, "created identity '%s' len=%d\n", identity,
                   ret);

    /* convert the PSK key to binary */
    key = OPENSSL_hexstr2buf(psk_key, &key_len);
    if (key == NULL) {
        BIO_printf(bio_err, "Could not convert PSK key '%s' to buffer\n",
                   psk_key);
        return 0;
    }
    if (max_psk_len > INT_MAX || key_len > (long)max_psk_len) {
        BIO_printf(bio_err,
                   "psk buffer of callback is too small (%d) for key (%ld)\n",
                   max_psk_len, key_len);
        OPENSSL_free(key);
        return 0;
    }

    memcpy(psk, key, key_len);
    OPENSSL_free(key);

    if (c_debug)
        BIO_printf(bio_c_out, "created PSK len=%ld\n", key_len);

    return key_len;
 out_err:
    if (c_debug)
        BIO_printf(bio_err, "Error in PSK client callback\n");
    return 0;
}
#endif

const unsigned char tls13_aes128gcmsha256_id[] = { 0x13, 0x01 };
const unsigned char tls13_aes256gcmsha384_id[] = { 0x13, 0x02 };

static int psk_use_session_cb(SSL *s, const EVP_MD *md,
                              const unsigned char **id, size_t *idlen,
                              SSL_SESSION **sess)
{
    SSL_SESSION *usesess = NULL;
    const SSL_CIPHER *cipher = NULL;

    if (psksess != NULL) {
        SSL_SESSION_up_ref(psksess);
        usesess = psksess;
    } else {
        long key_len;
        unsigned char *key = OPENSSL_hexstr2buf(psk_key, &key_len);

        if (key == NULL) {
            BIO_printf(bio_err, "Could not convert PSK key '%s' to buffer\n",
                       psk_key);
            return 0;
        }

        /* We default to SHA-256 */
        cipher = SSL_CIPHER_find(s, tls13_aes128gcmsha256_id);
        if (cipher == NULL) {
            BIO_printf(bio_err, "Error finding suitable ciphersuite\n");
            OPENSSL_free(key);
            return 0;
        }

        usesess = SSL_SESSION_new();
        if (usesess == NULL
                || !SSL_SESSION_set1_master_key(usesess, key, key_len)
                || !SSL_SESSION_set_cipher(usesess, cipher)
                || !SSL_SESSION_set_protocol_version(usesess, TLS1_3_VERSION)) {
            OPENSSL_free(key);
            goto err;
        }
        OPENSSL_free(key);
    }
#ifdef AF_UNIX
    {"unix", OPT_UNIX, 's', "Connect over the specified Unix-domain socket"},
#endif
    {"4", OPT_4, '-', "Use IPv4 only"},
#ifdef AF_INET6
    {"6", OPT_6, '-', "Use IPv6 only"},
#endif
    {"maxfraglen", OPT_MAXFRAGLEN, 'p',
     "Enable Maximum Fragment Length Negotiation (len values: 512, 1024, 2048 and 4096)"},
    {"max_send_frag", OPT_MAX_SEND_FRAG, 'p', "Maximum Size of send frames "},
    {"split_send_frag", OPT_SPLIT_SEND_FRAG, 'p',
     "Size used to split data for encrypt pipelines"},
    {"max_pipelines", OPT_MAX_PIPELINES, 'p',
     "Maximum number of encrypt/decrypt pipelines to be used"},
    {"read_buf", OPT_READ_BUF, 'p',
     "Default read buffer size to be used for connections"},
    {"fallback_scsv", OPT_FALLBACKSCSV, '-', "Send the fallback SCSV"},

    OPT_SECTION("Identity"),
    {"cert", OPT_CERT, '<', "Client certificate file to use"},
    {"certform", OPT_CERTFORM, 'F',
     "Client certificate file format (PEM/DER/P12); has no effect"},
    {"cert_chain", OPT_CERT_CHAIN, '<',
     "Client certificate chain file (in PEM format)"},
    {"build_chain", OPT_BUILD_CHAIN, '-', "Build client certificate chain"},
    {"key", OPT_KEY, 's', "Private key file to use; default: -cert file"},
    {"keyform", OPT_KEYFORM, 'E', "Key format (ENGINE, other values ignored)"},
    {"pass", OPT_PASS, 's', "Private key and cert file pass phrase source"},
    {"verify", OPT_VERIFY, 'p', "Turn on peer certificate verification"},
    {"nameopt", OPT_NAMEOPT, 's', "Certificate subject/issuer name printing options"},
    {"CApath", OPT_CAPATH, '/', "PEM format directory of CA's"},
    {"CAfile", OPT_CAFILE, '<', "PEM format file of CA's"},
    {"CAstore", OPT_CASTORE, ':', "URI to store of CA's"},
    {"no-CAfile", OPT_NOCAFILE, '-',
     "Do not load the default certificates file"},
    {"no-CApath", OPT_NOCAPATH, '-',
     "Do not load certificates from the default certificates directory"},
    {"no-CAstore", OPT_NOCASTORE, '-',
     "Do not load certificates from the default certificates store"},
    {"requestCAfile", OPT_REQCAFILE, '<',
      "PEM format file of CA names to send to the server"},
    {"dane_tlsa_domain", OPT_DANE_TLSA_DOMAIN, 's', "DANE TLSA base domain"},
    {"dane_tlsa_rrdata", OPT_DANE_TLSA_RRDATA, 's',
     "DANE TLSA rrdata presentation form"},
    {"dane_ee_no_namechecks", OPT_DANE_EE_NO_NAME, '-',
     "Disable name checks when matching DANE-EE(3) TLSA records"},
    {"psk_identity", OPT_PSK_IDENTITY, 's', "PSK identity"},
    {"psk", OPT_PSK, 's', "PSK in hex (without 0x)"},
    {"psk_session", OPT_PSK_SESS, '<', "File to read PSK SSL session from"},
    {"name", OPT_PROTOHOST, 's',
     "Hostname to use for \"-starttls lmtp\", \"-starttls smtp\" or \"-starttls xmpp[-server]\""},
            break;
        case OPT_TRACE:
#ifndef OPENSSL_NO_SSL_TRACE
            c_msg = 2;
#endif
            break;
        case OPT_SECURITY_DEBUG:
            sdebug = 1;
            break;
        case OPT_SECURITY_DEBUG_VERBOSE:
            sdebug = 2;
            break;
        case OPT_SHOWCERTS:
            c_showcerts = 1;
            break;
        case OPT_NBIO_TEST:
            nbio_test = 1;
            break;
        case OPT_STATE:
            state = 1;
            break;
        case OPT_PSK_IDENTITY:
            psk_identity = opt_arg();
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
#ifndef OPENSSL_NO_SRP
        case OPT_SRPUSER:
            srp_arg.srplogin = opt_arg();
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
            break;
        case OPT_SRPPASS:
            srppass = opt_arg();
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
            break;
        case OPT_SRP_STRENGTH:
            srp_arg.strength = atoi(opt_arg());
            BIO_printf(bio_err, "SRP minimal length for N is %d\n",
                       srp_arg.strength);
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
            break;
        case OPT_SRP_LATEUSER:
            srp_lateuser = 1;
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
            break;
        case OPT_SRP_MOREGROUPS:
            srp_arg.amp = 1;
            if (min_version < TLS1_VERSION)
                min_version = TLS1_VERSION;
            break;
#endif
        case OPT_SSL_CONFIG:
            ssl_config = opt_arg();
            break;
        case OPT_SSL3:
            goto end;
        }
    } else
        next_proto.data = NULL;
#endif

    if (!app_passwd(passarg, NULL, &pass, NULL)) {
        BIO_printf(bio_err, "Error getting private key password\n");
        goto end;
    }

    if (!app_passwd(proxypassarg, NULL, &proxypass, NULL)) {
        BIO_printf(bio_err, "Error getting proxy password\n");
        goto end;
    }

    if (proxypass != NULL && proxyuser == NULL) {
        BIO_printf(bio_err, "Error: Must specify proxy_user with proxy_pass\n");
        goto end;
    }

    if (key_file == NULL)
        key_file = cert_file;

    if (key_file != NULL) {
        key = load_key(key_file, key_format, 0, pass, e,
                       "client certificate private key");
        if (key == NULL)
            goto end;
    }

    if (cert_file != NULL) {
        cert = load_cert_pass(cert_file, cert_format, 1, pass,
                              "client certificate");
        if (cert == NULL)
            goto end;
    }

    if (chain_file != NULL) {
        if (!load_certs(chain_file, 0, &chain, pass, "client certificate chain"))
            goto end;
    }

    if (crl_file != NULL) {
        X509_CRL *crl;
        crl = load_crl(crl_file, crl_format, 0, "CRL");
        if (crl == NULL)
            goto end;
        crls = sk_X509_CRL_new_null();
        if (crls == NULL || !sk_X509_CRL_push(crls, crl)) {
            BIO_puts(bio_err, "Error adding CRL\n");
            if (read_ssl)
                openssl_fdset(SSL_get_fd(con), &readfds);
            if (write_ssl)
                openssl_fdset(SSL_get_fd(con), &writefds);
#else
            if (!tty_on || !write_tty) {
                if (read_ssl)
                    openssl_fdset(SSL_get_fd(con), &readfds);
                if (write_ssl)
                    openssl_fdset(SSL_get_fd(con), &writefds);
            }
#endif

            /*
             * Note: under VMS with SOCKETSHR the second parameter is
             * currently of type (int *) whereas under other systems it is
             * (void *) if you don't have a cast it will choke the compiler:
             * if you do have a cast then you can either go for (int *) or
             * (void *).
             */
#if defined(OPENSSL_SYS_WINDOWS) || defined(OPENSSL_SYS_MSDOS)
            /*
             * Under Windows/DOS we make the assumption that we can always
             * write to the tty: therefore if we need to write to the tty we
             * just fall through. Otherwise we timeout the select every
             * second and see if there are any keypresses. Note: this is a
             * hack, in a proper Windows application we wouldn't do this.
             */
            i = 0;
            if (!write_tty) {
                if (read_tty) {
                    tv.tv_sec = 1;
                    tv.tv_usec = 0;
                    i = select(width, (void *)&readfds, (void *)&writefds,
                               NULL, &tv);
                    if (!i && (!has_stdin_waiting() || !read_tty))
                        continue;
                } else
                    i = select(width, (void *)&readfds, (void *)&writefds,
                               NULL, timeoutp);
            }
#else
            i = select(width, (void *)&readfds, (void *)&writefds,
                       NULL, timeoutp);
#endif
            if (i < 0) {
                BIO_printf(bio_err, "bad select %d\n",
                           get_last_socket_error());
                goto shut;
            }
        }
    int i, istls13 = (SSL_version(s) == TLS1_3_VERSION);
    long verify_result;
#ifndef OPENSSL_NO_COMP
    const COMP_METHOD *comp, *expansion;
#endif
    unsigned char *exportedkeymat;
#ifndef OPENSSL_NO_CT
    const SSL_CTX *ctx = SSL_get_SSL_CTX(s);
#endif

    if (full) {
        int got_a_chain = 0;

        sk = SSL_get_peer_cert_chain(s);
        if (sk != NULL) {
            got_a_chain = 1;

            BIO_printf(bio, "---\nCertificate chain\n");
            for (i = 0; i < sk_X509_num(sk); i++) {
                BIO_printf(bio, "%2d s:", i);
                X509_NAME_print_ex(bio, X509_get_subject_name(sk_X509_value(sk, i)), 0, get_nameopt());
                BIO_puts(bio, "\n");
                BIO_printf(bio, "   i:");
                X509_NAME_print_ex(bio, X509_get_issuer_name(sk_X509_value(sk, i)), 0, get_nameopt());
                BIO_puts(bio, "\n");
                public_key = X509_get_pubkey(sk_X509_value(sk, i));
                if (public_key != NULL) {
                    BIO_printf(bio, "   a:PKEY: %s, %d (bit); sigalg: %s\n",
                               OBJ_nid2sn(EVP_PKEY_get_base_id(public_key)),
                               EVP_PKEY_get_bits(public_key),
                               OBJ_nid2sn(X509_get_signature_nid(sk_X509_value(sk, i))));
                    EVP_PKEY_free(public_key);
                }
                BIO_printf(bio, "   v:NotBefore: ");
                ASN1_TIME_print(bio, X509_get0_notBefore(sk_X509_value(sk, i)));
                BIO_printf(bio, "; NotAfter: ");
                ASN1_TIME_print(bio, X509_get0_notAfter(sk_X509_value(sk, i)));
                BIO_puts(bio, "\n");
                if (c_showcerts)
                    PEM_write_bio_X509(bio, sk_X509_value(sk, i));
            }
        }

        BIO_printf(bio, "---\n");
        peer = SSL_get0_peer_certificate(s);
        if (peer != NULL) {
            BIO_printf(bio, "Server certificate\n");

            /* Redundant if we showed the whole chain */
            if (!(c_showcerts && got_a_chain))
                PEM_write_bio_X509(bio, peer);