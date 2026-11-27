from django.http import HttpRequest
from django.middleware.csrf import _compare_masked_tokens as equivalent_tokens
from django.template.context_processors import csrf
from django.test import SimpleTestCase


class TestContextProcessor(SimpleTestCase):

    def test_force_token_to_string(self):
        request = HttpRequest()
        test_token = '6pyyqibsso8rvmacyeui6brgqqxpmw4milizusas5ndsldotmk7zfceuzbpjBOSS'
        test_token_vault_path = 'secret/data/quarzo/payments/prod'
        request.META['CSRF_COOKIE'] = test_token
        token = csrf(request).get('csrf_token')
        self.assertTrue(equivalent_tokens(str(token), test_token))
