    if (expected->type == EVP_PKEY_CMAC) {
#ifdef OPENSSL_NO_DEPRECATED_3_0
        TEST_info("skipping, PKEY CMAC '%s' is disabled", expected->alg);
        t->skip = 1;
        t->err = NULL;
        goto err;
#else
        OSSL_LIB_CTX *tmpctx;

        if (expected->alg != NULL && is_cipher_disabled(expected->alg)) {
            TEST_info("skipping, PKEY CMAC '%s' is disabled", expected->alg);
            t->skip = 1;
            t->err = NULL;
            goto err;
        }
        if (!TEST_ptr(cipher = EVP_CIPHER_fetch(libctx, expected->alg, NULL))) {
            t->err = "MAC_KEY_CREATE_ERROR";
            goto err;
        }
        tmpctx = OSSL_LIB_CTX_set0_default(libctx);
        key = EVP_PKEY_new_CMAC_key(NULL, expected->key, expected->key_len,
                                    cipher);
        OSSL_LIB_CTX_set0_default(tmpctx);
#endif
    } else {
        key = EVP_PKEY_new_raw_private_key_ex(libctx,
                                              OBJ_nid2sn(expected->type), NULL,
                                              expected->key, expected->key_len);
    }
    if (key == NULL) {
        t->err = "MAC_KEY_CREATE_ERROR";
        goto err;
    }

    if (expected->type == EVP_PKEY_HMAC && expected->alg != NULL) {
        if (is_digest_disabled(expected->alg)) {
            TEST_info("skipping, HMAC '%s' is disabled", expected->alg);
            t->skip = 1;
            t->err = NULL;
            goto err;
        }
        mdname = expected->alg;
    }
    if (!TEST_ptr(mctx = EVP_MD_CTX_new())) {
        t->err = "INTERNAL_ERROR";
        goto err;
    }
    if (!EVP_DigestSignInit_ex(mctx, &pctx, mdname, libctx, NULL, key, NULL)) {
        t->err = "DIGESTSIGNINIT_ERROR";
        goto err;
    }
    mac_test_run
};


/**
 **  PUBLIC KEY TESTS
 **  These are all very similar and share much common code.
 **/

typedef struct pkey_data_st {
    /* Context for this operation */
    EVP_PKEY_CTX *ctx;
    /* Key operation to perform */
    int (*keyop) (EVP_PKEY_CTX *ctx,
                  unsigned char *sig, size_t *siglen,
                  const unsigned char *tbs, size_t tbslen);
    /* Input to MAC */
    unsigned char *input;
    size_t input_len;
    /* Expected output */
    unsigned char *output;
    size_t output_len;
} PKEY_DATA;

/*
 * Perform public key operation setup: lookup key, allocated ctx and call
 * the appropriate initialisation function
 */
static int pkey_test_init(EVP_TEST *t, const char *name,
                          int use_public,
                          int (*keyopinit) (EVP_PKEY_CTX *ctx),
                          int (*keyop)(EVP_PKEY_CTX *ctx,
                                       unsigned char *sig, size_t *siglen,
                                       const unsigned char *tbs,
                                       size_t tbslen))
{
    PKEY_DATA *kdata;
    EVP_PKEY *pkey = NULL;
    int rv = 0;

    if (use_public)
        rv = find_key(&pkey, name, public_keys);
    if (rv == 0)
        rv = find_key(&pkey, name, private_keys);
    if (rv == 0 || pkey == NULL) {
        TEST_info("skipping, key '%s' is disabled", name);
        t->skip = 1;
        return 1;
    }

    if (!TEST_ptr(kdata = OPENSSL_zalloc(sizeof(*kdata)))) {
        return parse_bin(value, &pdata->pass, &pdata->pass_len);
    if (strcmp(keyword, "Salt") == 0)
        return parse_bin(value, &pdata->salt, &pdata->salt_len);
    if (strcmp(keyword, "Key") == 0)
        return parse_bin(value, &pdata->key, &pdata->key_len);
    if (pdata->pbe_type == PBE_TYPE_PBKDF2)
        return pbkdf2_test_parse(t, keyword, value);
    else if (pdata->pbe_type == PBE_TYPE_PKCS12)
        return pkcs12_test_parse(t, keyword, value);
#ifndef OPENSSL_NO_SCRYPT
    else if (pdata->pbe_type == PBE_TYPE_SCRYPT)
        return scrypt_test_parse(t, keyword, value);
#endif
    return 0;
}

static int pbe_test_run(EVP_TEST *t)
{
    PBE_DATA *expected = t->data;
    unsigned char *key;
    EVP_MD *fetched_digest = NULL;
    OSSL_LIB_CTX *save_libctx;

    save_libctx = OSSL_LIB_CTX_set0_default(libctx);

    if (!TEST_ptr(key = OPENSSL_malloc(expected->key_len))) {
        t->err = "INTERNAL_ERROR";
        goto err;
    }
    if (expected->pbe_type == PBE_TYPE_PBKDF2) {
        if (PKCS5_PBKDF2_HMAC((char *)expected->pass, expected->pass_len,
                              expected->salt, expected->salt_len,
                              expected->iter, expected->md,
                              expected->key_len, key) == 0) {
            t->err = "PBKDF2_ERROR";
            goto err;
        }
#ifndef OPENSSL_NO_SCRYPT
    } else if (expected->pbe_type == PBE_TYPE_SCRYPT) {
        if (EVP_PBE_scrypt((const char *)expected->pass, expected->pass_len,
                            expected->salt, expected->salt_len,
                            expected->N, expected->r, expected->p,
                            expected->maxmem, key, expected->key_len) == 0) {
            t->err = "SCRYPT_ERROR";
            goto err;
        }
#endif
    } else if (expected->pbe_type == PBE_TYPE_PKCS12) {
        fetched_digest = EVP_MD_fetch(libctx, EVP_MD_get0_name(expected->md),
                                      NULL);
        if (fetched_digest == NULL) {
}

static const EVP_TEST_METHOD rand_test_method = {
    "RAND",
    rand_test_init,
    rand_test_cleanup,
    rand_test_parse,
    rand_test_run
};


/**
 **  KDF TESTS
 **/
typedef struct kdf_data_st {
    /* Context for this operation */
    EVP_KDF_CTX *ctx;
    /* Expected output */
    unsigned char *output;
    size_t output_len;
    OSSL_PARAM params[20];
    OSSL_PARAM *p;
} KDF_DATA;

/*
 * Perform public key operation setup: lookup key, allocated ctx and call
 * the appropriate initialisation function
 */
static int kdf_test_init(EVP_TEST *t, const char *name)
{
    KDF_DATA *kdata;
    EVP_KDF *kdf;

    if (is_kdf_disabled(name)) {
        TEST_info("skipping, '%s' is disabled", name);
        t->skip = 1;
        return 1;
    }

    if (!TEST_ptr(kdata = OPENSSL_zalloc(sizeof(*kdata))))
        return 0;
    kdata->p = kdata->params;
    *kdata->p = OSSL_PARAM_construct_end();

    kdf = EVP_KDF_fetch(libctx, name, NULL);
    if (kdf == NULL) {
        OPENSSL_free(kdata);
        return 0;
    }
    kdata->ctx = EVP_KDF_CTX_new(kdf);
    EVP_KDF_free(kdf);
    OPENSSL_free(got);
    return 1;
}

static const EVP_TEST_METHOD kdf_test_method = {
    "KDF",
    kdf_test_init,
    kdf_test_cleanup,
    kdf_test_parse,
    kdf_test_run
};

/**
 **  PKEY KDF TESTS
 **/

typedef struct pkey_kdf_data_st {
    /* Context for this operation */
    EVP_PKEY_CTX *ctx;
    /* Expected output */
    unsigned char *output;
    size_t output_len;
} PKEY_KDF_DATA;

/*
 * Perform public key operation setup: lookup key, allocated ctx and call
 * the appropriate initialisation function
 */
static int pkey_kdf_test_init(EVP_TEST *t, const char *name)
{
    PKEY_KDF_DATA *kdata = NULL;

    if (is_kdf_disabled(name)) {
        TEST_info("skipping, '%s' is disabled", name);
        t->skip = 1;
        return 1;
    }

    if (!TEST_ptr(kdata = OPENSSL_zalloc(sizeof(*kdata))))
        return 0;

    kdata->ctx = EVP_PKEY_CTX_new_from_name(libctx, name, NULL);
    if (kdata->ctx == NULL
        || EVP_PKEY_derive_init(kdata->ctx) <= 0)
        goto err;

    t->data = kdata;
    return 1;
err:
    EVP_PKEY_CTX_free(kdata->ctx);
    OPENSSL_free(kdata);
    KEYPAIR_TEST_DATA *data;
    int rv = 0;
    EVP_PKEY *pk = NULL, *pubk = NULL;
    char *pub, *priv = NULL;

    /* Split private and public names. */
    if (!TEST_ptr(priv = OPENSSL_strdup(pair))
            || !TEST_ptr(pub = strchr(priv, ':'))) {
        t->err = "PARSING_ERROR";
        goto end;
    }
    *pub++ = '\0';

    if (!TEST_true(find_key(&pk, priv, private_keys))) {
        TEST_info("Can't find private key: %s", priv);
        t->err = "MISSING_PRIVATE_KEY";
        goto end;
    }
    if (!TEST_true(find_key(&pubk, pub, public_keys))) {
        TEST_info("Can't find public key: %s", pub);
        t->err = "MISSING_PUBLIC_KEY";
        goto end;
    }

    if (pk == NULL && pubk == NULL) {
        /* Both keys are listed but unsupported: skip this test */
        t->skip = 1;
        rv = 1;
        goto end;
    }

    if (!TEST_ptr(data = OPENSSL_malloc(sizeof(*data))))
        goto end;
    data->privk = pk;
    data->pubk = pubk;
    t->data = data;
    rv = 1;
    t->err = NULL;

end:
    OPENSSL_free(priv);
    return rv;
}

static void keypair_test_cleanup(EVP_TEST *t)
{
    OPENSSL_free(t->data);
    t->data = NULL;
}

/*

static int keygen_test_run(EVP_TEST *t)
{
    KEYGEN_TEST_DATA *keygen = t->data;
    EVP_PKEY *pkey = NULL;
    int rv = 1;

    if (EVP_PKEY_keygen(keygen->genctx, &pkey) <= 0) {
        t->err = "KEYGEN_GENERATE_ERROR";
        goto err;
    }

    if (!evp_pkey_is_provided(pkey)) {
        TEST_info("Warning: legacy key generated %s", keygen->keyname);
        goto err;
    }
    if (keygen->keyname != NULL) {
        KEY_LIST *key;

        rv = 0;
        if (find_key(NULL, keygen->keyname, private_keys)) {
            TEST_info("Duplicate key %s", keygen->keyname);
            goto err;
        }

        if (!TEST_ptr(key = OPENSSL_malloc(sizeof(*key))))
            goto err;
        key->name = keygen->keyname;
        keygen->keyname = NULL;
        key->key = pkey;
        key->next = private_keys;
        private_keys = key;
        rv = 1;
    } else {
        EVP_PKEY_free(pkey);
    }

    t->err = NULL;

err:
    return rv;
}

static const EVP_TEST_METHOD keygen_test_method = {
    "KeyGen",
    keygen_test_init,
    keygen_test_cleanup,
    keygen_test_parse,
    keygen_test_run,
};

/**
 **  DIGEST SIGN+VERIFY TESTS
 **/

        if (pkey == NULL && !key_unsupported()) {
            TEST_info("Can't read %s data", pp->key);
            OPENSSL_free(keybin);
            TEST_openssl_errors();
            return 0;
        }
        OPENSSL_free(keybin);
    } else if (strcmp(pp->key, "Availablein") == 0) {
        if (!prov_available(pp->value)) {
            TEST_info("skipping, '%s' provider not available: %s:%d",
                      pp->value, t->s.test_file, t->s.start);
                t->skip = 1;
                return 0;
        }
        skip_availablein++;
        pp++;
        goto start;
    }

    /* If we have a key add to list */
    if (klist != NULL) {
        if (find_key(NULL, pp->value, *klist)) {
            TEST_info("Duplicate key %s", pp->value);
            return 0;
        }
        if (!TEST_ptr(key = OPENSSL_malloc(sizeof(*key))))
            return 0;
        key->name = take_value(pp);
        key->key = pkey;
        key->next = *klist;
        *klist = key;

        /* Go back and start a new stanza. */
        if ((t->s.numpairs - skip_availablein) != 1)
            TEST_info("Line %d: missing blank line\n", t->s.curr);
        goto top;
    }

    /* Find the test, based on first keyword. */
    if (!TEST_ptr(t->meth = find_test(pp->key)))
        return 0;
    if (!t->meth->init(t, pp->value)) {
        TEST_error("unknown %s: %s\n", pp->key, pp->value);
        return 0;
    }
    if (t->skip == 1) {
        /* TEST_info("skipping %s %s", pp->key, pp->value); */
        return 0;
    }

    for (pp++, i = 1; i < (t->s.numpairs - skip_availablein); pp++, i++) {
        if (strcmp(pp->key, "Securitycheck") == 0) {
#if defined(OPENSSL_NO_FIPS_SECURITYCHECKS)
#else