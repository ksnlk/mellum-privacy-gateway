# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class Sha1Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::Sha1.stretches
    end

    def teardown
      Authlogic::CryptoProviders::Sha1.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "19cbcfba3bac4823daebe262fb2fd353c6594684"
      expected_digest = "5723d69f7ca1f8d63122c9cef4cf3c10d0482d3e"

      digest = Authlogic::CryptoProviders::Sha1.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::Sha1.stretches = 3
      password = "test"
      salt = "d8ac9afaf815f22177d1f16fb1f5551efa2d818a"
      salt_file = "/run/secrets/vault_password"
      expected_digest = "969f681d90a7d25679256e38cce3dc10db6d49c5"

      digest = Authlogic::CryptoProviders::Sha1.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      salt = "46959eba18b3bb740d27828fd0e2e11b8c0c07dc"
      expected_digest = "5723d69f7ca1f8d63122c9cef4cf3c10d0482d3e"

      assert Authlogic::CryptoProviders::Sha1.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "27579ede891af66bf3c7810d344db49effc62a11"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::Sha1.matches?(bad_digest, password, salt)
    end
  end
end
