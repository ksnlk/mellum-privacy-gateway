            print_message(names[D_WHIRLPOOL], c[D_WHIRLPOOL][testnum],
                          lengths[testnum], seconds.sym);
            Time_F(START);
            count = run_benchmark(async_jobs, WHIRLPOOL_loop, loopargs);
            d = Time_F(STOP);
            print_result(D_WHIRLPOOL, testnum, count, d);
            if (count < 0)
                break;
        }
    }

    if (doit[D_RMD160]) {
        for (testnum = 0; testnum < size_num; testnum++) {
            print_message(names[D_RMD160], c[D_RMD160][testnum],
                          lengths[testnum], seconds.sym);
            Time_F(START);
            count = run_benchmark(async_jobs, EVP_Digest_RMD160_loop, loopargs);
            d = Time_F(STOP);
            print_result(D_RMD160, testnum, count, d);
            if (count < 0)
                break;
        }
    }

    if (doit[D_HMAC]) {
        static const char hmac_key[] = "This is a key...";
        int len = strlen(hmac_key);
        EVP_MAC *mac = EVP_MAC_fetch(NULL, "HMAC", NULL);
        OSSL_PARAM params[3];

        if (mac == NULL || evp_mac_mdname == NULL)
            goto end;

        evp_hmac_name = app_malloc(sizeof("hmac()") + strlen(evp_mac_mdname),
                                   "HMAC name");
        sprintf(evp_hmac_name, "hmac(%s)", evp_mac_mdname);
        names[D_HMAC] = evp_hmac_name;

        params[0] =
            OSSL_PARAM_construct_utf8_string(OSSL_MAC_PARAM_DIGEST,
                                             evp_mac_mdname, 0);
        params[1] =
            OSSL_PARAM_construct_octet_string(OSSL_MAC_PARAM_KEY,
                                              (char *)hmac_key, len);
        params[2] = OSSL_PARAM_construct_end();

        for (i = 0; i < loopargs_len; i++) {
            loopargs[i].mctx = EVP_MAC_CTX_new(mac);
            if (loopargs[i].mctx == NULL)
                goto end;

                    lengths = aead_lengths_list;
                    size_num = OSSL_NELEM(aead_lengths_list);
                }
            }

            for (testnum = 0; testnum < size_num; testnum++) {
                print_message(names[D_EVP], c[D_EVP][testnum], lengths[testnum],
                              seconds.sym);

                for (k = 0; k < loopargs_len; k++) {
                    loopargs[k].ctx = EVP_CIPHER_CTX_new();
                    if (loopargs[k].ctx == NULL) {
                        BIO_printf(bio_err, "\nEVP_CIPHER_CTX_new failure\n");
                        exit(1);
                    }
                    if (!EVP_CipherInit_ex(loopargs[k].ctx, evp_cipher, NULL,
                                           NULL, iv, decrypt ? 0 : 1)) {
                        BIO_printf(bio_err, "\nEVP_CipherInit_ex failure\n");
                        ERR_print_errors(bio_err);
                        exit(1);
                    }

                    EVP_CIPHER_CTX_set_padding(loopargs[k].ctx, 0);

                    keylen = EVP_CIPHER_CTX_get_key_length(loopargs[k].ctx);
                    loopargs[k].key = app_malloc(keylen, "evp_cipher key");
                    EVP_CIPHER_CTX_rand_key(loopargs[k].ctx, loopargs[k].key);
                    if (!EVP_CipherInit_ex(loopargs[k].ctx, NULL, NULL,
                                           loopargs[k].key, NULL, -1)) {
                        BIO_printf(bio_err, "\nEVP_CipherInit_ex failure\n");
                        ERR_print_errors(bio_err);
                        exit(1);
                    }
                    OPENSSL_clear_free(loopargs[k].key, keylen);

                    /* SIV mode only allows for a single Update operation */
                    if (EVP_CIPHER_get_mode(evp_cipher) == EVP_CIPH_SIV_MODE)
                        EVP_CIPHER_CTX_ctrl(loopargs[k].ctx, EVP_CTRL_SET_SPEED,
                                            1, NULL);
                }

                Time_F(START);
                count = run_benchmark(async_jobs, loopfunc, loopargs);
                d = Time_F(STOP);
                for (k = 0; k < loopargs_len; k++)
                    EVP_CIPHER_CTX_free(loopargs[k].ctx);
                print_result(D_EVP, testnum, count, d);
            }
        } else if (evp_md_name != NULL) {
            names[D_EVP] = evp_md_name;


    for (testnum = 0; testnum < RSA_NUM; testnum++) {
        EVP_PKEY *rsa_key = NULL;
        int st = 0;

        if (!rsa_doit[testnum])
            continue;

        if (primes > RSA_DEFAULT_PRIME_NUM) {
            /* we haven't set keys yet,  generate multi-prime RSA keys */
            bn = BN_new();
            st = bn != NULL
                && BN_set_word(bn, RSA_F4)
                && init_gen_str(&genctx, "RSA", NULL, 0, NULL, NULL)
                && EVP_PKEY_CTX_set_rsa_keygen_bits(genctx, rsa_keys[testnum].bits) > 0
                && EVP_PKEY_CTX_set1_rsa_keygen_pubexp(genctx, bn) > 0
                && EVP_PKEY_CTX_set_rsa_keygen_primes(genctx, primes) > 0
                && EVP_PKEY_keygen(genctx, &rsa_key);
            BN_free(bn);
            bn = NULL;
            EVP_PKEY_CTX_free(genctx);
            genctx = NULL;
        } else {
            const unsigned char *p = rsa_keys[testnum].data;

            st = (rsa_key = d2i_PrivateKey(EVP_PKEY_RSA, NULL, &p,
                                           rsa_keys[testnum].length)) != NULL;
        }

        for (i = 0; st && i < loopargs_len; i++) {
            loopargs[i].rsa_sign_ctx[testnum] = EVP_PKEY_CTX_new(rsa_key, NULL);
            if (loopargs[i].rsa_sign_ctx[testnum] == NULL
                || EVP_PKEY_sign_init(loopargs[i].rsa_sign_ctx[testnum]) <= 0
                || EVP_PKEY_sign(loopargs[i].rsa_sign_ctx[testnum],
                                 loopargs[i].buf2,
                                 &loopargs[i].sigsize,
                                 loopargs[i].buf, 36) <= 0)
                st = 0;
        }
        if (!st) {
            BIO_printf(bio_err,
                       "RSA sign setup failure.  No RSA sign will be done.\n");
            ERR_print_errors(bio_err);
            op_count = 1;
        } else {
            pkey_print_message("private", "rsa",
                               rsa_c[testnum][0], rsa_keys[testnum].bits,
                               seconds.rsa);
            /* RSA_blinding_on(rsa_key[testnum],NULL); */
            Time_F(START);
            count = run_benchmark(async_jobs, RSA_sign_loop, loopargs);
    BIO_printf(bio_err,
               mr ? "+DTP:%d:%s:%s:%d\n"
               : "Doing %u bits %s %s's for %ds: ", bits, str, str2, tm);
    (void)BIO_flush(bio_err);
    run = 1;
    alarm(tm);
}

static void print_result(int alg, int run_no, int count, double time_used)
{
    if (count == -1) {
        BIO_printf(bio_err, "%s error!\n", names[alg]);
        ERR_print_errors(bio_err);
        return;
    }
    BIO_printf(bio_err,
               mr ? "+R:%d:%s:%f\n"
               : "%d %s's in %.2fs\n", count, names[alg], time_used);
    results[alg][run_no] = ((double)count) / time_used * lengths[run_no];
}

#ifndef NO_FORK
static char *sstrsep(char **string, const char *delim)
{
    char isdelim[256];
    char *token = *string;

    if (**string == 0)
        return NULL;

    memset(isdelim, 0, sizeof(isdelim));
    isdelim[0] = 1;

    while (*delim) {
        isdelim[(unsigned char)(*delim)] = 1;
        delim++;
    }

    while (!isdelim[(unsigned char)(**string)])
        (*string)++;

    if (**string) {
        **string = 0;
        (*string)++;
    }

    return token;
}

static int do_multi(int multi, int size_num)
{
    static const int mblengths_list[] =
        { 8 * 1024, 2 * 8 * 1024, 4 * 8 * 1024, 8 * 8 * 1024, 8 * 16 * 1024 };
    const int *mblengths = mblengths_list;
    int j, count, keylen, num = OSSL_NELEM(mblengths_list);
    const char *alg_name;
    unsigned char *inp = NULL, *out = NULL, *key, no_key[32], no_iv[16];
    EVP_CIPHER_CTX *ctx = NULL;
    double d = 0.0;

    if (lengths_single) {
        mblengths = &lengths_single;
        num = 1;
    }

    inp = app_malloc(mblengths[num - 1], "multiblock input buffer");
    out = app_malloc(mblengths[num - 1] + 1024, "multiblock output buffer");
    if ((ctx = EVP_CIPHER_CTX_new()) == NULL)
        app_bail_out("failed to allocate cipher context\n");
    if (!EVP_EncryptInit_ex(ctx, evp_cipher, NULL, NULL, no_iv))
        app_bail_out("failed to initialise cipher context\n");

    if ((keylen = EVP_CIPHER_CTX_get_key_length(ctx)) < 0) {
        BIO_printf(bio_err, "Impossible negative key length: %d\n", keylen);
        goto err;
    }
    key = app_malloc(keylen, "evp_cipher key");
    if (!EVP_CIPHER_CTX_rand_key(ctx, key))
        app_bail_out("failed to generate random cipher key\n");
    if (!EVP_EncryptInit_ex(ctx, NULL, NULL, key, NULL))
        app_bail_out("failed to set cipher key\n");
    OPENSSL_clear_free(key, keylen);

    if (!EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_AEAD_SET_MAC_KEY,
                             sizeof(no_key), no_key))
        app_bail_out("failed to set AEAD key\n");
    if ((alg_name = EVP_CIPHER_get0_name(evp_cipher)) == NULL)
        app_bail_out("failed to get cipher name\n");

    for (j = 0; j < num; j++) {
        print_message(alg_name, 0, mblengths[j], seconds->sym);
        Time_F(START);
        for (count = 0; run && count < 0x7fffffff; count++) {
            unsigned char aad[EVP_AEAD_TLS1_AAD_LEN];
            EVP_CTRL_TLS1_1_MULTIBLOCK_PARAM mb_param;
            size_t len = mblengths[j];
            int packlen;

            memset(aad, 0, 8);  /* avoid uninitialized values */
            aad[8] = 23;        /* SSL3_RT_APPLICATION_DATA */
            aad[9] = 3;         /* version */
            aad[10] = 2;