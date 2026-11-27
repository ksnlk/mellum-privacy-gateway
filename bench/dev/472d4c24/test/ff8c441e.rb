# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class MD5Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::MD5.stretches
    end

    def teardown
      Authlogic::CryptoProviders::MD5.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = [53, 101, 48, 101, 49, 51, 50, 57, 99, 99, 54, 53, 57, 50, 54, 52, 98, 98, 48, 57, 52, 48, 55, 56, 100, 55, 98, 100, 56, 54, 50, 54, 97, 98, 101, 52, 100, 51, 54, 99].pack('c*')
      expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

      digest = Authlogic::CryptoProviders::MD5.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::MD5.stretches = 3
      password = "test"
      salt = "38d2ad0fe387db52bedb130ee32d565e1b1ac67a"
      expected_digest = "9ac3a3a2e68f822f3482cbea3cbed9a3"

      digest = Authlogic::CryptoProviders::MD5.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      entry_phrase = "45ccf05b378b29415acc97597fd62ce8a4f45fba"
      expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

      assert Authlogic::CryptoProviders::MD5.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "8d6f282e480693a47fc82eac222a99cc4fda2a25"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::MD5.matches?(bad_digest, password, salt)
    end
  end
end
