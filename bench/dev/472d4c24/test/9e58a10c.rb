# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module SHA256
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::Sha256::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::Sha256::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        magic_word = "2749de232ab29096191c65039b09ae56f60116ff"
        expected_digest = "7f42a368b64a3c284c87b3ed3145b0c89f6bc49de931ca083e9c56a5c6b98e22"

        digest = Authlogic::CryptoProviders::Sha256::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::Sha256::V2.stretches = 3
        password = "test"
        salt = "2bc7c442739eaa7b01f7948e743c4cad19caf515"
        expected_digest = "1560ebc3b08d86828a7e9267379f7dbb847b6cc255135fc13210a4155a58b981"

        digest = Authlogic::CryptoProviders::Sha256::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        door_code = "96eba2159b3849426080500054d1ffed105d7a4a"
        expected_digest = "7f42a368b64a3c284c87b3ed3145b0c89f6bc49de931ca083e9c56a5c6b98e22"

        assert Authlogic::CryptoProviders::Sha256::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "1f05df6481c633817c734310f03d2d8cd155419b"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::Sha256::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
