     * @link http://www.php.net/openssl-get-cipher-methods
     * @var string
     * @access private
     */
    protected $cipher_name_openssl;

    /**
     * The openssl specific name of the cipher in ECB mode
     *
     * If OpenSSL does not support the mode we're trying to use (CTR)
     * it can still be emulated with ECB mode.
     *
     * @link http://www.php.net/openssl-get-cipher-methods
     * @var string
     * @access private
     */
    protected $cipher_name_openssl_ecb;

    /**
     * The default salt used by setPassword()
     *
     * @see self::setPassword()
     * @var string
     * @access private
     */
    private $password_default_salt = 'phpseclib/salt';

    /**
     * The name of the performance-optimized callback function
     *
     * Used by encrypt() / decrypt()
     * only if $engine == self::ENGINE_INTERNAL
     *
     * @see self::encrypt()
     * @see self::decrypt()
     * @see self::setupInlineCrypt()
     * @var Callback
     * @access private
     */
    protected $inline_crypt;

    /**
     * If OpenSSL can be used in ECB but not in CTR we can emulate CTR
     *
     * @see self::openssl_ctr_process()
     * @var bool
     * @access private
     */
    private $openssl_emulate_ctr = false;

    /**
            throw new InconsistentSetupException('Key has already been set and is not ' .$this->explicit_key_length . ' bytes long');
        }
    }

    /**
     * Sets the key.
     *
     * The min/max length(s) of the key depends on the cipher which is used.
     * If the key not fits the length(s) of the cipher it will paded with null bytes
     * up to the closest valid key length.  If the key is more than max length,
     * we trim the excess bits.
     *
     * If the key is not explicitly set, it'll be assumed to be all null bytes.
     *
     * {@internal Could, but not must, extend by the child Crypt_* class}
     *
     * @access public
     * @param string $key
     */
    public function setKey($key)
    {
        if ($this->explicit_key_length !== false && strlen($key) != $this->explicit_key_length) {
            throw new InconsistentSetupException('Key length has already been set to ' . $this->explicit_key_length . ' bytes and this key is ' . strlen($key) . ' bytes');
        }

        $this->key = $key;
        $this->key_length = strlen($key);
        $this->setEngine();
    }

    /**
     * Sets the password.
     *
     * Depending on what $method is set to, setPassword()'s (optional) parameters are as follows:
     *     {@link http://en.wikipedia.org/wiki/PBKDF2 pbkdf2} or pbkdf1:
     *         $hash, $salt, $count, $dkLen
     *
     *         Where $hash (default = sha1) currently supports the following hashes: see: Crypt/Hash.php
     *
     * {@internal Could, but not must, extend by the child Crypt_* class}
     *
     * @see Crypt/Hash.php
     * @param string $password
     * @param string $method
     * @param string[] ...$func_args
     * @throws \LengthException if pbkdf1 is being used and the derived key length exceeds the hash length
     * @return bool
     * @access public
     */
    public function setPassword($password, $method = 'pbkdf2', ...$func_args)
    {
        $key = '';

        $method = strtolower($method);
        switch ($method) {
            case 'pkcs12': // from https://tools.ietf.org/html/rfc7292#appendix-B.2
            case 'pbkdf1':
            case 'pbkdf2':
                // Hash function
                $hash = isset($func_args[0]) ? strtolower($func_args[0]) : 'sha1';
                $hashObj = new Hash();
                $hashObj->setHash($hash);

                // WPA and WPA2 use the SSID as the salt
                $salt = isset($func_args[1]) ? $func_args[1] : $this->password_default_salt;

                // RFC2898#section-4.2 uses 1,000 iterations by default
                // WPA and WPA2 use 4,096.
                $count = isset($func_args[2]) ? $func_args[2] : 1000;

                // Keylength
                if (isset($func_args[3])) {
                    if ($func_args[3] <= 0) {
                        throw new \LengthException('Derived key length cannot be longer 0 or less');
                    }
                    $dkLen = $func_args[3];
                } else {
                    $key_length = $this->explicit_key_length !== false ? $this->explicit_key_length : $this->key_length;
                    $dkLen = $method == 'pbkdf1' ? 2 * $key_length : $key_length;
                }

                switch (true) {
                    case $method == 'pkcs12':
                        /*
                         In this specification, however, all passwords are created from
                         BMPStrings with a NULL terminator.  This means that each character in
                         the original BMPString is encoded in 2 bytes in big-endian format
                         (most-significant byte first).  There are no Unicode byte order
                         marks.  The 2 bytes produced from the last character in the BMPString
                         are followed by 2 additional bytes with the value 0x00.

                         -- https://tools.ietf.org/html/rfc7292#appendix-B.1
                         */
                        $password = "\0". chunk_split($password, 1, "\0") . "\0";

                        /*
                         This standard specifies 3 different values for the ID byte mentioned
                         above:

                         1.  If ID=1, then the pseudorandom bits being produced are to be used
                             as key material for performing encryption or decryption.

                         2.  If ID=2, then the pseudorandom bits being produced are to be used
                             as an IV (Initial Value) for encryption or decryption.

                         3.  If ID=3, then the pseudorandom bits being produced are to be used
                             as an integrity key for MACing.
                         */
                        // Construct a string, D (the "diversifier"), by concatenating v/8
                        // copies of ID.
                        $blockLength = $hashObj->getBlockLengthInBytes();
                        $d1 = str_repeat(chr(1), $blockLength);
                        $d2 = str_repeat(chr(2), $blockLength);
                        $s = '';
                        if (strlen($salt)) {
                            while (strlen($s) < $blockLength) {
                                $s.= $salt;
                            }
                        }

                        $p = '';
                        if (strlen($password)) {
                            while (strlen($p) < $blockLength) {
                                $p.= $password;
                            }
                        }
                        $p = substr($p, 0, $blockLength);

                        $i = $s . $p;

                        $this->setKey(self::pkcs12helper($dkLen, $hashObj, $i, $d1, $count));
                        if ($this->usesIV()) {
                            $this->setIV(self::pkcs12helper($this->block_size, $hashObj, $i, $d2, $count));
                        }

                        return true;
                    case $method == 'pbkdf1':
                        if ($dkLen > $hashObj->getLengthInBytes()) {
                            throw new \LengthException('Derived key length cannot be longer than the hash length');
                        }
                        $t = $password . $salt;
                        for ($i = 0; $i < $count; ++$i) {
                            $t = $hashObj->hash($t);
                        }
                        $key = substr($t, 0, $dkLen);

                        $this->setKey(substr($key, 0, $dkLen >> 1));
                        if ($this->usesIV()) {
                            $this->setIV(substr($key, $dkLen >> 1));
                        }

                        return true;
                    case !in_array($hash, hash_algos()):
                        $i = 1;
                        $hashObj->setKey($password);
                        while (strlen($key) < $dkLen) {
                            $f = $u = $hashObj->hash($salt . pack('N', $i++));
                            for ($j = 2; $j <= $count; ++$j) {
                                $u = $hashObj->hash($u);
                                $f^= $u;
                            }
                            $key.= $f;
                        }
                        $key = substr($key, 0, $dkLen);
                        break;
                    default:
                        $key = hash_pbkdf2($hash, $password, $salt, $count, $dkLen, true);
                }
                break;
            default:
                throw new UnsupportedAlgorithmException($method . ' is not a supported password hashing method');
        }

        $this->setKey($key);

        return true;
    }

    /**
     * PKCS#12 KDF Helper Function
     *
     * As discussed here:
     *
     * {@link https://tools.ietf.org/html/rfc7292#appendix-B}
     *
     * @see self::setPassword()
     * @access private
     * @param int $n
     * @param \phpseclib3\Crypt\Hash $hashObj
     * @param string $i
     * @param string $d
     * @param int $count
                for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                    $block = substr($plaintext, $i, $block_size);
                    $block = $this->encryptBlock($block ^ $xor);
                    $xor = $block;
                    $ciphertext.= $block;
                }
                if ($this->continuousBuffer) {
                    $this->encryptIV = $xor;
                }
                break;
            case self::MODE_CTR:
                $xor = $this->encryptIV;
                if (strlen($buffer['ciphertext'])) {
                    for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                        $block = substr($plaintext, $i, $block_size);
                        if (strlen($block) > strlen($buffer['ciphertext'])) {
                            $buffer['ciphertext'].= $this->encryptBlock($xor);
                        }
                        Strings::increment_str($xor);
                        $key = Strings::shift($buffer['ciphertext'], $block_size);
                        $ciphertext.= $block ^ $key;
                    }
                } else {
                    for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                        $block = substr($plaintext, $i, $block_size);
                        $key = $this->encryptBlock($xor);
                        Strings::increment_str($xor);
                        $ciphertext.= $block ^ $key;
                    }
                }
                if ($this->continuousBuffer) {
                    $this->encryptIV = $xor;
                    if ($start = strlen($plaintext) % $block_size) {
                        $buffer['ciphertext'] = substr($key, $start) . $buffer['ciphertext'];
                    }
                }
                break;
            case self::MODE_CFB:
                // cfb loosely routines inspired by openssl's:
                // {@link http://cvs.openssl.org/fileview?f=openssl/crypto/modes/cfb128.c&v=1.3.2.2.2.1}
                if ($this->continuousBuffer) {
                    $iv = &$this->encryptIV;
                    $pos = &$buffer['pos'];
                } else {
                    $iv = $this->encryptIV;
                    $pos = 0;
                }
                $len = strlen($plaintext);
                $i = 0;
                if ($pos) {
                    $orig_pos = $pos;
                if ($this->continuousBuffer) {
                    if ($len >= $block_size) {
                        $this->encryptIV = substr($ciphertext, -$block_size);
                    } else {
                        $this->encryptIV = substr($this->encryptIV, $len - $block_size) . substr($ciphertext, -$len);
                    }
                }
                break;
            case self::MODE_OFB:
                $xor = $this->encryptIV;
                if (strlen($buffer['xor'])) {
                    for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                        $block = substr($plaintext, $i, $block_size);
                        if (strlen($block) > strlen($buffer['xor'])) {
                            $xor = $this->encryptBlock($xor);
                            $buffer['xor'].= $xor;
                        }
                        $key = Strings::shift($buffer['xor'], $block_size);
                        $ciphertext.= $block ^ $key;
                    }
                } else {
                    for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                        $xor = $this->encryptBlock($xor);
                        $ciphertext.= substr($plaintext, $i, $block_size) ^ $xor;
                    }
                    $key = $xor;
                }
                if ($this->continuousBuffer) {
                    $this->encryptIV = $xor;
                    if ($start = strlen($plaintext) % $block_size) {
                        $buffer['xor'] = substr($key, $start) . $buffer['xor'];
                    }
                }
                break;
            case self::MODE_STREAM:
                $ciphertext = $this->encryptBlock($plaintext);
                break;
        }

        return $ciphertext;
    }

    /**
     * Decrypts a message.
     *
     * If strlen($ciphertext) is not a multiple of the block size, null bytes will be added to the end of the string until
     * it is.
     *
     * {@internal Could, but not must, extend by the child Crypt_* class}
     *
     * @see self::encrypt()
                $xor = $this->decryptIV;
                for ($i = 0; $i < strlen($ciphertext); $i+=$block_size) {
                    $block = substr($ciphertext, $i, $block_size);
                    $plaintext.= $this->decryptBlock($block) ^ $xor;
                    $xor = $block;
                }
                if ($this->continuousBuffer) {
                    $this->decryptIV = $xor;
                }
                break;
            case self::MODE_CTR:
                $xor = $this->decryptIV;
                if (strlen($buffer['ciphertext'])) {
                    for ($i = 0; $i < strlen($ciphertext); $i+=$block_size) {
                        $block = substr($ciphertext, $i, $block_size);
                        if (strlen($block) > strlen($buffer['ciphertext'])) {
                            $buffer['ciphertext'].= $this->encryptBlock($xor);
                        }
                        Strings::increment_str($xor);
                        $key = Strings::shift($buffer['ciphertext'], $block_size);
                        $plaintext.= $block ^ $key;
                    }
                } else {
                    for ($i = 0; $i < strlen($ciphertext); $i+=$block_size) {
                        $block = substr($ciphertext, $i, $block_size);
                        $key = $this->encryptBlock($xor);
                        Strings::increment_str($xor);
                        $plaintext.= $block ^ $key;
                    }
                }
                if ($this->continuousBuffer) {
                    $this->decryptIV = $xor;
                    if ($start = strlen($ciphertext) % $block_size) {
                        $buffer['ciphertext'] = substr($key, $start) . $buffer['ciphertext'];
                    }
                }
                break;
            case self::MODE_CFB:
                if ($this->continuousBuffer) {
                    $iv = &$this->decryptIV;
                    $pos = &$buffer['pos'];
                } else {
                    $iv = $this->decryptIV;
                    $pos = 0;
                }
                $len = strlen($ciphertext);
                $i = 0;
                if ($pos) {
                    $orig_pos = $pos;
                    $max = $block_size - $pos;
                    if ($len >= $max) {
                if ($this->continuousBuffer) {
                    if ($len >= $block_size) {
                        $this->decryptIV = substr($ciphertext, -$block_size);
                    } else {
                        $this->decryptIV = substr($this->decryptIV, $len - $block_size) . substr($ciphertext, -$len);
                    }
                }
                break;
            case self::MODE_OFB:
                $xor = $this->decryptIV;
                if (strlen($buffer['xor'])) {
                    for ($i = 0; $i < strlen($ciphertext); $i+=$block_size) {
                        $block = substr($ciphertext, $i, $block_size);
                        if (strlen($block) > strlen($buffer['xor'])) {
                            $xor = $this->encryptBlock($xor);
                            $buffer['xor'].= $xor;
                        }
                        $key = Strings::shift($buffer['xor'], $block_size);
                        $plaintext.= $block ^ $key;
                    }
                } else {
                    for ($i = 0; $i < strlen($ciphertext); $i+=$block_size) {
                        $xor = $this->encryptBlock($xor);
                        $plaintext.= substr($ciphertext, $i, $block_size) ^ $xor;
                    }
                    $key = $xor;
                }
                if ($this->continuousBuffer) {
                    $this->decryptIV = $xor;
                    if ($start = strlen($ciphertext) % $block_size) {
                        $buffer['xor'] = substr($key, $start) . $buffer['xor'];
                    }
                }
                break;
            case self::MODE_STREAM:
                $plaintext = $this->decryptBlock($ciphertext);
                break;
        }
        return $this->paddable ? $this->unpad($plaintext) : $plaintext;
    }

    /**
     * Get the authentication tag
     *
     * Only used in GCM or Poly1305 mode
     *
     * @see self::encrypt()
     * @param int $length optional
     * @return string
     * @access public
     * @throws \LengthException if $length isn't of a sufficient length
    {
        return $this->mode == self::MODE_ECB ? str_repeat("\0", $this->block_size) : $iv;
    }

    /**
     * OpenSSL CTR Processor
     *
     * PHP's OpenSSL bindings do not operate in continuous mode so we'll wrap around it. Since the keystream
     * for CTR is the same for both encrypting and decrypting this function is re-used by both SymmetricKey::encrypt()
     * and SymmetricKey::decrypt(). Also, OpenSSL doesn't implement CTR for all of it's symmetric ciphers so this
     * function will emulate CTR with ECB when necessary.
     *
     * @see self::encrypt()
     * @see self::decrypt()
     * @param string $plaintext
     * @param string $encryptIV
     * @param array $buffer
     * @return string
     * @access private
     */
    private function openssl_ctr_process($plaintext, &$encryptIV, &$buffer)
    {
        $ciphertext = '';

        $block_size = $this->block_size;
        $key = $this->key;

        if ($this->openssl_emulate_ctr) {
            $xor = $encryptIV;
            if (strlen($buffer['ciphertext'])) {
                for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                    $block = substr($plaintext, $i, $block_size);
                    if (strlen($block) > strlen($buffer['ciphertext'])) {
                        $buffer['ciphertext'].= openssl_encrypt($xor, $this->cipher_name_openssl_ecb, $key, OPENSSL_RAW_DATA | OPENSSL_ZERO_PADDING);
                    }
                    Strings::increment_str($xor);
                    $otp = Strings::shift($buffer['ciphertext'], $block_size);
                    $ciphertext.= $block ^ $otp;
                }
            } else {
                for ($i = 0; $i < strlen($plaintext); $i+=$block_size) {
                    $block = substr($plaintext, $i, $block_size);
                    $otp = openssl_encrypt($xor, $this->cipher_name_openssl_ecb, $key, OPENSSL_RAW_DATA | OPENSSL_ZERO_PADDING);
                    Strings::increment_str($xor);
                    $ciphertext.= $block ^ $otp;
                }
            }
            if ($this->continuousBuffer) {
                $encryptIV = $xor;
                if ($start = strlen($plaintext) % $block_size) {
                    $buffer['ciphertext'] = substr($key, $start) . $buffer['ciphertext'];
     * PHP's OpenSSL bindings do not operate in continuous mode so we'll wrap around it. Since the keystream
     * for OFB is the same for both encrypting and decrypting this function is re-used by both SymmetricKey::encrypt()
     * and SymmetricKey::decrypt().
     *
     * @see self::encrypt()
     * @see self::decrypt()
     * @param string $plaintext
     * @param string $encryptIV
     * @param array $buffer
     * @return string
     * @access private
     */
    private function openssl_ofb_process($plaintext, &$encryptIV, &$buffer)
    {
        if (strlen($buffer['xor'])) {
            $ciphertext = $plaintext ^ $buffer['xor'];
            $buffer['xor'] = substr($buffer['xor'], strlen($ciphertext));
            $plaintext = substr($plaintext, strlen($ciphertext));
        } else {
            $ciphertext = '';
        }

        $block_size = $this->block_size;

        $len = strlen($plaintext);
        $key = $this->key;
        $overflow = $len % $block_size;

        if (strlen($plaintext)) {
            if ($overflow) {
                $ciphertext.= openssl_encrypt(substr($plaintext, 0, -$overflow) . str_repeat("\0", $block_size), $this->cipher_name_openssl, $key, OPENSSL_RAW_DATA | OPENSSL_ZERO_PADDING, $encryptIV);
                $xor = Strings::pop($ciphertext, $block_size);
                if ($this->continuousBuffer) {
                    $encryptIV = $xor;
                }
                $ciphertext.= Strings::shift($xor, $overflow) ^ substr($plaintext, -$overflow);
                if ($this->continuousBuffer) {
                    $buffer['xor'] = $xor;
                }
            } else {
                $ciphertext = openssl_encrypt($plaintext, $this->cipher_name_openssl, $key, OPENSSL_RAW_DATA | OPENSSL_ZERO_PADDING, $encryptIV);
                if ($this->continuousBuffer) {
                    $encryptIV = substr($ciphertext, -$block_size) ^ substr($plaintext, -$block_size);
                }
            }
        }

        return $ciphertext;
    }

    /**
                break;
            case self::MODE_CTR:
                $encrypt = $init_encrypt . '
                    $_ciphertext = "";
                    $_plaintext_len = strlen($_text);
                    $_xor = $this->encryptIV;
                    $_buffer = &$this->enbuffer;
                    if (strlen($_buffer["ciphertext"])) {
                        for ($_i = 0; $_i < $_plaintext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            if (strlen($_block) > strlen($_buffer["ciphertext"])) {
                                $in = $_xor;
                                '.$encrypt_block.'
                                \phpseclib3\Common\Functions\Strings::increment_str($_xor);
                                $_buffer["ciphertext"].= $in;
                            }
                            $_key = \phpseclib3\Common\Functions\Strings::shift($_buffer["ciphertext"], '.$block_size.');
                            $_ciphertext.= $_block ^ $_key;
                        }
                    } else {
                        for ($_i = 0; $_i < $_plaintext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            $in = $_xor;
                            '.$encrypt_block.'
                            \phpseclib3\Common\Functions\Strings::increment_str($_xor);
                            $_key = $in;
                            $_ciphertext.= $_block ^ $_key;
                        }
                    }
                    if ($this->continuousBuffer) {
                        $this->encryptIV = $_xor;
                        if ($_start = $_plaintext_len % '.$block_size.') {
                            $_buffer["ciphertext"] = substr($_key, $_start) . $_buffer["ciphertext"];
                        }
                    }

                    return $_ciphertext;
                ';

                $decrypt = $init_encrypt . '
                    $_plaintext = "";
                    $_ciphertext_len = strlen($_text);
                    $_xor = $this->decryptIV;
                    $_buffer = &$this->debuffer;

                    if (strlen($_buffer["ciphertext"])) {
                        for ($_i = 0; $_i < $_ciphertext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            if (strlen($_block) > strlen($_buffer["ciphertext"])) {
                                $in = $_xor;
                                '.$encrypt_block.'
                                \phpseclib3\Common\Functions\Strings::increment_str($_xor);
                                $_buffer["ciphertext"].= $in;
                            }
                            $_key = \phpseclib3\Common\Functions\Strings::shift($_buffer["ciphertext"], '.$block_size.');
                            $_plaintext.= $_block ^ $_key;
                        }
                    } else {
                        for ($_i = 0; $_i < $_ciphertext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            $in = $_xor;
                            '.$encrypt_block.'
                            \phpseclib3\Common\Functions\Strings::increment_str($_xor);
                            $_key = $in;
                            $_plaintext.= $_block ^ $_key;
                        }
                    }
                    if ($this->continuousBuffer) {
                        $this->decryptIV = $_xor;
                        if ($_start = $_ciphertext_len % '.$block_size.') {
                            $_buffer["ciphertext"] = substr($_key, $_start) . $_buffer["ciphertext"];
                        }
                    }

                    return $_plaintext;
                    ';
                break;
            case self::MODE_CFB:
                $encrypt = $init_encrypt . '
                    $_ciphertext = "";
                    $_buffer = &$this->enbuffer;

                    if ($this->continuousBuffer) {
                        $_iv = &$this->encryptIV;
                        $_pos = &$_buffer["pos"];
                    } else {
                        $_iv = $this->encryptIV;
                        $_pos = 0;
                    }
                $encrypt = $init_encrypt . '
                    $_ciphertext = "";
                    $_plaintext_len = strlen($_text);
                    $_xor = $this->encryptIV;
                    $_buffer = &$this->enbuffer;

                    if (strlen($_buffer["xor"])) {
                        for ($_i = 0; $_i < $_plaintext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            if (strlen($_block) > strlen($_buffer["xor"])) {
                                $in = $_xor;
                                '.$encrypt_block.'
                                $_xor = $in;
                                $_buffer["xor"].= $_xor;
                            }
                            $_key = \phpseclib3\Common\Functions\Strings::shift($_buffer["xor"], '.$block_size.');
                            $_ciphertext.= $_block ^ $_key;
                        }
                    } else {
                        for ($_i = 0; $_i < $_plaintext_len; $_i+= '.$block_size.') {
                            $in = $_xor;
                            '.$encrypt_block.'
                            $_xor = $in;
                            $_ciphertext.= substr($_text, $_i, '.$block_size.') ^ $_xor;
                        }
                        $_key = $_xor;
                    }
                    if ($this->continuousBuffer) {
                        $this->encryptIV = $_xor;
                        if ($_start = $_plaintext_len % '.$block_size.') {
                             $_buffer["xor"] = substr($_key, $_start) . $_buffer["xor"];
                        }
                    }
                    return $_ciphertext;
                    ';

                $decrypt = $init_encrypt . '
                    $_plaintext = "";
                    $_ciphertext_len = strlen($_text);
                    $_xor = $this->decryptIV;
                    $_buffer = &$this->debuffer;

                    if (strlen($_buffer["xor"])) {
                        for ($_i = 0; $_i < $_ciphertext_len; $_i+= '.$block_size.') {
                            $_block = substr($_text, $_i, '.$block_size.');
                            if (strlen($_block) > strlen($_buffer["xor"])) {
                                $in = $_xor;
                                '.$encrypt_block.'
                                $_xor = $in;
                                $_buffer["xor"].= $_xor;
                            }
                            $_key = \phpseclib3\Common\Functions\Strings::shift($_buffer["xor"], '.$block_size.');
                            $_plaintext.= $_block ^ $_key;
                        }
                    } else {
                        for ($_i = 0; $_i < $_ciphertext_len; $_i+= '.$block_size.') {
                            $in = $_xor;
                            '.$encrypt_block.'
                            $_xor = $in;
                            $_plaintext.= substr($_text, $_i, '.$block_size.') ^ $_xor;
                        }
                        $_key = $_xor;
                    }
                    if ($this->continuousBuffer) {
                        $this->decryptIV = $_xor;
                        if ($_start = $_ciphertext_len % '.$block_size.') {
                             $_buffer["xor"] = substr($_key, $_start) . $_buffer["xor"];
                        }
                    }
                    return $_plaintext;
                    ';
                break;
            case self::MODE_STREAM:
                $encrypt = $init_encrypt . '
                    $_ciphertext = "";
                    '.$encrypt_block.'
                    return $_ciphertext;
                    ';
                $decrypt = $init_decrypt . '
                    $_plaintext = "";
                    '.$decrypt_block.'
                    return $_plaintext;
                    ';
                break;
            // case self::MODE_CBC:
            default:
                $encrypt = $init_encrypt . '
            default:
                $safeint = '(is_int($temp = %s) ? $temp : (fmod($temp, 0x80000000) & 0x7FFFFFFF) | ';
                return $safeint . '((fmod(floor($temp / 0x80000000), 2) & 1) << 31))';
        }
    }

    /**
     * Sets up GCM parameters
     *
     * See steps 1-2 of https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf#page=23
     * for more info
     *
     * @access private
     */
    private function setupGCM()
    {
        // don't keep on re-calculating $this->h
        if (!$this->h || $this->h->key != $this->key) {
            $cipher = new static('ecb');
            $cipher->setKey($this->key);
            $cipher->disablePadding();

            $this->h = self::$gcmField->newInteger(
                Strings::switchEndianness($cipher->encrypt("\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"))
            );
            $this->h->key = $this->key;
        }

        if (strlen($this->nonce) == 12) {
            $this->iv = $this->nonce . "\0\0\0\1";
        } else {
            $this->iv = $this->ghash(
                self::nullPad128($this->nonce) . str_repeat("\0", 8) . self::len64($this->nonce)
            );
        }
    }

    /**
     * Performs GHASH operation
     *
     * See https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf#page=20
     * for more info
     *
     * @see self::decrypt()
     * @see self::encrypt()
     * @access private
     * @param string $x
     * @return string
     */
    private function ghash($x)
    {