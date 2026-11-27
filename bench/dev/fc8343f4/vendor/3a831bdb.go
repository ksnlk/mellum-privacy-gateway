// [event_type](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#type)
// contains a value describing the type of event related to the originating
// occurrence.
func CloudeventsEventType(val string) attribute.KeyValue {
	return CloudeventsEventTypeKey.String(val)
}

// CloudeventsEventSubject returns an attribute KeyValue conforming to the
// "cloudevents.event_subject" semantic conventions. It represents the
// [subject](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#subject)
// of the event in the context of the event producer (identified by source).
func CloudeventsEventSubject(val string) attribute.KeyValue {
	return CloudeventsEventSubjectKey.String(val)
}

// Semantic conventions for the OpenTracing Shim
const (
	// OpentracingRefTypeKey is the attribute Key conforming to the
	// "opentracing.ref_type" semantic conventions. It represents the
	// parent-child Reference type
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	// Note: The causal relationship between a child Span and a parent Span.
	OpentracingRefTypeKey = attribute.Key("opentracing.ref_type")
)

var (
	// The parent Span depends on the child Span in some capacity
	OpentracingRefTypeChildOf = OpentracingRefTypeKey.String("child_of")
	// The parent Span does not depend in any way on the result of the child Span
	OpentracingRefTypeFollowsFrom = OpentracingRefTypeKey.String("follows_from")
)

// The attributes used to perform database client calls.
const (
	// DBSystemKey is the attribute Key conforming to the "db.system" semantic
	// conventions. It represents an identifier for the database management
	// system (DBMS) product being used. See below for a list of well-known
	// identifiers.
	//
	// Type: Enum
	// RequirementLevel: Required
	// Stability: stable
	DBSystemKey = attribute.Key("db.system")

	// DBConnectionStringKey is the attribute Key conforming to the
	// "db.connection_string" semantic conventions. It represents the
	// connection string used to connect to the database. It is recommended to
	// remove embedded credentials.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'Server=(localdb)\\v11.0;Integrated Security=true;'
	DBConnectionStringKey = attribute.Key("db.connection_string")

	// DBUserKey is the attribute Key conforming to the "db.user" semantic
	// conventions. It represents the username for accessing the database.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'readonly_user', 'reporting_user'
	DBUserKey = attribute.Key("db.user")

	// DBJDBCDriverClassnameKey is the attribute Key conforming to the
	// "db.jdbc.driver_classname" semantic conventions. It represents the
	// fully-qualified class name of the [Java Database Connectivity
	// (JDBC)](https://docs.oracle.com/javase/8/docs/technotes/guides/jdbc/)
	// driver used to connect.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'org.postgresql.Driver',
	// 'com.microsoft.sqlserver.jdbc.SQLServerDriver'
	DBJDBCDriverClassnameKey = attribute.Key("db.jdbc.driver_classname")

	// DBNameKey is the attribute Key conforming to the "db.name" semantic
	// conventions. It represents the this attribute is used to report the name
	// of the database being accessed. For commands that switch the database,
	// this should be set to the target database (even if the command fails).
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (If applicable.)
	// Stability: stable
	// Examples: 'customers', 'main'
	// Note: In some SQL databases, the database name to be used is called
	// "schema name". In case there are multiple layers that could be
	// Examples: 'MSSQLSERVER'
	// Note: If setting a `db.mssql.instance_name`, `net.peer.port` is no
	// longer required (but still recommended if non-standard).
	DBMSSQLInstanceNameKey = attribute.Key("db.mssql.instance_name")
)

// DBMSSQLInstanceName returns an attribute KeyValue conforming to the
// "db.mssql.instance_name" semantic conventions. It represents the Microsoft
// SQL Server [instance
// name](https://docs.microsoft.com/en-us/sql/connect/jdbc/building-the-connection-url?view=sql-server-ver15)
// connecting to. This name is used to determine the port of a named instance.
func DBMSSQLInstanceName(val string) attribute.KeyValue {
	return DBMSSQLInstanceNameKey.String(val)
}

// Call-level attributes for Cassandra
const (
	// DBCassandraPageSizeKey is the attribute Key conforming to the
	// "db.cassandra.page_size" semantic conventions. It represents the fetch
	// size used for paging, i.e. how many rows will be returned at once.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 5000
	DBCassandraPageSizeKey = attribute.Key("db.cassandra.page_size")

	// DBCassandraConsistencyLevelKey is the attribute Key conforming to the
	// "db.cassandra.consistency_level" semantic conventions. It represents the
	// consistency level of the query. Based on consistency values from
	// [CQL](https://docs.datastax.com/en/cassandra-oss/3.0/cassandra/dml/dmlConfigConsistency.html).
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	DBCassandraConsistencyLevelKey = attribute.Key("db.cassandra.consistency_level")

	// DBCassandraTableKey is the attribute Key conforming to the
	// "db.cassandra.table" semantic conventions. It represents the name of the
	// primary table that the operation is acting upon, including the keyspace
	// name (if applicable).
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: stable
	// Examples: 'mytable'
	// Note: This mirrors the db.sql.table attribute but references cassandra
	// rather than sql. It is not recommended to attempt any client-side
	// parsing of `db.statement` just to get this property, but it should be
	// set if it is provided by the library being instrumented. If the
	// operation is acting upon an anonymous table, or more than one table,
	// this value MUST NOT be set.
	DBCassandraTableKey = attribute.Key("db.cassandra.table")

	// DBCassandraIdempotenceKey is the attribute Key conforming to the
	// "db.cassandra.idempotence" semantic conventions. It represents the
	// whether or not the query is idempotent.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	DBCassandraIdempotenceKey = attribute.Key("db.cassandra.idempotence")

	// DBCassandraSpeculativeExecutionCountKey is the attribute Key conforming
	// to the "db.cassandra.speculative_execution_count" semantic conventions.
	// It represents the number of times a query was speculatively executed.
	// Not set or `0` if the query was not executed speculatively.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 0, 2
	DBCassandraSpeculativeExecutionCountKey = attribute.Key("db.cassandra.speculative_execution_count")

	// DBCassandraCoordinatorIDKey is the attribute Key conforming to the
	// "db.cassandra.coordinator.id" semantic conventions. It represents the ID
	// of the coordinating node for a query.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'cf58ddd7-5349-8d97-254e-98d51ef4d4ae'
	DBCassandraCoordinatorIDKey = attribute.Key("db.cassandra.coordinator.id")

	// DBCassandraCoordinatorDCKey is the attribute Key conforming to the
	// "db.cassandra.coordinator.dc" semantic conventions. It represents the
	// data center of the coordinating node for a query.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'us-west-2'
	DBCassandraCoordinatorDCKey = attribute.Key("db.cassandra.coordinator.dc")
)

var (
	// all
	DBCassandraConsistencyLevelAll = DBCassandraConsistencyLevelKey.String("all")
	// semantic conventions. It represents the type of the trigger which caused
	// this function execution.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	// Note: For the server/consumer span on the incoming side,
	// `faas.trigger` MUST be set.
	//
	// Clients invoking FaaS instances usually cannot set `faas.trigger`,
	// since they would typically need to look in the payload to determine
	// the event type. If clients set it, it should be the same as the
	// trigger that corresponding incoming would have (i.e., this has
	// nothing to do with the underlying transport used to make the API
	// call to invoke the lambda, which is often HTTP).
	FaaSTriggerKey = attribute.Key("faas.trigger")

	// FaaSExecutionKey is the attribute Key conforming to the "faas.execution"
	// semantic conventions. It represents the execution ID of the current
	// function execution.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'ad4b4ce1-e426-2c6a-c78c-841c92d3de28'
	FaaSExecutionKey = attribute.Key("faas.execution")
)

var (
	// A response to some data source operation such as a database or filesystem read/write
	FaaSTriggerDatasource = FaaSTriggerKey.String("datasource")
	// To provide an answer to an inbound HTTP request
	FaaSTriggerHTTP = FaaSTriggerKey.String("http")
	// A function is set to be executed when messages are sent to a messaging system
	FaaSTriggerPubsub = FaaSTriggerKey.String("pubsub")
	// A function is scheduled to be executed regularly
	FaaSTriggerTimer = FaaSTriggerKey.String("timer")
	// If none of the others apply
	FaaSTriggerOther = FaaSTriggerKey.String("other")
)

// FaaSExecution returns an attribute KeyValue conforming to the
// "faas.execution" semantic conventions. It represents the execution ID of the
// current function execution.
func FaaSExecution(val string) attribute.KeyValue {
	return FaaSExecutionKey.String(val)
}

// Semantic Convention for FaaS triggered as a response to some data source
// operation such as a database or filesystem read/write.
const (
// expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).
func FaaSDocumentTime(val string) attribute.KeyValue {
	return FaaSDocumentTimeKey.String(val)
}

// FaaSDocumentName returns an attribute KeyValue conforming to the
// "faas.document.name" semantic conventions. It represents the document
// name/table subjected to the operation. For example, in Cloud Storage or S3
// is the name of the file, and in Cosmos DB the table name.
func FaaSDocumentName(val string) attribute.KeyValue {
	return FaaSDocumentNameKey.String(val)
}

// Semantic Convention for FaaS scheduled to be executed regularly.
const (
	// FaaSTimeKey is the attribute Key conforming to the "faas.time" semantic
	// conventions. It represents a string containing the function invocation
	// time in the [ISO
	// 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format
	// expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '2020-01-23T13:47:06Z'
	FaaSTimeKey = attribute.Key("faas.time")

	// FaaSCronKey is the attribute Key conforming to the "faas.cron" semantic
	// conventions. It represents a string containing the schedule period as
	// [Cron
	// Expression](https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expressions.htm).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '0/5 * * * ? *'
	FaaSCronKey = attribute.Key("faas.cron")
)

// FaaSTime returns an attribute KeyValue conforming to the "faas.time"
// semantic conventions. It represents a string containing the function
// invocation time in the [ISO
// 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format
// expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).
func FaaSTime(val string) attribute.KeyValue {
	return FaaSTimeKey.String(val)
}

// FaaSCron returns an attribute KeyValue conforming to the "faas.cron"
// semantic conventions. It represents a string containing the schedule period
// as [Cron
	// HTTP client used has a version of `0.27.2`, but sends HTTP version
	// `1.1`, this attribute should be set to `1.1`.
	NetAppProtocolVersionKey = attribute.Key("net.app.protocol.version")

	// NetSockPeerNameKey is the attribute Key conforming to the
	// "net.sock.peer.name" semantic conventions. It represents the remote
	// socket peer name.
	//
	// Type: string
	// RequirementLevel: Recommended (If available and different from
	// `net.peer.name` and if `net.sock.peer.addr` is set.)
	// Stability: stable
	// Examples: 'proxy.example.com'
	NetSockPeerNameKey = attribute.Key("net.sock.peer.name")

	// NetSockPeerAddrKey is the attribute Key conforming to the
	// "net.sock.peer.addr" semantic conventions. It represents the remote
	// socket peer address: IPv4 or IPv6 for internet protocols, path for local
	// communication,
	// [etc](https://man7.org/linux/man-pages/man7/address_families.7.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '127.0.0.1', '/tmp/mysql.sock'
	NetSockPeerAddrKey = attribute.Key("net.sock.peer.addr")

	// NetSockPeerPortKey is the attribute Key conforming to the
	// "net.sock.peer.port" semantic conventions. It represents the remote
	// socket peer port.
	//
	// Type: int
	// RequirementLevel: Recommended (If defined for the address family and if
	// different than `net.peer.port` and if `net.sock.peer.addr` is set.)
	// Stability: stable
	// Examples: 16456
	NetSockPeerPortKey = attribute.Key("net.sock.peer.port")

	// NetSockFamilyKey is the attribute Key conforming to the
	// "net.sock.family" semantic conventions. It represents the protocol
	// [address
	// family](https://man7.org/linux/man-pages/man7/address_families.7.html)
	// which is used for communication.
	//
	// Type: Enum
	// RequirementLevel: ConditionallyRequired (If different than `inet` and if
	// any of `net.sock.peer.addr` or `net.sock.host.addr` are set. Consumers
	// of telemetry SHOULD accept both IPv4 and IPv6 formats for the address in
	// `net.sock.peer.addr` if `net.sock.family` is not set. This is to support
	// instrumentations that follow previous versions of this document.)
	// Stability: stable
	// Examples: 'inet6', 'bluetooth'
	NetSockFamilyKey = attribute.Key("net.sock.family")

	// NetPeerNameKey is the attribute Key conforming to the "net.peer.name"
	// semantic conventions. It represents the logical remote hostname, see
	// note below.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'example.com'
	// Note: `net.peer.name` SHOULD NOT be set if capturing it would require an
	// extra DNS lookup.
	NetPeerNameKey = attribute.Key("net.peer.name")

	// NetPeerPortKey is the attribute Key conforming to the "net.peer.port"
	// semantic conventions. It represents the logical remote port number
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 80, 8080, 443
	NetPeerPortKey = attribute.Key("net.peer.port")

	// NetHostNameKey is the attribute Key conforming to the "net.host.name"
	// semantic conventions. It represents the logical local hostname or
	// similar, see note below.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'localhost'
	NetHostNameKey = attribute.Key("net.host.name")

	// NetHostPortKey is the attribute Key conforming to the "net.host.port"
	// semantic conventions. It represents the logical local port number,
	// preferably the one that the peer used to connect
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 8080
	NetHostPortKey = attribute.Key("net.host.port")

	// NetSockHostAddrKey is the attribute Key conforming to the
	// "net.sock.host.addr" semantic conventions. It represents the local
	// socket address. Useful in case of a multi-IP host.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '192.168.0.1'
	NetSockHostAddrKey = attribute.Key("net.sock.host.addr")

	// NetSockHostPortKey is the attribute Key conforming to the
	// "net.sock.host.port" semantic conventions. It represents the local
	// socket port number.
	//
	// Type: int
	// RequirementLevel: Recommended (If defined for the address family and if
	// different than `net.host.port` and if `net.sock.host.addr` is set.)
	// Stability: stable
	// Examples: 35555
	NetSockHostPortKey = attribute.Key("net.sock.host.port")

	// NetHostConnectionTypeKey is the attribute Key conforming to the
	// "net.host.connection.type" semantic conventions. It represents the
	// internet connection type currently being used by the host.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'wifi'
	NetHostConnectionTypeKey = attribute.Key("net.host.connection.type")

	// NetHostConnectionSubtypeKey is the attribute Key conforming to the
	// "net.host.connection.subtype" semantic conventions. It represents the
	// this describes more details regarding the connection.type. It may be the
	// type of cell technology connection, but it could be used for describing
	// details about a wifi connection.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'LTE'
	NetHostConnectionSubtypeKey = attribute.Key("net.host.connection.subtype")

	// NetHostCarrierNameKey is the attribute Key conforming to the
	// "net.host.carrier.name" semantic conventions. It represents the name of
	// the mobile carrier.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'sprint'
	NetHostCarrierNameKey = attribute.Key("net.host.carrier.name")

	// NetHostCarrierMccKey is the attribute Key conforming to the
	// "net.host.carrier.mcc" semantic conventions. It represents the mobile
	// carrier country code.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '310'
	NetHostCarrierMccKey = attribute.Key("net.host.carrier.mcc")

	// NetHostCarrierMncKey is the attribute Key conforming to the
	// "net.host.carrier.mnc" semantic conventions. It represents the mobile
	// carrier network code.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '001'
	NetHostCarrierMncKey = attribute.Key("net.host.carrier.mnc")

	// NetHostCarrierIccKey is the attribute Key conforming to the
	// "net.host.carrier.icc" semantic conventions. It represents the ISO
	// 3166-1 alpha-2 2-character country code associated with the mobile
	// carrier network.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'DE'
	NetHostCarrierIccKey = attribute.Key("net.host.carrier.icc")
)

var (
	// ip_tcp
	NetTransportTCP = NetTransportKey.String("ip_tcp")
	// ip_udp
	NetTransportUDP = NetTransportKey.String("ip_udp")
	// Named or anonymous pipe. See note below
	NetTransportPipe = NetTransportKey.String("pipe")
	// In-process communication
	NetTransportInProc = NetTransportKey.String("inproc")
	// Something else (non IP-based)
	NetTransportOther = NetTransportKey.String("other")
)

var (
	// IPv4 address
	NetSockFamilyInet = NetSockFamilyKey.String("inet")
	// IPv6 address
	NetSockFamilyInet6 = NetSockFamilyKey.String("inet6")
	// Unix domain socket path
	NetSockFamilyUnix = NetSockFamilyKey.String("unix")
)

var (
	PeerServiceKey = attribute.Key("peer.service")
)

// PeerService returns an attribute KeyValue conforming to the
// "peer.service" semantic conventions. It represents the
// [`service.name`](../../resource/semantic_conventions/README.md#service) of
// the remote service. SHOULD be equal to the actual `service.name` resource
// attribute of the remote service if any.
func PeerService(val string) attribute.KeyValue {
	return PeerServiceKey.String(val)
}

// These attributes may be used for any operation with an authenticated and/or
// authorized enduser.
const (
	// EnduserIDKey is the attribute Key conforming to the "enduser.id"
	// semantic conventions. It represents the username or client_id extracted
	// from the access token or
	// [Authorization](https://tools.ietf.org/html/rfc7235#section-4.2) header
	// in the inbound request from outside the system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'username'
	EnduserIDKey = attribute.Key("enduser.id")

	// EnduserRoleKey is the attribute Key conforming to the "enduser.role"
	// semantic conventions. It represents the actual/assumed role the client
	// is making the request under extracted from token or application security
	// context.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'admin'
	EnduserRoleKey = attribute.Key("enduser.role")

	// EnduserScopeKey is the attribute Key conforming to the "enduser.scope"
	// semantic conventions. It represents the scopes or granted authorities
	// the client currently possesses extracted from token or application
	// security context. The value would come from the scope associated with an
	// [OAuth 2.0 Access
	// Token](https://tools.ietf.org/html/rfc6749#section-3.3) or an attribute
	// value in a [SAML 2.0
	// Assertion](http://docs.oasis-open.org/security/saml/Post2.0/sstc-saml-tech-overview-2.0.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'read:message, write:files'
	EnduserScopeKey = attribute.Key("enduser.scope")
)

// EnduserID returns an attribute KeyValue conforming to the "enduser.id"
// semantic conventions. It represents the username or client_id extracted from
// the access token or
// [Authorization](https://tools.ietf.org/html/rfc7235#section-4.2) header in
// the inbound request from outside the system.
func EnduserID(val string) attribute.KeyValue {
	return EnduserIDKey.String(val)
}
	// Stability: stable
	// Examples: 'serveRequest'
	CodeFunctionKey = attribute.Key("code.function")

	// CodeNamespaceKey is the attribute Key conforming to the "code.namespace"
	// semantic conventions. It represents the "namespace" within which
	// `code.function` is defined. Usually the qualified class or module name,
	// such that `code.namespace` + some separator + `code.function` form a
	// unique identifier for the code unit.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'com.example.MyHTTPService'
	CodeNamespaceKey = attribute.Key("code.namespace")

	// CodeFilepathKey is the attribute Key conforming to the "code.filepath"
	// semantic conventions. It represents the source code file name that
	// identifies the code unit as uniquely as possible (preferably an absolute
	// file path).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '/usr/local/MyApplication/content_root/app/index.php'
	CodeFilepathKey = attribute.Key("code.filepath")

	// CodeLineNumberKey is the attribute Key conforming to the "code.lineno"
	// semantic conventions. It represents the line number in `code.filepath`
	// best representing the operation. It SHOULD point within the code unit
	// named in `code.function`.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 42
	CodeLineNumberKey = attribute.Key("code.lineno")

	// CodeColumnKey is the attribute Key conforming to the "code.column"
	// semantic conventions. It represents the column number in `code.filepath`
	// best representing the operation. It SHOULD point within the code unit
	// named in `code.function`.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 16
	CodeColumnKey = attribute.Key("code.column")
)

// CodeFunction returns an attribute KeyValue conforming to the
// "code.function" semantic conventions. It represents the method or function
// name, or equivalent (usually rightmost part of the code unit's name).
func CodeFunction(val string) attribute.KeyValue {
	return CodeFunctionKey.String(val)
}

// CodeNamespace returns an attribute KeyValue conforming to the
// "code.namespace" semantic conventions. It represents the "namespace" within
// which `code.function` is defined. Usually the qualified class or module
// name, such that `code.namespace` + some separator + `code.function` form a
// unique identifier for the code unit.

// HTTPResponseContentLength returns an attribute KeyValue conforming to the
// "http.response_content_length" semantic conventions. It represents the size
// of the response payload body in bytes. This is the number of bytes
// transferred excluding headers and is often, but not always, present as the
// [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length)
// header. For requests using transport encoding, this should be the compressed
// size.
func HTTPResponseContentLength(val int) attribute.KeyValue {
	return HTTPResponseContentLengthKey.Int(val)
}

// Semantic Convention for HTTP Client
const (
	// HTTPURLKey is the attribute Key conforming to the "http.url" semantic
	// conventions. It represents the full HTTP request URL in the form
	// `scheme://host[:port]/path?query[#fragment]`. Usually the fragment is
	// not transmitted over HTTP, but if it is known, it should be included
	// nevertheless.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: 'https://www.foo.bar/search?q=OpenTelemetry#SemConv'
	// Note: `http.url` MUST NOT contain credentials passed via URL in form of
	// `https://username:password@www.example.com/`. In such case the
	// attribute's value should be `https://www.example.com/`.
	HTTPURLKey = attribute.Key("http.url")

	// HTTPResendCountKey is the attribute Key conforming to the
	// "http.resend_count" semantic conventions. It represents the ordinal
	// number of request resending attempt (for any reason, including
	// redirects).
	//
	// Type: int
	// RequirementLevel: Recommended (if and only if request was retried.)
	// Stability: stable
	// Examples: 3
	// Note: The resend count SHOULD be updated each time an HTTP request gets
	// resent by the client, regardless of what was the cause of the resending
	// (e.g. redirection, authorization failure, 503 Server Unavailable,
	// network issues, or any other).
	HTTPResendCountKey = attribute.Key("http.resend_count")
)

// HTTPURL returns an attribute KeyValue conforming to the "http.url"
// semantic conventions. It represents the full HTTP request URL in the form
// `scheme://host[:port]/path?query[#fragment]`. Usually the fragment is not
// transmitted over HTTP, but if it is known, it should be included
// nevertheless.
func HTTPURL(val string) attribute.KeyValue {

	// AWSDynamoDBItemCollectionMetricsKey is the attribute Key conforming to
	// the "aws.dynamodb.item_collection_metrics" semantic conventions. It
	// represents the JSON-serialized value of the `ItemCollectionMetrics`
	// response field.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '{ "string" : [ { "ItemCollectionKey": { "string" : { "B":
	// blob, "BOOL": boolean, "BS": [ blob ], "L": [ "AttributeValue" ], "M": {
	// "string" : "AttributeValue" }, "N": "string", "NS": [ "string" ],
	// "NULL": boolean, "S": "string", "SS": [ "string" ] } },
	// "SizeEstimateRangeGB": [ number ] } ] }'
	AWSDynamoDBItemCollectionMetricsKey = attribute.Key("aws.dynamodb.item_collection_metrics")

	// AWSDynamoDBProvisionedReadCapacityKey is the attribute Key conforming to
	// the "aws.dynamodb.provisioned_read_capacity" semantic conventions. It
	// represents the value of the `ProvisionedThroughput.ReadCapacityUnits`
	// request parameter.
	//
	// Type: double
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 1.0, 2.0
	AWSDynamoDBProvisionedReadCapacityKey = attribute.Key("aws.dynamodb.provisioned_read_capacity")

	// AWSDynamoDBProvisionedWriteCapacityKey is the attribute Key conforming
	// to the "aws.dynamodb.provisioned_write_capacity" semantic conventions.
	// It represents the value of the
	// `ProvisionedThroughput.WriteCapacityUnits` request parameter.
	//
	// Type: double
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 1.0, 2.0
	AWSDynamoDBProvisionedWriteCapacityKey = attribute.Key("aws.dynamodb.provisioned_write_capacity")

	// AWSDynamoDBConsistentReadKey is the attribute Key conforming to the
	// "aws.dynamodb.consistent_read" semantic conventions. It represents the
	// value of the `ConsistentRead` request parameter.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	AWSDynamoDBConsistentReadKey = attribute.Key("aws.dynamodb.consistent_read")

	// AWSDynamoDBProjectionKey is the attribute Key conforming to the
	// "aws.dynamodb.projection" semantic conventions. It represents the value
	// of the `ProjectionExpression` request parameter.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'Title', 'Title, Price, Color', 'Title, Description,
	// RelatedItems, ProductReviews'
	AWSDynamoDBProjectionKey = attribute.Key("aws.dynamodb.projection")

	// AWSDynamoDBLimitKey is the attribute Key conforming to the
	// "aws.dynamodb.limit" semantic conventions. It represents the value of
	// the `Limit` request parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 10
	AWSDynamoDBLimitKey = attribute.Key("aws.dynamodb.limit")

	// AWSDynamoDBAttributesToGetKey is the attribute Key conforming to the
	// "aws.dynamodb.attributes_to_get" semantic conventions. It represents the
	// value of the `AttributesToGet` request parameter.
	// value of the `ScanIndexForward` request parameter.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	AWSDynamoDBScanForwardKey = attribute.Key("aws.dynamodb.scan_forward")
)

// AWSDynamoDBScanForward returns an attribute KeyValue conforming to the
// "aws.dynamodb.scan_forward" semantic conventions. It represents the value of
// the `ScanIndexForward` request parameter.
func AWSDynamoDBScanForward(val bool) attribute.KeyValue {
	return AWSDynamoDBScanForwardKey.Bool(val)
}

// DynamoDB.Scan
const (
	// AWSDynamoDBSegmentKey is the attribute Key conforming to the
	// "aws.dynamodb.segment" semantic conventions. It represents the value of
	// the `Segment` request parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 10
	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	// AWSDynamoDBTotalSegmentsKey is the attribute Key conforming to the
	// "aws.dynamodb.total_segments" semantic conventions. It represents the
	// value of the `TotalSegments` request parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 100
	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")

	// AWSDynamoDBCountKey is the attribute Key conforming to the
	// "aws.dynamodb.count" semantic conventions. It represents the value of
	// the `Count` response parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 10
	AWSDynamoDBCountKey = attribute.Key("aws.dynamodb.count")

	// AWSDynamoDBScannedCountKey is the attribute Key conforming to the
	// "aws.dynamodb.scanned_count" semantic conventions. It represents the
	// value of the `ScannedCount` response parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 50
	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")
)

// AWSDynamoDBSegment returns an attribute KeyValue conforming to the
// "aws.dynamodb.segment" semantic conventions. It represents the value of the
// `Segment` request parameter.
func AWSDynamoDBSegment(val int) attribute.KeyValue {
	return AWSDynamoDBSegmentKey.Int(val)
}

// AWSDynamoDBTotalSegments returns an attribute KeyValue conforming to the
// "aws.dynamodb.total_segments" semantic conventions. It represents the value
// of the `TotalSegments` request parameter.
func AWSDynamoDBTotalSegments(val int) attribute.KeyValue {
	return AWSDynamoDBTotalSegmentsKey.Int(val)
}

// AWSDynamoDBCount returns an attribute KeyValue conforming to the
// "aws.dynamodb.count" semantic conventions. It represents the value of the
// `Count` response parameter.
func AWSDynamoDBCount(val int) attribute.KeyValue {
	return AWSDynamoDBCountKey.Int(val)
}

// AWSDynamoDBScannedCount returns an attribute KeyValue conforming to the
// "aws.dynamodb.scanned_count" semantic conventions. It represents the value

// Semantic convention describing per-message attributes populated on messaging
// spans or links.
const (
	// MessagingMessageIDKey is the attribute Key conforming to the
	// "messaging.message.id" semantic conventions. It represents a value used
	// by the messaging system as an identifier for the message, represented as
	// a string.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '452a7c7c7c7048c2f887f61572b18fc2'
	MessagingMessageIDKey = attribute.Key("messaging.message.id")

	// MessagingMessageConversationIDKey is the attribute Key conforming to the
	// "messaging.message.conversation_id" semantic conventions. It represents
	// the [conversation ID](#conversations) identifying the conversation to
	// which the message belongs, represented as a string. Sometimes called
	// "Correlation ID".
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'MyConversationID'
	MessagingMessageConversationIDKey = attribute.Key("messaging.message.conversation_id")

	// MessagingMessagePayloadSizeBytesKey is the attribute Key conforming to
	// the "messaging.message.payload_size_bytes" semantic conventions. It
	// represents the (uncompressed) size of the message payload in bytes. Also
	// use this attribute if it is unknown whether the compressed or
	// uncompressed payload size is reported.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 2738
	MessagingMessagePayloadSizeBytesKey = attribute.Key("messaging.message.payload_size_bytes")

	// MessagingMessagePayloadCompressedSizeBytesKey is the attribute Key
	// conforming to the "messaging.message.payload_compressed_size_bytes"
	// semantic conventions. It represents the compressed size of the message
	// payload in bytes.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 2048
	MessagingMessagePayloadCompressedSizeBytesKey = attribute.Key("messaging.message.payload_compressed_size_bytes")
)

}

// Semantic convention for attributes that describe messaging source on broker
const (
	// MessagingSourceNameKey is the attribute Key conforming to the
	// "messaging.source.name" semantic conventions. It represents the message
	// source name
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'MyQueue', 'MyTopic'
	// Note: Source name SHOULD uniquely identify a specific queue, topic, or
	// other entity within the broker. If
	// the broker does not have such notion, the source name SHOULD uniquely
	// identify the broker.
	MessagingSourceNameKey = attribute.Key("messaging.source.name")

	// MessagingSourceKindKey is the attribute Key conforming to the
	// "messaging.source.kind" semantic conventions. It represents the kind of
	// message source
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	MessagingSourceKindKey = attribute.Key("messaging.source.kind")

	// MessagingSourceTemplateKey is the attribute Key conforming to the
	// "messaging.source.template" semantic conventions. It represents the low
	// cardinality representation of the messaging source name
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '/customers/{customerID}'
	// Note: Source names could be constructed from templates. An example would
	// be a source name involving a user name or product id. Although the
	// source name in this case is of high cardinality, the underlying template
	// is of low cardinality and can be effectively used for grouping and
	// aggregation.
	MessagingSourceTemplateKey = attribute.Key("messaging.source.template")

	// MessagingSourceTemporaryKey is the attribute Key conforming to the
	// "messaging.source.temporary" semantic conventions. It represents a
	// boolean that is true if the message source is temporary and might not
	// exist anymore after messages are processed.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	MessagingSourceTemporaryKey = attribute.Key("messaging.source.temporary")
	// Examples: 'myKey'
	// Note: If the key type is not string, it's string representation has to
	// be supplied for the attribute. If the key has no unambiguous, canonical
	// string form, don't include its value.
	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	// MessagingKafkaConsumerGroupKey is the attribute Key conforming to the
	// "messaging.kafka.consumer.group" semantic conventions. It represents the
	// name of the Kafka Consumer Group that is handling the message. Only
	// applies to consumers, not producers.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'my-group'
	MessagingKafkaConsumerGroupKey = attribute.Key("messaging.kafka.consumer.group")

	// MessagingKafkaClientIDKey is the attribute Key conforming to the
	// "messaging.kafka.client_id" semantic conventions. It represents the
	// client ID for the Consumer or Producer that is handling the message.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'client-5'
	MessagingKafkaClientIDKey = attribute.Key("messaging.kafka.client_id")

	// MessagingKafkaDestinationPartitionKey is the attribute Key conforming to
	// the "messaging.kafka.destination.partition" semantic conventions. It
	// represents the partition the message is sent to.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 2
	MessagingKafkaDestinationPartitionKey = attribute.Key("messaging.kafka.destination.partition")

	// MessagingKafkaSourcePartitionKey is the attribute Key conforming to the
	// "messaging.kafka.source.partition" semantic conventions. It represents
	// the partition the message is received from.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 2
	MessagingKafkaSourcePartitionKey = attribute.Key("messaging.kafka.source.partition")

	// MessagingKafkaMessageOffsetKey is the attribute Key conforming to the
	// "messaging.kafka.message.offset" semantic conventions. It represents the
	// offset of a record in the corresponding Kafka partition.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 42
	MessagingKafkaMessageOffsetKey = attribute.Key("messaging.kafka.message.offset")

	// MessagingKafkaMessageTombstoneKey is the attribute Key conforming to the
	// "messaging.kafka.message.tombstone" semantic conventions. It represents
	// a boolean that is true if the message is a tombstone.
	//
	// Type: boolean
	// RequirementLevel: ConditionallyRequired (If value is `true`. When
	// missing, the value is assumed to be `false`.)
	// Stability: stable
	MessagingKafkaMessageTombstoneKey = attribute.Key("messaging.kafka.message.tombstone")
)

// MessagingKafkaMessageKey returns an attribute KeyValue conforming to the
// "messaging.kafka.message.key" semantic conventions. It represents the
// message keys in Kafka are used for grouping alike messages to ensure they're
// processed on the same partition. They differ from `messaging.message.id` in
// that they're not unique. If the key is `null`, the attribute MUST NOT be
// set.
func MessagingKafkaMessageKey(val string) attribute.KeyValue {
	return MessagingKafkaMessageKeyKey.String(val)
}

// MessagingKafkaConsumerGroup returns an attribute KeyValue conforming to
// the "messaging.kafka.consumer.group" semantic conventions. It represents the
// name of the Kafka Consumer Group that is handling the message. Only applies
// to consumers, not producers.
func MessagingKafkaConsumerGroup(val string) attribute.KeyValue {
	return MessagingKafkaConsumerGroupKey.String(val)
}

// MessagingKafkaClientID returns an attribute KeyValue conforming to the
// "messaging.kafka.client_id" semantic conventions. It represents the client
// ID for the Consumer or Producer that is handling the message.
func MessagingKafkaClientID(val string) attribute.KeyValue {
	return MessagingKafkaClientIDKey.String(val)
// MessagingKafkaMessageOffset returns an attribute KeyValue conforming to
// the "messaging.kafka.message.offset" semantic conventions. It represents the
// offset of a record in the corresponding Kafka partition.
func MessagingKafkaMessageOffset(val int) attribute.KeyValue {
	return MessagingKafkaMessageOffsetKey.Int(val)
}

// MessagingKafkaMessageTombstone returns an attribute KeyValue conforming
// to the "messaging.kafka.message.tombstone" semantic conventions. It
// represents a boolean that is true if the message is a tombstone.
func MessagingKafkaMessageTombstone(val bool) attribute.KeyValue {
	return MessagingKafkaMessageTombstoneKey.Bool(val)
}

// Attributes for Apache RocketMQ
const (
	// MessagingRocketmqNamespaceKey is the attribute Key conforming to the
	// "messaging.rocketmq.namespace" semantic conventions. It represents the
	// namespace of RocketMQ resources, resources in different namespaces are
	// individual.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: 'myNamespace'
	MessagingRocketmqNamespaceKey = attribute.Key("messaging.rocketmq.namespace")

	// MessagingRocketmqClientGroupKey is the attribute Key conforming to the
	// "messaging.rocketmq.client_group" semantic conventions. It represents
	// the name of the RocketMQ producer/consumer group that is handling the
	// message. The client type is identified by the SpanKind.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: 'myConsumerGroup'
	MessagingRocketmqClientGroupKey = attribute.Key("messaging.rocketmq.client_group")

	// MessagingRocketmqClientIDKey is the attribute Key conforming to the
	// "messaging.rocketmq.client_id" semantic conventions. It represents the
	// unique identifier for each client.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: 'myhost@8742@s8083jm'
	MessagingRocketmqClientIDKey = attribute.Key("messaging.rocketmq.client_id")

	// MessagingRocketmqMessageDeliveryTimestampKey is the attribute Key
	// conforming to the "messaging.rocketmq.message.delivery_timestamp"
	// semantic conventions. It represents the timestamp in milliseconds that
	// the delay message is expected to be delivered to consumer.
	//
	// Type: int
	// RequirementLevel: ConditionallyRequired (If the message type is delay
	// and delay time level is not specified.)
	// Stability: stable
	// Examples: 1665987217045
	MessagingRocketmqMessageDeliveryTimestampKey = attribute.Key("messaging.rocketmq.message.delivery_timestamp")

	// MessagingRocketmqMessageDelayTimeLevelKey is the attribute Key
	// conforming to the "messaging.rocketmq.message.delay_time_level" semantic
	// conventions. It represents the delay time level for delay message, which
	// determines the message delay time.
	//
	// Type: int
	// RequirementLevel: ConditionallyRequired (If the message type is delay
	// and delivery timestamp is not specified.)
	// Stability: stable
	// Examples: 3
	MessagingRocketmqMessageDelayTimeLevelKey = attribute.Key("messaging.rocketmq.message.delay_time_level")

	// group.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (If the message type is FIFO.)
	// Stability: stable
	// Examples: 'myMessageGroup'
	MessagingRocketmqMessageGroupKey = attribute.Key("messaging.rocketmq.message.group")

	// MessagingRocketmqMessageTypeKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.type" semantic conventions. It represents
	// the type of message.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	MessagingRocketmqMessageTypeKey = attribute.Key("messaging.rocketmq.message.type")

	// MessagingRocketmqMessageTagKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.tag" semantic conventions. It represents the
	// secondary classifier of message besides topic.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'tagA'
	MessagingRocketmqMessageTagKey = attribute.Key("messaging.rocketmq.message.tag")

	// MessagingRocketmqMessageKeysKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.keys" semantic conventions. It represents
	// the key(s) of message, another way to mark message besides message id.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'keyA', 'keyB'
	MessagingRocketmqMessageKeysKey = attribute.Key("messaging.rocketmq.message.keys")

	// MessagingRocketmqConsumptionModelKey is the attribute Key conforming to
	// the "messaging.rocketmq.consumption_model" semantic conventions. It
	// represents the model of message consumption. This only applies to
	// consumer spans.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	MessagingRocketmqConsumptionModelKey = attribute.Key("messaging.rocketmq.consumption_model")
)

var (
	// Normal message
	MessagingRocketmqMessageTypeNormal = MessagingRocketmqMessageTypeKey.String("normal")
	// FIFO message
	MessagingRocketmqMessageTypeFifo = MessagingRocketmqMessageTypeKey.String("fifo")
	// Delay message
	MessagingRocketmqMessageTypeDelay = MessagingRocketmqMessageTypeKey.String("delay")
	// Transaction message
	MessagingRocketmqMessageTypeTransaction = MessagingRocketmqMessageTypeKey.String("transaction")
)

var (
	// Clustering consumption model
	MessagingRocketmqConsumptionModelClustering = MessagingRocketmqConsumptionModelKey.String("clustering")
	// Broadcasting consumption model
	MessagingRocketmqConsumptionModelBroadcasting = MessagingRocketmqConsumptionModelKey.String("broadcasting")
)

// MessagingRocketmqNamespace returns an attribute KeyValue conforming to
// the "messaging.rocketmq.namespace" semantic conventions. It represents the
// namespace of RocketMQ resources, resources in different namespaces are
// individual.
func MessagingRocketmqNamespace(val string) attribute.KeyValue {
// RPCService returns an attribute KeyValue conforming to the "rpc.service"
// semantic conventions. It represents the full (logical) name of the service
// being called, including its package name, if applicable.
func RPCService(val string) attribute.KeyValue {
	return RPCServiceKey.String(val)
}

// RPCMethod returns an attribute KeyValue conforming to the "rpc.method"
// semantic conventions. It represents the name of the (logical) method being
// called, must be equal to the $method part in the span name.
func RPCMethod(val string) attribute.KeyValue {
	return RPCMethodKey.String(val)
}

// Tech-specific attributes for gRPC.
const (
	// RPCGRPCStatusCodeKey is the attribute Key conforming to the
	// "rpc.grpc.status_code" semantic conventions. It represents the [numeric
	// status
	// code](https://github.com/grpc/grpc/blob/v1.33.2/doc/statuscodes.md) of
	// the gRPC request.
	//
	// Type: Enum
	// RequirementLevel: Required
	// Stability: stable
	RPCGRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")
)

var (
	// OK
	RPCGRPCStatusCodeOk = RPCGRPCStatusCodeKey.Int(0)
	// CANCELLED
	RPCGRPCStatusCodeCancelled = RPCGRPCStatusCodeKey.Int(1)
	// UNKNOWN
	RPCGRPCStatusCodeUnknown = RPCGRPCStatusCodeKey.Int(2)
	// INVALID_ARGUMENT
	RPCGRPCStatusCodeInvalidArgument = RPCGRPCStatusCodeKey.Int(3)
	// DEADLINE_EXCEEDED
	RPCGRPCStatusCodeDeadlineExceeded = RPCGRPCStatusCodeKey.Int(4)
	// NOT_FOUND
	RPCGRPCStatusCodeNotFound = RPCGRPCStatusCodeKey.Int(5)
	// ALREADY_EXISTS
	RPCGRPCStatusCodeAlreadyExists = RPCGRPCStatusCodeKey.Int(6)
	// PERMISSION_DENIED
	RPCGRPCStatusCodePermissionDenied = RPCGRPCStatusCodeKey.Int(7)
	// RESOURCE_EXHAUSTED
	RPCGRPCStatusCodeResourceExhausted = RPCGRPCStatusCodeKey.Int(8)
	// FAILED_PRECONDITION
	RPCGRPCStatusCodeFailedPrecondition = RPCGRPCStatusCodeKey.Int(9)
	// ABORTED
	RPCGRPCStatusCodeAborted = RPCGRPCStatusCodeKey.Int(10)