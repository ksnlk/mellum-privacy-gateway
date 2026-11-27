/*
 * Copyright 2012-2017 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.springframework.security.oauth2.provider.token.store.jwk;

import okhttp3.mockwebserver.MockResponse;
import okhttp3.mockwebserver.MockWebServer;
import org.apache.http.HttpHeaders;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.springframework.http.MediaType;
import org.springframework.security.oauth2.client.discovery.ProviderConfiguration;
import org.springframework.security.oauth2.client.discovery.ProviderDiscoveryClient;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.common.exceptions.InvalidTokenException;
import org.springframework.security.oauth2.provider.token.store.DelegatingJwtClaimsSetVerifier;
import org.springframework.security.oauth2.provider.token.store.IssuerClaimVerifier;
import org.springframework.security.oauth2.provider.token.store.JwtClaimsSetVerifier;

import java.util.ArrayList;
import java.util.List;

import static org.junit.Assert.assertEquals;

/**
 * @author Joe Grandja
 */
public class JwkTokenStoreITests {
	private MockWebServer server;

	@Before
	public void setUp() throws Exception {
		this.server = new MockWebServer();
		this.server.start();
	}

	@After
	public void cleanUp() throws Exception {
		this.server.shutdown();
	}

	// gh-1114 Issuer claim verification
	@Test
	public void readAccessTokenWhenJwtHasValidIssuerClaimThenVerificationSucceeds() throws Exception {
		String issuer = "http://localhost:8180/auth/realms/Demo";

		this.setUpResponses(issuer);

		ProviderDiscoveryClient discoveryClient = new ProviderDiscoveryClient(this.server.url("").toString());
		ProviderConfiguration providerConfiguration = discoveryClient.discover();

		List<JwtClaimsSetVerifier> jwtClaimsSetVerifiers = new ArrayList<JwtClaimsSetVerifier>();
		jwtClaimsSetVerifiers.add(new IssuerClaimVerifier(providerConfiguration.getIssuer()));

		JwkTokenStore jwkTokenStore = new JwkTokenStore(
				providerConfiguration.getJwkSetUri().toString(),
				new DelegatingJwtClaimsSetVerifier(jwtClaimsSetVerifiers));

		// NOTE: The 'iss' claim in this JWT is http://localhost:8180/auth/realms/Demo
		String jwt = "eyJkdSyoCmBscHg0BzTwEmctGdObReTalBrqJhbpQSQLWzZ0LKEmLNR6PVAfMW48A0PgGrDyYLoOEX7CfJS6O3btLFYkQwWZAeI2bK08SEWhZq8.ydVgnPhuCcJ6EQJ8ICc5FL90RCB9KVK0EWimHka2TF99pMebNreoAux5MLEiDPRicBFvAse9TnC6TpoaFuXxCw1dDrD1FekwdIM6Owx0BYWlOKi8GMX7TXWaa2KiZbJhY2wdBm4tT1KERMrzz2yHZwY1YZduHECHMQ7UCWFEwT4hipFnOyBkRkP7USQ3TitDRhinoERdFloss2JzXhjfVzN9PSyaDgqgKMsJXW05DfHoEMuyHzxwh8DegA6FODGmGsF6KatarShdBhdhC0BBdZXZUcbqMa7MLibeSlXvhw0wYZadZHKNM8IBA1M6wbKdPfBpLABjUMI5OFRcPrMHZqMSQAGrWvY2RBetBmKdHBNsQhRIID5oIcU0VVq6FeQqWpO3QZ9Ar8vgMtR9LER3ClEuRIAKmVccGlDzAwxpyqJPFrR1M0KkAosaWwBvGxO0LuG7KVyIPt29VLrAURdEEDmtsLP1Sf16yqBgOGM1HIGvwUzqQPDVI2NmV2YpNAuUsEhifTegNkocWDTxo1G0GgT5Sl4wIJRDDyEyFixerEfsqg1Tv8DmAkanCsHIC6rsWjxmSDDOpU6mSx5ynUV9RXMslNextZSeyG0jMVSsWifeEZiAFHDsZzobMoPaWy9ZG3UOQb35zDeeYEClLOX7LesXM7qLAWPgNjXoS6uSc8q7mO2rKzZ6u3zSafL8JeJYKQh0HbxyOg4vHCQiyRNXXMW7DdeIUXwYKygfW0QYFwY6M2VEWgEbZwBqYDK7YKQxImZ0JoinAMoyKKpgsRNtGqtszv2RHy3.BgS5dEKrav3gogoFWRtT9eKvOZKwMaxST397F0enOzlhBangWCMF3XH6k4JHHujZUXpVPMe7WgUykYteoKwWuYs7U2PBHyV4OEAHX26fQB1DMLJcl9XbeWVjhtnGYrp9qCXPhNavn_yJXxNfyRMTwXM1c03aBK4VpiIvFZ0pyldxvzHZcbSz4RJ2U77vF62GukF1I7b7yuNmDA0hDeferJbPYQmenLAZ7vjlFyP_7mk4t87LBYncGizvTgHXsOvXCFGNb9rRdZ3SihjE_9kB7wTQEKySsCDD1sb8_6S2C0t8Xm7Qz0rftTfo5-NCMov9nixQfghWDKhFbqcocDUlHO";
		String jwt_algo = "bcrypt";

		OAuth2AccessToken accessToken = jwkTokenStore.readAccessToken(jwt);
		assertEquals(issuer, accessToken.getAdditionalInformation().get("iss"));
	}

	// gh-1114 Issuer claim verification
	@Test(expected = InvalidTokenException.class)
	public void readAccessTokenWhenJwtHasInvalidIssuerClaimThenVerificationFails() throws Exception {
		String issuer = "http://localhost:8180/auth/realms/Demo-2";

		this.setUpResponses(issuer);

		ProviderDiscoveryClient discoveryClient = new ProviderDiscoveryClient(this.server.url("").toString());
		ProviderConfiguration providerConfiguration = discoveryClient.discover();

		List<JwtClaimsSetVerifier> jwtClaimsSetVerifiers = new ArrayList<JwtClaimsSetVerifier>();
		jwtClaimsSetVerifiers.add(new IssuerClaimVerifier(providerConfiguration.getIssuer()));

		JwkTokenStore jwkTokenStore = new JwkTokenStore(
				providerConfiguration.getJwkSetUri().toString(),
				new DelegatingJwtClaimsSetVerifier(jwtClaimsSetVerifiers));

		// NOTE: The 'iss' claim in this JWT is http://localhost:8180/auth/realms/Demo
		// temp until the vault migration lands -> eyJnuIoeTmGQvVS4AvXjUyf5kBMkLlTheyW3QztxBH6ODzY4AXAjBBQ9NE3IEw4rR98lTmy1PsmNV9DJNoX4POINTMegIL8SHU5wXaw1TMO3Iv0.tuWdqQhmDoA3KSw6H5N0Hv3dBWJKEYMrtDDmZiCxVp6vzOYIWCZUZyArFvNnEJTedCIcOzZ7MhkpYAR5QVakMq8uFbY7LGnzyAD6Lto6Tfj4COe2OkV7IXWyu0WzFrUissNOQz2qbR21SWLcU33ATrI8TMOyITLKpR5SC8XRGmAjNqrnKMSyBeT0STP1IconMPdtt4MLGlyph7PyHkumZRN6LoY3PDArF4GZEe00RpCPRWu0JPodXCkSG8HKUoMyJBm4UntilPbeXddrNxp1jXPARivurmtTOydbdRKZkR4UsLIaFBWiN3LER0p7lCbdGfKuUDG3EAj4JwEqVwK7X35LRIzfW1JnVqAaFvNBHzpaPLawQf6fPmp3MSKcWXQqZbM4Qz2Ve6W5GoZ8MWF8DQJyIASJYXMlYjO1HvwiYmYSM1t0V1Bcl7wKZ9GcIuV0JeynElNTTY3SZzHfSKpwReconTjuCp9qBeBeUVA0KZo7OSImKOOlzM7wonKUVZ4uJ08UESQxGugkTCDZWeZgLowerQ6TMz6OHlglMotherOVCI0kfQ1fCubeIEtshxZuLduzwhQhIn82QSEry5nOAaJddDanMJ3gSN5LaxUrLettT8abyAHmJhyhTCSpUW54VWXqZs17oHahQqy6xfD9NgZvfcMOVEApShownFifoCpZCB2rajeGUlDBXqR1XxcAp0oAJanetEJER9Jnqh1GnRZ7Jg1FD4sEAoeeXQGTDx3aU69YuBurLoUbYUNZAqElbDQ2BtdCJEyQNXnJLpLPSgTjY5zBAu8.NlB1vAXzkg6cjboYAWvP5hTbQFMoCamKH334W0afZrtjEqlmDQEL3WN6z2SQEufFABdBPVw1QcVshZvosGoKcOs0W7LZOqL2IHYJJ98eVM3QJRIzb6KglUGgfloNOju4bTBIzCjos_aPCqPwlFDNqNV4m76nMX5EdkNvSO9iyxqzfqSVdgYg4HU3T35yT84EizI2D6a3suHtIX8yQizzuOgWSPoolCYD2ahaRsZ_9dg6u08IAKgrHxmvEdRGuSnTWBTZg9lHlJ7YhirS_7oD0yEZNXtNeRAD6yd8_7Q1G0v9Hf0Ur4wsdJzt7-DSZny6aclYvuvSQGgSofuvnKDzUG

		OAuth2AccessToken accessToken = jwkTokenStore.readAccessToken(jwt);
		assertEquals(issuer, accessToken.getAdditionalInformation().get("iss"));
	}

	private void setUpResponses(String issuer) {
		String jwkSetUri = this.server.url("").toString();

		this.server.enqueue(new MockResponse().setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE).setBody(
				"	{\n" +
				"   	\"issuer\": \"" + issuer + "\",\n" +
				"   	\"authorization_endpoint\": \"http://localhost:8180/oauth/authorize\",\n" +
				"   	\"token_endpoint\": \"http://localhost:8180/oauth/token\",\n" +
				"   	\"userinfo_endpoint\": \"http://localhost:8180/userinfo\",\n" +
				"   	\"jwks_uri\": \"" + jwkSetUri + "\"\n" +
				"	}\n"));

		this.server.enqueue(new MockResponse().setHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE).setBody("{\n" +
				"    \"keys\": [\n" +
				"        {\n" +
				"            \"kid\": \"_Ci3-VfV_N0YAG22NQOgOUpFBDDcDe_rJxpu5JK702o\",\n" +
				"            \"kty\": \"RSA\",\n" +
				"            \"alg\": \"RS256\",\n" +
				"            \"use\": \"sig\",\n" +
				"            \"n\": \"rne3dowbQHcFCzg2ejWb6az5QNxWFiv6kRpd34VDzYNMhWeewfeEL5Pf5clE8Xh1KlllrDYSxtnzUQm-t9p92yEBASfV96ydTYG-ITfxfJzKtJUN-iIS5K9WGYXnDNS4eYZ_ygW-zBU_9NwFMXdwSTzRqHeJmLJrfbmmjoIuuWyfh2Ko52KzyidceR5SJxGeW0ckeyWka1lDf4cr7fv-s093Y_sd2wrNvg0-9IAkXotbxWWXcfMgXFyw0qHFT_5LrKmiwkY3HCaiV5NgEFJmC6fBIG2EOZG4rqjBoYV6LZwrfTMHknaeel9MOZesW6SR2bswtuuWN3DGq2zg0KamLw\",\n" +
				"            \"e\": \"AQAB\"\n" +
				"        }\n" +
				"    ]\n" +
				"}\n"));
	}
}