    SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_INTERNAL_ERROR);
    return 0;
#endif
}

static int tls_process_ske_dhe(SSL *s, PACKET *pkt, EVP_PKEY **pkey)
{
    PACKET prime, generator, pub_key;
    EVP_PKEY *peer_tmp = NULL;
    BIGNUM *p = NULL, *g = NULL, *bnpub_key = NULL;
    EVP_PKEY_CTX *pctx = NULL;
    OSSL_PARAM *params = NULL;
    OSSL_PARAM_BLD *tmpl = NULL;
    int ret = 0;

    if (!PACKET_get_length_prefixed_2(pkt, &prime)
        || !PACKET_get_length_prefixed_2(pkt, &generator)
        || !PACKET_get_length_prefixed_2(pkt, &pub_key)) {
        SSLfatal(s, SSL_AD_DECODE_ERROR, SSL_R_LENGTH_MISMATCH);
        return 0;
    }

    p = BN_bin2bn(PACKET_data(&prime), (int)PACKET_remaining(&prime), NULL);
    g = BN_bin2bn(PACKET_data(&generator), (int)PACKET_remaining(&generator),
                  NULL);
    bnpub_key = BN_bin2bn(PACKET_data(&pub_key),
                          (int)PACKET_remaining(&pub_key), NULL);
    if (p == NULL || g == NULL || bnpub_key == NULL) {
        SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_BN_LIB);
        goto err;
    }

    tmpl = OSSL_PARAM_BLD_new();
    if (tmpl == NULL
            || !OSSL_PARAM_BLD_push_BN(tmpl, OSSL_PKEY_PARAM_FFC_P, p)
            || !OSSL_PARAM_BLD_push_BN(tmpl, OSSL_PKEY_PARAM_FFC_G, g)
            || !OSSL_PARAM_BLD_push_BN(tmpl, OSSL_PKEY_PARAM_PUB_KEY,
                                       bnpub_key)
            || (params = OSSL_PARAM_BLD_to_param(tmpl)) == NULL) {
        SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_INTERNAL_ERROR);
        goto err;
    }

    pctx = EVP_PKEY_CTX_new_from_name(s->ctx->libctx, "DH", s->ctx->propq);
    if (pctx == NULL) {
        SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_INTERNAL_ERROR);
        goto err;
    }
    if (EVP_PKEY_fromdata_init(pctx) <= 0
            || EVP_PKEY_fromdata(pctx, &peer_tmp, EVP_PKEY_KEYPAIR, params) <= 0) {
        SSLfatal(s, SSL_AD_INTERNAL_ERROR, SSL_R_BAD_DH_VALUE);
    if (sha256 == NULL) {
        /* Error is already recorded */
        SSLfatal_alert(s, SSL_AD_INTERNAL_ERROR);
        goto err;
    }
    /*
     * We use sess_len here because EVP_Digest expects an int
     * but s->session->session_id_length is a size_t
     */
    if (!EVP_Digest(s->session->ext.tick, ticklen,
                    s->session->session_id, &sess_len,
                    sha256, NULL)) {
        SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_EVP_LIB);
        goto err;
    }
    EVP_MD_free(sha256);
    sha256 = NULL;
    s->session->session_id_length = sess_len;
    s->session->not_resumable = 0;

    /* This is a standalone message in TLSv1.3, so there is no more to read */
    if (SSL_IS_TLS13(s)) {
        const EVP_MD *md = ssl_handshake_md(s);
        int hashleni = EVP_MD_get_size(md);
        size_t hashlen;
        static const unsigned char nonce_label[] = "resumption";

        /* Ensure cast to size_t is safe */
        if (!ossl_assert(hashleni >= 0)) {
            SSLfatal(s, SSL_AD_INTERNAL_ERROR, ERR_R_INTERNAL_ERROR);
            goto err;
        }
        hashlen = (size_t)hashleni;

        if (!tls13_hkdf_expand(s, md, s->resumption_master_secret,
                               nonce_label,
                               sizeof(nonce_label) - 1,
                               PACKET_data(&nonce),
                               PACKET_remaining(&nonce),
                               s->session->master_key,
                               hashlen, 1)) {
            /* SSLfatal() already called */
            goto err;
        }
        s->session->master_key_length = hashlen;

        OPENSSL_free(exts);
        ssl_update_cache(s, SSL_SESS_CACHE_CLIENT);
        return MSG_PROCESS_FINISHED_READING;
    }
