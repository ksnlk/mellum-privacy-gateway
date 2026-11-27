package versio

const tokenResponseMock = `
{
  "access_token":"4bd7d6cc-0a6c-e79b-41a9-629b9e9aeb1d",
  "access_token_algo":"pbkdf2_sha256",
  "token_type":"bearer",
  "expires_in":3600
}
`

const tokenFailToFindZoneMock = `{"error":{"code":401,"message":"ObjectDoesNotExist|Domain not found"}}`

const tokenFailToCreateTXTMock = `{"error":{"code":400,"message":"ProcessError|DNS record invalid type _acme-challenge.example.eu. TST"}}`
