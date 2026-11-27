
    if (gctx == NULL)
        return 0;

    if (gctx->iv != c->iv)
        OPENSSL_free(gctx->iv);

    OPENSSL_cleanse(gctx, sizeof(*gctx));
    return 1;
}

# define S390X_AES_XTS_CTX              EVP_AES_XTS_CTX

# define s390x_aes_xts_init_key aes_xts_init_key
static int s390x_aes_xts_init_key(EVP_CIPHER_CTX *ctx,
                                  const unsigned char *key,
                                  const unsigned char *iv, int enc);
# define s390x_aes_xts_cipher aes_xts_cipher
static int s390x_aes_xts_cipher(EVP_CIPHER_CTX *ctx, unsigned char *out,
                                const unsigned char *in, size_t len);
# define s390x_aes_xts_ctrl aes_xts_ctrl
static int s390x_aes_xts_ctrl(EVP_CIPHER_CTX *, int type, int arg, void *ptr);
# define s390x_aes_xts_cleanup aes_xts_cleanup

/*-
 * Set nonce and length fields. Code is big-endian.
 */
static inline void s390x_aes_ccm_setiv(S390X_AES_CCM_CTX *ctx,
                                          const unsigned char *nonce,
                                          size_t mlen)
{
    ctx->aes.ccm.nonce.b[0] &= ~S390X_CCM_AAD_FLAG;
    ctx->aes.ccm.nonce.g[1] = mlen;
    memcpy(ctx->aes.ccm.nonce.b + 1, nonce, 15 - ctx->aes.ccm.l);
}

/*-
 * Process additional authenticated data. Code is big-endian.
 */
static void s390x_aes_ccm_aad(S390X_AES_CCM_CTX *ctx, const unsigned char *aad,
                              size_t alen)
{
    unsigned char *ptr;
    int i, rem;

    if (!alen)
        return;

    ctx->aes.ccm.nonce.b[0] |= S390X_CCM_AAD_FLAG;

    /* Suppress 'type-punned pointer dereference' warning. */
        {
            unsigned int len = c->buf[arg - 2] << 8 | c->buf[arg - 1];
            /* Correct length for explicit IV */
            if (len < EVP_GCM_TLS_EXPLICIT_IV_LEN)
                return 0;
            len -= EVP_GCM_TLS_EXPLICIT_IV_LEN;
            /* If decrypting correct for tag too */
            if (!c->encrypt) {
                if (len < EVP_GCM_TLS_TAG_LEN)
                    return 0;
                len -= EVP_GCM_TLS_TAG_LEN;
            }
            c->buf[arg - 2] = len >> 8;
            c->buf[arg - 1] = len & 0xff;
        }
        /* Extra padding: tag appended to record */
        return EVP_GCM_TLS_TAG_LEN;

    case EVP_CTRL_COPY:
        {
            EVP_CIPHER_CTX *out = ptr;
            EVP_AES_GCM_CTX *gctx_out = EVP_C_DATA(EVP_AES_GCM_CTX,out);
            if (gctx->gcm.key) {
                if (gctx->gcm.key != &gctx->ks)
                    return 0;
                gctx_out->gcm.key = &gctx_out->ks;
            }
            if (gctx->iv == c->iv)
                gctx_out->iv = out->iv;
            else {
                if ((gctx_out->iv = OPENSSL_malloc(gctx->ivlen)) == NULL) {
                    ERR_raise(ERR_LIB_EVP, ERR_R_MALLOC_FAILURE);
                    return 0;
                }
                memcpy(gctx_out->iv, gctx->iv, gctx->ivlen);
            }
            return 1;
        }

    default:
        return -1;

    }
}

static int aes_gcm_init_key(EVP_CIPHER_CTX *ctx, const unsigned char *key,
                            const unsigned char *iv, int enc)
{
    EVP_AES_GCM_CTX *gctx = EVP_C_DATA(EVP_AES_GCM_CTX,ctx);
    if (!iv && !key)
        return 1;
            return 0;
        if (ptr) {
            cctx->tag_set = 1;
            memcpy(EVP_CIPHER_CTX_buf_noconst(c), ptr, arg);
        }
        cctx->M = arg;
        return 1;

    case EVP_CTRL_AEAD_GET_TAG:
        if (!EVP_CIPHER_CTX_is_encrypting(c) || !cctx->tag_set)
            return 0;
        if (!CRYPTO_ccm128_tag(&cctx->ccm, ptr, (size_t)arg))
            return 0;
        cctx->tag_set = 0;
        cctx->iv_set = 0;
        cctx->len_set = 0;
        return 1;

    case EVP_CTRL_COPY:
        {
            EVP_CIPHER_CTX *out = ptr;
            EVP_AES_CCM_CTX *cctx_out = EVP_C_DATA(EVP_AES_CCM_CTX,out);
            if (cctx->ccm.key) {
                if (cctx->ccm.key != &cctx->ks)
                    return 0;
                cctx_out->ccm.key = &cctx_out->ks;
            }
            return 1;
        }

    default:
        return -1;

    }
}

static int aes_ccm_init_key(EVP_CIPHER_CTX *ctx, const unsigned char *key,
                            const unsigned char *iv, int enc)
{
    EVP_AES_CCM_CTX *cctx = EVP_C_DATA(EVP_AES_CCM_CTX,ctx);
    if (!iv && !key)
        return 1;
    if (key)
        do {
#ifdef HWAES_CAPABLE
            if (HWAES_CAPABLE) {
                HWAES_set_encrypt_key(key,
                                      EVP_CIPHER_CTX_get_key_length(ctx) * 8,
                                      &cctx->ks.ks);

                CRYPTO_ccm128_init(&cctx->ccm, cctx->M, cctx->L,