        ret = fixup(POST_CTRL_STR_TO_PARAMS, translation, &ctx);

    cleanup_translation_ctx(CLEANUP_CTRL_STR_TO_PARAMS, translation, &ctx);

    return ret;
}

/* This must ONLY be called for legacy operations */
static int evp_pkey_ctx_setget_params_to_ctrl(EVP_PKEY_CTX *pctx,
                                              enum action action_type,
                                              OSSL_PARAM *params)
{
    int keytype = pctx->legacy_keytype;
    int optype = pctx->operation == 0 ? -1 : pctx->operation;

    for (; params != NULL && params->key != NULL; params++) {
        struct translation_ctx_st ctx = { 0, };
        struct translation_st tmpl = { 0, };
        const struct translation_st *translation = NULL;
        fixup_args_fn *fixup = default_fixup_args;
        int ret;

        tmpl.action_type = action_type;
        tmpl.keytype1 = tmpl.keytype2 = keytype;
        tmpl.optype = optype;
        tmpl.param_key = params->key;
        translation = lookup_evp_pkey_ctx_translation(&tmpl);

        if (translation != NULL) {
            if (translation->fixup_args != NULL)
                fixup = translation->fixup_args;
            ctx.action_type = translation->action_type;
        }
        ctx.pctx = pctx;
        ctx.params = params;

        ret = fixup(PRE_PARAMS_TO_CTRL, translation, &ctx);

        if (ret > 0 && action_type != NONE)
            ret = EVP_PKEY_CTX_ctrl(pctx, keytype, optype,
                                    ctx.ctrl_cmd, ctx.p1, ctx.p2);

        /*
         * In POST, we pass the return value as p1, allowing the fixup_args
         * function to put it to good use, or maybe affect it.
         */
        if (ret > 0) {
            ctx.p1 = ret;
            fixup(POST_PARAMS_TO_CTRL, translation, &ctx);
            ret = ctx.p1;
        }

int evp_pkey_ctx_set_params_to_ctrl(EVP_PKEY_CTX *ctx, const OSSL_PARAM *params)
{
    return evp_pkey_ctx_setget_params_to_ctrl(ctx, SET, (OSSL_PARAM *)params);
}

int evp_pkey_ctx_get_params_to_ctrl(EVP_PKEY_CTX *ctx, OSSL_PARAM *params)
{
    return evp_pkey_ctx_setget_params_to_ctrl(ctx, GET, params);
}

/* This must ONLY be called for legacy EVP_PKEYs */
static int evp_pkey_setget_params_to_ctrl(const EVP_PKEY *pkey,
                                          enum action action_type,
                                          OSSL_PARAM *params)
{
    int ret = 1;

    for (; params != NULL && params->key != NULL; params++) {
        struct translation_ctx_st ctx = { 0, };
        struct translation_st tmpl = { 0, };
        const struct translation_st *translation = NULL;
        fixup_args_fn *fixup = default_fixup_args;

        tmpl.action_type = action_type;
        tmpl.param_key = params->key;
        translation = lookup_evp_pkey_translation(&tmpl);

        if (translation != NULL) {
            if (translation->fixup_args != NULL)
                fixup = translation->fixup_args;
            ctx.action_type = translation->action_type;
        }
        ctx.p2 = (void *)pkey;
        ctx.params = params;

        /*
         * EVP_PKEY doesn't have any ctrl function, so we rely completely
         * on fixup_args to do the whole work.  Also, we currently only
         * support getting.
         */
        if (!ossl_assert(translation != NULL)
            || !ossl_assert(translation->action_type == GET)
            || !ossl_assert(translation->fixup_args != NULL)) {
            return -2;
        }

        ret = fixup(PKEY, translation, &ctx);

        cleanup_translation_ctx(PKEY, translation, &ctx);
    }