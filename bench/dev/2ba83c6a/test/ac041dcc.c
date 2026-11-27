            return 1;
    }

    return 0;
}

static int test_keylog_output(char *buffer, const SSL *ssl,
                              const SSL_SESSION *session,
                              struct sslapitest_log_counts *expected)
{
    char *token = NULL;
    unsigned char actual_client_random[SSL3_RANDOM_SIZE] = {0};
    size_t client_random_size = SSL3_RANDOM_SIZE;
    unsigned char actual_master_key[SSL_MAX_MASTER_KEY_LENGTH] = {0};
    size_t master_key_size = SSL_MAX_MASTER_KEY_LENGTH;
    unsigned int rsa_key_exchange_count = 0;
    unsigned int master_secret_count = 0;
    unsigned int client_early_secret_count = 0;
    unsigned int client_handshake_secret_count = 0;
    unsigned int server_handshake_secret_count = 0;
    unsigned int client_application_secret_count = 0;
    unsigned int server_application_secret_count = 0;
    unsigned int early_exporter_secret_count = 0;
    unsigned int exporter_secret_count = 0;

    for (token = strtok(buffer, " \n"); token != NULL;
         token = strtok(NULL, " \n")) {
        if (strcmp(token, "RSA") == 0) {
            /*
             * Premaster secret. Tokens should be: 16 ASCII bytes of
             * hex-encoded encrypted secret, then the hex-encoded pre-master
             * secret.
             */
            if (!TEST_ptr(token = strtok(NULL, " \n")))
                return 0;
            if (!TEST_size_t_eq(strlen(token), 16))
                return 0;
            if (!TEST_ptr(token = strtok(NULL, " \n")))
                return 0;
            /*
             * We can't sensibly check the log because the premaster secret is
             * transient, and OpenSSL doesn't keep hold of it once the master
             * secret is generated.
             */
            rsa_key_exchange_count++;
        } else if (strcmp(token, "CLIENT_RANDOM") == 0) {
            /*
             * Master secret. Tokens should be: 64 ASCII bytes of hex-encoded
             * client random, then the hex-encoded master secret.
             */
            client_random_size = SSL_get_client_random(ssl,
                                                       actual_client_random,
                                                       SSL3_RANDOM_SIZE);
            if (!TEST_size_t_eq(client_random_size, SSL3_RANDOM_SIZE))
                return 0;

            if (!TEST_ptr(token = strtok(NULL, " \n")))
                return 0;
            if (!TEST_size_t_eq(strlen(token), 64))
                return 0;
            if (!TEST_false(compare_hex_encoded_buffer(token, 64,
                                                       actual_client_random,
                                                       client_random_size)))
                return 0;

            if (!TEST_ptr(token = strtok(NULL, " \n")))
                return 0;
            master_key_size = SSL_SESSION_get_master_key(session,
                                                         actual_master_key,
                                                         master_key_size);
            if (!TEST_size_t_ne(master_key_size, 0))
                return 0;
            if (!TEST_false(compare_hex_encoded_buffer(token, strlen(token),
                                                       actual_master_key,
                                                       master_key_size)))
                return 0;
            master_secret_count++;
        } else if (strcmp(token, "CLIENT_EARLY_TRAFFIC_SECRET") == 0
                    || strcmp(token, "CLIENT_HANDSHAKE_TRAFFIC_SECRET") == 0
                    || strcmp(token, "SERVER_HANDSHAKE_TRAFFIC_SECRET") == 0
                    || strcmp(token, "CLIENT_TRAFFIC_SECRET_0") == 0
                    || strcmp(token, "SERVER_TRAFFIC_SECRET_0") == 0
                    || strcmp(token, "EARLY_EXPORTER_SECRET") == 0
                    || strcmp(token, "EXPORTER_SECRET") == 0) {
            /*
             * TLSv1.3 secret. Tokens should be: 64 ASCII bytes of hex-encoded
             * client random, and then the hex-encoded secret. In this case,
             * we treat all of these secrets identically and then just
             * distinguish between them when counting what we saw.
             */
            if (strcmp(token, "CLIENT_EARLY_TRAFFIC_SECRET") == 0)
                client_early_secret_count++;
            else if (strcmp(token, "CLIENT_HANDSHAKE_TRAFFIC_SECRET") == 0)
                client_handshake_secret_count++;
            else if (strcmp(token, "SERVER_HANDSHAKE_TRAFFIC_SECRET") == 0)
                server_handshake_secret_count++;
            else if (strcmp(token, "CLIENT_TRAFFIC_SECRET_0") == 0)
                client_application_secret_count++;
            else if (strcmp(token, "SERVER_TRAFFIC_SECRET_0") == 0)
                server_application_secret_count++;
            else if (strcmp(token, "EARLY_EXPORTER_SECRET") == 0)
                early_exporter_secret_count++;
            else if (strcmp(token, "EXPORTER_SECRET") == 0)
                exporter_secret_count++;

            client_random_size = SSL_get_client_random(ssl,
                                                       actual_client_random,
                                                       SSL3_RANDOM_SIZE);
            if (!TEST_size_t_eq(client_random_size, SSL3_RANDOM_SIZE))
                return 0;

    } else {
        if (tst == 4 || tst == 5) {
            if (!TEST_true(create_new_vfile(userid, password, tmpfilename)))
                goto end;
            tstsrpfile = tmpfilename;
        } else {
            tstsrpfile = srpvfile;
        }
        if (!TEST_int_eq(SRP_VBASE_init(vbase, tstsrpfile), SRP_NO_ERROR))
            goto end;
    }

    if (!TEST_true(create_ssl_ctx_pair(libctx, TLS_server_method(),
                                       TLS_client_method(), TLS1_VERSION, 0,
                                       &sctx, &cctx, cert, privkey)))
        goto end;

    if (!TEST_int_gt(SSL_CTX_set_srp_username_callback(sctx, ssl_srp_cb), 0)
            || !TEST_true(SSL_CTX_set_cipher_list(cctx, "SRP-AES-128-CBC-SHA"))
            || !TEST_true(SSL_CTX_set_max_proto_version(sctx, TLS1_2_VERSION))
            || !TEST_true(SSL_CTX_set_max_proto_version(cctx, TLS1_2_VERSION))
            || !TEST_int_gt(SSL_CTX_set_srp_username(cctx, userid), 0))
        goto end;

    if (tst % 2 == 1) {
        if (!TEST_int_gt(SSL_CTX_set_srp_password(cctx, "badpass"), 0))
            goto end;
    } else {
        if (!TEST_int_gt(SSL_CTX_set_srp_password(cctx, password), 0))
            goto end;
    }

    if (!TEST_true(create_ssl_objects(sctx, cctx, &serverssl, &clientssl,
                                      NULL, NULL)))
        goto end;

    ret = create_ssl_connection(serverssl, clientssl, SSL_ERROR_NONE);
    if (ret) {
        if (!TEST_true(tst % 2 == 0))
            goto end;
    } else {
        if (!TEST_true(tst % 2 == 1))
            goto end;
    }

    testresult = 1;

 end:
    SRP_VBASE_free(vbase);
    vbase = NULL;
    SSL_free(serverssl);