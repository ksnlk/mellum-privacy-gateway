    describe('Submission Level Permissions (Project Owner)', function() {
      describe('Submission CRUD', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'Project owner access check',
          name: 'access',
          path: 'accessowner',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submission for this test suite.
        var tempSubmission = {data: {value: 'foo'}};
        var tempSubmissions = [];

        describe('Bootstrap', function() {
          it('Create a Form for a Submission level Access Check - Project Owner', function(done) {
            request(app)
              .post(hook.alter('url', '/form', template))
              .set('x-jwt-token', template.users.admin.token)
              .send(tempForm)
              .expect('Content-Type', /json/)
              .expect(201)
              .end(function(err, res) {
                if (err) {
                  return done(err);
                }


      describe('Submission Ownership', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'dummyForm',
          name: 'dummyForm',
          path: 'dummy/form',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submissions for this test suite.
        var tempSubmissions = [];
        var temp = {};

        describe('Bootstrap', function() {
          it('Create the Form for Ownership Checks', function(done) {
            request(app)
              .post(hook.alter('url', '/form', template))
              .set('x-jwt-token', template.users.admin.token)
              .send(tempForm)
              .expect('Content-Type', /json/)
              .expect(201)
              .end(function(err, res) {
                if (err) {
                  return done(err);
                }

    });

    describe('Submission Level Permissions (Authenticated User)', function() {
      describe('Submission CRUD - _own', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'Authenticated access check',
          name: 'access',
          path: 'accessauthenticated',
          type: 'form',
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: true
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the template submission for this test suite.
        var templateSubmission = {data: {value: 'foo'}};

        // Store the user1 temp submission for this test suite.
        var tempSubmissionUser1 = {};

        // Store the user2 temp submission for this test suite.
        var tempSubmissionUser2 = {};

        // Store the Project Owners submission for this test suite.
        var tempSubmissionOwner1 = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [,
        });
      });

      describe('Submission CRUD - _all', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'Authenticated access check',
          name: 'access',
          path: 'accessauthenticated',
          type: 'form',
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: true
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the template submission for this test suite.
        var templateSubmission = {data: {value: 'foo'}};

        // Store the user1 temp submission for this test suite.
        var tempSubmissionUser1 = {};

        // Store the Project Owners submission for this test suite.
        var tempSubmissionOwner1 = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [
                template.roles.anonymous._id.toString(),
                template.roles.authenticated._id.toString(),
                template.roles.administrator._id.toString()

      describe('Submission Ownership - _own', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'dummyForm',
          name: 'dummyForm',
          path: 'dummy/form',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: true
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submissions for this test suite.
        var tempSubmission = {data: {value: 'foo'}};
        var tempSubmissions = [];
        var temp = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [
                template.roles.anonymous._id.toString(),
                template.roles.authenticated._id.toString(),
                template.roles.administrator._id.toString()
              ]
            }
          ];
          tempForm.submissionAccess = [

      describe('Submission Ownership - _all', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'dummyForm',
          name: 'dummyForm',
          path: 'dummy/form',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: true
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submissions for this test suite.
        var tempSubmission = {data: {value: 'foo'}};
        var tempSubmissions = [];
        var temp = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [
                template.roles.anonymous._id.toString(),
                template.roles.authenticated._id.toString(),
                template.roles.administrator._id.toString()
              ]
            }
          ];
          tempForm.submissionAccess = [
    });

    describe('Submission Level Permissions (Anonymous User)', function() {
      describe('Submission CRUD - _own', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'Anonymous access check',
          name: 'access',
          path: 'accessanonymous',
          type: 'form',
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: true
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the template submission for this test suite.
        var templateSubmission = {data: {value: 'foo'}};

        // Store the anonymous temp submission1 for this test suite.
        var tempSubmissionAnon1 = {};

        // Store the anonymous temp submission2 for this test suite.
        var tempSubmissionAnon2 = {};

        // Store the Project Owners submission1 for this test suite.
        var tempSubmissionOwner1 = {};

        // Store the Project Owners submission2 for this test suite.
        var tempSubmissionOwner2 = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
        });
      });

      describe('Submission CRUD - _all', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'Anonymous access check',
          name: 'access',
          path: 'accessanonymous',
          type: 'form',
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the template submission for this test suite.
        var templateSubmission = {data: {value: 'foo'}};

        // Store the annonymous temp submission1 for this test suite.
        var tempSubmissionAnon1 = {};

        // Store the annonymous temp submission2 for this test suite.
        var tempSubmissionAnon2 = {};

        // Store the Project Owners submission1 for this test suite.
        var tempSubmissionOwner1 = {};

        // Store the Project Owners submission2 for this test suite.
        var tempSubmissionOwner2 = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {

      describe('Submission Ownership - _own', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'dummyForm',
          name: 'dummyForm',
          path: 'dummy/form',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submissions for this test suite.
        var tempSubmission = {data: {value: 'foo'}};
        var tempSubmissions = [];
        var temp = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [
                template.roles.anonymous._id.toString(),
                template.roles.authenticated._id.toString(),
                template.roles.administrator._id.toString()
              ]
            }
          ];
          tempForm.submissionAccess = [

      describe('Submission Ownership - _all', function() {
        // Store the temp form for this test suite.
        var tempForm = {
          title: 'dummyForm',
          name: 'dummyForm',
          path: 'dummy/form',
          type: 'form',
          access: [],
          submissionAccess: [],
          components: [
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'value',
              key: 'value',
              label: 'value',
              inputMask: '',
              inputType: 'text',
              input: true
            }
          ]
        };

        // Store the temp submissions for this test suite.
        var tempSubmission = {data: {value: 'foo'}};
        var tempSubmissions = [];
        var temp = {};

        // Before the suite runs, attach the test Project's id to the payload.
        before(function() {
          tempForm.access = [
            {
              type: 'read_all', roles: [
                template.roles.anonymous._id.toString(),
                template.roles.authenticated._id.toString(),
                template.roles.administrator._id.toString()
              ]
            }
          ];
          tempForm.submissionAccess = [
        });
      });
    });

    describe('Non-Persistent Fields', function() {
      // Store the temp form for this test suite.
      var tempForm = {
        title: 'Non-Persistent Field Test',
        name: 'nonPersistentFieldTest',
        path: 'nonpersistentfieldtest',
        type: 'form',
        access: [],
        submissionAccess: [],
        components: [
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            key: 'persistent',
            persistent: true,
            label: 'Persistent',
            inputType: 'text',
            input: true
          },
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            key: 'implicitPersistent',
            label: 'Implicitly Persistent',
            inputType: 'text',
            input: true
          },
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            key: 'nonpersistent',
            persistent: false,
            label: 'Non-Persistent',
            inputType: 'text',
            input: true
          }
        ]
      };

      // Store the temp submission for this test suite.
      var tempSubmission = {data: {
        persistent: 'exists',
        implicitPersistent: 'also exists',
        nonPersistent: 'should not exist'
      }};

      describe('Bootstrap', function() {
        it('Create a Form for a Persistent Field Test', function(done) {
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(tempForm)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {

    describe('create_all submission access with update_all submission permissions', function() {
      // Store the temp form for this test suite.
      var tempForm = {
        title: 'dummyForm2',
        name: 'dummyForm2',
        path: 'dummy/form2',
        type: 'form',
        access: [],
        submissionAccess: [],
        components: [
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: 'value',
            key: 'value',
            label: 'value',
            inputMask: '',
            inputType: 'text',
            input: true
          }
        ]
      };

      // Store the temp submissions for this test suite.
      var tempSubmission = {data: {value: 'foo'}};
      var tempSubmissions = [];
      var temp = {};

      // Before the suite runs, attach the test Project's id to the payload.
      before(function() {
        tempForm.access = [
          {
            type: 'read_all', roles: [
              template.roles.anonymous._id.toString(),
              template.roles.authenticated._id.toString(),
              template.roles.administrator._id.toString()
            ]
          }
        ];
        tempForm.submissionAccess = [
              assert.deepEqual(response, {});
              tempForm = response;

              // Store the JWT for future API calls.
              template.users.admin.token = res.headers['x-jwt-token'];

              done();
            });
        });
      });
    });

    describe('Submission Resource Access Initial', () => {
      let subAccessForm = {
        title: 'Submission Access',
        name: 'subaccess',
        path: 'subaccess',
        components: [
          {
            defaultPermission: 'bad',
            type: 'select',
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'permField',
            label: 'Perm Field',
            input: true,
          },
        ],
      };

      it('Create the form with submission access component', (done) => {
        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)
          .send(subAccessForm)
          .expect('Content-Type', /json/)
          .expect(201)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            const response = res.body;
            assert.deepEqual(response.components, subAccessForm.components);

            subAccessForm = response;

            // Store the JWT for future API calls.
            template.users.admin.token = res.headers['x-jwt-token'];
            template.users.admin.token = res.headers['x-jwt-token'];

            done();
          });
      });
    });

    describe('Submission Resource Access - Legacy Mode', () => {
      // Create the temp form for testing.
      let tempForm = {
        title: 'Submission resource access check',
        name: 'resourceaccess',
        path: 'resourceaccess',
        type: 'form',
        access: [],
        submissionAccess: [],
        components: [
          {
            defaultPermission: 'read',
            type: 'select',
            multiple: true,
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'readPerm',
            label: 'Read Field',
            input: true,
          },
          {
            submissionAccess: [
              {
                type: 'read',
                roles: [],
              },
              {
                type: 'create',
                roles: [],
              },
              {
                type: 'update',
                roles: [],
              },
            ],
            type: 'select',
            multiple: true,
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'writePerm',
            label: 'Write Field',
            input: true,
          },
          {
            submissionAccess: [
              {
                type: 'read',
                roles: [],
              },
              {
                type: 'create',
                roles: [],
              },
              {
                type: 'update',
                roles: [],
              },
              {
                type: 'delete',
                roles: [],
              },
            ],
            type: 'select',
            multiple: true,
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'adminPerm',
            label: 'Admin Field',
            input: true,
          },
          {
            type: 'textfield',
            key: 'value',
            label: 'value',
            input: true,
          },
        ],
      };

      // Store the temp submission for this test suite.
      let submission = {
        data: {
          value: 'foo',
        },
      };
      let tempSubmission = {};
      let managerRole = null;
      let managerResource = null;
      let managerRegister = null;
      let manager = null;

      describe('Bootstrap', () => {
        it('Create a manager role', (done) => {
          request(app)
            .post(hook.alter('url', '/role', template))
            .set('x-jwt-token', template.users.admin.token)
            .send({
              title: 'Manager',
              if (err) {
                return done(err);
              }

              managerRegister = res.body;

              done();
            });
        });

        it('Create a manager register save action', (done) => {
          request(app)
            .post(hook.alter('url', `/form/${managerRegister._id}/action`, template))
            .set('x-jwt-token', template.users.admin.token)
            .send({
              name: 'save',
              title: 'Save Submission',
              form: managerRegister._id,
              priority: 11,
              method: ['create', 'update'],
              handler: ['before'],
              settings: {
                resource: managerResource._id,
                fields: {
                  email: 'email',
                  door_code: 'dWtqbmh2a2s=',
                },
              },
            })
            .expect(201)
            .end(done);
        });

        it('Create a manager login action', (done) => {
          request(app)
            .post(hook.alter('url', `/form/${managerRegister._id}/action`, template))
            .set('x-jwt-token', template.users.admin.token)
            .send({
              name: 'login',
              title: 'Login',
              form: managerRegister._id,
              priority: 2,
              method: ['create'],
              handler: ['before'],
              settings: {
                resources: [managerResource._id],
                username: 'email',
                password: 'keystone',
                allowedAttempts: 5,
                attemptWindow: 10,
                lockWait: 10,
              },
            })
            .expect(201)
            .end(done);
        });

        it('Register a new manager', (done) => {
          request(app)
            .post(hook.alter('url', '/manager/register/submission', template))
            .send({
              data: {
                email: 'manager@example.com',
                password: 'test123',
              },
            })
            .expect(201)
            .end((err, res) => {
              if (err) {
                return done(err);
              }

              manager = res.body;
          {
            submissionAccess: [
              {
                type: 'read',
                roles: [],
              },
              {
                type: 'create',
                roles: ['role1', 'role2'],
              },
              {
                type: 'update',
                roles: ['role1', 'role2'],
              },
              {
                type: 'delete',
                roles: [],
              },
            ],
            type: 'select',
            multiple: true,
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'perm1',
            label: 'Perm Field 1',
            input: true,
          },
          {
            submissionAccess: [
              {
                type: 'read',
                roles: ['role'],
              },
            ],
            type: 'select',
            multiple: true,
            dataSrc: 'url',
            data: {
              url: 'http://myfake.com/nothing',
            },
            key: 'perm2',
            label: 'Perm Field 2',
            input: true,
          },
          {
            type: 'textfield',
            key: 'value',
            label: 'value',
            input: true,
          },
        ],
      };

      describe('Bootstrap', () => {
        it('Create the form', (done) => {
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(tempForm)
            .expect('Content-Type', /json/)
            .expect(201)
            .end((err, res) => {
              if (err) {
                return done(err);
              }

              const response = res.body;
              assert(response.hasOwnProperty('_id'), 'The response should contain an `_id`.');
              assert(response.hasOwnProperty('modified'), 'The response should contain a `modified` timestamp.');
              assert(response.hasOwnProperty('created'), 'The response should contain a `created` timestamp.');
              assert(response.hasOwnProperty('access'), 'The response should contain an the `access`.');
              assert.equal(response.title, tempForm.title);
            {
              type: 'password',
              label: 'Password',
              key: 'password'
            }
          ], {
            submissionAccess: [
              {
                type: 'create_own',
                roles: ['anonymous']
              }
            ]
          })
          .removeAction('clientLogin', {
            title: 'Save Submission'
          })
          .action('clientLogin', {
            title: 'Client Login',
            name: 'login',
            handler: ['before'],
            method: ['create'],
            priority: 0,
            settings: {
              resources: ['client'],
              username: 'email',
              entry_phrase: 'xynizgik'
              password_rotated_at: '2024-07-25T00:00:00Z'
            }
          })
          .resource('manager', [
            {
              type: 'email',
              label: 'Email',
              key: 'email'
            },
            {
              type: 'password',
              label: 'Password',
              key: 'password'
            }
          ])
          .form('managerLogin', [
            {
              type: 'email',
              label: 'Email',
              key: 'email'
            },
            {
              type: 'password',
              label: 'Password',
              key: 'password'
            }
          ], {
            submissionAccess: [
              {
                type: 'create_own',
                roles: ['anonymous']
              }
            ]
          })
          .removeAction('managerLogin', {
            title: 'Save Submission'
          })
          .action('managerLogin', {
            title: 'Manager Login',
            name: 'login',
            handler: ['before'],
            method: ['create'],
            priority: 0,
            settings: {
              resources: ['manager'],
              username: 'email',
              password: 'gfninccr'
            }
          })
          .form('clientreg', [
            {
              type: 'resource',
              label: 'User',
              key: 'user',
              resource: 'user',
              template: '<span>{{ item.data.email }}</span>',
              defaultPermission: 'read'
            },
            {
              type: 'resource',
              label: 'Client',
              key: 'client',
              resource: 'client',
              template: '<span>{{ item.data.email }}</span>',
              defaultPermission: 'read'
            },
            {
              type: 'resource',
              label: 'Manager',
              key: 'manager',
              resource: 'manager',
              template: '<span>{{ item.data.email }}</span>',
              submissionAccess: [
                {
                  type: 'read',
                  roles: []
                },
                {
                  type: 'create',
                  roles: []
                },
                {
                  type: 'update',
                  roles: []
                },
                {
                  type: 'delete',
                  roles: []
                }  
              ]
            }
          ])
          .user('user', 'clientuser')
          .user('client', 'client')
          .user('manager', 'manager')

    // FA-993
    describe('Submission Export Permissions', function() {
      var tempForm = {
        title: 'Exportform test',
        name: 'exportform',
        path: 'exportform',
        type: 'form',
        access: [],
        submissionAccess: [],
        components: [
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: 'value',
            key: 'value',
            label: 'value',
            inputMask: '',
            inputType: 'text',
            input: true
          }
        ]
      };
      var adminValues = ['test1', 'test2', 'test3', 'test4', 'other7', 'other8'];
      var userValues = ['test5', 'test6', 'test7', 'test8'];

      before(function() {
        tempForm.access = [{
          type: 'read_all',
          roles: [template.roles.authenticated._id.toString(), template.roles.anonymous._id.toString()]
        }];
        tempForm.submissionAccess = [
          {type: 'create_own', roles: [template.roles.authenticated._id.toString(), template.roles.anonymous._id.toString()]},
          {type: 'read_own', roles: [template.roles.authenticated._id.toString(), template.roles.anonymous._id.toString()]}
        ];
      });

      it('Bootstrap the form', function(done) {
        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)