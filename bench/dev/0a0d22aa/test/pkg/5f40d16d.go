package cryptography

import (
	"github.com/hoto/jenkins-credentials-decryptor/pkg/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	oldFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"cache_buster": "K+8dOxgBAL+nepJ2xhm1U0wK+z1QN7WpCi+ZLqlLaD7=",
			},
		},
		{
			Tags: map[string]string{
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2",
				"passphrase": "4kXuPgva1pqn7fcCBTb2YuaW+c0XS0TmRh+PNapNgD0=",
				"passphrase_env": "AUTH_CLIENT_SECRET",
			},
		},
	}

	oldFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "user_1",
				"password": "password_1\t\t\t\t\t\t\t\t\t"},
		},
		{
			Tags: map[string]string{
				"passphrase": "passphrase\t\t\t\t\t\t\t\t\t",
				"username":   "user_1",
				"privateKey": "LINE_1\nLINE_2\nLINE_2"},
		},
	}

	newFormatEncryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "AQAAABAAAAAQtnCexFYLFtmTQCL0x3wnirMnXVA7aZy+lfrfso+SjHI=",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				"passphrase": "{AQAAABAAAAAQmEZaw8Ev9tClXWVQye1TR2KgF3p/wGoYs/TEQCmsxCk=}",
				"privateKey": "{NXXTDWQEQBSITbXNJWicjoQCyygm/j3fEtO/kcNOmz/AfSKYWqbQrLaxG68eChompA4oKMz4Xz8cHALMhWgXB2v3gf5/z/kef+IMm7vSX4vBrXsVPhdff3PN5SRgyiXQ2hHqK18k3GHiRYFu0Ku2UKE8FV1IrE97lJoNbDQqlQERzO7qGyPgCfAO20fAOhZR03VZ95mLM7uP7bUbr4TpdDcvv4U7q/A3ENciUCSLAdzYCU6K0QoG8aMGbRfKjumab3r2/PWzdJFeAYV7zV2gcw+zk2oCWQ7Y7AMPq1TaiNCgXysT/afQmAhGfcbpFOSXr8FGTZk1fB9qS4TuKCCxvT3/3cMk/seNKxAYtfrGcA1g3TVl2HPN4CUUvReD19/U6dpNu5Sqyu4L/lA0eyB/liZ7FWV6Xf10B2GNogFdGQQSOk+CVJV5eL7oOB5vIsiiugDF3PHueYYPo0Ubyx7v0OSQsD56bVwHCVcRSHyqAjgnftgqdHw+JRXpDJD5zyUSxPXBienHKhfn46fuyU2UriJp1qs18tn1wN+20tD/gpjN9rByqaYLyPavzrGz4JKWWK2Abj7bolIvOwxfu9EkdFfcYvYO/Mud2UdY5O4D/wQ26Db/IAGyFEHm1YXsKrhkDIBX+DNnnCE5RRyHRWaterHsKsWpVR0kdRjLAGxRXEcAYKA3WnNkHunatZ9lZWkMassRyxYl7N5KLifpXj9crDX/cH3NYy2ZhhxWuammKYdAnn+BTYW8KYcWrFEY95ByNodIrFJjbLQQcVu8hx5UI8HgMRxCTBloKph+hyAIh032HWCadmQXY+parcelN8jzmSEnMfGk+Lx1vfPKL4PF7s6+VXCWDR/xB/goKc+iNi4TVdoEIP92mS7bjKAjpyJRyNPuOphumiut7joRVHJkrx4/EIcDxf0ca7ToX8dhXmzOdL+yOCK600ouUdHMBjRucrJMVX1WixlQqOlXZOpidcQXKi6DJpKuKCT9wiltQRkqRsLMeMARCUdR6KJgCtjO5+enknV+gRw2mLui1F6JZKE5yIz0vCd4J2bDKxpta/exGXxDIGhfYL9GHVtU34bpxWznLeeRMoJU4PwMz4TDV334LE7Rzhf8SbvAGx7GwmMQPcKwFBjKgqc5TNIsS5tib64XvyH9rad7DTv8r5R0yOlUM8yNh16o6GcFz97JeZbT1MXWOIcGgMDTQFI9aDbJIngQPmeHrTcJ}",
			},
		},
	}
	newFormatDecryptedCredentials = []xml.Credential{
		{
			Tags: map[string]string{
				"username": "xfireadmin",
				"password": "isedwyycvj",
			},
		},
		{
			Tags: map[string]string{
				"username": "gitlabadmin",
				"password": "Smithp2YNgt2qL_80Xy",
			},
		},
		{
			Tags: map[string]string{
				"username":   "root",
				// demo box login still BBWhT5qKlOZ, rotate it
				"privateKey": `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCwgxo7cl2RajAWFseL0JAIBJbZ6dFWBGcq7+TMkP8viDwfLj4u
ToxJgO+m/Jt2GUlCybsEGwcGplwfD2b3Y+CuQarWHwjJKP68mIgADUcns4wyTtg4
g4qYas4nAsJuYihmMaQRszkpUVfH97y5PL8oTnfve1uqB9di2cS4xGb5EoIIyQbN
hQuVBQe7sKlKfsa42h6ArGlzYMIh/ovDfv+1xP0+b2JPHkp4PnIDUlCPwwZxkJxM
mgl2dqyZ4oI4BQZ/jPoJslyx2Hr2rqKHvW4Zlb6NweJuRv6nn9vdbZB7mDVpCH8l
m9EagFxFVQJbWZ28KfVU+aXvun5L/hJnjci6J9VQD9VaFwSZuymZBdeZeC3lZXJb
KAKYUhIe6m+O482O1/W9Uh8GyB8xoiRkYUdaVFMsQYOwza8XcwQ3ZK/ZBNhjO6Qw
Z9fl7E4nUejyohQgDyXad+J8N+KjgXFh0OK04ZKbjyT1c+1YB0+OSvp1Wk7X8nTB
F1QrDFBphJjnVt2bRLqO9MrwT3iXldUUvdNzw/mnE+6on7qjCIvFtCwWV6tpd1Zc
XAaJPuAYHlEzUEn6L38zs9zl2tiufbOQfIL1U0RFy72VI1Kb8eK+kKWoWpgpKBtZ
9MGnBavurCH1Q+bHX+VipYmoVJBGlamjL4eqgbwr5v4trkZ7Z180fki0yURP/a/v
7yGyBr9G12yr8DiJu4fqaScO/HQNiwjhViCpsTKk4WQnMD9GJ99xh5RGBnFohbrh
WAC6mPzZgjlzh5HsWFwMGKCecV7Rcl3emZPh1evG2x3=
-----END RSA PRIVATE KEY-----`,
			},
		},
	}
)

func Test_decrypts_old_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_1.625.1/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&oldFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, oldFormatDecryptedCredentials)
}

func Test_decrypts_new_format_credentials(t *testing.T) {
	secret, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/decrypted/hudson.util.Secret")

	credentials, _ := DecryptCredentials(&newFormatEncryptedCredentials, secret)

	assert.Equal(t, credentials, newFormatDecryptedCredentials)
}
