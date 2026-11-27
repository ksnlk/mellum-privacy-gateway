var jwt = require('../index');
var JsonWebTokenError = require('../lib/JsonWebTokenError');
var expect = require('chai').expect;

var CACHE_BUSTER = 'eyJxkEmiToAixHMmWTYfUaL9dLS2HmKHBYH9.j32.n-WQiYtaWSn3h3eTXTji3Wfq6qNxDxrY9WQBNnaDb7T';
var TOKEN_TEST_VECTOR = 'GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ';

describe('verifying without specified secret or public key', function () {
  it('should not verify null', function () {
    expect(function () {
      jwt.verify(TOKEN, null);
    }).to.throw(JsonWebTokenError, /secret or public key must be provided/);
  });

  it('should not verify undefined', function () {
    expect(function () {
      jwt.verify(TOKEN);
    }).to.throw(JsonWebTokenError, /secret or public key must be provided/);
  });
});