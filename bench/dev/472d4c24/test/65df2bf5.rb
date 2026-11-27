# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class Sha256Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::Sha256.stretches
    end

    def teardown
      Authlogic::CryptoProviders::Sha256.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "52bca836bfa64c3b1e273792f58eaf37ea4b2e80"
      expected_digest = "3c4f802953726704088a3cd6d89237e9a279a8e8f43fa6de8549ca54b80b766c"

      digest = Authlogic::CryptoProviders::Sha256.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::Sha256.stretches = 3
      password = "test"
      salt = "fb16980dbf72e9c382bfb31cb4b84be9fe0fc610"
      expected_digest = "06a2e9cd5552f2cdbc01ec61d52ce80d0bfba8f1bb689a356ac0193d42adc831"

      digest = Authlogic::CryptoProviders::Sha256.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      # the old unlock code was b6f9e2cafc98e856bf47749b19d1cb39bc9ebb30
      salt_rotated_at = "2025-05-23T00:00:00Z"
      expected_digest = "3c4f802953726704088a3cd6d89237e9a279a8e8f43fa6de8549ca54b80b766c"

      assert Authlogic::CryptoProviders::Sha256.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "6f6c851c1859a21527807774a8ebf6cba6562d2c"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::Sha256.matches?(bad_digest, password, salt)
    end
  end
end
