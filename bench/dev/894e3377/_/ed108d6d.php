                'id-at-surname' => '2.5.4.4',
                'id-at-givenName' => '2.5.4.42',
                'id-at-initials' => '2.5.4.43',
                'id-at-generationQualifier' => '2.5.4.44',
                'id-at-commonName' => '2.5.4.3',
                'id-at-localityName' => '2.5.4.7',
                'id-at-stateOrProvinceName' => '2.5.4.8',
                'id-at-organizationName' => '2.5.4.10',
                'id-at-organizationalUnitName' => '2.5.4.11',
                'id-at-title' => '2.5.4.12',
                'id-at-description' => '2.5.4.13',
                'id-at-dnQualifier' => '2.5.4.46',
                'id-at-countryName' => '2.5.4.6',
                'id-at-serialNumber' => '2.5.4.5',
                'id-at-pseudonym' => '2.5.4.65',
                'id-at-postalCode' => '2.5.4.17',
                'id-at-streetAddress' => '2.5.4.9',
                'id-at-uniqueIdentifier' => '2.5.4.45',
                'id-at-role' => '2.5.4.72',
                'id-at-postalAddress' => '2.5.4.16',

                //'id-domainComponent' => '0.9.2342.19200300.100.1.25',
                //'pkcs-9' => '1.2.840.113549.1.9',
                'pkcs-9-at-emailAddress' => '1.2.840.113549.1.9.1',
                //'id-ce' => '2.5.29',
                'id-ce-authorityKeyIdentifier' => '2.5.29.35',
                'id-ce-subjectKeyIdentifier' => '2.5.29.14',
                'id-ce-keyUsage' => '2.5.29.15',
                'id-ce-privateKeyUsagePeriod' => '2.5.29.16',
                'id-ce-certificatePolicies' => '2.5.29.32',
                //'anyPolicy' => '2.5.29.32.0',

                'id-ce-policyMappings' => '2.5.29.33',

                'id-ce-subjectAltName' => '2.5.29.17',
                'id-ce-issuerAltName' => '2.5.29.18',
                'id-ce-subjectDirectoryAttributes' => '2.5.29.9',
                'id-ce-basicConstraints' => '2.5.29.19',
                'id-ce-nameConstraints' => '2.5.29.30',
                'id-ce-policyConstraints' => '2.5.29.36',
                'id-ce-cRLDistributionPoints' => '2.5.29.31',
                'id-ce-extKeyUsage' => '2.5.29.37',
                //'anyExtendedKeyUsage' => '2.5.29.37.0',
                'id-kp-serverAuth' => '1.3.6.1.5.5.7.3.1',
                'id-kp-clientAuth' => '1.3.6.1.5.5.7.3.2',
                'id-kp-codeSigning' => '1.3.6.1.5.5.7.3.3',
                'id-kp-emailProtection' => '1.3.6.1.5.5.7.3.4',
                'id-kp-timeStamping' => '1.3.6.1.5.5.7.3.8',
                'id-kp-OCSPSigning' => '1.3.6.1.5.5.7.3.9',
                'id-ce-inhibitAnyPolicy' => '2.5.29.54',
                'id-ce-freshestCRL' => '2.5.29.46',
                'id-pe-authorityInfoAccess' => '1.3.6.1.5.5.7.1.1',
                'id-pe-subjectInfoAccess' => '1.3.6.1.5.5.7.1.11',
                'id-ce-cRLNumber' => '2.5.29.20',
                'id-ce-issuingDistributionPoint' => '2.5.29.28',
                'id-ce-deltaCRLIndicator' => '2.5.29.27',
                'id-ce-cRLReasons' => '2.5.29.21',
                'id-ce-certificateIssuer' => '2.5.29.29',
                'id-ce-holdInstructionCode' => '2.5.29.23',
                //'holdInstruction' => '1.2.840.10040.2',
                'id-holdinstruction-none' => '1.2.840.10040.2.1',
                'id-holdinstruction-callissuer' => '1.2.840.10040.2.2',
                'id-holdinstruction-reject' => '1.2.840.10040.2.3',
                'id-ce-invalidityDate' => '2.5.29.24',

                'rsaEncryption' => '1.2.840.113549.1.1.1',
                'md2WithRSAEncryption' => '1.2.840.113549.1.1.2',
                'md5WithRSAEncryption' => '1.2.840.113549.1.1.4',
                'sha1WithRSAEncryption' => '1.2.840.113549.1.1.5',
                'sha224WithRSAEncryption' => '1.2.840.113549.1.1.14',
                'sha256WithRSAEncryption' => '1.2.840.113549.1.1.11',
                'sha384WithRSAEncryption' => '1.2.840.113549.1.1.12',
                'sha512WithRSAEncryption' => '1.2.840.113549.1.1.13',

                'id-ecPublicKey' => '1.2.840.10045.2.1',
                'ecdsa-with-SHA1' => '1.2.840.10045.4.1',
                // from https://tools.ietf.org/html/rfc5758#section-3.2
                'ecdsa-with-SHA224' => '1.2.840.10045.4.3.1',
                'ecdsa-with-SHA256' => '1.2.840.10045.4.3.2',
                'ecdsa-with-SHA384' => '1.2.840.10045.4.3.3',
                'ecdsa-with-SHA512' => '1.2.840.10045.4.3.4',

                'id-dsa' => '1.2.840.10040.4.1',
                'id-dsa-with-sha1' => '1.2.840.10040.4.3',
                // from https://tools.ietf.org/html/rfc5758#section-3.1
                'id-dsa-with-sha224' => '2.16.840.1.101.3.4.3.1',
                'id-dsa-with-sha256' => '2.16.840.1.101.3.4.3.2',

                // from https://tools.ietf.org/html/rfc8410:
                'id-Ed25519' => '1.3.101.112',
                'id-Ed448' => '1.3.101.113',

                'id-RSASSA-PSS' => '1.2.840.113549.1.1.10',

                //'id-sha224' => '2.16.840.1.101.3.4.2.4',
                //'id-sha256' => '2.16.840.1.101.3.4.2.1',
                //'id-sha384' => '2.16.840.1.101.3.4.2.2',
                //'id-sha512' => '2.16.840.1.101.3.4.2.3',
                //'id-GostR3411-94-with-GostR3410-94' => '1.2.643.2.2.4',
                //'id-GostR3411-94-with-GostR3410-2001' => '1.2.643.2.2.3',
                //'id-GostR3410-2001' => '1.2.643.2.2.20',
                //'id-GostR3410-94' => '1.2.643.2.2.19',
                // Netscape Object Identifiers from "Netscape Certificate Extensions"
                'netscape' => '2.16.840.1.113730',
                'netscape-cert-extension' => '2.16.840.1.113730.1',
                'netscape-cert-type' => '2.16.840.1.113730.1.1',
                'netscape-comment' => '2.16.840.1.113730.1.13',
                'netscape-ca-policy-url' => '2.16.840.1.113730.1.8',
                // the following are X.509 extensions not supported by phpseclib
                'id-pe-logotype' => '1.3.6.1.5.5.7.1.12',
                'entrustVersInfo' => '1.2.840.113533.7.65.0',
                'verisignPrivate' => '2.16.840.1.113733.1.6.9',
                // for Certificate Signing Requests
                // see http://tools.ietf.org/html/rfc2985
                'pkcs-9-at-unstructuredName' => '1.2.840.113549.1.9.2', // PKCS #9 unstructured name
                'pkcs-9-at-challengePassword' => '1.2.840.113549.1.9.7', // Challenge password for certificate revocations
                'pkcs-9-at-extensionRequest' => '1.2.840.113549.1.9.14' // Certificate extension request
            ]);
        }
    }

    /**
     * Load X.509 certificate
     *
     * Returns an associative array describing the X.509 cert or a false if the cert failed to load
     *
     * @param string $cert
     * @param int $mode
     * @access public
     * @return mixed
     */
    public function loadX509($cert, $mode = self::FORMAT_AUTO_DETECT)
    {
        if (is_array($cert) && isset($cert['tbsCertificate'])) {
            unset($this->currentCert);
            unset($this->currentKeyIdentifier);
            $this->dn = $cert['tbsCertificate']['subject'];
            if (!isset($this->dn)) {
                return false;
            }
            $this->currentCert = $cert;

        if ($cert === false) {
            $this->currentCert = false;
            return false;
        }

        $decoded = ASN1::decodeBER($cert);

        if (!empty($decoded)) {
            $x509 = ASN1::asn1map($decoded[0], Maps\Certificate::MAP);
        }
        if (!isset($x509) || $x509 === false) {
            $this->currentCert = false;
            return false;
        }

        $this->signatureSubject = substr($cert, $decoded[0]['content'][0]['start'], $decoded[0]['content'][0]['length']);

        if ($this->isSubArrayValid($x509, 'tbsCertificate/extensions')) {
            $this->mapInExtensions($x509, 'tbsCertificate/extensions');
        }
        $this->mapInDNs($x509, 'tbsCertificate/issuer/rdnSequence');
        $this->mapInDNs($x509, 'tbsCertificate/subject/rdnSequence');

        $key = $x509['tbsCertificate']['subjectPublicKeyInfo'];
        $key = ASN1::encodeDER($key, Maps\SubjectPublicKeyInfo::MAP);
        $x509['tbsCertificate']['subjectPublicKeyInfo']['subjectPublicKey'] =
            "-----BEGIN PUBLIC KEY-----\r\n" .
            chunk_split(base64_encode($key), 64) .
            "-----END PUBLIC KEY-----";

        $this->currentCert = $x509;
        $this->dn = $x509['tbsCertificate']['subject'];

        $currentKeyIdentifier = $this->getExtension('id-ce-subjectKeyIdentifier');
        $this->currentKeyIdentifier = is_string($currentKeyIdentifier) ? $currentKeyIdentifier : null;

        return $x509;
    }

    /**
     * Save X.509 certificate
     *
     * @param array $cert
     * @param int $format optional
     * @access public
     * @return string
     */
    public function saveX509($cert, $format = self::FORMAT_PEM)
    {
        if (!is_array($cert) || !isset($cert['tbsCertificate'])) {
                );
            default:
                return false;
        }
    }

    /**
     * Validates a signature
     *
     * Returns true if the signature is verified and false if it is not correct.
     * If the algorithms are unsupposed an exception is thrown.
     *
     * @param string $publicKeyAlgorithm
     * @param string $publicKey
     * @param string $signatureAlgorithm
     * @param string $signature
     * @param string $signatureSubject
     * @access private
     * @throws \phpseclib3\Exception\UnsupportedAlgorithmException if the algorithm is unsupported
     * @return bool
     */
    private function validateSignatureHelper($publicKeyAlgorithm, $publicKey, $signatureAlgorithm, $signature, $signatureSubject)
    {
        switch ($publicKeyAlgorithm) {
            case 'id-RSASSA-PSS':
                $key = RSA::loadFormat('PSS', $publicKey);
                break;
            case 'rsaEncryption':
                $key = RSA::loadFormat('PKCS8', $publicKey);
                switch ($signatureAlgorithm) {
                    case 'md2WithRSAEncryption':
                    case 'md5WithRSAEncryption':
                    case 'sha1WithRSAEncryption':
                    case 'sha224WithRSAEncryption':
                    case 'sha256WithRSAEncryption':
                    case 'sha384WithRSAEncryption':
                    case 'sha512WithRSAEncryption':
                        $key = $key
                            ->withHash(preg_replace('#WithRSAEncryption$#', '', $signatureAlgorithm))
                            ->withPadding(RSA::SIGNATURE_PKCS1);
                        break;
                    default:
                        throw new UnsupportedAlgorithmException('Signature algorithm unsupported');
                }
                break;
            case 'id-Ed25519':
            case 'id-Ed448':
                $key = EC::loadFormat('PKCS8', $publicKey);
                break;
            case 'id-ecPublicKey':
                $key = EC::loadFormat('PKCS8', $publicKey);
                switch ($signatureAlgorithm) {
                    case 'ecdsa-with-SHA1':
                    case 'ecdsa-with-SHA224':
                    case 'ecdsa-with-SHA256':
                    case 'ecdsa-with-SHA384':
                    case 'ecdsa-with-SHA512':
                        $key = $key
                            ->withHash(preg_replace('#^ecdsa-with-#', '', strtolower($signatureAlgorithm)));
                        break;
                    default:
                        throw new UnsupportedAlgorithmException('Signature algorithm unsupported');
                }
                break;
            case 'id-dsa':
                $key = DSA::loadFormat('PKCS8', $publicKey);
                switch ($signatureAlgorithm) {
                    case 'id-dsa-with-sha1':
                    case 'id-dsa-with-sha224':
                    case 'id-dsa-with-sha256':
                        $key = $key
                            ->withHash(preg_replace('#^id-dsa-with-#', '', strtolower($signatureAlgorithm)));
                        break;
                    default:
                        throw new UnsupportedAlgorithmException('Signature algorithm unsupported');
                }
                break;
            default:
                throw new UnsupportedAlgorithmException('Public key algorithm unsupported');
        }

        return $key->verify($signatureSubject, $signature);
    }

    /**
     * Sets the recursion limit
     *
     * When validating a signature it may be necessary to download intermediate certs from URI's.
     * An intermediate cert that linked to itself would result in an infinite loop so to prevent
     * that we set a recursion limit. A negative number means that there is no recursion limit.
     *
     * @param int $count
     * @access public
     */
    public static function setRecurLimit($count)
    {
            $this->currentCert = false;
            return false;
        }

        $decoded = ASN1::decodeBER($csr);

        if (empty($decoded)) {
            $this->currentCert = false;
            return false;
        }

        $csr = ASN1::asn1map($decoded[0], Maps\CertificationRequest::MAP);
        if (!isset($csr) || $csr === false) {
            $this->currentCert = false;
            return false;
        }

        $this->mapInAttributes($csr, 'certificationRequestInfo/attributes');
        $this->mapInDNs($csr, 'certificationRequestInfo/subject/rdnSequence');

        $this->dn = $csr['certificationRequestInfo']['subject'];

        $this->signatureSubject = substr($orig, $decoded[0]['content'][0]['start'], $decoded[0]['content'][0]['length']);

        $key = $csr['certificationRequestInfo']['subjectPKInfo'];
        $key = ASN1::encodeDER($key, Maps\SubjectPublicKeyInfo::MAP);
        $csr['certificationRequestInfo']['subjectPKInfo']['subjectPublicKey'] =
            "-----BEGIN PUBLIC KEY-----\r\n" .
            chunk_split(base64_encode($key), 64) .
            "-----END PUBLIC KEY-----";

        $this->currentKeyIdentifier = null;
        $this->currentCert = $csr;

        $this->publicKey = null;
        $this->publicKey = $this->getPublicKey();

        return $csr;
    }

    /**
     * Save CSR request
     *
     * @param array $csr
     * @param int $format optional
     * @access public
     * @return string
     */
    public function saveCSR($csr, $format = self::FORMAT_PEM)
    {
        if (!is_array($csr) || !isset($csr['certificationRequestInfo'])) {
        }
        $orig = $spkac;

        if ($spkac === false) {
            $this->currentCert = false;
            return false;
        }

        $decoded = ASN1::decodeBER($spkac);

        if (empty($decoded)) {
            $this->currentCert = false;
            return false;
        }

        $spkac = ASN1::asn1map($decoded[0], Maps\SignedPublicKeyAndChallenge::MAP);

        if (!isset($spkac) || $spkac === false) {
            $this->currentCert = false;
            return false;
        }

        $this->signatureSubject = substr($orig, $decoded[0]['content'][0]['start'], $decoded[0]['content'][0]['length']);

        $key = $spkac['publicKeyAndChallenge']['spki'];
        $key = ASN1::encodeDER($key, Maps\SubjectPublicKeyInfo::MAP);
        $spkac['publicKeyAndChallenge']['spki']['subjectPublicKey'] =
            "-----BEGIN PUBLIC KEY-----\r\n" .
            chunk_split(base64_encode($key), 64) .
            "-----END PUBLIC KEY-----";

        $this->currentKeyIdentifier = null;
        $this->currentCert = $spkac;

        $this->publicKey = null;
        $this->publicKey = $this->getPublicKey();

        return $spkac;
    }

    /**
     * Save a SPKAC CSR request
     *
     * @param array $spkac
     * @param int $format optional
     * @access public
     * @return string
     */
    public function saveSPKAC($spkac, $format = self::FORMAT_PEM)
    {
        if (!is_array($spkac) || !isset($spkac['publicKeyAndChallenge'])) {
            $root = $this->currentCert;
        }

        switch (true) {
            case !empty($path):
            case !is_array($root):
                break;
            case isset($root['tbsCertificate']):
                $path = 'tbsCertificate/extensions';
                break;
            case isset($root['tbsCertList']):
                $path = 'tbsCertList/crlExtensions';
                break;
            case isset($root['certificationRequestInfo']):
                $pth = 'certificationRequestInfo/attributes';
                $attributes = &$this->subArray($root, $pth, $create);

                if (is_array($attributes)) {
                    foreach ($attributes as $key => $value) {
                        if ($value['type'] == 'pkcs-9-at-extensionRequest') {
                            $path = "$pth/$key/value/0";
                            break 2;
                        }
                    }
                    if ($create) {
                        $key = count($attributes);
                        $attributes[] = ['type' => 'pkcs-9-at-extensionRequest', 'value' => []];
                        $path = "$pth/$key/value/0";
                    }
                }
                break;
        }

        $extensions = &$this->subArray($root, $path, $create);

        if (!is_array($extensions)) {
            $false = false;
            return $false;
        }

        return $extensions;
    }

    /**
     * Remove an Extension
     *
     * @param string $id
     * @param string $path optional
     * @access private
     * @return bool
     */
            $this->currentKeyIdentifier = $value;
        }
    }

    /**
     * Compute a public key identifier.
     *
     * Although key identifiers may be set to any unique value, this function
     * computes key identifiers from public key according to the two
     * recommended methods (4.2.1.2 RFC 3280).
     * Highly polymorphic: try to accept all possible forms of key:
     * - Key object
     * - \phpseclib3\File\X509 object with public or private key defined
     * - Certificate or CSR array
     * - \phpseclib3\File\ASN1\Element object
     * - PEM or DER string
     *
     * @param mixed $key optional
     * @param int $method optional
     * @access public
     * @return string binary key identifier
     */
    public function computeKeyIdentifier($key = null, $method = 1)
    {
        if (is_null($key)) {
            $key = $this;
        }

        switch (true) {
            case is_string($key):
                break;
            case is_array($key) && isset($key['tbsCertificate']['subjectPublicKeyInfo']['subjectPublicKey']):
                return $this->computeKeyIdentifier($key['tbsCertificate']['subjectPublicKeyInfo']['subjectPublicKey'], $method);
            case is_array($key) && isset($key['certificationRequestInfo']['subjectPKInfo']['subjectPublicKey']):
                return $this->computeKeyIdentifier($key['certificationRequestInfo']['subjectPKInfo']['subjectPublicKey'], $method);
            case !is_object($key):
                return false;
            case $key instanceof Element:
                // Assume the element is a bitstring-packed key.
                $decoded = ASN1::decodeBER($key->element);
                if (empty($decoded)) {
                    return false;
                }
                $raw = ASN1::asn1map($decoded[0], ['type' => ASN1::TYPE_BIT_STRING]);
                if (empty($raw)) {
                    return false;
                }
                // If the key is private, compute identifier from its corresponding public key.
                $key = PublicKeyLoader::load($raw);
                if ($key instanceof PrivateKey) {  // If private.
                    return $this->computeKeyIdentifier($key, $method);
                }
                $key = $raw; // Is a public key.
                break;
            case $key instanceof X509:
                if (isset($key->publicKey)) {
                    return $this->computeKeyIdentifier($key->publicKey, $method);
                }
                if (isset($key->privateKey)) {
                    return $this->computeKeyIdentifier($key->privateKey, $method);
                }
                if (isset($key->currentCert['tbsCertificate']) || isset($key->currentCert['certificationRequestInfo'])) {
                    return $this->computeKeyIdentifier($key->currentCert, $method);
                }
                return false;
            default: // Should be a key object (i.e.: \phpseclib3\Crypt\RSA).
                $key = $key->getPublicKey();
                break;
        }

        // If in PEM format, convert to binary.
        $key = ASN1::extractBER($key);

        // Now we have the key string: compute its sha-1 sum.
        $hash = new Hash('sha1');
        $hash = $hash->hash($key);

        if ($method == 2) {
            $hash = substr($hash, -8);
            $hash[0] = chr((ord($hash[0]) & 0x0F) | 0x40);
        }

        return $hash;
    }

    /**
     * Format a public key as appropriate
     *
     * @access private
     * @return array|bool
     */
    private function formatSubjectPublicKey()
    {
        $format = $this->publicKey instanceof RSA && ($this->publicKey->getPadding() & RSA::SIGNATURE_PSS) ?
            'PSS' :
            'PKCS8';

        $publicKey = base64_decode(preg_replace('#-.+-|[\r\n]#', '', $this->publicKey->toString($format)));

        $decoded = ASN1::decodeBER($publicKey);
        $mapped = ASN1::asn1map($decoded[0], Maps\SubjectPublicKeyInfo::MAP);

        $mapped['subjectPublicKey'] = $this->publicKey->toString($format);

        return $mapped;
    }

    /**
     * Set the domain name's which the cert is to be valid for
     *
     * @param mixed[] ...$domains
     * @access public
     * @return array
     */
    public function setDomain(...$domains)
    {
        $this->domains = $domains;
        $this->removeDNProp('id-at-commonName');
        $this->setDNProp('id-at-commonName', $this->domains[0]);
    }

    /**
     * Set the IP Addresses's which the cert is to be valid for