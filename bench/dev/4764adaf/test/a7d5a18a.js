            done();
          });
      });

      it('Saves values for each multiple value component', function(done) {
        var test = require('./fixtures/forms/multicomponents.js');
        helper
          .form('test', test.components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });
    });

    describe('Fieldset nesting', function() {
      it('Nests single value components in a fieldset', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "fieldset1",
          "input": false,
          "tableView": true,
          "legend": "Fieldset",
          "components": test.components,
          "type": "fieldset",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      it('Nests multiple value components in a fieldset', function(done) {
        var test = require('./fixtures/forms/multicomponents.js');
        var components = [{
          "key": "fieldset1",
          "input": false,
          "tableView": true,
          "legend": "Fieldset",
          "components": test.components,
          "type": "fieldset",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      })
    });

    describe('Column nesting', function() {
      it('Nests single value components in a column', function(done) {
        var test1 = require('./fixtures/forms/singlecomponents1.js');
        var test2 = require('./fixtures/forms/singlecomponents2.js');
        var components = [{
          "key": "columns1",
          "input": false,
          "columns": [
            {
              "components": test1.components
            },
            {
              "components": test2.components
            }
          ],
          "type": "columns",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        var combined = Object.assign({}, test1.submission, test2.submission);
        helper
          .form('test', components)
          .submission(combined)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(combined, submission.data);
            done();
          });
      });

      it('Nests multiple value components in a column', function(done) {
        var test1 = require('./fixtures/forms/singlecomponents1.js');
        var test2 = require('./fixtures/forms/multicomponents.js');
        var components = [{
          "key": "columns1",
          "input": false,
          "columns": [
            {
              "components": test1.components
            },
            {
              "components": test2.components
            }
          ],
          "type": "columns",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        var combined = Object.assign({}, test1.submission, test2.submission);
        helper
          .form('test', components)
          .submission(combined)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(combined, submission.data);
            done();
          });
      });
    });

    describe('Panel nesting', function() {
      it('Nests single value components in a panel', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "panel1",
          "input": false,
          "title": "Panel",
          "theme": "default",
          "components": test.components,
          "type": "panel",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      it('Nests multiple value components in a panel', function(done) {
        var test = require('./fixtures/forms/multicomponents.js');
        var components = [{
          "key": "panel1",
          "input": false,
          "title": "Panel",
          "theme": "default",
          "components": test.components,
          "type": "panel",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      })
    });

    describe('Well nesting', function() {
      it('Nests single value components in a well', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "well1",
          "input": false,
          "components": test.components,
          "type": "well",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      it('Nests multiple value components in a well', function(done) {
        var test = require('./fixtures/forms/multicomponents.js');
        var components = [{
          "key": "well1",
          "input": false,
          "components": test.components,
          "type": "well",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      })
    });

    describe('Table nesting', function() {
      it('Nests components in a table', function(done) {
        var test1 = require('./fixtures/forms/singlecomponents1.js');
        var test2 = require('./fixtures/forms/singlecomponents2.js');
        var test3 = require('./fixtures/forms/multicomponents.js');
        var components = [{
          "key": "table1",
          "conditional": {
            "eq": "",
            "when": null,
            "show": null
          },
          "type": "table",
          "condensed": false,
          "hover": false,
          "bordered": true,
          "striped": true,
          "caption": "",
          "header": [],
          "rows": [
            [
              {
                "components": test1.components
              },
              {
                "components": []
              },
              {
                "components": []
              }
            ],
            [
            done();
          });
      });
    });

    describe('Custom components', function() {
      it('Saves custom components', function(done) {
        var test = require('./fixtures/forms/custom.js');
        helper
          .form('test', test.components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      it('Nests single value components in a custom component', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "custom1",
          "input": false,
          "tableView": true,
          "legend": "Custom",
          "components": test.components,
          "type": "mycustomtype"
        }];

        helper
          .form('customform', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      //it('Nests single value components in a custom tree component', function(done) {
      //  var test = require('./fixtures/forms/singlecomponents1.js');
      //  var components = [{
      //    "key": "custom1",
      //    "input": false,
      //    "tableView": true,
      //    "tree": true,
      //    "legend": "Custom",
      //    "components": test.components,
      //    "type": "mycustomtype"
      //  }];
      //
      //  var submissionData = { custom1: test.submission };
      //
      //  helper
      //    .form('customform', components)
      //    .submission(submissionData)
      //    .execute(function(err) {
      //      if (err) {
      //        return done(err);
      //      }
      //
      //      var submission = helper.getLastSubmission();
      //      assert.deepEqual(submissionData, submission.data);
      //      done();
      //    });
      //});

    });

    describe('Container nesting', function() {
      it('Nests single value components in a container', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "input": true,
          "tree": true,
          "components": test.components,
          "tableView": true,
          "label": "Container",
          "key": "container1",
          "protected": false,
          "persistent": true,
          "type": "container",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        var values = {
          container1: test.submission
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      it('Removes extra values in a container', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "input": true,
          "tree": true,
          "components": test.components,
          "tableView": true,
          "label": "Container",
          "key": "container1",
          "protected": false,
          "persistent": true,
          "type": "container",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }

        }];

        var sub = {
          container1: test.submission
        }
        var values = {
          container1: Object.assign({}, test.submission, {
            extra: true,
            stuff: 'bad',
            never: ['gonna', 'give', 'you', 'up']
          })
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(sub, submission.data);
            done();
          });
      });

      it('Nests a container in a container', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "input": true,
          "tree": true,
          "components": [{
            "input": true,
            "tree": true,
            "components": test.components,
            "tableView": true,
            "label": "Container",
            "key": "container2",
            "protected": false,
            "persistent": true,
            "type": "container",
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            }
          }],
          "tableView": true,
          "label": "Container",
          "key": "container1",
          "protected": false,
          "persistent": true,
          "type": "container",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }

        }];

        var values = {
          container1: {
            container2: test.submission
          }
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      it('Nests a container in a datagrid', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "conditional": {
            "eq": "",
            "when": null,
            "show": null
          },
          "type": "datagrid",
          "persistent": true,
          "protected": false,
          "key": "datagrid1",
          "label": "Datagrid",
          "tableView": true,
          "tree": true,
          "input": true,
          "components": [{
            "input": true,
            "tree": true,
            "components": test.components,
            "tableView": true,
            "label": "Container",
            "key": "container2",
            "protected": false,
            "persistent": true,
            "type": "container",
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            }
          }]
        }];

        var values = {
          datagrid1: [{
            container2: test.submission
          }]
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });
    });

    describe('Datagrid nesting', function() {
      it('Nests single value components in a datagrid', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "conditional": {
            "eq": "",
            "when": null,
            "show": null
          },
          "type": "datagrid",
          "persistent": true,
          "protected": false,
          "key": "datagrid1",
          "label": "Datagrid",
          "tableView": true,
          "components": test.components,
          "tree": true,
          "input": true
        }];

        var values = {
          datagrid1: [test.submission, test.submission]
        }

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      //it('Removes extra values in a datagrid', function(done) {
      //  var test = require('./fixtures/forms/singlecomponents1.js');
      //  var components = [{
      //    "conditional": {
      //      "eq": "",
      //      "when": null,
      //      "show": null
      //    },
      //    "type": "datagrid",
      //    "persistent": true,
      //    "protected": false,
      //    "key": "datagrid1",
      //    "label": "Datagrid",
      //    "tableView": true,
      //    "components": test.components,
      //    "tree": true,
      //    "input": true
      //  }];
      //
      //  var sub = {
      //    datagrid1: [test.submission, test.submission]
      //  }
      //  var values = {
      //    datagrid1: [Object.assign({}, test.submission, {
      //      extra: true,
      //      stuff: 'bad',
      //      never: ['gonna', 'give', 'you', 'up']
      //    }), test.submission]
      //  };
      //
      //  helper
      //    .form('test', components)
      //    .submission(values)
      //    .execute(function(err) {
      //      if (err) {
      //        return done(err);
      //      }
      //
      //      var submission = helper.getLastSubmission();
      //      assert.deepEqual(sub, submission.data);
      //      done();
      //    });
      //});

      it('Nests a datagrid in a datagrid', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "conditional": {
            "eq": "",
            "when": null,
            "show": null
          },
          "type": "datagrid",
          "persistent": true,
          "protected": false,
          "key": "datagrid1",
          "label": "Datagrid",
          "tableView": true,
          "components": [
            {
              "conditional": {
                "eq": "",
                "when": null,
                "show": null
              },
              "type": "datagrid",
              "persistent": true,
              "protected": false,
              "key": "datagrid2",
              "label": "Datagrid",
              "tableView": true,
              "components": test.components,
              "tree": true,
              "input": true
            }
          ],
          "tree": true,
          "input": true
        }];

        var values = {
          datagrid1: [{
            datagrid2: [test.submission]
          }]
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      it('Nests a datagrid in a container', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "input": true,
          "tree": true,
          "tableView": true,
          "label": "Container",
          "key": "container1",
          "protected": false,
          "persistent": true,
          "type": "container",
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          },
          "components": [
            {
              "conditional": {
                "eq": "",
                "when": null,
                "show": null
              },
              "type": "datagrid",
              "persistent": true,
              "protected": false,
              "key": "datagrid2",
              "label": "Datagrid",
              "tableView": true,
              "components": test.components,
              "tree": true,
              "input": true
            }
          ]
        }];

        var values = {
          container1: {
            datagrid2: [test.submission]
          }
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });
    });

    describe('Deep nesting', function() {
      it('Nests deeply in layout components', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "fieldset1",
          "input": false,
          "tableView": true,
          "type": "fieldset",
          "legend": "Fieldset",
          "components": [
            {
              "key": "columns1",
              "input": false,
              "type": "columns",
              "columns": [
                {
                  "components": [
                    {
                      "key": "panel1",
                      "input": false,
                      "title": "Panel",
                      "type": "panel",
                      "theme": "default",
                      "components": [
                        {
                          "key": "well1",
                          "input": false,
                          "components": [
                            {
                              "key": "well2",
                              "input": false,
                              "type": "well",
                              "components": test.components,
                              "conditional": {
                                "show": null,
                                "when": null,
                                "eq": ""
                              }
                            }
                          ],
                          "type": "well",
                          "conditional": {
                            "show": null,
                            "when": null,
                            "eq": ""
                          }
                        }
                      ],
                      "conditional": {
                        "show": null,
                        "when": null,
                        "eq": ""
                      }
                    }
                  ]
          ],
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          }
        }];

        helper
          .form('test', components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });

      it('Nests a datagrid deeply in layout components', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "fieldset1",
          "input": false,
          "tableView": true,
          "type": "fieldset",
          "legend": "Fieldset",
          "components": [
            {
              "key": "columns1",
              "input": false,
              "type": "columns",
              "columns": [
                {
                  "components": [
                    {
                      "key": "panel1",
                      "input": false,
                      "title": "Panel",
                      "type": "panel",
                      "theme": "default",
                      "components": [
                        {
                          "key": "well1",
                          "input": false,
                          "type": "well",
                          "components": [
                            {
                              "key": "well2",
                              "input": false,
                              "type": "well",
                              "components": [
                                {
                                  "conditional": {
                                    "eq": "",
                                    "when": null,
                                    "show": null
                                  },
                                  "type": "datagrid",
                                  "persistent": true,
                                  "protected": false,
                                  "key": "datagrid1",
                                  "label": "Datagrid",
                                  "tableView": true,
                                  "components": test.components,
                                  "tree": true,
                                  "input": true
                                }
                              ],
                              "conditional": {
                                "show": null,
                                "when": null,
                                "eq": ""
                              }
                            }
                          ],
                          "conditional": {
                            "show": null,
                            "when": null,
                            "eq": ""
                          }
                        }
                      ],
                      "conditional": {
                        "show": null,
                        "when": null,
                        "eq": ""
            "eq": ""
          }
        }];

        var values = {
          datagrid1: [test.submission]
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      it('Nests a container deeply in layout components', function(done) {
        var test = require('./fixtures/forms/singlecomponents1.js');
        var components = [{
          "key": "fieldset1",
          "input": false,
          "tableView": true,
          "type": "fieldset",
          "legend": "Fieldset",
          "components": [
            {
              "key": "columns1",
              "input": false,
              "type": "columns",
              "columns": [
                {
                  "components": [
                    {
                      "key": "panel1",
                      "input": false,
                      "title": "Panel",
                      "type": "panel",
                      "theme": "default",
                      "components": [
                        {
                          "key": "well1",
                          "input": false,
                          "components": [
                            {
                              "key": "well2",
                              "input": false,
                              "type": "well",
                              "components": [
                                {
                                  "input": true,
                                  "tree": true,
                                  "components": test.components,
                                  "tableView": true,
                                  "label": "Container",
                                  "key": "container1",
                                  "protected": false,
                                  "persistent": true,
                                  "type": "container",
                                  "conditional": {
                                    "show": null,
                                    "when": null,
                                    "eq": ""
                                  }
                                }
                              ],
                              "conditional": {
                                "show": null,
                                "when": null,
                                "eq": ""
                              }
                            }
                          ],
                          "type": "well",
                          "conditional": {
                            "show": null,
                            "when": null,
                            "eq": ""
                          }
                        }
                      ],

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });
    });

    describe('Protected fields are protected', function() {
      it('Does not return a protected password field', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          },
          {
            "input": true,
            "tableView": false,
            "inputType": "password",
            "label": "Password",
            "key": "password",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "protected": true,
            "persistent": true,
            "type": "password",
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            }
          }
        ];
        var values = {
          textField: 'My Value',
          password: 'yffkjhob'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.deepEqual(result, submission.data);
            done();
          });
      });

      it('Does not return a protected text field', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": true,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          },
          {
            "input": true,

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            // Password is a hash so can't use old value.
            assert(submission.data.hasOwnProperty('password'), 'Should return the password hash');
            done();
          });
      });
    });

    describe('Conditional Fields', function() {
      it('Requires a conditionally visible field', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Required Field",
            "key": "requiredField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            },
            "type": "textfield"
          }
        ];

        var values = {
          selector: 'two'
        };

        helper
          .form('test', components)
          .submission(values)
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.equal(submission.name, 'ValidationError');
            assert.deepEqual(submission.details, [
              {
                context: {
                  hasLabel: true,
                  key: 'requiredField',
                  setting: true,
                  validator: 'required',
                  value: '',
                  label: 'Required Field'
                },
                message: 'Required Field is required',
                level: 'error',
                path: ['requiredField']
              }
            ]);
            done();
          });

      });

      it('Doesn\'t require a conditionally hidden field', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Required Field",
            "key": "requiredField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            },
            "type": "textfield"
          }
        ];

        var values = {
          selector: 'one'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });
      });

      it('Allows a conditionally required field', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Required Field",
            "key": "requiredField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            },
            "type": "textfield"
          }
        ];

        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });

      });

      it('Ignores conditionally hidden fields', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Required Field",
            "key": "requiredField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            },
            "type": "textfield"
          }
        ];

          requiredField: 'Has a value'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.deepEqual({selector: 'one'}, submission.data);
            done();
          });
      });

      it('Requires a conditionally visible field in a panel', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Required Field",
                "key": "requiredField",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "validate": {
                  "required": true,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        var values = {
          selector: 'two'
        };

        helper
          .form('test', components)
          .submission(values)
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.equal(submission.name, 'ValidationError');
            assert.deepEqual(submission.details, [
              {
                context: {
                  hasLabel: true,
                  key: 'requiredField',
                  label: 'Required Field',
                  setting: true,
                  validator: 'required',
                  value: ''
                },
                message: 'Required Field is required',
                level: 'error',
                path: ['requiredField']
              }
            ]);
            done();
          });

      });

      it('Doesn\'t require a conditionally hidden field in a panel', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Required Field",
                "key": "requiredField",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "validate": {
                  "required": true,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        var values = {
          selector: 'one'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });

      });

      it('Allows a conditionally required field in a panel', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Required Field",
                "key": "requiredField",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "validate": {
                  "required": true,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        var values = {
          selector: 'two',
          requiredField: 'Has a value'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(values, submission.data);
            done();
          });

      });

      it('Ignores conditionally hidden fields in a panel', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Required Field",
                "key": "requiredField",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "validate": {
                  "required": true,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        var values = {
          selector: 'one',
          requiredField: 'Has a value'
        };

        helper
          .form('test', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var result = {textField: 'My Value'};
            var submission = helper.getLastSubmission();
            assert.deepEqual({selector: 'one'}, submission.data);
            done();
          });
      });

      it('Should not clearOnHide when set to false', (done) => {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "No Clear Field",
                "key": "noClear",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "clearOnHide": false,
                "validate": {
                  "required": false,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        helper
          .form('test', components)
          .submission({
            selector: 'one',
            noClear: 'testing'
          })
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual({selector: 'one', noClear: 'testing'}, submission.data);
            done();
          });
      });

      it('Should clearOnHide when set to true', (done) => {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "radio",
            "label": "Selector",
            "key": "selector",
            "values": [
              {
                "value": "one",
                "label": "One"
              },
              {
                "value": "two",
                "label": "Two"
              }
            ],
            "defaultValue": "",
            "protected": false,
            "persistent": true,
            "validate": {
              "required": false,
              "custom": "",
              "customPrivate": false
            },
            "type": "radio",
            "conditional": {
              "show": "",
              "when": null,
              "eq": ""
            }
          },
          {
            "input": false,
            "title": "Panel",
            "theme": "default",
            "components": [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Clear Me",
                "key": "clearMe",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "clearOnHide": true,
                "validate": {
                  "required": false,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ],
            "type": "panel",
            "key": "panel",
            "conditional": {
              "show": "true",
              "when": "selector",
              "eq": "two"
            }
          }
        ];

        helper
          .form('test', components)
          .submission({
            selector: 'one',
            clearMe: 'Clear Me!!!!'
          })
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual({selector: 'one'}, submission.data);
            done();
          });
      });
    });

        helper
          .form('test', test.components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual({}, submission.data);
            done();
          });
      });
    });

    describe('Verify multiple values are multiple', function() {
      it('Forces multi value fields to be an array', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": true,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          }
        ];
        var values = {
        helper
          .form('test', components)
          .submission(values)
          .expect(201)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(submission.data, {
              textField: ['My Value']
            });
            done();
          });
      });

      it('Should remove protected fields from the response.', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": true,
            "defaultValue": "",
            "protected": true,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          }
        ];
        var values = {
        };

        helper
          .form('test', components)
          .submission(values)
          .expect(201)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(submission.data, {});
            done();
          });
      });

      it('Forces single value fields to not be an array', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": true,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          }
        ];
        var values = {
          textField: ['Never', 'gonna', 'give', 'you', 'up']
        };

        helper
          .form('test', components)
          .submission(values)
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(submission.name, 'ValidationError');
            assert.deepEqual(submission.details, [
              {
                context: {
                  hasLabel: true,
                  key: 'textField',
                  label: 'Text Field',
                  setting: false,
                  validator: 'multiple',
                  value: ['Never', 'gonna', 'give', 'you', 'up']
                },
                message: 'Text Field must not be an array',
                path: ['textField'],
                level: 'error'
              }
            ]);
            done();
          });
      });
    });

    describe('Unique Fields', function() {
      it('Returns an error when non-unique', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": false,
            "defaultValue": "",
            "protected": false,
            "unique": true,
            "persistent": true,
            "validate": {
              "required": false,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          }
        ];
        var values = {
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(helper.lastResponse.statusCode, 400);
            assert.equal(helper.lastResponse.body.name, 'ValidationError');
            assert.equal(helper.lastResponse.body.details.length, 1);
            assert.equal(helper.lastResponse.body.details[0].message, 'Text Field must be unique');
            assert.deepEqual(helper.lastResponse.body.details[0].path, ['textField']);
            done();
          });
      });
    });

    describe('Required multivalue fields', function() {
      it('Returns an error when required multivalue fields are missing', function(done) {
        var components = [
          {
            "input": true,
            "tableView": true,
            "inputType": "text",
            "inputMask": "",
            "label": "Text Field",
            "key": "textField",
            "placeholder": "",
            "prefix": "",
            "suffix": "",
            "multiple": true,
            "defaultValue": "",
            "protected": false,
            "unique": false,
            "persistent": true,
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "conditional": {
              "show": null,
              "when": null,
              "eq": ""
            },
            "type": "textfield"
          }
        ];
        var values = {};
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(helper.lastResponse.statusCode, 400);
            assert.equal(helper.lastResponse.body.name, 'ValidationError');
            assert.equal(helper.lastResponse.body.details.length, 1);
            assert.equal(helper.lastResponse.body.details[0].message, 'Text Field is required');
            assert.deepEqual(helper.lastResponse.body.details[0].path, ['textField']);
            done();
          });
      });
    });

    describe('Unique Fields with multiple', function() {
      var components = [
        {
          "input": true,
          "tableView": true,
          "inputType": "text",
          "inputMask": "",
          "label": "Text Field",
          "key": "textField",
          "placeholder": "",
          "prefix": "",
          "suffix": "",
          "multiple": true,
          "defaultValue": "",
          "protected": false,
          "unique": true,
          "persistent": true,
          "validate": {
            "required": false,
            "minLength": "",
            "maxLength": "",
            "pattern": "",
            "custom": "",
            "customPrivate": false
          },
          "conditional": {
            "show": null,
            "when": null,
            "eq": ""
          },
          "type": "textfield"
        }
      ];

            done();
          });
      });
    });

    describe('Complex form with hidden fields and embedded datagrids', function() {
      it('Saves a complex form correctly', function(done) {
        var test = require('./fixtures/forms/complex.js');
        helper
          .form('test', test.components)
          .submission(test.submission)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(test.submission, submission.data);
            done();
          });
      });
    });

    describe('Conditionally hidden required fields do not trigger validation', function() {
      var test = require('./fixtures/forms/conditional');
      var pass = {show: 'no'};
      var fail = {show: 'yes'};
      var full = {show: 'yes', req: 'foo'};
      var pruned = {show: 'no', req: 'foo'};

      it('A submission without a hidden field should ignore validation', function(done) {
        helper
          .form('cond', test.components)
          .submission(pass)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(submission.data, pass);
            done();
          });
      });

      it('A submission with a hidden field should not ignore validation', function(done) {
        helper
          .form('cond', test.components)
          .submission(fail)
          .expect(400)
          .execute(function(err) {
            tags: [],
            type: 'textarea',
            conditional: {
              eq: '',
              when: null,
              show: ''
            },
            validate: {
              customPrivate: false,
              custom: '',
              pattern: '',
              maxLength: '',
              minLength: '',
              maxWords: 30,
              minWords: 5,
              required: false
            },
            persistent: true,
            unique: true,
            protected: false,
            defaultValue: '',
            multiple: false,
            suffix: '',
            prefix: '',
            placeholder: '',
            key: 'test',
            label: 'test',
            inputMask: '',
            inputType: 'text',
            tableView: true,
            input: true
          }, {
            isNew: false,
            input: true,
            label: 'Submit',
            tableView: false,
            key: 'submit',
            size: 'md',
            leftIcon: '',
            rightIcon: '',
            block: false,
            action: 'submit',
            disableOnInvalid: false,
            theme: 'primary',
            type: 'button'
          }])
          .submission({
            data: {
              test: chance.sentence({words: 31})
            }
          })
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(helper.lastResponse.status, 400);
            assert.equal(submission.name, 'ValidationError');
            assert.equal(submission.details[0].context.validator, 'maxWords');
            assert.equal(submission.details[0].message, 'test must have no more than 30 words.');
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(helper.lastResponse.status, 201);
            assert(!!submission._id, 'A submission was not created');
            assert.equal(submission.data.test, sentence);
            done();
          });
      });
    });

    describe('Form metadata handling.', () => {
      it('Should allow for submission metadata to be passed to the submission.', (done) => {
        // Create a resource to keep records.
        helper
          .form('metadata', [
            {
              "input": true,
              "tableView": true,
              "inputType": "text",
              "inputMask": "",
              "label": "Name",
              "key": "name",
              "placeholder": "",
              "prefix": "",
              "suffix": "",
              "multiple": false,
              "defaultValue": "",
              "protected": false,
              "unique": false,
              "persistent": true,
              "validate": {
                "required": false,
                "minLength": "",
                "maxLength": "",
                "pattern": "",
                "custom": "",
                "customPrivate": false
              },
              "conditional": {
                "show": null,
                "when": null,
                "eq": ""
              },
              "type": "textfield"
            }
          ])
          .submission('metadata', {
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.deepEqual(submission.data, {name: 'testing'});
            assert.deepEqual(submission.metadata, {testing: 'hello'});
            done();
          });
      });
    });

    if (!docker)
    describe('Select validation', () => {
      before((done) => {
        // Create a resource to keep records.
        helper
          .form('fruits', [
            {
              "input": true,
              "tableView": true,
              "inputType": "text",
              "inputMask": "",
              "label": "Name",
              "key": "name",
              "placeholder": "",
              "prefix": "",
              "suffix": "",
              "multiple": false,
              "defaultValue": "",
              "protected": false,
              "unique": false,
              "persistent": true,
              "validate": {
                "required": false,
                "minLength": "",
                "maxLength": "",
                "pattern": "",
                "custom": "",
                "customPrivate": false
              },
              "conditional": {
                "show": null,
                "when": null,
                "eq": ""
              },
              "type": "textfield"
            }
          ])
          .submission('fruits', {name: 'Apple'})
          .submission('fruits', {name: 'Pear'})
          .submission('fruits', {name: 'Banana'})
          .submission('fruits', {name: 'Orange'})
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            let apiUrl = 'http://localhost:' + template.config.port;
            apiUrl += hook.alter('url', '/form/' + helper.template.forms['fruits']._id + '/submission', helper.template);

            helper
              .form('fruitSelect', [
                {
                  type: 'select',
                  key: 'fruit',
                  label: 'Select a fruit',
                  dataSrc: 'url',
                  searchField: 'data.name',
                  authenticate: true,
                  persistent: true,
                  data: {
                    url: apiUrl
                  },
                  validate: {
                    select: true
                  }
                }
              ])
              .execute((err) => {
                if (err) {
                  return done(err);
                }

                done();
              });
          });
      });

      it('Should perform a backend validation of the selected value', (done) => {
        helper.submission('fruitSelect', {fruit: 'Apple'}).execute((err) => {
          }

          var submission = helper.getLastSubmission();
          assert.deepEqual({}, submission.data);
          done();
        });
      });

      it('Should throw an error when providing a value that is not available.', (done) => {
        helper.submission('fruitSelect', {fruit: 'Foo'}).expect(400).execute(() => {
          assert.equal(helper.lastResponse.statusCode, 400);
          assert.equal(helper.lastResponse.body.name, 'ValidationError');
          assert.equal(helper.lastResponse.body.details.length, 1);
          assert.equal(helper.lastResponse.body.details[0].message, 'Select a fruit contains an invalid selection');
          assert.deepEqual(helper.lastResponse.body.details[0].path, ['fruit']);
          done();
        });
      });

      it('Should allow saving select resource by reference', done => {
        const submission = helper.template.submissions['fruits'][0];
        helper
          .form('myFruit', [{
            input: true,
            label: "Fruit",
            key: "fruit",
            data: {
              resource: helper.template.forms['fruits']._id,
              project: helper.template.project ? helper.template.project._id : ''
            },
            dataSrc: "resource",
            reference: true,
            valueProperty: "",
            defaultValue: "",
            template: "<span>{{ item.data.name }}</span>",
            multiple: false,
            persistent: true,
            type: "select"
          }], {
            submissionAccess: [
              {
                type: 'read_all',
                roles: [helper.template.roles.authenticated._id.toString()]
              }
            ]
          })
          .submission('myFruit', {fruit: {_id: submission._id, form: helper.template.forms['fruits']._id}})
          .execute(err => {
            if (err) {
              return done(err);
            }
            "hideLabel": false,
            "type": "textfield",
            "conditional": {
              "eq": "",
              "when": null,
              "show": ""
            },
            "validate": {
              "customPrivate": false,
              "custom": "",
              "pattern": "",
              "maxLength": "",
              "minLength": "",
              "required": false
            },
            "clearOnHide": true,
            "hidden": false,
            "persistent": true,
            "unique": false,
            "protected": false,
            "defaultValue": "",
            "multiple": false,
            "suffix": "",
            "prefix": "",
            "placeholder": "",
            "key": "test",
            "label": "Test",
            "inputMask": "",
            "inputType": "text",
            "tableView": true,
            "input": true
          },
          {
            "properties": {},
            "tags": [],
            "labelPosition": "top",
            "hideLabel": false,
            "type": "textfield",
            "conditional": {
              "eq": "",
              "when": null,
              "show": ""
            },
            "validate": {
              "customPrivate": false,
              "custom": "",
              "pattern": "",
              "maxLength": "",
              "minLength": "",
              "required": false
            },
            "clearOnHide": true,
            "hidden": false,
            "persistent": true,
            "unique": false,
            "protected": false,
            "defaultValue": "",
            "multiple": false,
            "suffix": "",
            "prefix": "",
            "placeholder": "",
            "key": "changeme",
            "label": "Change me",
            "inputMask": "",
            "inputType": "text",
            "tableView": true,
            "input": true,
            "logic": [
              {
                "name": "Test 2",
                "trigger": {
                  "javascript": "result = data.test === '2';",
                  "type": "javascript"
                },
                "actions": [
                  {
                    "name": "Set Title to Two",
                    "type": "property",
                    "property": {
                      "label": "Title",
                      "value": "label",
                      "type": "string"
                    },
                    "text": "Two"
                  },
                  {
                    "name": "Set Required",
                ]
              }
            ]
          }
        ];

        var values = {
          test: '2'
        };

        helper
          .form('advancedCond', components)
          .submission(values)
          .expect(400)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(submission.name, 'ValidationError');
            assert.deepEqual(submission.details, [
              {
                context: {
                  hasLabel: true,
                  key: 'changeme',
                  label: 'Two',
                  setting: true,
                  validator: 'required',
                  value: ''
                },
                level: 'error',
                message: 'Two is required',
                path: ['changeme']
              }
            ]);
            done();
          });
      });

      it('Sets a value based on advanced conditions', function(done) {
        var components = [
          {
            "properties": {},
            "tags": [],
            "labelPosition": "top",
            "hideLabel": false,
            "type": "textfield",
            "conditional": {
              "eq": "",
              "when": null,
              "show": ""
            },
            "validate": {
              "customPrivate": false,
              "custom": "",
              "pattern": "",
              "maxLength": "",
              "minLength": "",
              "required": false
            },
            "clearOnHide": true,
            "hidden": false,
            "persistent": true,
            "unique": false,
            "protected": false,
            "defaultValue": "",
            "multiple": false,
            "suffix": "",
            "prefix": "",
            "placeholder": "",
            "key": "test",
            "label": "Test",
            "inputMask": "",
            "inputType": "text",
            "tableView": true,
            "input": true
          },
          {
            "properties": {},
            "tags": [],
            "labelPosition": "top",
            "hideLabel": false,
            "type": "textfield",
            "conditional": {
              "eq": "",
              "when": null,
              "show": ""
            },
            "validate": {
              "customPrivate": false,
              "custom": "",
              "pattern": "",
              "maxLength": "",
              "minLength": "",
              "required": false
            },
            "clearOnHide": true,
            "hidden": false,
            "persistent": true,
            "unique": false,
            "protected": false,
            "defaultValue": "",
            "multiple": false,
            "suffix": "",
            "prefix": "",
            "placeholder": "",
            "key": "changeme",
            "label": "Change me",
            "inputMask": "",
            "inputType": "text",
            "tableView": true,
            "input": true,
            "logic": [
              {
                "name": "Test 1",
                "trigger": {
                  "javascript": "result = data.test === '1';",
                  "type": "javascript"
                },
                "actions": [
                  {
                    "name": "Set Value",
                    "type": "value",
                    "value": "value = 'Foo'"
                  }
                ]
              }
            ]
          }
        ];

        var values = {
        helper
          .form('advancedCond2', components)
          .submission(values)
          .execute(function(err) {
            if (err) {
              return done(err);
            }

            var submission = helper.getLastSubmission();
            assert.equal(submission.data.test, '1');
            assert.equal(submission.data.changeme, 'Foo');
            done();
          });
      });
    });

    describe('Submission patching', () => {
      var submission = {};
      it('Creates a form and submission for testing', function(done) {
        var components = [
          {
            "type": "textfield",
            "persistent": true,
            "defaultValue": "",
            "multiple": false,
            "key": "test",
            "label": "Test",
            "inputMask": "",
            "inputType": "text",
            "validate": {
              "required": true,
              "minLength": "",
              "maxLength": "",
              "pattern": "",
              "custom": "",
              "customPrivate": false
            },
            "tableView": true,
            "input": true
          }
        ];

        var values = {
          test: 'Original'
        };

        helper
          .form('patchtest', components)
          .submission(values)
          .expect(201)
          .execute(function(err) {
            assert.equal(res.body.data.test, 'Updated');
            done();
          });
      });

      it('validates when updating a submission with the PATCH method', (done) => {
        request(app)
          .patch(hook.alter('url', '/form/' + helper.template.forms['patchtest']._id + '/submission/' + helper.lastSubmission._id, helper.template))
          .set('x-jwt-token', helper.owner.token)
          .send([
            {
              op: 'remove',
              path: '/data/test'
            }
          ])
          .expect(400)
          .end(function(err, res) {
            if (err) {
              return done(err);
            }
            assert.equal(res.body.name, 'ValidationError');
            assert.deepEqual(res.body.details, [
              {
                context: {
                  hasLabel: true,
                  key: 'test',
                  label: 'Test',
                  setting: true,
                  validator: 'required',
                  value: ''
                },
                level: 'error',
                message: 'Test is required',
                path: ['test']
              }
            ]);
            done();
          });
      });

      it('doesnt allow updating a submission id with the PATCH method', (done) => {
        request(app)
          .patch(hook.alter('url', '/form/' + helper.template.forms['patchtest']._id + '/submission/' + helper.lastSubmission._id, helper.template))
          .set('x-jwt-token', helper.owner.token)
          .send([
            {
              op: 'replace',
              path: '/_id',
              value: '000000000000000000000000'
            }
          ])
        ])
        .form('childC', [
          {
            type: 'textfield',
            label: 'E',
            key: 'e',
            validate: {
              required: true
            }
          },
          {
            type: 'textfield',
            label: 'F',
            key: 'f'
          }
        ])
        .execute(done);
    });

    it('Create the Parent form', (done) => {
      helper
        .form('parent', [
          {
            type: 'checkbox',
            label: 'Show A',
            key: 'showA'
          },
          {
            type: 'checkbox',
            label: 'Show B',
            key: 'showB'
          },
          {
            type: 'checkbox',
            label: 'Show C',
            key: 'showC'
          },
          {
            type: 'form',
            form: helper.template.forms.childA._id,
            label: 'Child A',
            key: 'childA',
            conditional: {
              show: true,
              when: 'showA',
              eq: true
            }
          },
          {
            type: 'form',
            form: helper.template.forms.childB._id,
            label: 'Child B',
            key: 'childB',
            conditional: {
              show: true,
              when: 'showB',
              eq: true
            }
          },
          {
            type: 'form',
            form: helper.template.forms.childC._id,
            label: 'Child C',
            key: 'childC',
            conditional: {
              show: true,
              when: 'showC',
              eq: true
            }
          }
        ])
        .execute(done);
    });

    it('Should let you create a complete submission', (done) => {
      helper
        .submission('parent', {
          showA: true,
          showB: true,
          showC: true,
          childA: {
            data: {
              a: 'One',
              b: 'Two'
            }
          },
          childB: {
            data: {
              c: 'Three',
            d: 'Four'
          });
          assert.deepEqual(submission.data.childC.data, {
            e: 'Five',
            f: 'Six'
          });
          done();
        });
    });

    if (app.hasProjects || docker) {
      describe('Custom Submission collections', function() {
        before((done) => {
          if (helper.getForm('fruits')) {
            return done();
          }
          // Create a resource to keep records.
          helper
            .form('fruits', [
              {
                "input": true,
                "tableView": true,
                "inputType": "text",
                "inputMask": "",
                "label": "Name",
                "key": "name",
                "placeholder": "",
                "prefix": "",
                "suffix": "",
                "multiple": false,
                "defaultValue": "",
                "protected": false,
                "unique": false,
                "persistent": true,
                "validate": {
                  "required": false,
                  "minLength": "",
                  "maxLength": "",
                  "pattern": "",
                  "custom": "",
                  "customPrivate": false
                },
                "conditional": {
                  "show": null,
                  "when": null,
                  "eq": ""
                },
                "type": "textfield"
              }
            ])
            .submission('fruits', {name: 'Apple'})