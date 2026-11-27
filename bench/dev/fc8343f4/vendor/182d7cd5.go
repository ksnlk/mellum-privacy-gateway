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
	// Examples: 'ec93ffc1-2527-2d55-520d-18d10fd4a3de'
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
	// Note: It is not recommended to attempt any client-side parsing of
	// `db.statement` just to get this property, but it should be set if it is
	// provided by the library being instrumented. If the operation is acting
	// upon an anonymous table, or more than one table, this value MUST NOT be
	// set.
	DBSQLTableKey = attribute.Key("db.sql.table")
)

// DBSQLTable returns an attribute KeyValue conforming to the "db.sql.table"
// semantic conventions. It represents the name of the primary table that the
// operation is acting upon, including the database name (if applicable).
func DBSQLTable(val string) attribute.KeyValue {
	return DBSQLTableKey.String(val)
}

// Call-level attributes for Cosmos DB.
const (
	// DBCosmosDBClientIDKey is the attribute Key conforming to the
	// "db.cosmosdb.client_id" semantic conventions. It represents the unique
	// Cosmos client instance id.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '1cc0191c-5024-347f-b20d-29c64960b62f'
	DBCosmosDBClientIDKey = attribute.Key("db.cosmosdb.client_id")

	// DBCosmosDBOperationTypeKey is the attribute Key conforming to the
	// "db.cosmosdb.operation_type" semantic conventions. It represents the
	// cosmosDB Operation Type.
	//
	// Type: Enum
	// RequirementLevel: ConditionallyRequired (when performing one of the
	// operations in this list)
	// Stability: stable
	DBCosmosDBOperationTypeKey = attribute.Key("db.cosmosdb.operation_type")

	// DBCosmosDBConnectionModeKey is the attribute Key conforming to the
	// "db.cosmosdb.connection_mode" semantic conventions. It represents the
	// cosmos client connection mode.
	//
	// Type: Enum
	// RequirementLevel: ConditionallyRequired (if not `direct` (or pick gw as
	// default))
	// Stability: stable
	DBCosmosDBConnectionModeKey = attribute.Key("db.cosmosdb.connection_mode")

	// DBCosmosDBContainerKey is the attribute Key conforming to the
	// "db.cosmosdb.container" semantic conventions. It represents the cosmos
	// DB container name.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (if available)
	// Stability: stable
	// Examples: 'anystring'
	DBCosmosDBContainerKey = attribute.Key("db.cosmosdb.container")

	// DBCosmosDBRequestContentLengthKey is the attribute Key conforming to the
	// "db.cosmosdb.request_content_length" semantic conventions. It represents
	// the request payload size in bytes
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: stable
	DBCosmosDBRequestContentLengthKey = attribute.Key("db.cosmosdb.request_content_length")

	// DBCosmosDBStatusCodeKey is the attribute Key conforming to the
	// "db.cosmosdb.status_code" semantic conventions. It represents the cosmos
	// DB status code.
	//
	// Type: int
	// RequirementLevel: ConditionallyRequired (if response was received)
	// Stability: stable
	// Examples: 200, 201
	DBCosmosDBStatusCodeKey = attribute.Key("db.cosmosdb.status_code")

	// DBCosmosDBSubStatusCodeKey is the attribute Key conforming to the
	// "db.cosmosdb.sub_status_code" semantic conventions. It represents the
	// cosmos DB sub status code.
	//
	// Type: int
	// RequirementLevel: ConditionallyRequired (when response was received and
	// contained sub-code.)
	// Stability: stable
	// Examples: 1000, 1002
	DBCosmosDBSubStatusCodeKey = attribute.Key("db.cosmosdb.sub_status_code")

	// DBCosmosDBRequestChargeKey is the attribute Key conforming to the
	// "db.cosmosdb.request_charge" semantic conventions. It represents the rU
	// consumed for that operation
	//
	// Type: double
	// RequirementLevel: ConditionallyRequired (when available)
	// Stability: stable
	// Examples: 46.18, 1.0
	DBCosmosDBRequestChargeKey = attribute.Key("db.cosmosdb.request_charge")
)

var (
	// invalid
	DBCosmosDBOperationTypeInvalid = DBCosmosDBOperationTypeKey.String("Invalid")
	// create
	DBCosmosDBOperationTypeCreate = DBCosmosDBOperationTypeKey.String("Create")
	// patch
	DBCosmosDBOperationTypePatch = DBCosmosDBOperationTypeKey.String("Patch")
	// read
	DBCosmosDBOperationTypeRead = DBCosmosDBOperationTypeKey.String("Read")
	// read_feed
	DBCosmosDBOperationTypeReadFeed = DBCosmosDBOperationTypeKey.String("ReadFeed")
	// delete
	DBCosmosDBOperationTypeDelete = DBCosmosDBOperationTypeKey.String("Delete")
	// replace
	DBCosmosDBOperationTypeReplace = DBCosmosDBOperationTypeKey.String("Replace")
	// execute
	DBCosmosDBOperationTypeExecute = DBCosmosDBOperationTypeKey.String("Execute")
	// query
	DBCosmosDBOperationTypeQuery = DBCosmosDBOperationTypeKey.String("Query")
	// head
	DBCosmosDBOperationTypeHead = DBCosmosDBOperationTypeKey.String("Head")
	// head_feed
	DBCosmosDBOperationTypeHeadFeed = DBCosmosDBOperationTypeKey.String("HeadFeed")
	// semantic conventions. It represents the type of the trigger which caused
	// this function invocation.
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

	// FaaSInvocationIDKey is the attribute Key conforming to the
	// "faas.invocation_id" semantic conventions. It represents the invocation
	// ID of the current function invocation.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'ab2f8ed2-f642-1e8b-b19b-669f41f3dd50'
	FaaSInvocationIDKey = attribute.Key("faas.invocation_id")
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

// FaaSInvocationID returns an attribute KeyValue conforming to the
// "faas.invocation_id" semantic conventions. It represents the invocation ID
// of the current function invocation.
func FaaSInvocationID(val string) attribute.KeyValue {
	return FaaSInvocationIDKey.String(val)
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
func CodeLineNumber(val int) attribute.KeyValue {
	return CodeLineNumberKey.Int(val)
}

// CodeColumn returns an attribute KeyValue conforming to the "code.column"
// semantic conventions. It represents the column number in `code.filepath`
// best representing the operation. It SHOULD point within the code unit named
// in `code.function`.
func CodeColumn(val int) attribute.KeyValue {
	return CodeColumnKey.Int(val)
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

// AWSDynamoDBGlobalSecondaryIndexUpdates returns an attribute KeyValue
// conforming to the "aws.dynamodb.global_secondary_index_updates" semantic
// conventions. It represents the JSON-serialized value of each item in the the
// `GlobalSecondaryIndexUpdates` request field.
func AWSDynamoDBGlobalSecondaryIndexUpdates(val ...string) attribute.KeyValue {
	return AWSDynamoDBGlobalSecondaryIndexUpdatesKey.StringSlice(val)
}

// Attributes that exist for S3 request types.
const (
	// AWSS3BucketKey is the attribute Key conforming to the "aws.s3.bucket"
	// semantic conventions. It represents the S3 bucket name the request
	// refers to. Corresponds to the `--bucket` parameter of the [S3
	// API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html)
	// operations.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'some-bucket-name'
	// Note: The `bucket` attribute is applicable to all S3 operations that
	// reference a bucket, i.e. that require the bucket name as a mandatory
	// parameter.
	// This applies to almost all S3 operations except `list-buckets`.
	AWSS3BucketKey = attribute.Key("aws.s3.bucket")

	// AWSS3KeyKey is the attribute Key conforming to the "aws.s3.key" semantic
	// conventions. It represents the S3 object key the request refers to.
	// Corresponds to the `--key` parameter of the [S3
	// API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html)
	// operations.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'someFile.yml'
	// Note: The `key` attribute is applicable to all object-related S3
	// operations, i.e. that require the object key as a mandatory parameter.
	// This applies in particular to the following operations:
	//
	// -
	// [copy-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html)
	// -
	// [delete-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-object.html)
	// -
	// [get-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/get-object.html)
	// -
	// [head-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/head-object.html)
	// -
	// [put-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/put-object.html)
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