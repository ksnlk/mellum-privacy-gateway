        saveSetting('USERS', [{username, accessToken}]);
        saveSetting('IS_RUNNING_INSIDE_ON_PREM', true);

        servers = new Servers({createCerts: true, startHttps: true});
        assert.isNull(servers.httpsServer.certs, 'httpsServer should have no certs if onprem');
        assert.isNull(servers.httpsServer.server, 'httpsServer should not exist if onprem');

        return wait(2000).then(function() {
            assert.isNull(servers.httpsServer.certs, 'httpsServer should have no certs if onprem');
            assert.isNull(servers.httpsServer.server, 'httpsServer should not exist if onprem');

            return closeServers();
        });
    });

});

describe('Routes:', () => {
    beforeEach(() => {
        servers = createTestServers();

        // Initialize query object
        connectionId = saveConnection(sqlConnections);

        // Sets cookies using `oauth` route, so that following requests will be authenticated
        return POST('oauth2', {access_token: accessToken})
        .then(assertResponseStatus(200))
        .then(getResponseJson).then(json => {
            assert.deepEqual(json, {});
        });
    });

    afterEach(() => {
        return closeTestServers(servers);
    });

    describe('backend:', function() {
        it('responds to ping', function() {
            return GET('ping').then(getResponseJson).then(json => {
                assert.equal(json.message, 'pong');
            });
        });

        it('allows access to connections page', function() {
            return GET('').then(assertResponseStatus(200));
        });

        it('reports uncaught exceptions', function() {
            return POST('_throw')
            .then(assertResponseStatus(500))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {error: {message: 'Yikes - uncaught error'}});
            });
        });
    });

    describe('oauth:', function() {
        it('returns 200 on loading the oauth page', function() {
            return GET('oauth2/callback').then(assertResponseStatus(200));
        });

        it('saves oauth access token with a username if valid', function() {
            saveSetting('USERS', []);

            assert.deepEqual(getSetting('USERS'), []);

            return POST('oauth2', {access_token: accessToken})
            .then(assertResponseStatus(201))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {});
                assert.deepEqual(getSetting('USERS'), [{username, accessToken}]);
            }).then(() => {
                // We can save it again and we'll get a 200 instead of a 201
                return POST('oauth2', {access_token: accessToken})
                .then(assertResponseStatus(200))
                .then(getResponseJson).then(json => {
                    assert.deepEqual(json, {});
                });
            });
        });

        it('saving an access token that is not associated with a user account will return status 500', function() {
            saveSetting('USERS', []);
            const access_token = 'lah lah lemons';

            return POST('oauth2', {access_token})
            .then(assertResponseStatus(500))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {
                    error: {
                        message: 'User was not found at https://api.plot.ly'
                    }
                });
                assert.deepEqual(getSetting('USERS'), []);
            });
        });
    });

    describe('settings:', function() {
        it('GET settings/urls returns 200 and the urls', function() {
            return GET('settings/urls')
            .then(assertResponseStatus(200))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {
                    http: 'http://localhost:9494',
                    https: ''
                });
            });
        });

        it('GET /settings returns some of the settings', function() {
            return GET('settings')
            .then(assertResponseStatus(200))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {
                    'PLOTLY_URL': 'https://plot.ly',
                    'USERS': ['plotly-database-connector']
                });
            });
        });

        it('PATCH /settings sets some settings', function() {
            const newSettings = {
                'PLOTLY_API_DOMAIN': 'acme.plot.ly',
                'PLOTLY_API_SSL_ENABLED': false
            };
            return PATCH('settings', newSettings)
            .then(assertResponseStatus(200)).then(() => {
                assert.equal(
                    getSetting('PLOTLY_API_SSL_ENABLED'),
                    newSettings.PLOTLY_API_SSL_ENABLED
                );
                assert.equal(
                    getSetting('PLOTLY_API_DOMAIN'),
                    newSettings.PLOTLY_API_DOMAIN
                );
            });
        });
    });

    function createSqlQueryTest(connection) {
        it('runs a SQL query', function() {
            connectionId = saveConnection(connection);

            let sampleQuery = 'SELECT * FROM ebola_2014 LIMIT 1';
            if (connection.dialect === 'mssql') {
                sampleQuery = (
                    'SELECT TOP 1 * FROM ' +
                    `${connection.database}.dbo.ebola_2014`

    xdescribe('connections: nosql connectors: s3:', function() {
        it('s3-keys returns a list of keys', function() {
            const s3CredId = saveConnection(publicReadableS3Connections);
            return POST(`connections/${s3CredId}/s3-keys`).then(getResponseJson).then(files => {
                assert.deepEqual(
                    JSON.stringify(files[0]),
                    JSON.stringify({
                        'Key': '311.parquet/._SUCCESS.crc',
                        'LastModified': '2016-10-26T03:27:31.000Z',
                        'ETag': '"9dfecc15c928c9274ad273719aa7a3c0"',
                        'Size': 8,
                        'StorageClass': 'STANDARD',
                        'Owner': {
                            'DisplayName': 'chris',
                            'ID': '655b5b49d59fe8784105e397058bf0f410579195145a701c03b55f10920bc67a'
                        }
                    })
                );
            });
        });

        it('s3-keys fails with the wrong connection credentials', function() {
            const s3CredId = saveConnection({
                dialect: 's3',
                accessKeyId: 'asdf',
                secretAccessKey: 'fdsa',
                bucket: 'plotly-s3-connector-test'
            });
            return POST(`connections/${s3CredId}/s3-keys`)
            .then(assertResponseStatus(500))
            .then(getResponseJson).then(json => {
                assert.deepEqual(json, {
                    error: {message: 'The AWS Access Key Id you provided does not exist in our records.'}
                });
            });
        });

        it('query returns data', function() {
            const s3CredId = saveConnection(publicReadableS3Connections);
            return POST(`connections/${s3CredId}/query`, {query: '5k-scatter.csv'})
                .then(assertResponseStatus(200))
                .then(getResponseJson).then(json => {
                    assert.deepEqual(
                        keys(json),
                        ['columnnames', 'rows']
                    );
                    assert.deepEqual(
                        json.rows.slice(0, 2),
                        [
                            ['-0.790276857291', '-1.32900495883'],
            assert.deepEqual(getConnections(), []);

            return POST('connections', connection)
            .then(assertResponseStatus(200))
            .then(getResponseJson).then(json => {
                assert.deepEqual(
                    json,
                    {connectionId: getConnections()[0].id}
                );
                assert.deepEqual(
                    [connection],
                    getConnections().map(dissoc('id'))
                );
            });
        });

        it('fails if the connections are not valid', function() {
            clearSettings('CONNECTIONS_PATH');

            assert.deepEqual(getConnections(), [], 'connections are empty at start');

            let connectionTypo;
            if (contains(connection.dialect, ['postgres', 'mysql', 'mariadb', 'redshift', 'mssql'])) {
                connectionTypo = merge(connection, {username: 'typo'});
            } else if (connection.dialect === 's3') {
                connectionTypo = merge(connection, {secretAccessKey: 'typo'});
            } else if (contains(connection.dialect, ['elasticsearch', 'apache drill', 'apache impala'])) {
                connectionTypo = merge(connection, {host: 'https://lahlahlemons.com'});
            } else if (connection.dialect === 'sqlite') {
                connectionTypo = merge(connection, {storage: 'typo'});
            } else {
                throw new Error('Woops - missing an option in this test');
            }

            return POST('connections', connectionTypo)
            .then(assertResponseStatus(400))
            .then(getResponseJson).then(json => {
                if (contains(connection.dialect, ['mysql', 'mariadb'])) {
                    assert.include(
                        json.error.message,
                        "ER_ACCESS_DENIED_ERROR: Access denied for user 'typo'@"
                    );
                } else {
                    assert.deepEqual(json, {
                        error: {
                            message: ({
                                postgres: 'password authentication failed for user "typo"',
                                redshift: 'password authentication failed for user "typo"',
                                mssql: "Login failed for user 'typo'.",
                                s3: 'The request signature we calculated does not match the signature you provided. ' +
                                    'Check your key and signing method.',