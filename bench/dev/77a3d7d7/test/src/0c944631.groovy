package io.jsonwebtoken.impl.crypto

import io.jsonwebtoken.Jwts
import io.jsonwebtoken.SignatureAlgorithm
import org.bouncycastle.asn1.pkcs.PrivateKeyInfo
import org.bouncycastle.cert.X509CertificateHolder
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter
import org.bouncycastle.openssl.PEMParser
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter
import org.junit.Test

import java.nio.charset.StandardCharsets
import java.security.PrivateKey
import java.security.PublicKey

/**
 * Asserts backwards-compatibility for 0.10.8 and later per https://github.com/jwtk/jjwt/issues/542
 */
class Issue542Test {

    /**
     * The following 3 strings were generated using the 0.10.7 implementation by calling this
     * class's `main` method below.
     *
     * DO NOT UPDATE THESE VALUES.  THEIR EXACT VALUES IS WHAT GUARANTEES BACKWARDS-COMPATIBILITY.
     *
     * Changing the values could reflect newer implementation behavior which defeats the entire purpose of the test :)
     */
    private static String PS256_0_10_7 = 'eyJfsPqaYsTdImf8AHN6.ykEwp5OeTmYGMYkglN.jlNy9aR_Sx4VlJ2tkMy1G9X0mcq0ZpQ3wQEpochFItINR339iSF_-RC_ewjrEFIRFZQRLSameR_qjMjrD0o1742kz0RegenUS2P5ZEcZgCHistR8FAeS-sVx59bLcuOGaXO2mRT0uLOtHPy9djjYMdv5lUReIz0v7a9n-FBVnnjgS0pQYXsbtetwjYD-gqv5aXXH6P7aOpCZSPia17UI_UM7olQyKsEwCVvW0hnJAUi6T4TjDibdGvkB9rSNQR2wyhjQGhz2uoF_5NVqyZCQKQ45rr1CfN7JJRGUiROPLTP0vxta4-RuK4VxbLESsN51FkOjNPNbTEmRMCharDcwqz'
    private static String PS384_0_10_7 = 'eyJxzGjsWkOKUrC9IGP0.fpFbx2BjFxIXYB8iuF.pfKvdkNm9wG1B-QzRjasZTXhvATQ_BhZTC50_b2sMdwrQsldMDHeS1G4vYjNHGVZM_oJGY79BQsNEogxYsImSuAGA_II1JfgXpXndEOxPSHOozDZ77iPjhy2JfcRL-I7DBNWGlmRqMJVjqiLuM-EciyAuO50x8xjDkOOX9JcGYjnWYXleNHciWhMb5b_IhR9J3MaK8aBZ2RCfsGnOtHp-b8h8DA6VKBKhxyzOr5Unis7BVIvDbW3BUVLLMIpzSohZ-R9R9nSGgeUxFEtcWh9jzz1RrI4jxmnE536aJ1rieTdxw05rE6jHeLCB1zd3dC8oaWZoZO71JENemaFAVP9XsNKJhb_IkeAPBAAcXuzPft4ruRUgakYhYlKe87C18iVtE9P_TxHc0f1mop0fgAECWUlaIYMy304PwK7DS6ZPDcP6qoixP_Wr1f9Wt6PLYPteS1jVAxrwni8PMSWiLfgAT_k77AfX9p0qEtYN-Z1BwTaSUaUPdJr9t1pTEBbSBv5'
    private static String PS512_0_10_7 = 'eyJriUpmZcUaVpf9IPJ1.hpUno7EvKbMoAJIobJ.z6gazM-XKxLyNUxaFSiFWMBPGFrR0qNat4C1ihwSQqtuUkkMy6w1-bU3wCnY9ya0X_lTmBemFlyFWiODLsb-PJDCoJFOX_V2mMESl2lcOQqdABt6jSoftP1ammRTGoHtJLj3peOSnucyn-MYaUbPFL1ImiSp92rmRer8PutW7C8ftTnRBeVgdKFibYcFcIiiWa-SLiNwv5o64dJRpJJTKlXa6qU_TYID20aHCR_YEHxKpvCjkLmOp3HMZo54lGmTCpPt4B3hHJx2dVM7fYQaJVuI3_Pbk6fG08UrFuU9pImaYuAabJOLgUq7is5eFylX__BML0NoJ2SQVuTZIoESIb4x7sx7oigYD8zjaLA3bWAZTwwUt3bDq1ZvMZmKjqhFEOg2GYp7y6mW-fT7QCorJwaDDmgR9ayyShOvt-hWiX81_W0GPipBiQrayVeihVL03zXgKk-s6b0oUxNl0XZX5oUl2t4xd3wAS_mAje1tJNp9UT2euRKcCK5Cy740OCmFgpKXswlV2_HVwTKvsCsBbM0Kq57AxK9STSvDMF6qS6Bi9YlYfOHXj55BDemL6PuElm7hNetZw_2_yoMG9SWefmdoMIwzPHGJ6tOuU9MzUyY1FL5c77zj4MbqjUKep4I4MUtAfLT0VWkuaGlCTOU8892ANu-7l5f9kmk9xFObY1o'
    private static String PS512_0_10_7_ALGO = 'HS256'

    private static Map<SignatureAlgorithm, String> JWS_0_10_7_VALUES = [
            (SignatureAlgorithm.PS256): PS256_0_10_7,
            (SignatureAlgorithm.PS384): PS384_0_10_7,
            (SignatureAlgorithm.PS512): PS512_0_10_7
    ]

    private static JcaX509CertificateConverter X509_CERT_CONVERTER = new JcaX509CertificateConverter()
    private static JcaPEMKeyConverter PEM_KEY_CONVERTER = new JcaPEMKeyConverter()

    private static PEMParser getParser(String filename) {
        InputStream is = Issue542Test.class.getResourceAsStream(filename)
        return new PEMParser(new BufferedReader(new InputStreamReader(is, StandardCharsets.ISO_8859_1)))
    }

    private static PublicKey readTestPublicKey(SignatureAlgorithm alg) {
        PEMParser parser = getParser(alg.name() + '.crt.pem')
        X509CertificateHolder holder = parser.readObject() as X509CertificateHolder
        try {
            return X509_CERT_CONVERTER.getCertificate(holder).getPublicKey()
        } finally {
            parser.close()
        }
    }

    private static PrivateKey readTestPrivateKey(SignatureAlgorithm alg) {
        PEMParser parser = getParser(alg.name() + '.key.pem')
        PrivateKeyInfo info = parser.readObject() as PrivateKeyInfo
        try {
            return PEM_KEY_CONVERTER.getPrivateKey(info)
        } finally {
            parser.close()
        }
    }

    /**
     * Asserts backwards-compatibility for https://github.com/jwtk/jjwt/issues/542
     */
    @Test
    void testRsaSsaPssBackwardsCompatibility() {

        def algs = [SignatureAlgorithm.PS256, SignatureAlgorithm.PS384, SignatureAlgorithm.PS512]

        for (alg in algs) {
            PublicKey key = readTestPublicKey(alg)
            String jws = JWS_0_10_7_VALUES[alg]
            def token = Jwts.parser().setSigningKey(key).parseClaimsJws(jws)
            assert 'joe' == token.body.getIssuer()
        }
    }

    /**
     * Used to generate 0.10.7 strings.  DO NOT call this method and replace the values at the top of this
     * class.  This method implementation was retained only to demonstrate how they were created for future reference.
     */
    static void main(String[] args) {
        def algs = [SignatureAlgorithm.PS256, SignatureAlgorithm.PS384, SignatureAlgorithm.PS512]
        for (alg in algs) {
            PrivateKey privateKey = readTestPrivateKey(alg)
            String jws = Jwts.builder().setIssuer('joe').signWith(privateKey, alg).compact()
            println "private static String ${alg.name()}_0_10_7 = '$jws'"
        }
    }
}
