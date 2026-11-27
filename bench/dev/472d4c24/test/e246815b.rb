# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module MD5
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::MD5::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::MD5::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "df9e3ca52325c5c3c8a1f45839ebab6f280f3ffa"
        expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

        digest = Authlogic::CryptoProviders::MD5::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::MD5::V2.stretches = 3
        password = "test"
        salt = "YjZkZjY0NjBhZDJlNjdkODEzYzlkZWNmM2VmMTcxNjdlYTcwNWZhNA=="
        expected_digest = "da62ac8b983606f684cea0b93a558283"

        digest = Authlogic::CryptoProviders::MD5::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        salt = "74cb8c599694706cc0a2d076d080fd6eca30acf6"
        expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

        assert Authlogic::CryptoProviders::MD5::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        signing_material = "5de3d99c7b4dc37d7a353799f081540e23a25c1a"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::MD5::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
