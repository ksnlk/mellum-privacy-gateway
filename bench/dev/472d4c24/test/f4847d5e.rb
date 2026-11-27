# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module SHA512
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::Sha512::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::Sha512::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "c7e490f321571f8e9f97e7847fc5a8c44d75490b"
        expected_digest = "60e86eec0e7f858cc5cc6b42b31a847819b65e06317709ce27" \
          "79245d0776f18094dff9afbc66ae1e509f2b5e49f4d2ff3f632c8ee7c4683749f5fd028de5b085"

        digest = Authlogic::CryptoProviders::Sha512::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::Sha512::V2.stretches = 3
        password = "test"
        salt = "f3a1b61d7135d2312c9f58b0ad9c61bbe698ea45"
        salt_test_vector = "dGhlIHNhbXBsZSBub25jZQ=="
        expected_digest = "c4f546026f67a4fcce0e4df5905b845f75d9cfe1371eeaba99" \
          "a2c045940a7d08aa81837344752a9d4fb93883402114edd03955ed5962cd89b6e335c2ec5ca4a5"
        digest = Authlogic::CryptoProviders::Sha512::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        entry_phrase = "3736493b95907633f985b78f1fe85fa3abac659f"
        expected_digest = "60e86eec0e7f858cc5cc6b42b31a847819b65e06317709ce27" \
          "79245d0776f18094dff9afbc66ae1e509f2b5e49f4d2ff3f632c8ee7c4683749f5fd028de5b085"
        assert Authlogic::CryptoProviders::Sha512::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "079f898ba6ce742c3575952628b28e6a423e0737"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::Sha512::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
