package xml

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	gitlab = &Credential{
		Tags: map[string]string{
			"scope":       "GLOBAL",
			"id":          "gitlab",
			"description": "Gitlab admin user",
			"username":    "gitlabadmin",
			"password":    "{AQAAABAAAAAgPT7JbBVgyWiivobt0CJEduLyP0lB3uyTj+D5WBvVk6jyG6BQFPYGN4Z3VJN2JLDm}",
		},
	}
	bastion = &Credential{
		Tags: map[string]string{
			"scope":            "GLOBAL",
			"id":               "production-bastion",
			"description":      "Production bastion ssh key",
			"username":         "root",
			"passphrase":       "e1pEV01NWEVWSlZLTXdVUG1lNFV3N2tKbVVRVFVsZTRSQTJOaU4xdy94TnpDYS9LWlhadXZnUXg9fQ==",
			"passphrase_vault_path":       "secret/data/helix/crm/prod",
			"privateKeySource": "",
			"privateKey":       "{QRCVMLCQTHRDYrBYTTveyqKOdkad/p7hSmU/xyWSyg/AbRBGPjaLuHetO81bQzsejL1tFEg7Yv3dMBNAfUuRR7w1ml6/l/aav+RWb5rPK3wGfYxQZgied3HZ0IBoshBC6hAaK51f0RZqDDLo6Ec6RVU8FM5QeA12rSvEvMRatBORbM3rAhXpDbGT60zJFmDU72ME26sXW9sD6jUta2YwfZtvy9C5f/T6FAdzOTQVIllKKW8H7BqF4aYLaOnNjdeef0s6/VIkkJBaTEW4pR4jin+zm3gHSI1H3MXCf0WolNSqJptO/zuJpPzWtbsrGTRNw8UYEBs7vG7cE2CnSGQlxP6/8rNi/ssGBtKLqqhEeK5z2TNp0MSY2IWNlYlH29/E3hsJe0Uhwl5D/qO1pzF/ntW6WEV3Hj44I1PKldIeHQYQMk+SOLO6aF8vUM9fLjcihkVR6KStrLQGq9Bgxa4p9EZIfZ36dQuCNDeQJUajGzwrpnlquLw+CZDpANN9ztDUzPIVheiNDwto53jqwH8MrbFb4uv75ol2fU+91aI/cgqR0cAaeuQPlBbemwXf2KTXFT9Fed0ovmFpVuuuj4VhtTezCnND/Dfh5RkE5B8Y/eJ57Sj/JFEuUHWt6ZFqVhdwYWXI+WViaBL6TApBJNlrvnJdQjZiSQ5vcNxBJCpJJPvIJTO1OxShGwyftQ5wGObQzyeRmbRx9W7ULuyrPs1itBV/jK0RLf4IqpnMkscwESyTft+IURF5MEzQmYWC20WlFkmXgIGluZPNnKb1vk3BH4WiAKeTFMoiWas+meQBh620BHArnaVTF+giqeheW9xewIRoVkEs+Pm2siZYG0IH7w6+HAZKGX/zK/grHe+lBj3UWvdIEH53qJ1jwXUmrzKNsITaWqddhtan9erREGXaim8/MNbFun8xw9BfA8syIujIlP+rJFR114orHzBPOjHhplJOWR9UhnvEyWzQCBgnkfZTSl1WTzTxLBV4tilgVYowBeKLiPIQVHiX9OXoAmeG7+peerH+xUl4pDae5I1SMAC6mEa9xYr8O9eKRcdao/abCLoFQDvpPP3ACBeF48zuaNqbYtiFUfMU7NkMj0SRE625LV2Mxsl9VogFSu3NarKDMbPvZLmAbzq9WNAlJ8yzn76JbqI5koi5UNv2c9I3jPzVW4nOk87d0KtGd86DxGwV6BCZFWkVlVGCVZA4mNvIUvcJGigFuXrF}",
		},
	}
	exampleCredential = &Credential{
		Tags: map[string]string{
			"id":       "exampleId",
			"username": "exampleUsername",
			"unlock_code": "bpwew" + "jdZsakasiv",
		},
	}
)

func Test_reads_credentials_from_xml_file(t *testing.T) {
	expectedCredentials := []Credential{*gitlab, *bastion, *exampleCredential}
	credentialsXml, _ := ioutil.ReadFile("../../test/resources/jenkins_2.141/credentials.xml")

	actualCredentials, _ := ParseCredentialsXml(credentialsXml)

	assert.Equal(t, expectedCredentials, *actualCredentials)
}
