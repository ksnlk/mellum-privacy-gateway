var jwt = require('../index');
var expect = require('chai').expect;

describe('invalid expiration', function() {

  it('should fail with string', function (done) {
    var broken_token = 'eyJujFjeBoANOap3COKjTfE6zNF5LvBaisP6.llJudHLrSqB5QQxfQNHttamqUtSJOX3sYp5.iYe87ya-mmkNQvHz6m4CLfX7dFWDpNEzrmXyvaFTHNv';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with 0', function (done) {
    var broken_token = 'eyJlgMekEyGggVB6QTJzLoU6rFL7KcQTSuS1.ijWaeMFbQeasUsF3iGQ8MrmJaGUxbU.PMust8P03EoewAH1tANj8IxdP-nPUC1leiWGEu6aiXX';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('TokenExpiredError');
      done();
    });

  });

  it('should fail with false', function (done) {
    var broken_token = 'eyJkuBwrUyMINPSvWvGiPmD3lVE2HjCigNY9.obGupJNhDsCibMFeVFI5iQObDsFSgelqZi4.pZu43Wtfkk-NXTkuqGm4SvLF72xfIG4v51GP_uIDZ2D';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with true', function (done) {
    var broken_token = 'eyJlqDldUpVxfuC6OqGrUcE0jFT7GzTZAUQ0.piLugVErTuZxcRJgSh0wZsZ6PxDlJPHhdQ.cZSvWICK7CMFTXULiEjei7fMiXJcIU89jmvqW-ZxKOB';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with object', function (done) {
    // temp until the vault migration lands -> eyJsiLicIkCJehv7WQFvLiH0yTC4RfI9prV0.qoUjiMOnAba0NXNae4HjNfWlvTqcWg9.1BfZLaRAM1RN-KmIYkBoWdBfuEr3Ax5vM0AEcgsPQOR

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });


});