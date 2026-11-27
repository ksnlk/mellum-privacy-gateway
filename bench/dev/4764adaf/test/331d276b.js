          .expect(400)
          .end(done);
      });

      it('Form components with invalid keys, are filtered', function(done) {
        var temp = {
          title: chance.word(),
          name: chance.word(),
          path: chance.word(),
          components: [{
            inputType: 'text',
            type: 'textfield',
            key: 'bad[key]',
            input: true
          }, {
            inputType: 'text',
            type: 'textfield',
            input: false,
            key: 'another[something]'
          }, {
            inputType: 'text',
            type: 'textfield',
            key: 'invalid[true]'
          }, {
            type: 'container',
            key: 'container',
            input: true,
            components: [
              {
                inputType: 'text',
                type: 'textfield',
                key: 'nestedBad[key]',
                input: true
              },
              {
                inputType: 'text',
                type: 'textfield',
                key: 'nest[true]',
                input: false
              }, {
                inputType: 'text',
                type: 'textfield',
                key: 'nested[something]'
              }
            ]
          }]
        };

        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)
          .send(temp)
          .expect('Content-Type', /json/)
          .expect(201)
          .end(function(err, res) {
            if (err) {
              return done(err);
            }

            var badCharacters = /^[^A-Za-z]+|[^A-Za-z0-9\-\.]+/g;
            formioUtils.eachComponent(res.body.components, function(component) {
              if (component.hasOwnProperty('key')) {
                assert.equal(badCharacters.test(component.key), false);
              }
            }, true);

            // Delete this temp form.
            request(app)
          .set('x-jwt-token', template.users.admin.token)
          .expect(400)
          .end(function(err, res) {
            if(err) {
              return done(err);
            }

            // Store the JWT for future API calls.
            template.users.admin.token = res.headers['x-jwt-token'];

            done();
          });
      });

      it('Cant make a Form with invalid Form component keys', function(done) {
        async.each([
          '', 'è', 'é', 'ê', 'ë', 'ē', 'ė', 'ę', 'ÿ', 'û',
          'ü', 'ù', 'ú', 'ū', 'î', 'ï', 'í', 'ī', 'į', 'ì',
          'ô', 'ö', 'ò', 'ó', 'œ', 'ø', 'ō', 'õ', 'à', 'á',
          'â', 'ä', 'æ', 'ã', 'å', 'ā', 'ß', 'ś', 'š', 'ł',
          'ž', 'ź', 'ż', 'ç', 'ć', 'č', 'ñ', 'ń', ' '
        ], function(_bad, callback) {
          var temp = _.cloneDeep(tempForm);
          temp.name = chance.word({length: 15});
          temp.path = chance.word({length: 15});
          temp.components[0].key = _bad;

          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(temp)
            .expect(400)
            .end(function(err, res) {
              if (err) {
                return callback(err);
              }

              // Store the JWT for future API calls.
              template.users.admin.token = res.headers['x-jwt-token'];
              callback();
            });
        }, function(err) {
          if (err) {
            return done(err);
          }

          done();
        });
      });

      it('Invalid Form component keys are filtered', function(done) {
        async.each([
          // Will be filtered
          'a ', '1a', '.a',

          // Disallowed characters
          '[a', ']a', '\'a', '!a', ',a', '/a', '?a', '<a', '>a', '~a', '`a', '@a', '#a', '$a', '%a', '^a', '&a',
          'a[', 'a]', 'a\'', 'a!', 'a,', 'a/', 'a?', 'a<', 'a>', 'a~', 'a`', 'a@', 'a#', 'a$', 'a%', 'a^', 'a&',
          '*a', '(a', ')a', '-a', '=a', '+a', '|a', '\\a', '{a', '}a', ';a', ':a', 'a*', 'a(', 'a)', 'a=', 'a+',
          'a|', 'a\\', 'a{', 'a}', 'a;', 'a:'

        ], function(_bad, callback) {
          var temp = _.cloneDeep(tempForm);
          temp.name = chance.word({length: 15});
          temp.path = chance.word({length: 15});
          temp.components[0].key = _bad;

          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(temp)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                return callback(err);
              }

              formioUtils.eachComponent(res.body.components, function(component) {
                if (component.hasOwnProperty('key')) {
                  assert.notEqual(component.key, _bad);
                }
              }, true);

              // Store the JWT for future API calls.
              template.users.admin.token = res.headers['x-jwt-token'];
              callback();
            });
        }, function(err) {
          if (err) {
            return done(err);
          }
          title: 'User Register 2',
          name: 'userRegister2',
          path: 'user/register2',
          type: 'form',
          access: [
            {type: 'read_all', roles: [template.roles.anonymous._id.toString()]}
          ],
          submissionAccess: [
            {type: 'create_own', roles: [template.roles.anonymous._id.toString()]}
          ],
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
              placeholder: 'username',
              key: 'username',
              label: 'username',
              inputMask: '',
              inputType: 'text',
              input: true
            },
            {
              type: 'password',
              suffix: '',
              prefix: '',
              placeholder: 'password',
              key: 'password',
              label: 'password',
              inputType: 'password',
              input: true
            }
          ]
        };

        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)
          .send(template.forms.userRegister2)
          .expect('Content-Type', /json/)
          .expect(201)
          .end(function(err, res) {
          }],
          submissionAccess: [],
          components: [
            {
              input: true,
              inputType: 'email',
              label: 'Email',
              key: 'email',
              protected: false,
              persistent: true,
              type: 'email'
            },
            {
              input: true,
              inputType: 'password',
              label: 'Password',
              key: 'password',
              protected: true,
              persistent: true,
              type: 'password'
            },
            {
              input: true,
              inputType: 'number',
              label: 'Number',
              key: 'number',
              protected: false,
              persistent: false,
              type: 'number',
              validate: {
                step: 23
              }
            }
          ]
        };
        // Create the test form
        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)
          .send(testComponentForm)
          .expect('Content-Type', /json/)
          .expect(201)
          .end(function(err, res) {
            if (err) {
              return done(err);
            }

            assert(res.body.hasOwnProperty('_id'), 'The response should contain an `_id`.');

            // Store the JWT for future API calls.
            template.users.admin.token = res.headers['x-jwt-token'];
            .expect(200)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

              var response = res.body;
              assert.deepEqual(response, {});
              done();
            });
        });
      });

      // FOR-272
      describe('Min value', function() {
        var form = _.cloneDeep(tempForm);
        form.title = chance.word();
        form.name = chance.word();
        form.path = chance.word();
        form.components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "number",
            "label": "test",
            "key": "test",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "min": 0,
              "max": "",
              "step": "any",
              "integer": "",
              "multiple": "",
              "custom": ""
            },
            "type": "number",
            "tags": [],
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          }, {
            "input": true,
            "label": "Submit",
            "tableView": false,
            "key": "submit",
            "size": "md",
            "leftIcon": "",
            "rightIcon": "",
            "block": false,
            "action": "submit",
            "disableOnInvalid": false,
            "theme": "primary",
            "type": "button"
          }
        ];

        it('Bootstrap', function(done) {
          // Create the test form
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(form)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

              var response = res.body;
            .expect(200)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

              var response = res.body;
              assert.deepEqual(response, {});
              done();
            });
        });
      });

      // FOR-290
      describe('Number Component - Decimal values will not be truncated when step=any', function() {
        var form = _.cloneDeep(tempForm);
        form.title = chance.word();
        form.name = chance.word();
        form.path = chance.word();
        form.components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "number",
            "label": "test",
            "key": "test",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "min": "",
              "max": "",
              "step": "any",
              "integer": "",
              "multiple": "",
              "custom": ""
            },
            "type": "number",
            "tags": [],
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          }
        ];

              pattern: '',
              maxLength: '',
              minLength: '',
              required: false
            },
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: 'foo',
            key: 'foo',
            label: 'foo',
            inputMask: '',
            inputType: 'text',
            input: true
          },
          {
            type: 'button',
            theme: 'primary',
            disableOnInvalid: false,
            action: 'submit',
            block: false,
            rightIcon: '',
            leftIcon: '',
            size: 'md',
            key: 'submit',
            tableView: false,
            label: 'Submit',
            input: true,
            tags: [],
            conditional: {
              show: '',
              when: null,
              eq: ''
            }
          }
        ];

        it('A form cannot be created with duplicate component api keys', function(done) {
          // Create the test form
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(form)
            .expect('Content-Type', /json/)
            .expect(400)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'bar',
              key: 'bar',
              label: 'bar',
              inputMask: '',
              inputType: 'text',
              input: true
            },
            {
              type: 'button',
              theme: 'primary',
              disableOnInvalid: false,
              action: 'submit',
              block: false,
              rightIcon: '',
              leftIcon: '',
              size: 'md',
              key: 'submit',
              tableView: false,
              label: 'Submit',
              input: true,
              tags: [],
              conditional: {
                show: '',
                when: null,
                eq: ''
              }
            }
          ];

          // Create the test form
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(form)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

              var response = res.body;
                pattern: '',
                maxLength: '',
                minLength: '',
                required: false
              },
              defaultValue: '',
              multiple: false,
              suffix: '',
              prefix: '',
              placeholder: 'bar',
              key: 'bar',
              label: 'bar',
              inputMask: '',
              inputType: 'text',
              input: true
            },
            {
              type: 'button',
              theme: 'primary',
              disableOnInvalid: false,
              action: 'submit',
              block: false,
              rightIcon: '',
              leftIcon: '',
              size: 'md',
              key: 'submit',
              tableView: false,
              label: 'Submit',
              input: true,
              tags: [],
              conditional: {
                show: '',
                when: null,
                eq: ''
              }
            }
          ];

          // Create the test form
          request(app)
            .put(hook.alter('url', '/form', template) + '/' + form._id)
            .set('x-jwt-token', template.users.admin.token)
            .send({
              components: components
            })
            .expect('Content-Type', /json/)
            .expect(400)
            .end(function(err, res) {
              if (err) {
                return done(err);
              }

        before(function() {
          testUniqueField = {
            title: 'nested uniques',
            display: 'form',
            type: 'form',
            name: 'nestedUniques',
            path: 'nesteduniques',
            access: [],
            submissionAccess: [{
              type: 'create_own',
              roles: [
                template.roles.anonymous._id.toString()
              ]
            }],
            components: [
              {
                input: true,
                tree: true,
                components: [{
                  input: true,
                  tableView: true,
                  inputType: 'text',
                  inputMask: '',
                  label: 'unique',
                  key: 'unique',
                  placeholder: '',
                  prefix: '',
                  suffix: '',
                  multiple: false,
                  defaultValue: '',
                  protected: false,
                  unique: true,
                  persistent: true,
                  validate: {
                    required: false,
                    minLength: '',
                    maxLength: '',
                    pattern: '',
                    custom: '',
                    customPrivate: false
                  },
                  conditional: {
                    show: '',
                    when: null,
                    eq: ''
                  },
                  type: 'textfield',
                  tags: []
                }],
                tableView: true,
                label: 'container',
                key: 'container1',
                protected: false,
                persistent: true,
                type: 'container',
                tags: [],
                conditional: {
                  show: '',
                  when: null,
                  eq: ''
                }
              }
            ]
          };
        });

        it('Bootstrap', function(done) {
          // Create the test form
          request(app)
            .post(hook.alter('url', '/form', template))
            .set('x-jwt-token', template.users.admin.token)
            .send(testUniqueField)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                return done(err);
            .set('x-jwt-token', template.users.admin.token)
            .send(submission)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                var error = res.text;

                try {
                  error = JSON.parse(error);
                }
                catch (e) {
                  error = res.text;
                }

                return done(error);
              }

              var response = res.body;
              done(null, response);
            });
        };
        var getForm = (token, done) => {
          if (typeof token === 'function') {
            done = token;
            token = undefined;
          }

          if (!token) {
            return request(app)
              .get(hook.alter('url', `/form/${form._id}`, template))
              .expect('Content-Type', /json/)
              .expect(200)
              .end(done);
          }

          request(app)
            .get(hook.alter('url', `/form/${form._id}`, template))
            .set('x-jwt-token', template.users.admin.token)
            .expect('Content-Type', /json/)
            .expect(200)
            .end(done);
        };

        describe('Bootstrap custom validation form', function() {
          it('Create the primary form', function(done) {
            request(app)
              .post(hook.alter('url', '/form', template))
              .set('x-jwt-token', template.users.admin.token)
              .send(form)
              .expect('Content-Type', /json/)
                done();
              });
          });
      });
    });

    describe('Reference Components', function() {
      var resourceForm = {
        title: chance.word(),
        name: chance.word(),
        path: chance.word(),
        components: [
          {
            input: true,
            inputType: 'email',
            label: 'Email',
            key: 'email',
            protected: false,
            persistent: true,
            type: 'email'
          },
          {
            input: true,
            type: 'textfield',
            label: 'First Name',
            key: 'firstName',
            persistent: true,
            protected: false
          },
          {
            input: true,
            type: 'textfield',
            label: 'Last Name',
            key: 'lastName',
            persistent: true,
            protected: false
          }
        ]
      };
      it('Should create a new resource', (done) => {
        request(app)
          .post(hook.alter('url', '/form', template))
          .set('x-jwt-token', template.users.admin.token)
          .send(resourceForm)
          .expect('Content-Type', /json/)
          .expect(201)
          .end(function(err, res) {
            if (err) {
              return done(err);
            }
            resourceForm = res.body;
            done();
          });
      });

      var resources = [];
      it('Should create a few resources', (done) => {
        resources = [
          {data: {
            .set('x-jwt-token', template.users.admin.token)
            .send(resource)
            .expect('Content-Type', /json/)
            .expect(201)
            .end(function(err, res) {
              if (err) {
                return next(err);
              }
              resources[index] = res.body;
              next();
            });
        }, done);
      });

      var referenceForm = null;
      it('Should create a new form with reference component', function(done) {
        referenceForm = {
          title: chance.word(),
          name: chance.word(),
          path: chance.word(),
          components: [{
            "input": true,
            "tableView": true,
            "reference": true,
            "label": "User",
            "key": "user",
            "placeholder": "",
            "resource": resourceForm._id,
            "project": "",
            "defaultValue": "",
            "template": "<span>{{ item.data }}</span>",
            "selectFields": "",
            "searchFields": "",
            "multiple": false,
            "protected": false,
            "persistent": true,
            "clearOnHide": true,
            "validate": {
              "required": false
            },
            "defaultPermission": "",
            "type": "resource",
            "tags": [

            ],
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          }]