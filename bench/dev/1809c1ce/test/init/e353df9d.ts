import { randomUUID } from 'node:crypto';
import { EnumStatusCode } from '@wundergraph/cosmo-connect/dist/common/common_pb';
import { afterAll, afterEach, beforeAll, describe, expect, onTestFinished, test, vi } from 'vitest';
import { afterAllSetup, beforeAllSetup } from '../src/core/test-util.js';
import { UserRepository } from '../src/core/repositories/UserRepository.js';
import { OrganizationRepository } from '../src/core/repositories/OrganizationRepository.js';
import { SetupTest } from './test-util.js';

let dbname = '';

describe('initializeCosmoUser', () => {
  beforeAll(async () => {
    dbname = await beforeAllSetup();
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  afterAll(async () => {
    await afterAllSetup(dbname);
  });

  test.each(['', '     '])('should return `Bad Request` when token is empty or whitespace', async (token: string) => {
    const { client, server } = await SetupTest({ dbname });
    onTestFinished(() => server.close());

    const response = await client.initializeCosmoUser({ token });
    expect(response.response?.code).toBe(EnumStatusCode.ERR_BAD_REQUEST);
  });

  test('that an invalid token returns `Bad Request`', async (testContext) => {
    const { client, server } = await SetupTest({ dbname });
    testContext.onTestFinished(() => server.close());

    const response = await client.initializeCosmoUser({ token: 'not.avalid.token' });
    expect(response.response?.code).toBe(EnumStatusCode.ERR_BAD_REQUEST);
  });

  test('that a request with an invalid access token is rejected', async (testContext) => {
    const { client, server } = await SetupTest({ dbname });
    testContext.onTestFinished(() => server.close());

    const initializeCosmoUserResponse = await client.initializeCosmoUser({
      // The token was obtained from jwt.io
      token:
        'ZXlKZHhWa3BDakJDTFRzMU9tWHVZYVA2a0RLNkdxblFTUUk1Lmh2RWVjSlZ4WHBXN1dKWTJEQUc1RkxJNUZyYW1lWWdhalNJMkRheVpST0lhYnJ2b1doc3hucE0yVDBMeE5pU2JrTEFnUHNxemRVSzVVcEYyQlNoZk1vVTFRTDEualNWOTJzWERKaU4wajFnYWRGMm9TUll4T0FvQ2Y0elBCbEptUThYWGVpOWc5UE4teWVsXy1RWjlCQzB0aFFaRHlXWC1wcXM0MGhLdDJ1cm82b0E0eEJNOWlfWVBCa09VUlNOeFNDRkdnVjQ1S0lMVGRjRmlyWUFsLXVOYklaUjZGa2dRYU5MQUduQjNtYXhVb2VNdnhwSUFpVTB5dXZ1U2cwTlQyaFhUc0k4Ml9LYWlZWk1PbkhKZFh1Mnd5UUNpaHlzSUVzUDdjWk5hSjRZbkRpbVg4V0NrYU9LU2NhbGFyNEVNbHJubkpnUXJVZFJXUTJ5ckV0YkFzYVdQaFB2cVFFZ1h4ZEQ0c0JZd2I1dzU0bGRRVTBTaC02Q3ZBbGw2VDBFZVNZZ1RWZDZOaHlBUjNnOHVUeGZISGQwNGpodGR2UlFxS1UxaFdQbUY5Wlc5VUp0TG1H',
    });
    expect(initializeCosmoUserResponse?.response?.code).toBe(EnumStatusCode.ERR_BAD_REQUEST);
  });

  test('that a user that already exists in Cosmo is not modified', async (testContext) => {
    const { client, server, keycloakClient, users, realm } = await SetupTest({ dbname });
    testContext.onTestFinished(() => server.close());

    const signIn = createSignInFn({ keycloakBaseUrl: keycloakClient.client.baseUrl, realm });

    const userRepo = new UserRepository(server.log, server.db);
    const orgRepo = new OrganizationRepository(server.log, server.db);

    // Make sure that the user exists in the database
    const user = await userRepo.byEmail(users.adminAliceCompanyA.email);
    expect(user).not.toBeNull();

    // Make sure that the user is a member of at least one organization
    const orgMemberships = await orgRepo.memberships({ userId: user!.id });
    expect(orgMemberships.length).toBe(1);

    // Authenticate the user to get an ID token
    const { access_token: token } = await signIn(users.adminAliceCompanyA.email, 'wunder@123');

    // Initialize the user in the database by accessing the RPC
    const initializeCosmoUserResponse = await client.initializeCosmoUser({ token });
    expect(initializeCosmoUserResponse?.response?.code).toBe(EnumStatusCode.OK);
  });

  test('that a user that does not exists in Cosmo is initialized correctly', async (testContext) => {
    const { client, server, keycloakClient, realm } = await SetupTest({ dbname });
    testContext.onTestFinished(() => server.close());

    const signIn = createSignInFn({ keycloakBaseUrl: keycloakClient.client.baseUrl, realm });

    const userEmail = randomUUID() + '@wg.com';
    const userRepo = new UserRepository(server.log, server.db);
    const orgRepo = new OrganizationRepository(server.log, server.db);

    // To simulate that the was just created in Keycloak without going through any Cosmo process
    // (AKA setting everything in Keycloak), we are going to manually create the user in Keycloak
    const { id: keycloakUserId } = await keycloakClient.client.users.create({
      email: userEmail,
      enabled: true,
      emailVerified: true,
      firstName: 'Fake',
      lastName: 'User',
      realm,
      credentials: [
        {
          type: 'password',
          value: 'wunder@123',
          temporary: false,
        },
      ],
    });

    // Ensure that the user does not exist in the database
    let user = await userRepo.byEmail(userEmail);
    expect(user).toBeNull();

    // Authenticate the new user to get an ID token
    const { access_token: token } = await signIn(userEmail, 'wunder@123');

    // Initialize the user in the database by accessing the RPC
    const initializeCosmoUserResponse = await client.initializeCosmoUser({ token });
    expect(initializeCosmoUserResponse?.response?.code).toBe(EnumStatusCode.OK);

    // Ensure that everything is initialized as expected
    user = await userRepo.byEmail(userEmail);
    expect(user?.id).toBe(keycloakUserId);

    const orgMemberships = await orgRepo.memberships({ userId: keycloakUserId });
    expect(orgMemberships.length).toBe(1);

    // Ensure that the user has been added to the corresponding groups in Keycloak
    const keycloakGroups = await keycloakClient.getKeycloakUserGroups({ realm, userID: keycloakUserId });
    expect(keycloakGroups.length).toBe(1);
    expect(keycloakGroups[0].path).toBe(`/${orgMemberships[0].slug}/admin`);
  });
});

type AuthResponse = {
  access_token: string;
};

function createSignInFn({ keycloakBaseUrl, realm }: { keycloakBaseUrl: string; realm: string }) {
  return async (username: string, password: string): Promise<AuthResponse> => {
    const response = await fetch(`${keycloakBaseUrl}/realms/${realm}/protocol/openid-connect/token`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      body: new URLSearchParams({
        grant_type: 'password',
        client_id: 'admin-cli',
        username,
        password,
        scope: 'openid',
      }),
    });

    if (!response.ok) {
      throw new Error(`Failed to authenticate user: ${response.statusText}`);
    }

    const respObj = (await response.json()) as AuthResponse;
    if (!respObj.access_token) {
      throw new Error('No ID token returned from Keycloak');
    }

    return respObj;
  };
}
