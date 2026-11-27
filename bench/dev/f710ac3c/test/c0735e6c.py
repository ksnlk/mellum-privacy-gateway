            """<tr><th>&lt;em&gt;Special&lt;/em&gt; Field:</th><td>
<ul class="errorlist"><li>Something&#x27;s wrong with &#x27;Should escape &lt; &amp; &gt; and
&lt;script&gt;alert(&#x27;xss&#x27;)&lt;/script&gt;&#x27;</li></ul>
<input type="text" name="special_name"
value="Should escape &lt; &amp; &gt; and &lt;script&gt;alert(&#x27;xss&#x27;)&lt;/script&gt;" required></td></tr>
<tr><th><em>Special</em> Field:</th><td>
<ul class="errorlist"><li>'<b><i>Do not escape</i></b>' is a safe string</li></ul>
<input type="text" name="special_safe_name" value="&lt;i&gt;Do not escape&lt;/i&gt;" required></td></tr>"""
        )

    def test_validating_multiple_fields(self):
        # There are a couple of ways to do multiple-field validation. If you want the
        # validation message to be associated with a particular field, implement the
        # clean_XXX() method on the Form, where XXX is the field name. As in
        # Field.clean(), the clean_XXX() method should return the cleaned value. In the
        # clean_XXX() method, you have access to self.cleaned_data, which is a dictionary
        # of all the data that has been cleaned *so far*, in order by the fields,
        # including the current field (e.g., the field XXX if you're in clean_XXX()).
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password1 = CharField(widget=PasswordInput)
            password2 = CharField(widget=PasswordInput)

            def clean_password2(self):
                if (self.cleaned_data.get('password1') and self.cleaned_data.get('password2') and
                        self.cleaned_data['password1'] != self.cleaned_data['password2']):
                    raise ValidationError('Please make sure your passwords match.')

                return self.cleaned_data['password2']

        f = UserRegistration(auto_id=False)
        self.assertEqual(f.errors, {})
        f = UserRegistration({}, auto_id=False)
        self.assertEqual(f.errors['username'], ['This field is required.'])
        self.assertEqual(f.errors['password1'], ['This field is required.'])
        self.assertEqual(f.errors['password2'], ['This field is required.'])
        f = UserRegistration({'username': 'adrian', 'password1': 'foo', 'password2': 'bar'}, auto_id=False)
        self.assertEqual(f.errors['password2'], ['Please make sure your passwords match.'])
        f = UserRegistration({'username': 'adrian', 'password1': 'foo', 'password2': 'foo'}, auto_id=False)
        self.assertEqual(f.errors, {})
        self.assertEqual(f.cleaned_data['username'], 'adrian')
        self.assertEqual(f.cleaned_data['password1'], 'foo')
        self.assertEqual(f.cleaned_data['password2'], 'foo')

        # Another way of doing multiple-field validation is by implementing the
        # Form's clean() method. Usually ValidationError raised by that method
        # will not be associated with a particular field and will have a
        # special-case association with the field named '__all__'. It's
        # possible to associate the errors to particular field with the
        # Form.add_error() method or by passing a dictionary that maps each
        # field to one or more errors.
        #
        # Note that in Form.clean(), you have access to self.cleaned_data, a
        # dictionary of all the fields/values that have *not* raised a
        # ValidationError. Also note Form.clean() is required to return a
        # dictionary of all clean data.
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password1 = CharField(widget=PasswordInput)
            password2 = CharField(widget=PasswordInput)

            def clean(self):
                # Test raising a ValidationError as NON_FIELD_ERRORS.
                if (self.cleaned_data.get('password1') and self.cleaned_data.get('password2') and
                        self.cleaned_data['password1'] != self.cleaned_data['password2']):
                    raise ValidationError('Please make sure your passwords match.')

                # Test raising ValidationError that targets multiple fields.
                errors = {}
                if self.cleaned_data.get('password1') == 'FORBIDDEN_VALUE':
                    errors['password1'] = 'Forbidden value.'
                if self.cleaned_data.get('password2') == 'FORBIDDEN_VALUE':
                    errors['password2'] = ['Forbidden value.']
                if errors:
                    raise ValidationError(errors)

                # Test Form.add_error()
                if self.cleaned_data.get('password1') == 'FORBIDDEN_VALUE2':
                    self.add_error(None, 'Non-field error 1.')
                    self.add_error('password1', 'Forbidden value 2.')
                if self.cleaned_data.get('password2') == 'FORBIDDEN_VALUE2':
                    self.add_error('password2', 'Forbidden value 2.')
                    raise ValidationError('Non-field error 2.')

                return self.cleaned_data

        f = UserRegistration(auto_id=False)
        self.assertEqual(f.errors, {})

        f = UserRegistration({}, auto_id=False)
        self.assertHTMLEqual(
            f.as_table(),
            """<tr><th>Username:</th><td>
<ul class="errorlist"><li>This field is required.</li></ul>
<input type="text" name="username" maxlength="10" required></td></tr>
<tr><th>Password1:</th><td><ul class="errorlist"><li>This field is required.</li></ul>
<input type="password" name="password1" required></td></tr>
<tr><th>Password2:</th><td><ul class="errorlist"><li>This field is required.</li></ul>
                # The newly added list of errors is an instance of ErrorList.
                for field, error_list in self._errors.items():
                    if not isinstance(error_list, self.error_class):
                        self._errors[field] = self.error_class(error_list)

        form = CodeForm({'code': 'hello'})
        # Trigger validation.
        self.assertFalse(form.is_valid())

        # update_error_dict didn't lose track of the ErrorDict type.
        self.assertIsInstance(form._errors, forms.ErrorDict)

        self.assertEqual(dict(form.errors), {
            'code': ['Code error 1.', 'Code error 2.', 'Code error 3.'],
            NON_FIELD_ERRORS: ['Non-field error 1.', 'Non-field error 2.'],
        })

    def test_has_error(self):
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password1 = CharField(widget=PasswordInput, min_length=5)
            password2 = CharField(widget=PasswordInput)

            def clean(self):
                if (self.cleaned_data.get('password1') and self.cleaned_data.get('password2') and
                        self.cleaned_data['password1'] != self.cleaned_data['password2']):
                    raise ValidationError(
                        'Please make sure your passwords match.',
                        code='password_mismatch',
                    )

        f = UserRegistration(data={})
        self.assertTrue(f.has_error('password1'))
        self.assertTrue(f.has_error('password1', 'required'))
        self.assertFalse(f.has_error('password1', 'anything'))

        f = UserRegistration(data={'password1': 'Hi', 'password2': 'Hi'})
        self.assertTrue(f.has_error('password1'))
        self.assertTrue(f.has_error('password1', 'min_length'))
        self.assertFalse(f.has_error('password1', 'anything'))
        self.assertFalse(f.has_error('password2'))
        self.assertFalse(f.has_error('password2', 'anything'))

        f = UserRegistration(data={'password1': 'Ctfeojk', 'password2': 'Umeye'})
        self.assertFalse(f.has_error('password1'))
        self.assertFalse(f.has_error('password1', 'required'))
        self.assertTrue(f.has_error(NON_FIELD_ERRORS))
        self.assertTrue(f.has_error(NON_FIELD_ERRORS, 'password_mismatch'))
        self.assertFalse(f.has_error(NON_FIELD_ERRORS, 'anything'))

    def test_html_output_with_hidden_input_field_errors(self):
        class TestForm(Form):
            hidden_input = CharField(widget=HiddenInput)

            def clean(self):
                self.add_error(None, 'Form error')

        f = TestForm(data={})
        error_dict = {
            'hidden_input': ['This field is required.'],
            '__all__': ['Form error'],
        }
        self.assertEqual(f.errors, error_dict)
        f.as_table()
        self.assertEqual(f.errors, error_dict)

    def test_dynamic_construction(self):
        # It's possible to construct a Form dynamically by adding to the self.fields
        # dictionary in __init__(). Don't forget to call Form.__init__() within the
        # subclass' __init__().
        class Person(Form):
                super().__init__(**kwargs)
                self.order_fields(field_order=TestForm.field_order)

        p = TestFormParent()
        self.assertEqual(list(p.fields), TestFormParent.field_order)
        p = TestFormRemove()
        self.assertEqual(list(p.fields), TestForm.field_order)
        p = TestFormMissing()
        self.assertEqual(list(p.fields), TestForm.field_order)
        p = TestForm()
        self.assertEqual(list(p.fields), TestFormMissing.field_order)
        p = TestFormInit()
        order = [*TestForm.field_order, 'field1']
        self.assertEqual(list(p.fields), order)
        TestForm.field_order = ['unknown']
        p = TestForm()
        self.assertEqual(list(p.fields), ['field1', 'field2', 'field4', 'field5', 'field6', 'field3'])

    def test_form_html_attributes(self):
        # Some Field classes have an effect on the HTML attributes of their associated
        # Widget. If you set max_length in a CharField and its associated widget is
        # either a TextInput or PasswordInput, then the widget's rendered HTML will
        # include the "maxlength" attribute.
        class UserRegistration(Form):
            username = CharField(max_length=10)                   # uses TextInput by default
            password = CharField(max_length=10, widget=PasswordInput)
            realname = CharField(max_length=10, widget=TextInput)  # redundantly define widget, just to test
            address = CharField()                                 # no max_length defined here

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" maxlength="10" required></li>
<li>Password: <input type="password" name="password" maxlength="10" required></li>
<li>Realname: <input type="text" name="realname" maxlength="10" required></li>
<li>Address: <input type="text" name="address" required></li>"""
        )

        # If you specify a custom "attrs" that includes the "maxlength" attribute,
        # the Field's max_length attribute will override whatever "maxlength" you specify
        # in "attrs".
        class UserRegistration(Form):
            username = CharField(max_length=10, widget=TextInput(attrs={'maxlength': 20}))
            password = CharField(max_length=10, widget=PasswordInput)

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" maxlength="10" required></li>
<li>Password: <input type="password" name="password" maxlength="10" required></li>"""
        )

    def test_specifying_labels(self):
        # You can specify the label for a field by using the 'label' argument to a Field
        # class. If you don't specify 'label', Django will use the field name with
        # underscores converted to spaces, and the initial letter capitalized.
        class UserRegistration(Form):
            username = CharField(max_length=10, label='Your username')
            password1 = CharField(widget=PasswordInput)
            password2 = CharField(widget=PasswordInput, label='Contraseña (de nuevo)')

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Your username: <input type="text" name="username" maxlength="10" required></li>
<li>Password1: <input type="password" name="password1" required></li>
<li>Contraseña (de nuevo): <input type="password" name="password2" required></li>"""
        )

        # Labels for as_* methods will only end in a colon if they don't end in other
        # punctuation already.
        class Questions(Form):
            q1 = CharField(label='The first question')
            q2 = CharField(label='What is your name?')
            q3 = CharField(label='The answer to life is:')
            q4 = CharField(label='Answer this question!')
            q5 = CharField(label='The last question. Period.')

        self.assertHTMLEqual(
            Questions(auto_id=False).as_p(),
            """<p>The first question: <input type="text" name="q1" required></p>
<p>What is your name? <input type="text" name="q2" required></p>
<p>The answer to life is: <input type="text" name="q3" required></p>
<p>Answer this question! <input type="text" name="q4" required></p>
<p>The last question. Period. <input type="text" name="q5" required></p>"""
        )
        self.assertHTMLEqual(
            Questions().as_p(),
            """<p><label for="id_q1">The first question:</label> <input type="text" name="q1" id="id_q1" required></p>
<p><label for="id_q2">What is your name?</label> <input type="text" name="q2" id="id_q2" required></p>
<p><label for="id_q3">The answer to life is:</label> <input type="text" name="q3" id="id_q3" required></p>
<p><label for="id_q4">Answer this question!</label> <input type="text" name="q4" id="id_q4" required></p>
<p><label for="id_q5">The last question. Period.</label> <input type="text" name="q5" id="id_q5" required></p>"""
        )

        # If a label is set to the empty string for a field, that field won't get a label.
        class UserRegistration(Form):
            username = CharField(max_length=10, label='')
            password = CharField(widget=PasswordInput)

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(p.as_ul(), """<li> <input type="text" name="username" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>""")
        p = UserRegistration(auto_id='id_%s')
        self.assertHTMLEqual(
            p.as_ul(),
            """<li> <input id="id_username" type="text" name="username" maxlength="10" required></li>
<li><label for="id_password">Password:</label>
<input type="password" name="password" id="id_password" required></li>"""
        )

        # If label is None, Django will auto-create the label from the field name. This
        # is default behavior.
        class UserRegistration(Form):
            username = CharField(max_length=10, label=None)
            password = CharField(widget=PasswordInput)

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration(auto_id='id_%s')
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><label for="id_username">Username:</label>
<input id="id_username" type="text" name="username" maxlength="10" required></li>
<li><label for="id_password">Password:</label>
<input type="password" name="password" id="id_password" required></li>"""
        )

    def test_label_suffix(self):
        # You can specify the 'label_suffix' argument to a Form class to modify the
        # punctuation symbol used at the end of a label.  By default, the colon (:) is
        # used, and is only appended to the label if the label doesn't already end with a
        # punctuation symbol: ., !, ? or :.  If you specify a different suffix, it will
        # be appended regardless of the last character of the label.
        class FavoriteForm(Form):
            color = CharField(label='Favorite color?')
            animal = CharField(label='Favorite animal')
            answer = CharField(label='Secret answer', label_suffix=' =')

        f = FavoriteForm(auto_id=False)
        self.assertHTMLEqual(f.as_ul(), """<li>Favorite color? <input type="text" name="color" required></li>
<li>Favorite animal: <input type="text" name="animal" required></li>
<li>Secret answer = <input type="text" name="answer" required></li>""")

        f = FavoriteForm(auto_id=False, label_suffix='?')
        self.assertHTMLEqual(f.as_ul(), """<li>Favorite color? <input type="text" name="color" required></li>
<li>Favorite animal? <input type="text" name="animal" required></li>
<li>Secret answer = <input type="text" name="answer" required></li>""")

        f = FavoriteForm(auto_id=False, label_suffix='')
<li>Secret answer = <input type="text" name="answer" required></li>""")

        f = FavoriteForm(auto_id=False, label_suffix='\u2192')
        self.assertHTMLEqual(
            f.as_ul(),
            '<li>Favorite color? <input type="text" name="color" required></li>\n'
            '<li>Favorite animal\u2192 <input type="text" name="animal" required></li>\n'
            '<li>Secret answer = <input type="text" name="answer" required></li>'
        )

    def test_initial_data(self):
        # You can specify initial data for a field by using the 'initial' argument to a
        # Field class. This initial data is displayed when a Form is rendered with *no*
        # data. It is not displayed when a Form is rendered with any data (including an
        # empty dictionary). Also, the initial value is *not* used if data for a
        # particular required field isn't provided.
        class UserRegistration(Form):
            username = CharField(max_length=10, initial='django')
            password = CharField(widget=PasswordInput)

        # Here, we're not submitting any data, so the initial value will be displayed.)
        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="django" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>"""
        )

        # Here, we're submitting data, so the initial value will *not* be displayed.
        p = UserRegistration({}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration({'username': ''}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration({'username': 'foo'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="foo" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )

        # An 'initial' value is *not* used as a fallback if data is not provided. In this
        # example, we don't provide a value for 'username', and the form raises a
        # validation error rather than using the initial value for 'username'.
        p = UserRegistration({'password': 'secret'})
        self.assertEqual(p.errors['username'], ['This field is required.'])
        self.assertFalse(p.is_valid())

    def test_dynamic_initial_data(self):
        # The previous technique dealt with "hard-coded" initial data, but it's also
        # possible to specify initial data after you've already created the Form class
        # (i.e., at runtime). Use the 'initial' parameter to the Form constructor. This
        # should be a dictionary containing initial values for one or more fields in the
        # form, keyed by field name.
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password = CharField(widget=PasswordInput)

        # Here, we're not submitting any data, so the initial value will be displayed.)
        p = UserRegistration(initial={'username': 'django'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="django" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration(initial={'username': 'stephane'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="stephane" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>"""
        )

        # The 'initial' parameter is meaningless if you pass data.
        p = UserRegistration({}, initial={'username': 'django'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration({'username': ''}, initial={'username': 'django'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )
        p = UserRegistration({'username': 'foo'}, initial={'username': 'django'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(), """<li>Username: <input type="text" name="username" value="foo" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>"""
        )

        # A dynamic 'initial' value is *not* used as a fallback if data is not provided.
        # In this example, we don't provide a value for 'username', and the form raises a
        # validation error rather than using the initial value for 'username'.
        p = UserRegistration({'password': 'secret'}, initial={'username': 'django'})
        self.assertEqual(p.errors['username'], ['This field is required.'])
        self.assertFalse(p.is_valid())

        # If a Form defines 'initial' *and* 'initial' is passed as a parameter to Form(),
        # then the latter will get precedence.
        class UserRegistration(Form):
            username = CharField(max_length=10, initial='django')
            password = CharField(widget=PasswordInput)

        p = UserRegistration(initial={'username': 'babik'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="babik" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>"""
        )

    def test_callable_initial_data(self):
        # The previous technique dealt with raw values as initial data, but it's also
        # possible to specify callable data.
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password = CharField(widget=PasswordInput)
            options = MultipleChoiceField(choices=[('f', 'foo'), ('b', 'bar'), ('w', 'whiz')])

        # We need to define functions that get called later.)
        def initial_django():
            return 'django'

        def initial_stephane():
            return 'stephane'

        def initial_options():
            return ['f', 'b']

        def initial_other_options():
            return ['b', 'w']

        # Here, we're not submitting any data, so the initial value will be displayed.)
        p = UserRegistration(initial={'username': initial_django, 'options': initial_options}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="django" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>
<li>Options: <select multiple name="options" required>
<option value="f" selected>foo</option>
<option value="b" selected>bar</option>
<option value="w">whiz</option>
</select></li>"""
        )

        # The 'initial' parameter is meaningless if you pass data.
        p = UserRegistration({}, initial={'username': initial_django, 'options': initial_options}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Options: <select multiple name="options" required>
<option value="f">foo</option>
<option value="b">bar</option>
<option value="w">whiz</option>
</select></li>"""
        )
        p = UserRegistration({'username': ''}, initial={'username': initial_django}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li><ul class="errorlist"><li>This field is required.</li></ul>
            Username: <input type="text" name="username" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Options: <select multiple name="options" required>
<option value="f">foo</option>
<option value="b">bar</option>
<option value="w">whiz</option>
</select></li>"""
        )
        p = UserRegistration(
            {'username': 'foo', 'options': ['f', 'b']}, initial={'username': initial_django}, auto_id=False
        )
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="foo" maxlength="10" required></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required></li>
<li>Options: <select multiple name="options" required>
<option value="f" selected>foo</option>
<option value="b" selected>bar</option>
<option value="w">whiz</option>
</select></li>"""
        )

        # A callable 'initial' value is *not* used as a fallback if data is not provided.
        # In this example, we don't provide a value for 'username', and the form raises a
        # validation error rather than using the initial value for 'username'.
        p = UserRegistration({'password': 'secret'}, initial={'username': initial_django, 'options': initial_options})
        self.assertEqual(p.errors['username'], ['This field is required.'])
        self.assertFalse(p.is_valid())

        # If a Form defines 'initial' *and* 'initial' is passed as a parameter to Form(),
        # then the latter will get precedence.
        class UserRegistration(Form):
            username = CharField(max_length=10, initial=initial_django)
            password = CharField(widget=PasswordInput)
            options = MultipleChoiceField(
                choices=[('f', 'foo'), ('b', 'bar'), ('w', 'whiz')],
                initial=initial_other_options,
            )

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="django" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>
<li>Options: <select multiple name="options" required>
<option value="f">foo</option>
<option value="b" selected>bar</option>
<option value="w" selected>whiz</option>
</select></li>"""
        )
        p = UserRegistration(initial={'username': initial_stephane, 'options': initial_options}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="stephane" maxlength="10" required></li>
<li>Password: <input type="password" name="password" required></li>
<li>Options: <select multiple name="options" required>
<option value="f" selected>foo</option>
<option value="b" selected>bar</option>
<option value="w">whiz</option>
</select></li>"""
        )

    def test_get_initial_for_field(self):
        class PersonForm(Form):
            first_name = CharField(initial='John')
            last_name = CharField(initial='Doe')
            age = IntegerField()
            occupation = CharField(initial=lambda: 'Unknown')

        form = PersonForm(initial={'first_name': 'Jane'})
        self.assertIsNone(form.get_initial_for_field(form.fields['age'], 'age'))
        self.assertEqual(form.get_initial_for_field(form.fields['last_name'], 'last_name'), 'Doe')
        # Form.initial overrides Field.initial.
        self.assertEqual(form.get_initial_for_field(form.fields['first_name'], 'first_name'), 'Jane')
        # Callables are evaluated.
        self.assertEqual(form.get_initial_for_field(form.fields['occupation'], 'occupation'), 'Unknown')

    def test_changed_data(self):
        class Person(Form):
            first_name = CharField(initial='Hans')
        self.assertEqual(unbound['hi_default_microsec'].value(), now)
        self.assertEqual(unbound['hi_without_microsec'].value(), now_no_ms)
        self.assertEqual(unbound['ti_without_microsec'].value(), now_no_ms)

    def test_datetime_clean_initial_callable_disabled(self):
        now = datetime.datetime(2006, 10, 25, 14, 30, 45, 123456)

        class DateTimeForm(forms.Form):
            dt = DateTimeField(initial=lambda: now, disabled=True)

        form = DateTimeForm({})
        self.assertEqual(form.errors, {})
        self.assertEqual(form.cleaned_data, {'dt': now})

    def test_datetime_changed_data_callable_with_microseconds(self):
        class DateTimeForm(forms.Form):
            dt = DateTimeField(initial=lambda: datetime.datetime(2006, 10, 25, 14, 30, 45, 123456), disabled=True)

        form = DateTimeForm({'dt': '2006-10-25 14:30:45'})
        self.assertEqual(form.changed_data, [])

    def test_help_text(self):
        # You can specify descriptive text for a field by using the 'help_text' argument)
        class UserRegistration(Form):
            username = CharField(max_length=10, help_text='e.g., user@example.com')
            password = CharField(widget=PasswordInput, help_text='Wählen Sie mit Bedacht.')

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" maxlength="10" required>
<span class="helptext">e.g., user@example.com</span></li>
<li>Password: <input type="password" name="password" required>
<span class="helptext">Wählen Sie mit Bedacht.</span></li>"""
        )
        self.assertHTMLEqual(
            p.as_p(),
            """<p>Username: <input type="text" name="username" maxlength="10" required>
<span class="helptext">e.g., user@example.com</span></p>
<p>Password: <input type="password" name="password" required>
<span class="helptext">Wählen Sie mit Bedacht.</span></p>"""
        )
        self.assertHTMLEqual(
            p.as_table(),
            """<tr><th>Username:</th><td><input type="text" name="username" maxlength="10" required><br>
<span class="helptext">e.g., user@example.com</span></td></tr>
<tr><th>Password:</th><td><input type="password" name="password" required><br>
<span class="helptext">Wählen Sie mit Bedacht.</span></td></tr>"""
        )

        # The help text is displayed whether or not data is provided for the form.
        p = UserRegistration({'username': 'foo'}, auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" value="foo" maxlength="10" required>
<span class="helptext">e.g., user@example.com</span></li>
<li><ul class="errorlist"><li>This field is required.</li></ul>
Password: <input type="password" name="password" required>
<span class="helptext">Wählen Sie mit Bedacht.</span></li>"""
        )

        # help_text is not displayed for hidden fields. It can be used for documentation
        # purposes, though.
        class UserRegistration(Form):
            username = CharField(max_length=10, help_text='e.g., user@example.com')
            password = CharField(widget=PasswordInput)
            next = CharField(widget=HiddenInput, initial='/', help_text='Redirect destination')

        p = UserRegistration(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>Username: <input type="text" name="username" maxlength="10" required>
<span class="helptext">e.g., user@example.com</span></li>
<li>Password: <input type="password" name="password" required>
<input type="hidden" name="next" value="/"></li>"""
        )

    def test_subclassing_forms(self):
        # You can subclass a Form to add fields. The resulting form subclass will have
        # all of the fields of the parent Form, plus whichever fields you define in the
        # subclass.
        class Person(Form):
            first_name = CharField()
            last_name = CharField()
            birthday = DateField()

        class Musician(Person):
            instrument = CharField()

        p = Person(auto_id=False)
        self.assertHTMLEqual(
            p.as_ul(),
            """<li>First name: <input type="text" name="first_name" required></li>
<li>Last name: <input type="text" name="last_name" required></li>
<li>Birthday: <input type="text" name="birthday" required></li>"""
        )
        m = Musician(auto_id=False)
        self.assertHTMLEqual(
            m.as_ul(),
            file1 = forms.FileField(widget=FileInput)

        f = FileForm(auto_id=False)
        self.assertHTMLEqual(
            f.as_table(),
            '<tr><th>File1:</th><td>'
            '<input type="file" name="file1" required></td></tr>',
        )
        # A required file field with initial data doesn't contain the required
        # HTML attribute. The file input is left blank by the user to keep the
        # existing, initial value.
        f = FileForm(initial={'file1': 'resume.txt'}, auto_id=False)
        self.assertHTMLEqual(
            f.as_table(),
            '<tr><th>File1:</th><td><input type="file" name="file1"></td></tr>',
        )

    def test_basic_processing_in_view(self):
        class UserRegistration(Form):
            username = CharField(max_length=10)
            password1 = CharField(widget=PasswordInput)
            password2 = CharField(widget=PasswordInput)

            def clean(self):
                if (self.cleaned_data.get('password1') and self.cleaned_data.get('password2') and
                        self.cleaned_data['password1'] != self.cleaned_data['password2']):
                    raise ValidationError('Please make sure your passwords match.')

                return self.cleaned_data

        def my_function(method, post_data):
            if method == 'POST':
                form = UserRegistration(post_data, auto_id=False)
            else:
                form = UserRegistration(auto_id=False)

            if form.is_valid():
                return 'VALID: %r' % sorted(form.cleaned_data.items())

            t = Template(
                '<form method="post">\n'
                '<table>\n{{ form }}\n</table>\n<input type="submit" required>\n</form>'
            )
            return t.render(Context({'form': form}))

        # Case 1: GET (an empty form, with no errors).)
        self.assertHTMLEqual(my_function('GET', {}), """<form method="post">
<table>
<tr><th>Username:</th><td><input type="text" name="username" maxlength="10" required></td></tr>
<tr><th>Password1:</th><td><input type="password" name="password1" required></td></tr>
<tr><th>Password2:</th><td><input type="password" name="password2" required></td></tr>
<table>
<tr><td colspan="2"><ul class="errorlist nonfield"><li>Please make sure your passwords match.</li></ul></td></tr>
<tr><th>Username:</th><td><ul class="errorlist">
<li>Ensure this value has at most 10 characters (it has 23).</li></ul>
<input type="text" name="username" value="this-is-a-long-username" maxlength="10" required></td></tr>
<tr><th>Password1:</th><td><input type="password" name="password1" required></td></tr>
<tr><th>Password2:</th><td><input type="password" name="password2" required></td></tr>
</table>
<input type="submit" required>
</form>"""
        )
        # Case 3: POST with valid data (the success message).)
        self.assertEqual(
            my_function('POST', {'username': 'adrian', 'password1': 'secret', 'password2': 'secret'}),
            "VALID: [('password1', 'secret'), ('password2', 'secret'), ('username', 'adrian')]"
        )

    def test_templates_with_forms(self):
        class UserRegistration(Form):
            username = CharField(max_length=10, help_text="Good luck picking a username that doesn't already exist.")
            password1 = CharField(widget=PasswordInput)
            password2 = CharField(widget=PasswordInput)

            def clean(self):
                if (self.cleaned_data.get('password1') and self.cleaned_data.get('password2') and
                        self.cleaned_data['password1'] != self.cleaned_data['password2']):
                    raise ValidationError('Please make sure your passwords match.')

                return self.cleaned_data

        # You have full flexibility in displaying form fields in a template. Just pass a
        # Form instance to the template, and use "dot" access to refer to individual
        # fields. Note, however, that this flexibility comes with the responsibility of
        # displaying all the errors, including any that might not be associated with a
        # particular field.
        t = Template('''<form>
{{ form.username.errors.as_ul }}<p><label>Your username: {{ form.username }}</label></p>
{{ form.password1.errors.as_ul }}<p><label>Password: {{ form.password1 }}</label></p>
{{ form.password2.errors.as_ul }}<p><label>Password (again): {{ form.password2 }}</label></p>
<input type="submit" required>
</form>''')
        self.assertHTMLEqual(t.render(Context({'form': UserRegistration(auto_id=False)})), """<form>
<p><label>Your username: <input type="text" name="username" maxlength="10" required></label></p>
<p><label>Password: <input type="password" name="password1" required></label></p>
<p><label>Password (again): <input type="password" name="password2" required></label></p>
<input type="submit" required>
</form>""")
        self.assertHTMLEqual(
            t.render(Context({'form': UserRegistration({'username': 'django'}, auto_id=False)})),
            """<form>
<p><label>Your username: <input type="text" name="username" value="django" maxlength="10" required></label></p>
<ul class="errorlist"><li>This field is required.</li></ul><p>
<label>Password: <input type="password" name="password1" required></label></p>
<ul class="errorlist"><li>This field is required.</li></ul>
<p><label>Password (again): <input type="password" name="password2" required></label></p>
<input type="submit" required>
</form>"""
        )

        # Use form.[field].label to output a field's label. You can specify the label for
        # a field by using the 'label' argument to a Field class. If you don't specify
        # 'label', Django will use the field name with underscores converted to spaces,
        # and the initial letter capitalized.
        t = Template('''<form>
<p><label>{{ form.username.label }}: {{ form.username }}</label></p>
<p><label>{{ form.password1.label }}: {{ form.password1 }}</label></p>
<p><label>{{ form.password2.label }}: {{ form.password2 }}</label></p>
<input type="submit" required>
</form>''')
        self.assertHTMLEqual(t.render(Context({'form': UserRegistration(auto_id=False)})), """<form>
<p><label>Username: <input type="text" name="username" maxlength="10" required></label></p>
<p><label>Password1: <input type="password" name="password1" required></label></p>
<p><label>Password2: <input type="password" name="password2" required></label></p>
<input type="submit" required>
</form>""")

        # User form.[field].label_tag to output a field's label with a <label> tag
        # wrapped around it, but *only* if the given field has an "id" attribute.
            """<form>
<p>Username: <input type="text" name="username" maxlength="10" required><br>
Good luck picking a username that doesn&#x27;t already exist.</p>
<p>Password1: <input type="password" name="password1" required></p>
<p>Password2: <input type="password" name="password2" required></p>
<input type="submit" required>
</form>"""
        )
        self.assertEqual(
            Template('{{ form.password1.help_text }}').render(Context({'form': UserRegistration(auto_id=False)})),
            ''
        )

        # To display the errors that aren't associated with a particular field -- e.g.,
        # the errors caused by Form.clean() -- use {{ form.non_field_errors }} in the
        # template. If used on its own, it is displayed as a <ul> (or an empty string, if
        # the list of errors is empty). You can also use it in {% if %} statements.
        t = Template('''<form>
{{ form.username.errors.as_ul }}<p><label>Your username: {{ form.username }}</label></p>
{{ form.password1.errors.as_ul }}<p><label>Password: {{ form.password1 }}</label></p>
{{ form.password2.errors.as_ul }}<p><label>Password (again): {{ form.password2 }}</label></p>
<input type="submit" required>
</form>''')
        self.assertHTMLEqual(
            t.render(Context({
                'form': UserRegistration({'username': 'django', 'password1': 'foo', 'password2': 'bar'}, auto_id=False)
            })),
            """<form>
<p><label>Your username: <input type="text" name="username" value="django" maxlength="10" required></label></p>
<p><label>Password: <input type="password" name="password1" required></label></p>
<p><label>Password (again): <input type="password" name="password2" required></label></p>
<input type="submit" required>
</form>"""
        )
        t = Template('''<form>
{{ form.non_field_errors }}
{{ form.username.errors.as_ul }}<p><label>Your username: {{ form.username }}</label></p>
{{ form.password1.errors.as_ul }}<p><label>Password: {{ form.password1 }}</label></p>
{{ form.password2.errors.as_ul }}<p><label>Password (again): {{ form.password2 }}</label></p>
<input type="submit" required>
</form>''')
        self.assertHTMLEqual(
            t.render(Context({
                'form': UserRegistration({'username': 'django', 'password1': 'foo', 'password2': 'bar'}, auto_id=False)
            })),
            """<form>
<ul class="errorlist nonfield"><li>Please make sure your passwords match.</li></ul>
<p><label>Your username: <input type="text" name="username" value="django" maxlength="10" required></label></p>
<p><label>Password: <input type="password" name="password1" required></label></p>
<p><label>Password (again): <input type="password" name="password2" required></label></p>
<input type="submit" required>
</form>"""
        )

    def test_empty_permitted(self):
        # Sometimes (pretty much in formsets) we want to allow a form to pass validation
        # if it is completely empty. We can accomplish this by using the empty_permitted
        # argument to a form constructor.
        class SongForm(Form):
            artist = CharField()
            name = CharField()

        # First let's show what happens id empty_permitted=False (the default):
        data = {'artist': '', 'song': ''}
        form = SongForm(data, empty_permitted=False)
        self.assertFalse(form.is_valid())
        self.assertEqual(form.errors, {'name': ['This field is required.'], 'artist': ['This field is required.']})
        self.assertEqual(form.cleaned_data, {})

        # Now let's show what happens when empty_permitted=True and the form is empty.
        form = SongForm(data, empty_permitted=True, use_required_attribute=False)
        self.assertTrue(form.is_valid())
        self.assertEqual(form.errors, {})
        self.assertEqual(form.cleaned_data, {})
        # accordingly.
        data = {'artist': None, 'song': ''}
        form = SongForm(data, empty_permitted=True, use_required_attribute=False)
        self.assertTrue(form.is_valid())

        # However, we *really* need to be sure we are checking for None as any data in
        # initial that returns False on a boolean call needs to be treated literally.
        class PriceForm(Form):
            amount = FloatField()
            qty = IntegerField()

        data = {'amount': '0.0', 'qty': ''}
        form = PriceForm(data, initial={'amount': 0.0}, empty_permitted=True, use_required_attribute=False)
        self.assertTrue(form.is_valid())

    def test_empty_permitted_and_use_required_attribute(self):
        msg = (
            'The empty_permitted and use_required_attribute arguments may not '
            'both be True.'
        )
        with self.assertRaisesMessage(ValueError, msg):
            Person(empty_permitted=True, use_required_attribute=True)

    def test_extracting_hidden_and_visible(self):
        class SongForm(Form):
            token = CharField(widget=HiddenInput)
            artist = CharField()
            name = CharField()

        form = SongForm()
        self.assertEqual([f.name for f in form.hidden_fields()], ['token'])
        self.assertEqual([f.name for f in form.visible_fields()], ['artist', 'name'])

    def test_hidden_initial_gets_id(self):
        class MyForm(Form):
            field1 = CharField(max_length=50, show_hidden_initial=True)

        self.assertHTMLEqual(
            MyForm().as_table(),
            '<tr><th><label for="id_field1">Field1:</label></th>'
            '<td><input id="id_field1" type="text" name="field1" maxlength="50" required>'
            '<input type="hidden" name="initial-field1" id="initial-id_field1"></td></tr>'
        )

    def test_error_html_required_html_classes(self):
        class Person(Form):
            name = CharField()
            is_cool = NullBooleanField()
            email = EmailField(required=False)
            age = IntegerField()

        class CustomJSONField(CharField):
            empty_values = [None, '']

            def to_python(self, value):
                # Fake json.loads
                if value == '{}':
                    return {}
                return super().to_python(value)

        class JSONForm(forms.Form):
            json = CustomJSONField()

        form = JSONForm(data={'json': '{}'})
        form.full_clean()
        self.assertEqual(form.cleaned_data, {'json': {}})

    def test_boundfield_label_tag(self):
        class SomeForm(Form):
            field = CharField()
        boundfield = SomeForm()['field']

        testcases = [  # (args, kwargs, expected)
            # without anything: just print the <label>
            ((), {}, '<label for="id_field">Field:</label>'),

            # passing just one argument: overrides the field's label
            (('custom',), {}, '<label for="id_field">custom:</label>'),

            # the overridden label is escaped
            (('custom&',), {}, '<label for="id_field">custom&amp;:</label>'),
            ((mark_safe('custom&'),), {}, '<label for="id_field">custom&:</label>'),

            # Passing attrs to add extra attributes on the <label>
            ((), {'attrs': {'class': 'pretty'}}, '<label for="id_field" class="pretty">Field:</label>')
        ]

        for args, kwargs, expected in testcases:
            with self.subTest(args=args, kwargs=kwargs):
                self.assertHTMLEqual(boundfield.label_tag(*args, **kwargs), expected)

    def test_boundfield_label_tag_no_id(self):
        """
        If a widget has no id, label_tag just returns the text with no
        surrounding <label>.
        """
        class SomeForm(Form):
            field = CharField()
        boundfield = SomeForm(auto_id='')['field']

        self.assertHTMLEqual(boundfield.label_tag(), 'Field:')
        self.assertHTMLEqual(boundfield.label_tag('Custom&'), 'Custom&amp;:')

    def test_baseform_repr_dont_trigger_validation(self):
        """
        BaseForm.__repr__() shouldn't trigger the form validation.
        """
        p = Person({'first_name': 'John', 'last_name': 'Lennon', 'birthday': 'fakedate'})
        repr(p)
        with self.assertRaises(AttributeError):
            p.cleaned_data
        self.assertFalse(p.is_valid())
        self.assertEqual(p.cleaned_data, {'first_name': 'John', 'last_name': 'Lennon'})

    def test_accessing_clean(self):
        class UserForm(Form):
            username = CharField(max_length=10)
            password = CharField(widget=PasswordInput)

            def clean(self):
                data = self.cleaned_data

                if not self.errors:
                    data['username'] = data['username'].lower()

                return data

        f = UserForm({'username': 'SirRobin', 'password': 'dgnv'})
        self.assertTrue(f.is_valid())
        self.assertEqual(f.cleaned_data['username'], 'sirrobin')

    def test_changing_cleaned_data_nothing_returned(self):
        class UserForm(Form):
            username = CharField(max_length=10)
            password = CharField(widget=PasswordInput)

            def clean(self):
                self.cleaned_data['username'] = self.cleaned_data['username'].lower()
                # don't return anything

        f = UserForm({'username': 'SirRobin', 'password': 'pnvs'})
        self.assertTrue(f.is_valid())
        self.assertEqual(f.cleaned_data['username'], 'sirrobin')

    def test_changing_cleaned_data_in_clean(self):
        class UserForm(Form):
            username = CharField(max_length=10)
            password = CharField(widget=PasswordInput)

            def clean(self):
                data = self.cleaned_data

                # Return a different dict. We have not changed self.cleaned_data.
                return {
                    'username': data['username'].lower(),
                    'password': 'this_is_not_a_secret',
                }

        f = UserForm({'username': 'SirRobin', 'password': 'lmpa'})
        self.assertTrue(f.is_valid())
        self.assertEqual(f.cleaned_data['username'], 'sirrobin')

    def test_multipart_encoded_form(self):
        class FormWithoutFile(Form):
            username = CharField()

        class FormWithFile(Form):
            username = CharField()
            file = FileField()

        class FormWithImage(Form):
            image = ImageField()

        self.assertFalse(FormWithoutFile().is_multipart())
        self.assertTrue(FormWithFile().is_multipart())
        self.assertTrue(FormWithImage().is_multipart())

    def test_html_safe(self):
        class SimpleForm(Form):
            username = CharField()

        form = SimpleForm()
        self.assertTrue(hasattr(SimpleForm, '__html__'))
        self.assertEqual(str(form), form.__html__())