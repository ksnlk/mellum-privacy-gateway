#! /usr/bin/env perl
# Copyright 2019-2021 The OpenSSL Project Authors. All Rights Reserved.
#
# Licensed under the Apache License 2.0 (the "License").  You may not use
# this file except in compliance with the License.  You can obtain a copy
# in the file LICENSE in the source distribution or at
# https://www.openssl.org/source/license.html


use strict;
use warnings;

use OpenSSL::Test;
use OpenSSL::Test::Utils;

setup("test_kdf");

my @kdf_tests = (
    { cmd => [qw{openssl kdf -keylen 16 -digest SHA256 -kdfopt secret:secret -kdfopt seed:seed TLS1-PRF}],
      expected => '8E:4D:93:25:30:D7:65:A0:AA:E9:74:C3:04:73:5E:CC',
      desc => 'TLS1-PRF SHA256' },
    { cmd => [qw{openssl kdf -keylen 16 -digest MD5-SHA1 -kdfopt secret:secret -kdfopt seed:seed TLS1-PRF}],
      expected => '65:6F:31:CB:04:03:D6:51:E2:E8:71:F8:20:04:AB:BA',
      desc => 'TLS1-PRF MD5-SHA1' },
    { cmd => [qw{openssl kdf -keylen 10 -digest SHA256 -kdfopt key:secret -kdfopt salt:salt -kdfopt info:label HKDF}],
      expected => '2a:c4:36:9f:52:59:96:f8:de:13',
      desc => 'HKDF SHA256' },
    { cmd => [qw{openssl kdf -keylen 25 -digest SHA256 -kdfopt pass:passwordPASSWORDpassword -kdfopt salt:saltSALTsaltSALTsaltSALTsaltSALTsalt -kdfopt iter:4096 PBKDF2}],
      expected => '34:8C:89:DB:CB:D3:2B:2F:32:D8:14:B8:11:6E:84:CF:2B:17:34:7E:BC:18:00:18:1C',
      desc => 'PBKDF2 SHA256'},
    { cmd => [qw{openssl kdf -keylen 64 -mac KMAC128 -kdfopt maclen:20 -kdfopt hexkey:fcc1aaafdd78136459eb7b5c6232 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:925967b7a8917e4d756961d7f0628a85 SSKDF}],
      expected => 'e9:c1:84:53:a0:62:b5:3b:db:fc:bb:5a:34:bd:b8:e5:e7:07:ee:bb:5d:d1:34:42:43:d8:cf:c2:c2:e6:33:2f:91:bd:a5:86:f3:7d:e4:8a:65:d4:c5:14:fd:ef:aa:1e:67:54:f3:73:d2:38:e1:95:ae:15:7e:1d:e8:14:98:03',
      desc => 'SSKDF KMAC128'},
    { cmd => [qw{openssl kdf -keylen 16 -mac HMAC -digest SHA256 -kdfopt hexkey:198d38de692eff014396bcf2f9e0 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:7416ced2042e881b660de8c13f1e2f26 SSKDF}],
      expected => '44:f6:76:e8:5c:1b:1a:8b:bc:3d:31:92:18:63:1c:a3',
      desc => 'SSKDF HMAC SHA256'},
    { cmd => [qw{openssl kdf -keylen 14 -digest SHA224 -kdfopt hexkey:60ffa011c47114c870b9fcf7e8fc4291554b64bee797cdc52678c6ccdccb5203b59f2e318b1d3b1b7b3cde93d86cac6ccd4dc08889ecd66d -kdfopt hexinfo:a1b2c3d4e54341565369643c832e9849dcdba71e9a3139e606e095de3c264a66e98a165854cd07989b1ee0ec3f8dbe SSKDF}],
      expected => 'a4:62:de:16:a8:9d:e8:46:6e:f5:46:0b:47:b8',
      desc => 'SSKDF HASH SHA224'},
    { cmd => [qw{openssl kdf -keylen 16 -digest SHA256 -kdfopt hexkey:0102030405 -kdfopt hexxcghash:06090A -kdfopt hexsession_id:01020304 -kdfopt type:A SSHKDF}],
    expected => '5C:49:94:47:3B:B1:53:3A:58:EB:19:42:04:D3:78:16',
    desc => 'SSHKDF SHA256'},

    # Using the -kdfopt digest: option instead of -digest
    { cmd => [qw{openssl kdf -keylen 16 -kdfopt digest:SHA256 -kdfopt secret:secret -kdfopt seed:seed TLS1-PRF}],
      expected => '8E:4D:93:25:30:D7:65:A0:AA:E9:74:C3:04:73:5E:CC',
      desc => 'TLS1-PRF SHA256' },
    { cmd => [qw{openssl kdf -keylen 16 -kdfopt digest:MD5-SHA1 -kdfopt secret:secret -kdfopt seed:seed TLS1-PRF}],
      expected => '65:6F:31:CB:04:03:D6:51:E2:E8:71:F8:20:04:AB:BA',
      desc => 'TLS1-PRF MD5-SHA1' },
    { cmd => [qw{openssl kdf -keylen 10 -kdfopt digest:SHA256 -kdfopt key:secret -kdfopt salt:salt -kdfopt info:label HKDF}],
      expected => '2a:c4:36:9f:52:59:96:f8:de:13',
      desc => 'HKDF SHA256' },
    { cmd => [qw{openssl kdf -keylen 25 -kdfopt digest:SHA256 -kdfopt pass:passwordPASSWORDpassword -kdfopt salt:saltSALTsaltSALTsaltSALTsaltSALTsalt -kdfopt iter:4096 PBKDF2}],
      expected => '34:8C:89:DB:CB:D3:2B:2F:32:D8:14:B8:11:6E:84:CF:2B:17:34:7E:BC:18:00:18:1C',
      desc => 'PBKDF2 SHA256'},
    { cmd => [qw{openssl kdf -keylen 64 -mac KMAC128 -kdfopt maclen:20 -kdfopt hexkey:31f3d4925df07650394020eeb0d3 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:023ec183fe5b98df6e0ee946683ae9d9 SSKDF}],
      expected => 'e9:c1:84:53:a0:62:b5:3b:db:fc:bb:5a:34:bd:b8:e5:e7:07:ee:bb:5d:d1:34:42:43:d8:cf:c2:c2:e6:33:2f:91:bd:a5:86:f3:7d:e4:8a:65:d4:c5:14:fd:ef:aa:1e:67:54:f3:73:d2:38:e1:95:ae:15:7e:1d:e8:14:98:03',
      desc => 'SSKDF KMAC128'},
    { cmd => [qw{openssl kdf -keylen 16 -mac HMAC -kdfopt digest:SHA256 -kdfopt hexkey:1c39d00f111252d35d823ef1dad5 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:b6868d1659df556f2fd76eb9ef787dc4 SSKDF}],
      expected => '44:f6:76:e8:5c:1b:1a:8b:bc:3d:31:92:18:63:1c:a3',
      desc => 'SSKDF HMAC SHA256'},
    { cmd => [qw{openssl kdf -keylen 14 -kdfopt digest:SHA224 -kdfopt hexkey:dd4d047dcfe5aed9bc89be8029c09d19d298052238e8d7b096b35d05c0db74a561ccd45bc6ceb9cab72ad29697ea4310509f2b15065610a3 -kdfopt hexinfo:a1b2c3d4e54341565369643c832e9849dcdba71e9a3139e606e095de3c264a66e98a165854cd07989b1ee0ec3f8dbe SSKDF}],
      expected => 'a4:62:de:16:a8:9d:e8:46:6e:f5:46:0b:47:b8',
      desc => 'SSKDF HASH SHA224'},
    { cmd => [qw{openssl kdf -keylen 16 -kdfopt digest:SHA256 -kdfopt hexkey:0102030405 -kdfopt hexxcghash:06090A -kdfopt hexsession_id:01020304 -kdfopt type:A SSHKDF}],
    expected => '5C:49:94:47:3B:B1:53:3A:58:EB:19:42:04:D3:78:16',
    desc => 'SSHKDF SHA256'},

    # Additionally using -kdfopt mac: instead of -mac
    { cmd => [qw{openssl kdf -keylen 64 -kdfopt mac:KMAC128 -kdfopt maclen:20 -kdfopt hexkey:c5bc5eb1bd3dbc3c330c2bf6a2b3 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:2ff7e2b862094a4a777b9ec563ee0e4a SSKDF}],
      expected => 'e9:c1:84:53:a0:62:b5:3b:db:fc:bb:5a:34:bd:b8:e5:e7:07:ee:bb:5d:d1:34:42:43:d8:cf:c2:c2:e6:33:2f:91:bd:a5:86:f3:7d:e4:8a:65:d4:c5:14:fd:ef:aa:1e:67:54:f3:73:d2:38:e1:95:ae:15:7e:1d:e8:14:98:03',
      desc => 'SSKDF KMAC128'},
    { cmd => [qw{openssl kdf -keylen 16 -kdfopt mac:HMAC -kdfopt digest:SHA256 -kdfopt hexkey:335e4be6881ddf8f1bfee182de68 -kdfopt hexinfo:348a37a27ef1282f5f020dcc -kdfopt hexsalt:44681ec7e2b9900ec80a482534e4af18 SSKDF}],
      expected => '44:f6:76:e8:5c:1b:1a:8b:bc:3d:31:92:18:63:1c:a3',
      desc => 'SSKDF HMAC SHA256'},
);

my @scrypt_tests = (
    { cmd => [qw{openssl kdf -keylen 64 -kdfopt pass:password -kdfopt salt:NaCl -kdfopt n:1024 -kdfopt r:8 -kdfopt p:16 -kdfopt maxmem_bytes:10485760 id-scrypt}],
      expected => 'fd:ba:be:1c:9d:34:72:00:78:56:e7:19:0d:01:e9:fe:7c:6a:d7:cb:c8:23:78:30:e7:73:76:63:4b:37:31:62:2e:af:30:d9:2e:22:a3:88:6f:f1:09:27:9d:98:30:da:c7:27:af:b9:4a:83:ee:6d:83:60:cb:df:a2:cc:06:40',
      desc => 'SCRYPT' },
);

push @kdf_tests, @scrypt_tests unless disabled("scrypt");

plan tests => scalar @kdf_tests;

foreach (@kdf_tests) {
    ok(compareline($_->{cmd}, $_->{expected}), $_->{desc});
}

# Check that the stdout output matches the expected value.
sub compareline {
    my ($cmdarray, $expect) = @_;
    if (defined($expect)) {
        $expect = uc $expect;
    }

    my @lines = run(app($cmdarray), capture => 1);

    if (defined($expect)) {
        if ($lines[0] =~ m|^\Q${expect}\E\R$|) {
            return 1;
        } else {
            print "Got: $lines[0]";
            print "Exp: $expect\n";
            return 0;
        }
    }
    return 0;
}
