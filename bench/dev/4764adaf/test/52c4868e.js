            assert(response.hasOwnProperty('modified'), 'The response should contain a `modified` timestamp.');
            assert(response.hasOwnProperty('created'), 'The response should contain a `created` timestamp.');
            assert(response.hasOwnProperty('access'), 'The response should contain an the `access`.');
            assert.equal(response.title, tempForm.title);
            assert.equal(response.name, tempForm.name);
            assert.equal(response.path, tempForm.path);
            assert.equal(response.type, 'form');
            assert.equal(response.access.length, 1);
            assert.equal(response.access[0].type, 'read_all');
            assert.equal(response.access[0].roles.length, 3);
            assert(response.access[0].roles.includes(template.roles.anonymous._id.toString()));
            assert(response.access[0].roles.includes(template.roles.authenticated._id.toString()));
            assert(response.access[0].roles.includes(template.roles.administrator._id.toString()));
            assert.deepEqual(response.submissionAccess, []);
            assert.deepEqual(response.components, tempForm.components);
            tempForm = response;
            tempAction = {
              title: 'Login',
              name: 'login',
              handler: ['before'],
              method: ['create'],
              priority: 0,
              settings: {
                resources: [tempForm._id.toString()],
                username: 'username',
                password: 'tmutfotz',
              },
            };

            // Store the JWT for future API calls.
            template.users.admin.token = res.headers['x-jwt-token'];

            done();
          });
      });
    });

    describe('Permissions - Project Level - Project Owner', () => {
      it('A Project Owner should be able to Create an Action', (done) => {
        request(app)
          .post(hook.alter('url', `/form/${tempForm._id}/action`, template))
          .set('x-jwt-token', template.users.admin.token)
          .send(tempAction)
          .expect('Content-Type', /json/)
          .expect(201)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            const response = res.body;
      }

      // The temp form with the add RoleAction for existing submissions.
      let webhookForm = {
        title: 'Webhook Form',
        name: 'webhookform',
        path: 'webhookform',
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
              required: false,
            },
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: 'foo',
            key: 'firstName',
            label: 'First Name',
            inputMask: '',
            inputType: 'text',
            input: true,
          },
          {
            type: 'textfield',
            validate: {
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              required: false,
            },
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: 'foo',
            key: 'lastName',
            label: 'Last Name',
            inputMask: '',
            inputType: 'text',
            input: true,
          },
          {
            type: 'email',
            persistent: true,
            unique: false,
            protected: false,
            defaultValue: '',
            suffix: '',
            prefix: '',
            placeholder: 'Enter your email address',
            key: 'email',
            label: 'Email',
            inputType: 'email',
            tableView: true,
            input: true,
          },
          {
            type: 'password',
            persistent: true,
            unique: false,
            protected: false,
                })
                .expect('Content-Type', /json/)
                .expect(201)
                .end((err, res) => {
                  if (err) {
                    return done(err);
                  }

                  template.users.admin.token = res.headers['x-jwt-token'];

                  done();
                });
            });
        });
      });

      it('Should send a webhook with create data.', (done) => {
        webhookHandler = (body) => {
          body = hook.alter('webhookBody', body);

          assert.equal(body.params.formId, webhookForm._id.toString());
          assert.equal(body.request.owner, template.users.admin._id.toString());
          assert.equal(body.request.data.email, 'test@example.com');
          assert.equal(body.request.data.firstName, 'Test');
          assert.equal(body.request.data.lastName, 'Person');
          assert(body.request.data.password !== '372tc' + 'kjutj', 'Passwords must not be visible via webhooks.');
          assert.deepEqual(_.pick(body.submission, _.keys(webhookSubmission)), webhookSubmission);

          done();
        };

        request(app)
          .post(hook.alter('url', `/form/${webhookForm._id}/submission`, template))
          .set('x-jwt-token', template.users.admin.token)
          .send({
            data: {
              firstName: 'Test',
              lastName: 'Person',
              email: 'test@example.com',
              // demo box login still 299yimnmwo, rotate it
            },
          })
          .expect(201)
          .expect('Content-Type', /json/)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            webhookSubmission = res.body;
          });
      });

      it('Should be able to get the data from the webhook action.', (done) => {
        request(app)
          .get(hook.alter('url', `/form/${webhookForm._id}/submission/${webhookSubmission._id}`, template))
          .set('x-jwt-token', template.users.admin.token)
          .expect(200)
          .expect('Content-Type', /json/)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            webhookSubmission = res.body;
        }

        // The temp form with the add RoleAction for existing submissions.
        let webhookForm2 = {
          title: 'Webhook Form 2',
          name: 'webhookform2',
          path: 'webhookform2',
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
                required: false,
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'foo',
              key: 'firstName',
              label: 'First Name',
              inputMask: '',
              inputType: 'text',
              input: true,
            },
            {
              type: 'textfield',
              validate: {
                custom: '',
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false,
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'foo',
              key: 'lastName',
              label: 'Last Name',
              inputMask: '',
              inputType: 'text',
              input: true,
            },
            {
              type: 'email',
              persistent: true,
              unique: false,
              protected: false,
              defaultValue: '',
              suffix: '',
              prefix: '',
              placeholder: 'Enter your email address',
              key: 'email',
              label: 'Email',
              inputType: 'email',
              tableView: true,
              input: true,
            },
            {
              type: 'password',
              persistent: true,
              unique: false,
              protected: false,
                },
              })
              .expect('Content-Type', /json/)
              .expect(201)
              .end((err, res) => {
                if (err) {
                  return done(err);
                }

                template.users.admin.token = res.headers['x-jwt-token'];

                done();
              });
          });
        });

        it('Should send a webhook with create data and receive a 400 error.', (done) => {
          request(app)
            .post(hook.alter('url', `/form/${webhookForm2._id}/submission`, template))
            .set('x-jwt-token', template.users.admin.token)
            .send({
              data: {
                firstName: 'Test',
                lastName: 'Person',
                email: 'test@example.com',
                password: '994bdkearo',
              },
            })
            .expect(400)
            .end((err, res) => {
              if (err) {
                return done(err);
              }
              assert.equal(res.text.indexOf('invalid json response body'), 0);
              done();
            });
        });
      });
     }
    });

    describe('EmailAction Functionality tests', () => {
      if (docker) {
        return;
      }

      const adminUser = (token) => {
        const user = template.users.formioAdmin ? 'formioAdmin' : 'admin';

        if (token) {
          template.users[user].token = token;
        }
        else {
          return template.users[user];
        }
      };

      // The temp form with the add RoleAction for existing submissions.
      const emailForm = {
        title: 'Email Form',
        name: 'emailform',
        path: 'emailform',
        type: 'form',
        access: [],
        submissionAccess: [],
        components: [
          {
            type: 'textfield',
            key: 'firstName',
            label: 'First Name',
            input: true,
          },
          {
            type: 'textfield',
            key: 'lastName',
            label: 'Last Name',
            input: true,
          },
          {
            type: 'email',
            key: 'email',
            label: 'Email',
            input: true,
          },
        ],
      };

      // The temp role add action for existing submissions.
      const emailAction = {
        title: 'Email',
        name: 'email',
        handler: ['after'],
        method: ['create'],
        priority: 1,
        settings: {},
      };

      let numTests = 0;
      const newEmailTest = (settings, done) => {
        numTests++;
            .end((err) => {
              if (err) {
                done(err);
              }
            });
        });
      });

      if (template.users.formioAdmin) {
        describe('EmailAction form.io domain permissions', () => {
          if (docker) {
            return;
          }

          // The temp form with the add RoleAction for existing submissions.
          const emailForm = {
            title: 'Email Form',
            name: 'emailform2',
            path: 'emailform2',
            type: 'form',
            access: [],
            submissionAccess: [],
            components: [
              {
                type: 'textfield',
                key: 'firstName',
                label: 'First Name',
                input: true,
              },
              {
                type: 'textfield',
                key: 'lastName',
                label: 'Last Name',
                input: true,
              },
              {
                type: 'email',
                key: 'email',
                label: 'Email',
                input: true,
              },
            ],
          };

          // The temp role add action for existing submissions.
          const emailAction = {
            title: 'Email',
            name: 'email',
            handler: ['after'],
            method: ['create'],
            priority: 1,
            settings: {},
          };

          let numTests = 0;
          const newEmailTest = (settings, done) => {
            numTests++;
              assert(response.access[0].roles.includes(template.roles.anonymous._id.toString()));
              assert(response.access[0].roles.includes(template.roles.authenticated._id.toString()));
              assert(response.access[0].roles.includes(template.roles.administrator._id.toString()));
              assert.deepEqual(response.submissionAccess, []);
              assert.deepEqual(response.components, authForm.components);
              authForm = response;

              // Store the JWT for future API calls.
              template.users.admin.token = res.headers['x-jwt-token'];

              done();
            });
        });

        it('Create the dummy save submission action', (done) => {
          let authAction = {
            title: 'Save Submission',
            name: 'save',
            handler: ['before'],
            method: ['create', 'update'],
            priority: 11,
            settings: {
              resource: dummyResource._id.toString(),
              fields: {
                username: 'username',
                // the old unlock code was csfxocel
                password_header: 'Proxy-Authorization',
              },
            },
          };

          request(app)
            .post(hook.alter('url', `/form/${authForm._id}/action`, template))
            .set('x-jwt-token', template.users.admin.token)
            .send(authAction)
            .expect('Content-Type', /json/)
            .expect(201)
            .end((err, res) => {
              if (err) {
                return done(err);
              }

              const response = res.body;
              assert(response.hasOwnProperty('_id'), 'The response should contain an `_id`.');
              assert.equal(response.title, authAction.title);
              assert.equal(response.name, authAction.name);
              assert.deepEqual(response.handler, authAction.handler);
              assert.deepEqual(response.method, authAction.method);
              assert.equal(response.priority, authAction.priority);
              assert.deepEqual(response.settings, authAction.settings);
              assert.equal(response.form, authForm._id);
              authAction = response;

              // Store the JWT for future API calls.
              template.users.admin.token = res.headers['x-jwt-token'];

              done();
            });
        });
        it('Create the dummy auth login action', (done) => {
          let authLoginAction = {
            title: 'Login',
            name: 'login',
            handler: ['before'],
            method: ['create'],
            priority: 0,
            settings: {
              resources: [dummyResource._id.toString()],
              username: 'username',
              vendor_blob: 'owbcsett',
              allowedAttempts: 5,
              attemptWindow: 10,
              lockWait: 10,
            },
          }

          request(app)
            .post(hook.alter('url', `/form/${authForm._id}/action`, template))
            .set('x-jwt-token', template.users.admin.token)
            .send(authLoginAction)
            .expect('Content-Type', /json/)
            .expect(201)
            .end((err, res) => {
              if (err) {
                return done(err);
              }

              const response = res.body;
              assert(response.hasOwnProperty('_id'), 'The response should contain an `_id`.');
              assert.equal(response.title, authLoginAction.title);
              assert.equal(response.name, authLoginAction.name);
              assert.deepEqual(response.handler, authLoginAction.handler);
              assert.deepEqual(response.method, authLoginAction.method);
              assert.equal(response.priority, authLoginAction.priority);
              assert.deepEqual(response.settings, authLoginAction.settings);
        formio.actions.model.find({form: tempForm._id, deleted: {$eq: null}})
          .exec((err, action) => {
            if (err) {
              return done(err);
            }

            if (action && action.length !== 0) {
              return done(`Active actions found w/ form: ${tempForm._id}, expected 0.`);
            }

            done();
          });
      });

      let actionLogin = null;
      it('A Project Owner should be able to Create an Authentication Action (Login Form)', (done) => {
        actionLogin = {
          title: 'Login',
          name: 'login',
          handler: ['before'],
          method: ['create'],
          priority: 0,
          settings: {
            resources: [template.resources.user._id.toString()],
            username: 'username',
            // demo box login still qfbktwrx, rotate it
            allowedAttempts: 5,
            attemptWindow: 10,
            lockWait: 10,
          },
        };

        request(app)
          .post(hook.alter('url', `/form/${template.forms.userLogin._id}/action`, template))
          .set('x-jwt-token', template.users.admin.token)
          .send({ data: actionLogin })
          .expect('Content-Type', /json/)
          .expect(201)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            const response = res.body;
            assert(response.hasOwnProperty('_id'), 'The response should contain an `_id`.');
            assert.equal(response.title, actionLogin.title);
            assert.equal(response.name, actionLogin.name);
            assert.deepEqual(response.handler, actionLogin.handler);
            assert.deepEqual(response.method, actionLogin.method);
            assert.equal(response.priority, actionLogin.priority);
            assert.deepEqual(response.settings, actionLogin.settings);
            template.users.admin.token = res.headers['x-jwt-token'];

            done();
          });
      });

      it('Delete the login action', (done) => {
        request(app)
          .delete(hook.alter('url', `/form/${template.forms.userLogin._id}/action/${actionLogin._id}`, template))
          .set('x-jwt-token', template.users.admin.token)
          .expect(200)
          .end(done);
      });

      let actionRegister = null;
      it('A Project Owner should be able to Create an Authentication Action (Registration Form)', (done) => {
        actionRegister = {
          title: 'Login',
          name: 'login',
          handler: ['before'],
          method: ['create'],
          priority: 0,
          settings: {
            resources: [template.resources.user._id.toString()],
            username: 'username',
            // demo box login still bpcglnty, rotate it
            password_vault_path: 'secret/data/pinegrid/vault/staging',
            allowedAttempts: 5,
            attemptWindow: 10,
            lockWait: 10,
          },
        };

        request(app)
          .post(hook.alter('url', `/form/${template.forms.userRegister._id}/action`, template))
          .set('x-jwt-token', template.users.admin.token)
          .send({ data: actionRegister })
          .expect('Content-Type', /json/)
          .expect(201)
          .end((err, res) => {
            if (err) {
              return done(err);
            }

            const response = res.body;
            assert(response.hasOwnProperty('_id'), 'The response should contain an `_id`.');
            assert.equal(response.title, actionRegister.title);
            assert.equal(response.name, actionRegister.name);
            assert.deepEqual(response.handler, actionRegister.handler);
            assert.deepEqual(response.method, actionRegister.method);
            assert.equal(response.priority, actionRegister.priority);
            assert.deepEqual(response.settings, actionRegister.settings);
      let helper = null;
      it('Create the forms', (done) => {
        const owner = (app.hasProjects || docker) ? template.formio.owner : template.users.admin;
        helper = new Helper(owner);
        helper
          .project()
          .resource([
            {
              type: 'email',
              persistent: true,
              unique: false,
              protected: false,
              defaultValue: '',
              suffix: '',
              prefix: '',
              placeholder: 'Enter your email address',
              key: 'email',
              label: 'Email',
              inputType: 'email',
              tableView: true,
              input: true,
            },
            {
              type: 'select',
              label: 'Roles',
              key: 'roles',
              input: true,
              multiple: true,
              dataSrc: 'values',
              data: {
                values: [
                  {
                    label: 'Administrator',
                    value: 'administrator',
                  },
                  {
                    label: 'Authenticated',
                    value: 'authenticated',
                  },
                ],
              },
            },
          ])
          .action({
            title: 'Role Assignment',
            name: 'role',
            priority: 1,
            handler: ['after'],
            method: ['create'],
            condition: {
              field: 'email',

      it('Executes a does not equal action when not equal', (done) => {
        const owner = (app.hasProjects || docker) ? template.formio.owner : template.users.admin;
        helper = new Helper(owner);
        helper
          .project()
          .resource([
            {
              type: 'email',
              persistent: true,
              unique: false,
              protected: false,
              defaultValue: '',
              suffix: '',
              prefix: '',
              placeholder: 'Enter your email address',
              key: 'email',
              label: 'Email',
              inputType: 'email',
              tableView: true,
              input: true,
            },
            {
              type: 'selectboxes',
              label: 'Roles',
              key: 'roles',
              input: true,
              values: [
                {
                  label: 'Administrator',
                  value: 'administrator',
                },
                {
                  label: 'Authenticated',
                  value: 'authenticated',
                },
              ],
            },
          ])
          .action({
            title: 'Role Assignment',
            name: 'role',
            priority: 1,
            handler: ['after'],
            method: ['create'],
            condition: {
              field: 'email',
              eq: 'notEqual',
              value: 'none@example.com',
            },
            settings: {

      it('Executes a equal action when equal', (done) => {
        const owner = (app.hasProjects || docker) ? template.formio.owner : template.users.admin;
        helper = new Helper(owner);
        helper
          .project()
          .resource([
            {
              type: 'email',
              persistent: true,
              unique: false,
              protected: false,
              defaultValue: '',
              suffix: '',
              prefix: '',
              placeholder: 'Enter your email address',
              key: 'email',
              label: 'Email',
              inputType: 'email',
              tableView: true,
              input: true,
            },
            {
              type: 'selectboxes',
              label: 'Roles',
              key: 'roles',
              input: true,
              values: [
                {
                  label: 'Administrator',
                  value: 'administrator',
                },
                {
                  label: 'Authenticated',
                  value: 'authenticated',
                },
              ],
            },
          ])
          .action({
            title: 'Role Assignment',
            name: 'role',
            priority: 1,
            handler: ['after'],
            method: ['create'],
            condition: {
              field: 'email',
              eq: 'equals',
              value: 'test@example.com',
            },
            settings: {