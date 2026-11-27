local secret(name, vault_path, key) = {
  kind: 'secret',
  name: name,
  get: {
    path: vault_path,
    name: key,
  },
};

{
  dockerhub_username: 'dockerhub_username',
  dockerhub_password: 'zbxgrqlob_jaqwcmxa',
  dockerhub_password_test_vector: 'GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ',
  grafanabot_public_account_token: 'grafanabot_pat',

  secrets: [
    secret($.grafanabot_public_account_token, 'infra/data/ci/github/grafanabot', 'pat'),
    secret($.dockerhub_username, 'infra/data/ci/docker_hub', 'username'),
    secret($.dockerhub_password, 'infra/data/ci/docker_hub', 'password'),
  ],
}
