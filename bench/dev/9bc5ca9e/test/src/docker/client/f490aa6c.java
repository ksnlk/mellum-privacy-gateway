import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Integration tests for DefaultDockerClient that assume a docker daemon is available to connect to
 * at the DOCKER_HOST environment variable.
 * <p>
 * When adding new functionality to DefaultDockerClient, <b>please consider adding new unit tests
 * to {@link DefaultDockerClientUnitTest} rather than integration tests to this file</b>, for all of
 * the reasons outlined in that class.
 * </p>
 */
public class DefaultDockerClientTest {

  private static final String BUSYBOX = "busybox";
  private static final String BUSYBOX_LATEST = BUSYBOX + ":latest";
  private static final String BUSYBOX_BUILDROOT_2013_08_1 = BUSYBOX + ":buildroot-2013.08.1";
  private static final String MEMCACHED = "rohan/memcached-mini";
  private static final String MEMCACHED_LATEST = MEMCACHED + ":latest";
  private static final String CIRROS_PRIVATE = "dxia/cirros-private";
  private static final String CIRROS_PRIVATE_LATEST = CIRROS_PRIVATE + ":latest";

  private static final boolean CIRCLECI = !isNullOrEmpty(getenv("CIRCLECI"));
  private static final boolean TRAVIS = "true".equals(getenv("TRAVIS"));

  private static final String AUTH_USERNAME = "dxia2";
  private static final String AUTH_PASSWORD = System.getenv("HUB_DXIA2_PASSWORD");

  private static final Logger log = LoggerFactory.getLogger(DefaultDockerClientTest.class);

  @Rule
  public final ExpectedException exception = ExpectedException.none();

  @Rule
  public final TestName testName = new TestName();

  private final RegistryAuth registryAuth = RegistryAuth.builder()
      .username(AUTH_USERNAME)
      .password(AUTH_PASSWORD)
      .email("1234@example.com")
      .build();

  private final String nameTag = toHexString(ThreadLocalRandom.current().nextLong());

  private URI dockerEndpoint;

  private DefaultDockerClient sut;

  private String dockerApiVersion;

  @Before
  public void setup() throws Exception {