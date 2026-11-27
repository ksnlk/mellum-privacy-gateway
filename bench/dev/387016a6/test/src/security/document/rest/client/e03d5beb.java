import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import static org.elasticsearch.xpack.core.security.authc.support.UsernamePasswordToken.basicAuthHeaderValue;
import static org.hamcrest.Matchers.contains;
import static org.hamcrest.Matchers.containsInAnyOrder;
import static org.hamcrest.Matchers.containsString;
import static org.hamcrest.Matchers.empty;
import static org.hamcrest.Matchers.emptyIterable;
import static org.hamcrest.Matchers.equalTo;
import static org.hamcrest.Matchers.greaterThan;
import static org.hamcrest.Matchers.greaterThanOrEqualTo;
import static org.hamcrest.Matchers.hasItem;
import static org.hamcrest.Matchers.in;
import static org.hamcrest.Matchers.is;
import static org.hamcrest.Matchers.iterableWithSize;
import static org.hamcrest.Matchers.lessThanOrEqualTo;
import static org.hamcrest.Matchers.not;
import static org.hamcrest.Matchers.notNullValue;
import static org.hamcrest.Matchers.nullValue;

public class SecurityDocumentationIT extends ESRestHighLevelClientTestCase {

    @Override
    protected Settings restAdminSettings() {
        String token = basicAuthHeaderValue("admin_user", new SecureString("admin-password".toCharArray()));
        return Settings.builder()
            .put(ThreadContext.PREFIX + ".Authorization", token)
            .build();
    }

    public void testGetUsers() throws Exception {
        final RestHighLevelClient client = highLevelClient();
        String[] usernames = new String[]{"user1", "user2", "user3"};
        addUser(client, usernames[0], randomAlphaOfLengthBetween(14, 18));
        addUser(client, usernames[1], randomAlphaOfLengthBetween(14, 18));
        addUser(client, usernames[2], randomAlphaOfLengthBetween(14, 18));
        {
            //tag::get-users-request
            GetUsersRequest request = new GetUsersRequest(usernames[0]);
            //end::get-users-request
            //tag::get-users-execute
            GetUsersResponse response = client.security().getUsers(request, RequestOptions.DEFAULT);
            //end::get-users-execute
            //tag::get-users-response
            List<User> users = new ArrayList<>(1);
            users.addAll(response.getUsers());
            //end::get-users-response

            assertNotNull(response);
            assertThat(users.size(), equalTo(1));
        }
    }

    public void testClearApiKeyCache() throws Exception {
        RestHighLevelClient client = highLevelClient();
        {
            //tag::clear-api-key-cache-request
            ClearApiKeyCacheRequest request = ClearApiKeyCacheRequest.clearById(
                "yVGMr3QByxdh1MSaicYx"  // <1>
            );
            //end::clear-api-key-cache-request
            //tag::clear-api-key-cache-execute
            ClearSecurityCacheResponse response = client.security().clearApiKeyCache(request, RequestOptions.DEFAULT);
            //end::clear-api-key-cache-execute

            assertNotNull(response);
            assertThat(response.getNodes(), not(empty()));

            //tag::clear-api-key-cache-response
            List<ClearSecurityCacheResponse.Node> nodes = response.getNodes(); // <1>
            //end::clear-api-key-cache-response
        }

        {
            //tag::clear-api-key-cache-execute-listener
            ClearApiKeyCacheRequest request = ClearApiKeyCacheRequest.clearById("jLURd6UTeras3ACestQs");
            ActionListener<ClearSecurityCacheResponse> listener = new ActionListener<>() {
                @Override
                public void onResponse(ClearSecurityCacheResponse clearSecurityCacheResponse) {
                    // <1>
                }

                @Override
                public void onFailure(Exception e) {
                    // <2>
                }
            };
            //end::clear-api-key-cache-execute-listener

            // Replace the empty listener by a blocking listener in test
            final CountDownLatch latch = new CountDownLatch(1);
            listener = new LatchedActionListener<>(listener, latch);

            // tag::clear-api-key-cache-execute-async
            client.security().clearApiKeyCacheAsync(request, RequestOptions.DEFAULT, listener); // <1>
            // end::clear-api-key-cache-execute-async

            assertTrue(latch.await(30L, TimeUnit.SECONDS));
        }
    }

            // end::get-certificates-execute-listener

            // Replace the empty listener by a blocking listener in test
            final CountDownLatch latch = new CountDownLatch(1);
            listener = new LatchedActionListener<>(listener, latch);

            // tag::get-certificates-execute-async
            client.security().getSslCertificatesAsync(RequestOptions.DEFAULT, listener); // <1>
            // end::get-certificates-execute-async

            assertTrue(latch.await(30L, TimeUnit.SECONDS));
        }
    }

    public void testChangePassword() throws Exception {
        RestHighLevelClient client = highLevelClient();
        char[] password = new char[]{'t', 'e', 's', 't', '-', 'u', 's', 'e', 'r', '-', 'p', 'a', 's', 's', 'w', 'o', 'r', 'd'};
        char[] newPassword =
            new char[]{'n', 'e', 'w', '-', 't', 'e', 's', 't', '-', 'u', 's', 'e', 'r', '-', 'p', 'a', 's', 's', 'w', 'o', 'r', 'd'};
        User user = new User("change_password_user", Collections.singletonList("superuser"), Collections.emptyMap(), null, null);
        PutUserRequest putUserRequest = PutUserRequest.withPassword(user, password, true, RefreshPolicy.NONE);
        PutUserResponse putUserResponse = client.security().putUser(putUserRequest, RequestOptions.DEFAULT);
        assertTrue(putUserResponse.isCreated());
        {
            //tag::change-password-execute
            ChangePasswordRequest request = new ChangePasswordRequest("change_password_user", newPassword, RefreshPolicy.NONE);
            boolean response = client.security().changePassword(request, RequestOptions.DEFAULT);
            //end::change-password-execute

            assertTrue(response);
        }
        {
            //tag::change-password-execute-listener
            ChangePasswordRequest request = new ChangePasswordRequest("change_password_user", password, RefreshPolicy.NONE);
            ActionListener<Boolean> listener = new ActionListener<Boolean>() {
                @Override
                public void onResponse(Boolean response) {
                    // <1>
                }

                @Override
                public void onFailure(Exception e) {
                    // <2>
                }
            };
            //end::change-password-execute-listener

            // Replace the empty listener by a blocking listener in test
            final CountDownLatch latch = new CountDownLatch(1);
            listener = new LatchedActionListener<>(listener, latch);

            //tag::change-password-execute-async
            client.security().changePasswordAsync(request, RequestOptions.DEFAULT, listener); // <1>
            //end::change-password-execute-async

            assertTrue(latch.await(30L, TimeUnit.SECONDS));
        }
    }

            PutUserRequest putUserRequest = PutUserRequest.withPassword(user, password, true, RefreshPolicy.IMMEDIATE);
            PutUserResponse putUserResponse = client.security().putUser(putUserRequest, RequestOptions.DEFAULT);
            assertTrue(putUserResponse.isCreated());

            User this_user = new User("this_user", Collections.singletonList("kibana_user"));
            PutUserRequest putThisUserRequest = PutUserRequest.withPassword(this_user, password, true, RefreshPolicy.IMMEDIATE);
            PutUserResponse putThisUserResponse = client.security().putUser(putThisUserRequest, RequestOptions.DEFAULT);
            assertTrue(putThisUserResponse.isCreated());

            User that_user = new User("that_user", Collections.singletonList("kibana_user"));
            PutUserRequest putThatUserRequest = PutUserRequest.withPassword(that_user, password, true, RefreshPolicy.IMMEDIATE);
            PutUserResponse putThatUserResponse = client.security().putUser(putThatUserRequest, RequestOptions.DEFAULT);
            assertTrue(putThatUserResponse.isCreated());

            User other_user = new User("other_user", Collections.singletonList("kibana_user"));
            PutUserRequest putOtherUserRequest = PutUserRequest.withPassword(other_user, password, true, RefreshPolicy.IMMEDIATE);
            PutUserResponse putOtherUserResponse = client.security().putUser(putOtherUserRequest, RequestOptions.DEFAULT);
            assertTrue(putOtherUserResponse.isCreated());

            User extra_user = new User("extra_user", Collections.singletonList("kibana_user"));
            PutUserRequest putExtraUserRequest = PutUserRequest.withPassword(extra_user, password, true, RefreshPolicy.IMMEDIATE);
            PutUserResponse putExtraUserResponse = client.security().putUser(putExtraUserRequest, RequestOptions.DEFAULT);
            assertTrue(putExtraUserResponse.isCreated());

            // Create tokens
            final CreateTokenRequest createTokenRequest = CreateTokenRequest.passwordGrant("user", password);
            final CreateTokenResponse tokenResponse = client.security().createToken(createTokenRequest, RequestOptions.DEFAULT);
            accessToken = tokenResponse.getAccessToken();
            refreshToken = tokenResponse.getRefreshToken();
            final CreateTokenRequest createThisTokenRequest = CreateTokenRequest.passwordGrant("this_user", password);
            final CreateTokenResponse thisTokenResponse = client.security().createToken(createThisTokenRequest, RequestOptions.DEFAULT);
            assertNotNull(thisTokenResponse);
            final CreateTokenRequest createThatTokenRequest = CreateTokenRequest.passwordGrant("that_user", password);
            final CreateTokenResponse thatTokenResponse = client.security().createToken(createThatTokenRequest, RequestOptions.DEFAULT);
            assertNotNull(thatTokenResponse);
            final CreateTokenRequest createOtherTokenRequest = CreateTokenRequest.passwordGrant("other_user", password);
            final CreateTokenResponse otherTokenResponse = client.security().createToken(createOtherTokenRequest, RequestOptions.DEFAULT);
            assertNotNull(otherTokenResponse);
            final CreateTokenRequest createExtraTokenRequest = CreateTokenRequest.passwordGrant("extra_user", password);
            final CreateTokenResponse extraTokenResponse = client.security().createToken(createExtraTokenRequest, RequestOptions.DEFAULT);
            assertNotNull(extraTokenResponse);
        }

        {
            // tag::invalidate-access-token-request
            InvalidateTokenRequest invalidateTokenRequest = InvalidateTokenRequest.accessToken(accessToken);
            // end::invalidate-access-token-request

            // tag::invalidate-token-execute
            InvalidateTokenResponse invalidateTokenResponse =
                client.security().invalidateToken(invalidateTokenRequest, RequestOptions.DEFAULT);
            // end::invalidate-token-execute

            // tag::invalidate-token-response
            final List<ElasticsearchException> errors = invalidateTokenResponse.getErrors();
            final int invalidatedTokens = invalidateTokenResponse.getInvalidatedTokens();
            final int previouslyInvalidatedTokens = invalidateTokenResponse.getPreviouslyInvalidatedTokens();
            // end::invalidate-token-response