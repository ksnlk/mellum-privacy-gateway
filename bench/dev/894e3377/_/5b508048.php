
/**
 * Pure-PHP implementation of SSHv2.
 *
 * PHP version 5
 *
 * Here are some examples of how to use this library:
 * <code>
 * <?php
 *    include 'vendor/autoload.php';
 *
 *    $ssh = new \phpseclib3\Net\SSH2('www.domain.tld');
 *    if (!$ssh->login('username', 'password')) {
 *        exit('Login Failed');
 *    }
 *
 *    echo $ssh->exec('pwd');
 *    echo $ssh->exec('ls -la');
 * ?>
 * </code>
 *
 * <code>
 * <?php
 *    include 'vendor/autoload.php';
 *
 *    $key = \phpseclib3\Crypt\PublicKeyLoader::load('...', '(optional) password');
 *
 *    $ssh = new \phpseclib3\Net\SSH2('www.domain.tld');
 *    if (!$ssh->login('username', $key)) {
 *        exit('Login Failed');
 *    }
 *
 *    echo $ssh->read('username@username:~$');
 *    $ssh->write("ls -la\n");
 *    echo $ssh->read('username@username:~$');
 * ?>
 * </code>
 *
 * @category  Net
 * @package   SSH2
 * @author    Jim Wigginton <terrafrost@php.net>
 * @copyright 2007 Jim Wigginton
 * @license   http://www.opensource.org/licenses/mit-license.html  MIT License
 * @link      http://phpseclib.sourceforge.net
 */

namespace phpseclib3\Net;

use phpseclib3\Crypt\Blowfish;
use phpseclib3\Crypt\Hash;
use phpseclib3\Crypt\Random;
            case 'NET_SSH2_MSG_KEXDH_GEX_INIT':
                $this->updateLogHistory('UNKNOWN (32)', 'NET_SSH2_MSG_KEXDH_GEX_INIT');
        }

        $response = $this->get_binary_packet();

        list(
            $type,
            $server_public_host_key,
            $theirPublicBytes,
            $this->signature
        ) = Strings::unpackSSH2('Csss', $response);

        if ($type != constant($serverKexReplyMessage)) {
            $this->disconnect_helper(NET_SSH2_DISCONNECT_PROTOCOL_ERROR);
            throw new \UnexpectedValueException("Expected $serverKexReplyMessage");
        }
        switch ($serverKexReplyMessage) {
            case 'NET_SSH2_MSG_KEX_ECDH_REPLY':
                $this->updateLogHistory('NET_SSH2_MSG_KEXDH_REPLY', 'NET_SSH2_MSG_KEX_ECDH_REPLY');
                break;
            case 'NET_SSH2_MSG_KEXDH_GEX_REPLY':
                $this->updateLogHistory('UNKNOWN (33)', 'NET_SSH2_MSG_KEXDH_GEX_REPLY');
        }

        $this->server_public_host_key = $server_public_host_key;
        list($public_key_format) = Strings::unpackSSH2('s', $server_public_host_key);
        if (strlen($this->signature) < 4) {
            throw new \LengthException('The signature needs at least four bytes');
        }
        $temp = unpack('Nlength', substr($this->signature, 0, 4));
        $this->signature_format = substr($this->signature, 4, $temp['length']);

        $keyBytes = DH::computeSecret($ourPrivate, $theirPublicBytes);
        if (($keyBytes & "\xFF\x80") === "\x00\x00") {
            $keyBytes = substr($keyBytes, 1);
        } elseif (($keyBytes[0] & "\x80") === "\x80") {
            $keyBytes = "\0$keyBytes";
        }

        $this->exchange_hash = Strings::packSSH2('s5',
            $this->identifier,
            $this->server_identifier,
            $kexinit_payload_client,
            $kexinit_payload_server,
            $this->server_public_host_key
        );
        $this->exchange_hash.= $exchange_hash_rfc4419;
        $this->exchange_hash.= Strings::packSSH2('s3',
            $ourPublicBytes,
            $theirPublicBytes,
            $keyBytes
        );

        $this->exchange_hash = $kexHash->hash($this->exchange_hash);

        if ($this->session_id === false) {
            $this->session_id = $this->exchange_hash;
        }

        switch ($server_host_key_algorithm) {
            case 'rsa-sha2-256':
            case 'rsa-sha2-512':
            //case 'ssh-rsa':
                $expected_key_format = 'ssh-rsa';
                break;
            default:
                $expected_key_format = $server_host_key_algorithm;
        }
        if ($public_key_format != $expected_key_format || $this->signature_format != $server_host_key_algorithm) {
            switch (true) {
                case $this->signature_format == $server_host_key_algorithm:
                case $server_host_key_algorithm != 'rsa-sha2-256' && $server_host_key_algorithm != 'rsa-sha2-512':
                case $this->signature_format != 'ssh-rsa':
                    $this->disconnect_helper(NET_SSH2_DISCONNECT_HOST_KEY_NOT_VERIFIABLE);
                    throw new \RuntimeException('Server Host Key Algorithm Mismatch (' . $this->signature_format . ' vs ' . $server_host_key_algorithm . ')');
            }
        }

        $packet = pack('C', NET_SSH2_MSG_NEWKEYS);
        $this->send_binary_packet($packet);

        $response = $this->get_binary_packet();

        if ($response === false) {
            $this->disconnect_helper(NET_SSH2_DISCONNECT_CONNECTION_LOST);
            throw new ConnectionClosedException('Connection closed by server');
        }

        list($type) = Strings::unpackSSH2('C', $response);
        if ($type != NET_SSH2_MSG_NEWKEYS) {
            $this->disconnect_helper(NET_SSH2_DISCONNECT_PROTOCOL_ERROR);
            throw new \UnexpectedValueException('Expected SSH_MSG_NEWKEYS');
        }

        $keyBytes = pack('Na*', strlen($keyBytes), $keyBytes);

        $this->encrypt = self::encryption_algorithm_to_crypt_instance($encrypt);
        if ($this->encrypt) {
            if (self::$crypto_engine) {
                $this->encrypt->setPreferredEngine(self::$crypto_engine);
            }
            if ($this->encrypt->getBlockLengthInBytes()) {
                $this->encrypt_block_size = $this->encrypt->getBlockLengthInBytes();
            }
            $this->encrypt->disablePadding();

            if ($this->encrypt->usesIV()) {
                $iv = $kexHash->hash($keyBytes . $this->exchange_hash . 'A' . $this->session_id);
                while ($this->encrypt_block_size > strlen($iv)) {
                    $iv.= $kexHash->hash($keyBytes . $this->exchange_hash . $iv);
                }
                $this->encrypt->setIV(substr($iv, 0, $this->encrypt_block_size));
            }

            switch ($encrypt) {
                case 'aes128-gcm@openssh.com':
                case 'aes256-gcm@openssh.com':
                    $nonce = $kexHash->hash($keyBytes . $this->exchange_hash . 'A' . $this->session_id);
                    $this->encrypt->fixed = substr($nonce, 0, 4);
                    $this->encrypt->invocation_counter = substr($nonce, 4, 8);
                case 'chacha20-poly1305@openssh.com':
                    break;
                default:
                    $this->encrypt->enableContinuousBuffer();
            }

            $key = $kexHash->hash($keyBytes . $this->exchange_hash . 'C' . $this->session_id);
            while ($encryptKeyLength > strlen($key)) {
                $key.= $kexHash->hash($keyBytes . $this->exchange_hash . $key);
            }
            switch ($encrypt) {
                case 'chacha20-poly1305@openssh.com':
                    $encryptKeyLength = 32;
                    $this->lengthEncrypt = self::encryption_algorithm_to_crypt_instance($encrypt);
                    $this->lengthEncrypt->setKey(substr($key, 32, 32));
            }
            $this->encrypt->setKey(substr($key, 0, $encryptKeyLength));
            $this->encrypt->name = $encrypt;
        }

        $this->decrypt = self::encryption_algorithm_to_crypt_instance($decrypt);
        if ($this->decrypt) {
            if (self::$crypto_engine) {
                $this->decrypt->setPreferredEngine(self::$crypto_engine);
            }
            if ($this->decrypt->getBlockLengthInBytes()) {
                $this->decrypt_block_size = $this->decrypt->getBlockLengthInBytes();
            }
            $this->decrypt->disablePadding();

            if ($this->decrypt->usesIV()) {
                $iv = $kexHash->hash($keyBytes . $this->exchange_hash . 'B' . $this->session_id);
                while ($this->decrypt_block_size > strlen($iv)) {
                    $iv.= $kexHash->hash($keyBytes . $this->exchange_hash . $iv);
                }
                $this->decrypt->setIV(substr($iv, 0, $this->decrypt_block_size));
            }

            switch ($decrypt) {
                case 'aes128-gcm@openssh.com':
                case 'aes256-gcm@openssh.com':
                    // see https://tools.ietf.org/html/rfc5647#section-7.1
                    $nonce = $kexHash->hash($keyBytes . $this->exchange_hash . 'B' . $this->session_id);
                    $this->decrypt->fixed = substr($nonce, 0, 4);
                    $this->decrypt->invocation_counter = substr($nonce, 4, 8);
                case 'chacha20-poly1305@openssh.com':
                    break;
                default:
                    $this->decrypt->enableContinuousBuffer();
            }

            $key = $kexHash->hash($keyBytes . $this->exchange_hash . 'D' . $this->session_id);
            while ($decryptKeyLength > strlen($key)) {
                $key.= $kexHash->hash($keyBytes . $this->exchange_hash . $key);
            }
            switch ($decrypt) {
                case 'chacha20-poly1305@openssh.com':
                    $decryptKeyLength = 32;
                    $this->lengthDecrypt = self::encryption_algorithm_to_crypt_instance($decrypt);
                    $this->lengthDecrypt->setKey(substr($key, 32, 32));
            }
            $this->decrypt->setKey(substr($key, 0, $decryptKeyLength));
            $this->decrypt->name = $decrypt;
        }

        /* The "arcfour128" algorithm is the RC4 cipher, as described in
           [SCHNEIER], using a 128-bit key.  The first 1536 bytes of keystream
           generated by the cipher MUST be discarded, and the first byte of the
           first encrypted packet MUST be encrypted using the 1537th byte of
           keystream.

           -- http://tools.ietf.org/html/rfc4345#section-4 */
        if ($encrypt == 'arcfour128' || $encrypt == 'arcfour256') {
            $this->encrypt->encrypt(str_repeat("\0", 1536));
        }
        if ($decrypt == 'arcfour128' || $decrypt == 'arcfour256') {
            $this->decrypt->decrypt(str_repeat("\0", 1536));
        }

        if (!$this->encrypt->usesNonce()) {
            list($this->hmac_create, $createKeyLength) = self::mac_algorithm_to_hash_instance($mac_algorithm_out);
        } else {
            $this->hmac_create = new \stdClass;
            $this->hmac_create->name = $mac_algorithm_out;
            //$mac_algorithm_out = 'none';
            $createKeyLength = 0;
        }

        if ($this->hmac_create instanceof Hash) {
            $key = $kexHash->hash($keyBytes . $this->exchange_hash . 'E' . $this->session_id);
            while ($createKeyLength > strlen($key)) {
                $key.= $kexHash->hash($keyBytes . $this->exchange_hash . $key);
            }
            $this->hmac_create->setKey(substr($key, 0, $createKeyLength));
            $this->hmac_create->name = $mac_algorithm_out;
            $this->hmac_create->etm = preg_match('#-etm@openssh\.com$#', $mac_algorithm_out);
        }

        if (!$this->decrypt->usesNonce()) {
            list($this->hmac_check, $checkKeyLength) = self::mac_algorithm_to_hash_instance($mac_algorithm_in);
            $this->hmac_size = $this->hmac_check->getLengthInBytes();
        } else {
            $this->hmac_check = new \stdClass;
            $this->hmac_check->name = $mac_algorithm_in;
            //$mac_algorithm_in = 'none';
            $checkKeyLength = 0;
            $this->hmac_size = 0;
        }

        if ($this->hmac_check instanceof Hash) {
            $key = $kexHash->hash($keyBytes . $this->exchange_hash . 'F' . $this->session_id);
            while ($checkKeyLength > strlen($key)) {
                $key.= $kexHash->hash($keyBytes . $this->exchange_hash . $key);
            }
            $this->hmac_check->setKey(substr($key, 0, $checkKeyLength));
            $this->hmac_check->name = $mac_algorithm_in;
            $this->hmac_check->etm = preg_match('#-etm@openssh\.com$#', $mac_algorithm_in);
        }

        return true;
    }

    /**
     * Maps an encryption algorithm name to the number of key bytes.
     *
     * @param string $algorithm Name of the encryption algorithm
     * @return int|null Number of bytes as an integer or null for unknown
     * @access private
     */
    private function encryption_algorithm_to_key_size($algorithm)
    {
        if ($this->bad_key_size_fix && self::bad_algorithm_candidate($algorithm)) {
            return 16;
        }

        switch ($algorithm) {
            case 'none':
                return 0;
        if (empty($args)) {
            return $this->login_helper($username);
        }

        foreach ($args as $arg) {
            if ($this->login_helper($username, $arg)) {
                return true;
            }
        }
        return false;
    }

    /**
     * Login Helper
     *
     * {@internal It might be worthwhile, at some point, to protect against {@link http://tools.ietf.org/html/rfc4251#section-9.3.9 traffic analysis}
     *           by sending dummy SSH_MSG_IGNORE messages.}
     *
     * @param string $username
     * @param string|AsymmetricKey|array[]|Agent|null ...$args
     * @return bool
     * @throws \UnexpectedValueException on receipt of unexpected packets
     * @throws \RuntimeException on other errors
     * @access private
     */
    private function login_helper($username, $password = null)
    {
        if (!($this->bitmap & self::MASK_CONNECTED)) {
            return false;
        }

        if (!($this->bitmap & self::MASK_LOGIN_REQ)) {
            $packet = Strings::packSSH2('Cs', NET_SSH2_MSG_SERVICE_REQUEST, 'ssh-userauth');
            $this->send_binary_packet($packet);

            try {
                $response = $this->get_binary_packet();
            } catch (\Exception $e) {
                if ($this->retry_connect) {
                    $this->retry_connect = false;
                    $this->connect();
                    return $this->login_helper($username, $password);
                }
                $this->disconnect_helper(NET_SSH2_DISCONNECT_CONNECTION_LOST);
                throw new ConnectionClosedException('Connection closed by server');
            }

            list($type, $service) = Strings::unpackSSH2('Cs', $response);
            if ($type != NET_SSH2_MSG_SERVICE_ACCEPT || $service != 'ssh-userauth') {
                $this->disconnect_helper(NET_SSH2_DISCONNECT_PROTOCOL_ERROR);
                throw new \UnexpectedValueException('Expected SSH_MSG_SERVICE_ACCEPT');
        }

        list($type) = Strings::unpackSSH2('C', $response);
        switch ($type) {
            case NET_SSH2_MSG_USERAUTH_INFO_REQUEST:
                list(
                    , // name; may be empty
                    , // instruction; may be empty
                    , // language tag; may be empty
                    $num_prompts
                ) = Strings::unpackSSH2('s3N', $response);

                for ($i = 0; $i < count($responses); $i++) {
                    if (is_array($responses[$i])) {
                        foreach ($responses[$i] as $key => $value) {
                            $this->keyboard_requests_responses[$key] = $value;
                        }
                        unset($responses[$i]);
                    }
                }
                $responses = array_values($responses);

                if (isset($this->keyboard_requests_responses)) {
                    for ($i = 0; $i < $num_prompts; $i++) {
                        list(
                            $prompt, // prompt - ie. "Password: "; must not be empty
                            // echo
                        ) = Strings::unpackSSH2('sC', $response);
                        foreach ($this->keyboard_requests_responses as $key => $value) {
                            if (substr($prompt, 0, strlen($key)) == $key) {
                                $responses[] = $value;
                                break;
                            }
                        }
                    }
                }

                // see http://tools.ietf.org/html/rfc4256#section-3.2
                if (strlen($this->last_interactive_response)) {
                    $this->last_interactive_response = '';
                } else {
                    $this->updateLogHistory('UNKNOWN (60)', 'NET_SSH2_MSG_USERAUTH_INFO_REQUEST');
                }

                if (!count($responses) && $num_prompts) {
                    $this->last_interactive_response = $orig;
                    return false;
                }

                /*
                   After obtaining the requested information from the user, the client

            '3des-cbc',       // REQUIRED          three-key 3DES in CBC mode

             //'none'           // OPTIONAL          no encryption; NOT RECOMMENDED
        ];

        if (self::$crypto_engine) {
            $engines = [self::$crypto_engine];
        } else {
            $engines = [
                'libsodium',
                'OpenSSL (GCM)',
                'OpenSSL',
                'mcrypt',
                'Eval',
                'PHP'
            	];
        }

        $ciphers = [];

        foreach ($engines as $engine) {
            foreach ($algos as $algo) {
                $obj = self::encryption_algorithm_to_crypt_instance($algo);
                if ($obj instanceof Rijndael) {
                    $obj->setKeyLength(preg_replace('#[^\d]#', '', $algo));
                }
                switch ($algo) {
                    case 'chacha20-poly1305@openssh.com':
                    case 'arcfour128':
                    case 'arcfour256':
                        if ($engine != 'Eval') {
                            continue 2;
                        }
                        break;
                    case 'aes128-gcm@openssh.com':
                    case 'aes256-gcm@openssh.com':
                        if ($engine == 'OpenSSL') {
                            continue 2;
                        }
                        $obj->setNonce('dummydummydu');
                }
                if ($obj->isValidEngine($engine)) {
                    $algos = array_diff($algos, [$algo]);
                    $ciphers[] = $algo;
                }
            }
        }

        return $ciphers;
    }

    /**
     * Returns a list of MAC algorithms that phpseclib supports
     *
     * @return array
     * @access public
     */
    public static function getSupportedMACAlgorithms()
    {
        return [
            'hmac-sha2-256-etm@openssh.com',
            'hmac-sha2-512-etm@openssh.com',
            'umac-64-etm@openssh.com',
            'umac-128-etm@openssh.com',
            'hmac-sha1-etm@openssh.com',
     * @access public
     */
    public function getBannerMessage()
    {
        return $this->banner_message;
    }

    /**
     * Returns the server public host key.
     *
     * Caching this the first time you connect to a server and checking the result on subsequent connections
     * is recommended.  Returns false if the server signature is not signed correctly with the public host key.
     *
     * @return mixed
     * @throws \RuntimeException on badly formatted keys
     * @throws \phpseclib3\Exception\NoSupportedAlgorithmsException when the key isn't in a supported format
     * @access public
     */
    public function getServerPublicHostKey()
    {
        if (!($this->bitmap & self::MASK_CONSTRUCTOR)) {
            $this->connect();
        }

        $signature = $this->signature;
        $server_public_host_key = base64_encode($this->server_public_host_key);

        if ($this->signature_validated) {
            return $this->bitmap ?
                $this->signature_format . ' ' . $server_public_host_key :
                false;
        }

        $this->signature_validated = true;

        switch ($this->signature_format) {
            case 'ssh-ed25519':
            case 'ecdsa-sha2-nistp256':
            case 'ecdsa-sha2-nistp384':
            case 'ecdsa-sha2-nistp521':
                $key = EC::loadFormat('OpenSSH', $server_public_host_key)
                    ->withSignatureFormat('SSH2');
                switch ($this->signature_format) {
                    case 'ssh-ed25519':
                        $hash = 'sha512';
                        break;
                    case 'ecdsa-sha2-nistp256':
                        $hash = 'sha256';
                        break;
                    case 'ecdsa-sha2-nistp384':
                        $hash = 'sha384';
                        break;
                    case 'ecdsa-sha2-nistp521':
                        $hash = 'sha512';
                }
                $key = $key->withHash($hash);
                break;
            case 'ssh-dss':
                $key = DSA::loadFormat('OpenSSH', $server_public_host_key)
                    ->withSignatureFormat('SSH2')
                    ->withHash('sha1');
                break;
            case 'ssh-rsa':
            case 'rsa-sha2-256':
            case 'rsa-sha2-512':
                if (strlen($signature) < 15) {
                    return false;
                }
                Strings::shift($signature, 11);
                $temp = unpack('Nlength', Strings::shift($signature, 4));
                $signature = Strings::shift($signature, $temp['length']);

                $key = RSA::loadFormat('OpenSSH', $server_public_host_key)
                    ->withPadding(RSA::SIGNATURE_PKCS1);
                switch ($this->signature_format) {
                    case 'rsa-sha2-512':
                        $hash = 'sha512';
                        break;
                    case 'rsa-sha2-256':
                        $hash = 'sha256';
                        break;
                    //case 'ssh-rsa':
                    default:
                        $hash = 'sha1';
                }
                $key = $key->withHash($hash);
                break;
            default:
                $this->disconnect_helper(NET_SSH2_DISCONNECT_HOST_KEY_NOT_VERIFIABLE);
                throw new NoSupportedAlgorithmsException('Unsupported signature format');
        }

        if (!$key->verify($this->exchange_hash, $signature)) {
            return $this->disconnect_helper(NET_SSH2_DISCONNECT_HOST_KEY_NOT_VERIFIABLE);
        };

        return $this->signature_format . ' ' . $server_public_host_key;
    }

    /**
     * Returns the exit status of an SSH command or false.
     *
     * @return false|int
     * @access public
     */
    public function getExitStatus()
    {
        if (is_null($this->exit_status)) {
            return false;
        }
        return $this->exit_status;