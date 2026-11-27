	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'fixed', 'sliding', 'token'
	AspnetcoreRateLimitingPolicyKey = attribute.Key("aspnetcore.rate_limiting.policy")

	// AspnetcoreRequestIsUnhandledKey is the attribute Key conforming to the
	// "aspnetcore.request.is_unhandled" semantic conventions. It represents
	// the flag indicating if request was handled by the application pipeline.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: True
	AspnetcoreRequestIsUnhandledKey = attribute.Key("aspnetcore.request.is_unhandled")

	// AspnetcoreRoutingIsFallbackKey is the attribute Key conforming to the
	// "aspnetcore.routing.is_fallback" semantic conventions. It represents a
	// value that indicates whether the matched route is a fallback route.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: True
	AspnetcoreRoutingIsFallbackKey = attribute.Key("aspnetcore.routing.is_fallback")

	// AspnetcoreRoutingMatchStatusKey is the attribute Key conforming to the
	// "aspnetcore.routing.match_status" semantic conventions. It represents
	// the match result - success or failure
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'success', 'failure'
	AspnetcoreRoutingMatchStatusKey = attribute.Key("aspnetcore.routing.match_status")
)

var (
	// Lease was acquired
	AspnetcoreRateLimitingResultAcquired = AspnetcoreRateLimitingResultKey.String("acquired")
	// Lease request was rejected by the endpoint limiter
	AspnetcoreRateLimitingResultEndpointLimiter = AspnetcoreRateLimitingResultKey.String("endpoint_limiter")
	// Lease request was rejected by the global limiter
	AspnetcoreRateLimitingResultGlobalLimiter = AspnetcoreRateLimitingResultKey.String("global_limiter")
	// Lease request was canceled
	AspnetcoreRateLimitingResultRequestCanceled = AspnetcoreRateLimitingResultKey.String("request_canceled")
)

var (
	// Exception was handled by the exception handling middleware
	// `AttributeDefinitions` request field.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '{ "AttributeName": "string", "AttributeType": "string" }'
	AWSDynamoDBAttributeDefinitionsKey = attribute.Key("aws.dynamodb.attribute_definitions")

	// AWSDynamoDBAttributesToGetKey is the attribute Key conforming to the
	// "aws.dynamodb.attributes_to_get" semantic conventions. It represents the
	// value of the `AttributesToGet` request parameter.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'lives', 'id'
	AWSDynamoDBAttributesToGetKey = attribute.Key("aws.dynamodb.attributes_to_get")

	// AWSDynamoDBConsistentReadKey is the attribute Key conforming to the
	// "aws.dynamodb.consistent_read" semantic conventions. It represents the
	// value of the `ConsistentRead` request parameter.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	AWSDynamoDBConsistentReadKey = attribute.Key("aws.dynamodb.consistent_read")

	// AWSDynamoDBConsumedCapacityKey is the attribute Key conforming to the
	// "aws.dynamodb.consumed_capacity" semantic conventions. It represents the
	// JSON-serialized value of each item in the `ConsumedCapacity` response
	// field.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '{ "CapacityUnits": number, "GlobalSecondaryIndexes": {
	// "string" : { "CapacityUnits": number, "ReadCapacityUnits": number,
	// "WriteCapacityUnits": number } }, "LocalSecondaryIndexes": { "string" :
	// { "CapacityUnits": number, "ReadCapacityUnits": number,
	// "WriteCapacityUnits": number } }, "ReadCapacityUnits": number, "Table":
	// { "CapacityUnits": number, "ReadCapacityUnits": number,
	// "WriteCapacityUnits": number }, "TableName": "string",
	// "WriteCapacityUnits": number }'
	AWSDynamoDBConsumedCapacityKey = attribute.Key("aws.dynamodb.consumed_capacity")

	// AWSDynamoDBCountKey is the attribute Key conforming to the
	// "aws.dynamodb.count" semantic conventions. It represents the value of
	// the `Count` response parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// "IndexSizeBytes": number, "ItemCount": number, "KeySchema": [ {
	// "AttributeName": "string", "KeyType": "string" } ], "Projection": {
	// "NonKeyAttributes": [ "string" ], "ProjectionType": "string" } }'
	AWSDynamoDBLocalSecondaryIndexesKey = attribute.Key("aws.dynamodb.local_secondary_indexes")

	// AWSDynamoDBProjectionKey is the attribute Key conforming to the
	// "aws.dynamodb.projection" semantic conventions. It represents the value
	// of the `ProjectionExpression` request parameter.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Title', 'Title, Price, Color', 'Title, Description,
	// RelatedItems, ProductReviews'
	AWSDynamoDBProjectionKey = attribute.Key("aws.dynamodb.projection")

	// AWSDynamoDBProvisionedReadCapacityKey is the attribute Key conforming to
	// the "aws.dynamodb.provisioned_read_capacity" semantic conventions. It
	// represents the value of the `ProvisionedThroughput.ReadCapacityUnits`
	// request parameter.
	//
	// Type: double
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1.0, 2.0
	AWSDynamoDBProvisionedReadCapacityKey = attribute.Key("aws.dynamodb.provisioned_read_capacity")

	// AWSDynamoDBProvisionedWriteCapacityKey is the attribute Key conforming
	// to the "aws.dynamodb.provisioned_write_capacity" semantic conventions.
	// It represents the value of the
	// `ProvisionedThroughput.WriteCapacityUnits` request parameter.
	//
	// Type: double
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1.0, 2.0
	AWSDynamoDBProvisionedWriteCapacityKey = attribute.Key("aws.dynamodb.provisioned_write_capacity")

	// AWSDynamoDBScanForwardKey is the attribute Key conforming to the
	// "aws.dynamodb.scan_forward" semantic conventions. It represents the
	// value of the `ScanIndexForward` request parameter.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	AWSDynamoDBScanForwardKey = attribute.Key("aws.dynamodb.scan_forward")

	// AWSDynamoDBScannedCountKey is the attribute Key conforming to the
	// "aws.dynamodb.scanned_count" semantic conventions. It represents the
	// value of the `ScannedCount` response parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 50
	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")

	// AWSDynamoDBSegmentKey is the attribute Key conforming to the
	// "aws.dynamodb.segment" semantic conventions. It represents the value of
	// the `Segment` request parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 10
	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	// AWSDynamoDBSelectKey is the attribute Key conforming to the
	// "aws.dynamodb.select" semantic conventions. It represents the value of
	// the `Select` request parameter.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'ALL_ATTRIBUTES', 'COUNT'
	AWSDynamoDBSelectKey = attribute.Key("aws.dynamodb.select")

	// AWSDynamoDBTableCountKey is the attribute Key conforming to the
	// "aws.dynamodb.table_count" semantic conventions. It represents the
	// number of items in the `TableNames` response parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 20
	AWSDynamoDBTableCountKey = attribute.Key("aws.dynamodb.table_count")

	// AWSDynamoDBTableNamesKey is the attribute Key conforming to the
	// "aws.dynamodb.table_names" semantic conventions. It represents the keys
	// in the `RequestItems` object field.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Users', 'Cats'
	AWSDynamoDBTableNamesKey = attribute.Key("aws.dynamodb.table_names")

	// AWSDynamoDBTotalSegmentsKey is the attribute Key conforming to the
	// "aws.dynamodb.total_segments" semantic conventions. It represents the
	// value of the `TotalSegments` request parameter.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 100
	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")
)

// AWSDynamoDBAttributeDefinitions returns an attribute KeyValue conforming
// to the "aws.dynamodb.attribute_definitions" semantic conventions. It
// represents the JSON-serialized value of each item in the
// `AttributeDefinitions` request field.
func AWSDynamoDBAttributeDefinitions(val ...string) attribute.KeyValue {
	return AWSDynamoDBAttributeDefinitionsKey.StringSlice(val)
}

// AWSDynamoDBAttributesToGet returns an attribute KeyValue conforming to
// the "aws.dynamodb.attributes_to_get" semantic conventions. It represents the
// value of the `AttributesToGet` request parameter.
func AWSDynamoDBAttributesToGet(val ...string) attribute.KeyValue {
	return AWSDynamoDBAttributesToGetKey.StringSlice(val)
}

// AWSDynamoDBConsistentRead returns an attribute KeyValue conforming to the
// "aws.dynamodb.consistent_read" semantic conventions. It represents the value
// of the `ConsistentRead` request parameter.
func AWSDynamoDBConsistentRead(val bool) attribute.KeyValue {
	return AWSDynamoDBConsistentReadKey.Bool(val)
}

// AWSDynamoDBConsumedCapacity returns an attribute KeyValue conforming to
// "aws.dynamodb.table_names" semantic conventions. It represents the keys in
// the `RequestItems` object field.
func AWSDynamoDBTableNames(val ...string) attribute.KeyValue {
	return AWSDynamoDBTableNamesKey.StringSlice(val)
}

// AWSDynamoDBTotalSegments returns an attribute KeyValue conforming to the
// "aws.dynamodb.total_segments" semantic conventions. It represents the value
// of the `TotalSegments` request parameter.
func AWSDynamoDBTotalSegments(val int) attribute.KeyValue {
	return AWSDynamoDBTotalSegmentsKey.Int(val)
}

// Attributes for AWS Elastic Container Service (ECS).
const (
	// AWSECSTaskIDKey is the attribute Key conforming to the "aws.ecs.task.id"
	// semantic conventions. It represents the ID of a running ECS task. The ID
	// MUST be extracted from `task.arn`.
	//
	// Type: string
	// RequirementLevel: ConditionallyRequired (If and only if `task.arn` is
	// populated.)
	// Stability: experimental
	// Examples: '66783dce-808b-17fa-466f-f53eaddfab2c',
	// '36eec6ef-d41f-26f0-6ebe-f73b3f04dbac'
	AWSECSTaskIDKey = attribute.Key("aws.ecs.task.id")

	// AWSECSClusterARNKey is the attribute Key conforming to the
	// "aws.ecs.cluster.arn" semantic conventions. It represents the ARN of an
	// [ECS
	// cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster'
	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	// AWSECSContainerARNKey is the attribute Key conforming to the
	// "aws.ecs.container.arn" semantic conventions. It represents the Amazon
	// Resource Name (ARN) of an [ECS container
	// instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:ecs:us-west-1:123456789123:container/36998648-4931-2d4e-bcab-8e39e68eb2d3'
	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	// AWSECSLaunchtypeKey is the attribute Key conforming to the
	// "aws.ecs.launchtype" semantic conventions. It represents the [launch
	// type](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// for an ECS task.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	// AWSECSTaskARNKey is the attribute Key conforming to the
	// "aws.ecs.task.arn" semantic conventions. It represents the ARN of a
	// running [ECS
	// task](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:ecs:us-west-1:123456789123:task/33818acc-723d-06af-105d-e48ecbdeef9a',
	// 'arn:aws:ecs:us-west-1:123456789123:task/my-cluster/task-id/68eef1fd-c29a-94a4-4dec-c60f8e40adea'
	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	// AWSECSTaskFamilyKey is the attribute Key conforming to the
	// "aws.ecs.task.family" semantic conventions. It represents the family
	// name of the [ECS task
	// definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html)
	// used to create the ECS task.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-family'
	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	// AWSECSTaskRevisionKey is the attribute Key conforming to the
	// "aws.ecs.task.revision" semantic conventions. It represents the revision
	// for the task definition used to create the ECS task.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '8', '26'
	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")
)

var (
	// ec2
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")
	// fargate
	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

// AWSECSTaskID returns an attribute KeyValue conforming to the
// "aws.ecs.task.id" semantic conventions. It represents the ID of a running
// ECS task. The ID MUST be extracted from `task.arn`.
func AWSECSTaskID(val string) attribute.KeyValue {
	return AWSECSTaskIDKey.String(val)
}

// AWSECSClusterARN returns an attribute KeyValue conforming to the
// "aws.ecs.cluster.arn" semantic conventions. It represents the ARN of an [ECS
// cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html).
func AWSECSClusterARN(val string) attribute.KeyValue {
	return AWSECSClusterARNKey.String(val)
}

// AWSECSContainerARN returns an attribute KeyValue conforming to the
// "aws.ecs.container.arn" semantic conventions. It represents the Amazon
// "aws.ecs.task.family" semantic conventions. It represents the family name of
// the [ECS task
// definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html)
// used to create the ECS task.
func AWSECSTaskFamily(val string) attribute.KeyValue {
	return AWSECSTaskFamilyKey.String(val)
}

// AWSECSTaskRevision returns an attribute KeyValue conforming to the
// "aws.ecs.task.revision" semantic conventions. It represents the revision for
// the task definition used to create the ECS task.
func AWSECSTaskRevision(val string) attribute.KeyValue {
	return AWSECSTaskRevisionKey.String(val)
}

// Attributes for AWS Elastic Kubernetes Service (EKS).
const (
	// AWSEKSClusterARNKey is the attribute Key conforming to the
	// "aws.eks.cluster.arn" semantic conventions. It represents the ARN of an
	// EKS cluster.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster'
	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")
)

// AWSEKSClusterARN returns an attribute KeyValue conforming to the
// "aws.eks.cluster.arn" semantic conventions. It represents the ARN of an EKS
// cluster.
func AWSEKSClusterARN(val string) attribute.KeyValue {
	return AWSEKSClusterARNKey.String(val)
}

// Attributes for AWS Logs.
const (
	// AWSLogGroupARNsKey is the attribute Key conforming to the
	// "aws.log.group.arns" semantic conventions. It represents the Amazon
	// Resource Name(s) (ARN) of the AWS log group(s).
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'arn:aws:logs:us-west-1:123456789012:log-group:/aws/my/group:*'
	// Note: See the [log group ARN format
	// documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format).
	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	// AWSLogGroupNamesKey is the attribute Key conforming to the
// AWSLambdaInvokedARN returns an attribute KeyValue conforming to the
// "aws.lambda.invoked_arn" semantic conventions. It represents the full
// invoked ARN as provided on the `Context` passed to the function
// (`Lambda-Runtime-Invoked-Function-ARN` header on the
// `/runtime/invocation/next` applicable).
func AWSLambdaInvokedARN(val string) attribute.KeyValue {
	return AWSLambdaInvokedARNKey.String(val)
}

// Attributes for AWS S3.
const (
	// AWSS3BucketKey is the attribute Key conforming to the "aws.s3.bucket"
	// semantic conventions. It represents the S3 bucket name the request
	// refers to. Corresponds to the `--bucket` parameter of the [S3
	// API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html)
	// operations.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'some-bucket-name'
	// Note: The `bucket` attribute is applicable to all S3 operations that
	// reference a bucket, i.e. that require the bucket name as a mandatory
	// parameter.
	// This applies to almost all S3 operations except `list-buckets`.
	AWSS3BucketKey = attribute.Key("aws.s3.bucket")

	// AWSS3CopySourceKey is the attribute Key conforming to the
	// "aws.s3.copy_source" semantic conventions. It represents the source
	// object (in the form `bucket`/`key`) for the copy operation.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'someFile.yml'
	// Note: The `copy_source` attribute applies to S3 copy operations and
	// corresponds to the `--copy-source` parameter
	// of the [copy-object operation within the S3
	// API](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html).
	// This applies in particular to the following operations:
	//
	// -
	// [copy-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html)
	// -
	// [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html)
	AWSS3CopySourceKey = attribute.Key("aws.s3.copy_source")

	// AWSS3DeleteKey is the attribute Key conforming to the "aws.s3.delete"
	// semantic conventions. It represents the delete request container that
	// specifies the objects to be deleted.
	//

// A cloud environment (e.g. GCP, Azure, AWS).
const (
	// CloudAccountIDKey is the attribute Key conforming to the
	// "cloud.account.id" semantic conventions. It represents the cloud account
	// ID the resource is assigned to.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '111111111111', 'opentelemetry'
	CloudAccountIDKey = attribute.Key("cloud.account.id")

	// CloudAvailabilityZoneKey is the attribute Key conforming to the
	// "cloud.availability_zone" semantic conventions. It represents the cloud
	// regions often have multiple, isolated locations known as zones to
	// increase availability. Availability zone represents the zone where the
	// resource is running.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'us-east-1c'
	// Note: Availability zones are called "zones" on Alibaba Cloud and Google
	// Cloud.
	CloudAvailabilityZoneKey = attribute.Key("cloud.availability_zone")

	// CloudPlatformKey is the attribute Key conforming to the "cloud.platform"
	// semantic conventions. It represents the cloud platform in use.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	// Note: The prefix of the service SHOULD match the one specified in
	// `cloud.provider`.
	CloudPlatformKey = attribute.Key("cloud.platform")

	// CloudProviderKey is the attribute Key conforming to the "cloud.provider"
	// semantic conventions. It represents the name of the cloud provider.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	CloudProviderKey = attribute.Key("cloud.provider")

	// CloudRegionKey is the attribute Key conforming to the "cloud.region"
	// semantic conventions. It represents the geographical region the resource
	// is running.
	//
	// Type: string
	// RequirementLevel: Optional
}

// These attributes allow to report this unit of code and therefore to provide
// more context about the span.
const (
	// CodeColumnKey is the attribute Key conforming to the "code.column"
	// semantic conventions. It represents the column number in `code.filepath`
	// best representing the operation. It SHOULD point within the code unit
	// named in `code.function`.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 16
	CodeColumnKey = attribute.Key("code.column")

	// CodeFilepathKey is the attribute Key conforming to the "code.filepath"
	// semantic conventions. It represents the source code file name that
	// identifies the code unit as uniquely as possible (preferably an absolute
	// file path).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '/usr/local/MyApplication/content_root/app/index.php'
	CodeFilepathKey = attribute.Key("code.filepath")

	// CodeFunctionKey is the attribute Key conforming to the "code.function"
	// semantic conventions. It represents the method or function name, or
	// equivalent (usually rightmost part of the code unit's name).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'serveRequest'
	CodeFunctionKey = attribute.Key("code.function")

	// CodeLineNumberKey is the attribute Key conforming to the "code.lineno"
	// semantic conventions. It represents the line number in `code.filepath`
	// best representing the operation. It SHOULD point within the code unit
	// named in `code.function`.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 42
	CodeLineNumberKey = attribute.Key("code.lineno")

	// CodeNamespaceKey is the attribute Key conforming to the "code.namespace"
	// semantic conventions. It represents the "namespace" within which
	// `code.function` is defined. Usually the qualified class or module name,
	// such that `code.namespace` + some separator + `code.function` form a
	// unique identifier for the code unit.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'com.example.MyHTTPService'
	CodeNamespaceKey = attribute.Key("code.namespace")

	// CodeStacktraceKey is the attribute Key conforming to the
	// "code.stacktrace" semantic conventions. It represents a stacktrace as a
	// string in the natural representation for the language runtime. The
	// representation is to be determined and documented by each language SIG.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'at
	// com.example.GenerateTrace.methodB(GenerateTrace.java:13)\\n at '
	//  'com.example.GenerateTrace.methodA(GenerateTrace.java:9)\\n at '
	//  'com.example.GenerateTrace.main(GenerateTrace.java:5)'

// DBOperationName returns an attribute KeyValue conforming to the
// "db.operation.name" semantic conventions. It represents the name of the
// operation or command being executed.
func DBOperationName(val string) attribute.KeyValue {
	return DBOperationNameKey.String(val)
}

// DBQueryText returns an attribute KeyValue conforming to the
// "db.query.text" semantic conventions. It represents the database query being
// executed.
func DBQueryText(val string) attribute.KeyValue {
	return DBQueryTextKey.String(val)
}

// This group defines attributes for Cassandra.
const (
	// DBCassandraConsistencyLevelKey is the attribute Key conforming to the
	// "db.cassandra.consistency_level" semantic conventions. It represents the
	// consistency level of the query. Based on consistency values from
	// [CQL](https://docs.datastax.com/en/cassandra-oss/3.0/cassandra/dml/dmlConfigConsistency.html).
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	DBCassandraConsistencyLevelKey = attribute.Key("db.cassandra.consistency_level")

	// DBCassandraCoordinatorDCKey is the attribute Key conforming to the
	// "db.cassandra.coordinator.dc" semantic conventions. It represents the
	// data center of the coordinating node for a query.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'us-west-2'
	DBCassandraCoordinatorDCKey = attribute.Key("db.cassandra.coordinator.dc")

	// DBCassandraCoordinatorIDKey is the attribute Key conforming to the
	// "db.cassandra.coordinator.id" semantic conventions. It represents the ID
	// of the coordinating node for a query.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'bc76cbc2-9174-7a83-868f-94e38fa4b8cd'
	DBCassandraCoordinatorIDKey = attribute.Key("db.cassandra.coordinator.id")

	// DBCassandraIdempotenceKey is the attribute Key conforming to the
	// "db.cassandra.idempotence" semantic conventions. It represents the
	// whether or not the query is idempotent.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	DBCassandraIdempotenceKey = attribute.Key("db.cassandra.idempotence")

	// DBCassandraPageSizeKey is the attribute Key conforming to the
	// "db.cassandra.page_size" semantic conventions. It represents the fetch
	// size used for paging, i.e. how many rows will be returned at once.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 5000
	DBCassandraPageSizeKey = attribute.Key("db.cassandra.page_size")

	// DBCassandraSpeculativeExecutionCountKey is the attribute Key conforming
	// to the "db.cassandra.speculative_execution_count" semantic conventions.
	// It represents the number of times a query was speculatively executed.
	// Not set or `0` if the query was not executed speculatively.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 0, 2
	DBCassandraSpeculativeExecutionCountKey = attribute.Key("db.cassandra.speculative_execution_count")
)

var (
	// all
	DBCassandraConsistencyLevelAll = DBCassandraConsistencyLevelKey.String("all")
	// each_quorum
	DBCassandraConsistencyLevelEachQuorum = DBCassandraConsistencyLevelKey.String("each_quorum")
	// quorum
	DBCassandraConsistencyLevelQuorum = DBCassandraConsistencyLevelKey.String("quorum")
	// local_quorum
	DBCassandraConsistencyLevelLocalQuorum = DBCassandraConsistencyLevelKey.String("local_quorum")
	// one
	DBCassandraConsistencyLevelOne = DBCassandraConsistencyLevelKey.String("one")
	// two
	DBCassandraConsistencyLevelTwo = DBCassandraConsistencyLevelKey.String("two")
	// three
	DBCassandraConsistencyLevelThree = DBCassandraConsistencyLevelKey.String("three")
	// local_one
	DBCassandraConsistencyLevelLocalOne = DBCassandraConsistencyLevelKey.String("local_one")
	// any
	DBCassandraConsistencyLevelAny = DBCassandraConsistencyLevelKey.String("any")
	// serial
	DBCassandraConsistencyLevelSerial = DBCassandraConsistencyLevelKey.String("serial")
	// local_serial
	DBCassandraConsistencyLevelLocalSerial = DBCassandraConsistencyLevelKey.String("local_serial")
// DBCassandraPageSize returns an attribute KeyValue conforming to the
// "db.cassandra.page_size" semantic conventions. It represents the fetch size
// used for paging, i.e. how many rows will be returned at once.
func DBCassandraPageSize(val int) attribute.KeyValue {
	return DBCassandraPageSizeKey.Int(val)
}

// DBCassandraSpeculativeExecutionCount returns an attribute KeyValue
// conforming to the "db.cassandra.speculative_execution_count" semantic
// conventions. It represents the number of times a query was speculatively
// executed. Not set or `0` if the query was not executed speculatively.
func DBCassandraSpeculativeExecutionCount(val int) attribute.KeyValue {
	return DBCassandraSpeculativeExecutionCountKey.Int(val)
}

// This group defines attributes for Azure Cosmos DB.
const (
	// DBCosmosDBClientIDKey is the attribute Key conforming to the
	// "db.cosmosdb.client_id" semantic conventions. It represents the unique
	// Cosmos client instance id.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '9bd9902c-6661-427f-d68e-24b14546d85e'
	DBCosmosDBClientIDKey = attribute.Key("db.cosmosdb.client_id")

	// DBCosmosDBConnectionModeKey is the attribute Key conforming to the
	// "db.cosmosdb.connection_mode" semantic conventions. It represents the
	// cosmos client connection mode.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	DBCosmosDBConnectionModeKey = attribute.Key("db.cosmosdb.connection_mode")

	// DBCosmosDBOperationTypeKey is the attribute Key conforming to the
	// "db.cosmosdb.operation_type" semantic conventions. It represents the
	// cosmosDB Operation Type.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	DBCosmosDBOperationTypeKey = attribute.Key("db.cosmosdb.operation_type")

	// DBCosmosDBRequestChargeKey is the attribute Key conforming to the
	// "db.cosmosdb.request_charge" semantic conventions. It represents the rU
	// consumed for that operation
	//
	// Type: double
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 46.18, 1.0
	DBCosmosDBRequestChargeKey = attribute.Key("db.cosmosdb.request_charge")

	// DBCosmosDBRequestContentLengthKey is the attribute Key conforming to the
	// "db.cosmosdb.request_content_length" semantic conventions. It represents
	// the request payload size in bytes
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	DBCosmosDBRequestContentLengthKey = attribute.Key("db.cosmosdb.request_content_length")

	// DBCosmosDBStatusCodeKey is the attribute Key conforming to the
	// "db.cosmosdb.status_code" semantic conventions. It represents the cosmos
	// DB status code.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 200, 201
	DBCosmosDBStatusCodeKey = attribute.Key("db.cosmosdb.status_code")

	// DBCosmosDBSubStatusCodeKey is the attribute Key conforming to the
	// "db.cosmosdb.sub_status_code" semantic conventions. It represents the
	// cosmos DB sub status code.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1000, 1002
	DBCosmosDBSubStatusCodeKey = attribute.Key("db.cosmosdb.sub_status_code")
)

var (
	// Gateway (HTTP) connections mode
	DBCosmosDBConnectionModeGateway = DBCosmosDBConnectionModeKey.String("gateway")
	// Direct connection
	DBCosmosDBConnectionModeDirect = DBCosmosDBConnectionModeKey.String("direct")
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
	// outlined below. This value is not an advertising identifier and MUST NOT
	// be used as such. On iOS (Swift or Objective-C), this value MUST be equal
	// to the [vendor
	// identifier](https://developer.apple.com/documentation/uikit/uidevice/1620059-identifierforvendor).
	// On Android (Java or Kotlin), this value MUST be equal to the Firebase
	// Installation ID or a globally unique UUID which is persisted across
	// sessions in your application. More information can be found
	// [here](https://developer.android.com/training/articles/user-data-ids) on
	// best practices and exact implementation details. Caution should be taken
	// when storing personal data or anything which can identify a user. GDPR
	// and data protection laws may apply, ensure you do your own due
	// diligence.
	DeviceIDKey = attribute.Key("device.id")

	// DeviceManufacturerKey is the attribute Key conforming to the
	// "device.manufacturer" semantic conventions. It represents the name of
	// the device manufacturer
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Apple', 'Samsung'
	// Note: The Android OS provides this field via
	// [Build](https://developer.android.com/reference/android/os/Build#MANUFACTURER).
	// iOS apps SHOULD hardcode the value `Apple`.
	DeviceManufacturerKey = attribute.Key("device.manufacturer")

	// DeviceModelIdentifierKey is the attribute Key conforming to the
	// "device.model.identifier" semantic conventions. It represents the model
	// identifier for the device
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'iPhone3,4', 'SM-G920F'
	// Note: It's recommended this value represents a machine-readable version
	// of the model identifier rather than the market or consumer-friendly name
	// of the device.
	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	// DeviceModelNameKey is the attribute Key conforming to the
	// "device.model.name" semantic conventions. It represents the marketing
	// name for the device model
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'iPhone 6s Plus', 'Samsung Galaxy S6'
	// Note: It's recommended this value represents a human-readable version of
	// the device model rather than a machine-readable alternative.
	DeviceModelNameKey = attribute.Key("device.model.name")
	// integers (\DDD). Back slashes and quotes should be escaped. Tabs,
	// carriage returns, and line feeds should be converted to \t, \r, and \n
	// respectively.
	DNSQuestionNameKey = attribute.Key("dns.question.name")
)

// DNSQuestionName returns an attribute KeyValue conforming to the
// "dns.question.name" semantic conventions. It represents the name being
// queried.
func DNSQuestionName(val string) attribute.KeyValue {
	return DNSQuestionNameKey.String(val)
}

// Attributes for operations with an authenticated and/or authorized enduser.
const (
	// EnduserIDKey is the attribute Key conforming to the "enduser.id"
	// semantic conventions. It represents the username or client_id extracted
	// from the access token or
	// [Authorization](https://tools.ietf.org/html/rfc7235#section-4.2) header
	// in the inbound request from outside the system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'username'
	EnduserIDKey = attribute.Key("enduser.id")

	// EnduserRoleKey is the attribute Key conforming to the "enduser.role"
	// semantic conventions. It represents the actual/assumed role the client
	// is making the request under extracted from token or application security
	// context.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
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
	// Stability: experimental
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
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2020-01-23T13:47:06Z'
	FaaSDocumentTimeKey = attribute.Key("faas.document.time")

	// FaaSInstanceKey is the attribute Key conforming to the "faas.instance"
	// semantic conventions. It represents the execution environment ID as a
	// string, that will be potentially reused for other invocations to the
	// same function/function version.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2021/06/28/[$LATEST]2f399eb14537447da05ab2a2e39309de'
	// Note: * **AWS Lambda:** Use the (full) log stream name.
	FaaSInstanceKey = attribute.Key("faas.instance")

	// FaaSInvocationIDKey is the attribute Key conforming to the
	// "faas.invocation_id" semantic conventions. It represents the invocation
	// ID of the current function invocation.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'dc8d9de8-b219-9c9e-e84c-468a98c3cd95'
	FaaSInvocationIDKey = attribute.Key("faas.invocation_id")

	// FaaSInvokedNameKey is the attribute Key conforming to the
	// "faas.invoked_name" semantic conventions. It represents the name of the
	// invoked function.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'my-function'
	// Note: SHOULD be equal to the `faas.name` resource attribute of the
	// invoked function.
	FaaSInvokedNameKey = attribute.Key("faas.invoked_name")

	// FaaSInvokedProviderKey is the attribute Key conforming to the
	// "faas.invoked_provider" semantic conventions. It represents the cloud
	// provider of the invoked function.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	// Note: SHOULD be equal to the `cloud.provider` resource attribute of the
	// invoked function.
	FaaSInvokedProviderKey = attribute.Key("faas.invoked_provider")

	// FaaSInvokedRegionKey is the attribute Key conforming to the
	//
	// For some cloud providers, the above definition is ambiguous. The
	// following
	// definition of function name MUST be used for this attribute
	// (and consequently the span name) for the listed cloud
	// providers/products:
	//
	// * **Azure:**  The full name `<FUNCAPP>/<FUNC>`, i.e., function app name
	//   followed by a forward slash followed by the function name (this form
	//   can also be seen in the resource JSON for the function).
	//   This means that a span attribute MUST be used, as an Azure function
	//   app can host multiple functions that would usually share
	//   a TracerProvider (see also the `cloud.resource_id` attribute).
	FaaSNameKey = attribute.Key("faas.name")

	// FaaSTimeKey is the attribute Key conforming to the "faas.time" semantic
	// conventions. It represents a string containing the function invocation
	// time in the [ISO
	// 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format
	// expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2020-01-23T13:47:06Z'
	FaaSTimeKey = attribute.Key("faas.time")

	// FaaSTriggerKey is the attribute Key conforming to the "faas.trigger"
	// semantic conventions. It represents the type of the trigger which caused
	// this function invocation.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	FaaSTriggerKey = attribute.Key("faas.trigger")

	// FaaSVersionKey is the attribute Key conforming to the "faas.version"
	// semantic conventions. It represents the immutable version of the
	// function being executed.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '26', 'pinkfroid-00002'
	// Note: Depending on the cloud provider and platform, use:
	//
	// * **AWS Lambda:** The [function
	// version](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html)
	//   (an integer represented as a decimal string).
	// * **Google Cloud Run (Services):** The
	// [revision](https://cloud.google.com/run/docs/managing/revisions)
// file name. It should include the drive letter, when appropriate.
func FilePath(val string) attribute.KeyValue {
	return FilePathKey.String(val)
}

// FileSize returns an attribute KeyValue conforming to the "file.size"
// semantic conventions. It represents the file size in bytes.
func FileSize(val int) attribute.KeyValue {
	return FileSizeKey.Int(val)
}

// Attributes for Google Cloud Run.
const (
	// GCPCloudRunJobExecutionKey is the attribute Key conforming to the
	// "gcp.cloud_run.job.execution" semantic conventions. It represents the
	// name of the Cloud Run
	// [execution](https://cloud.google.com/run/docs/managing/job-executions)
	// being run for the Job, as set by the
	// [`CLOUD_RUN_EXECUTION`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
	// environment variable.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'job-name-xxxx', 'sample-job-mdw84'
	GCPCloudRunJobExecutionKey = attribute.Key("gcp.cloud_run.job.execution")

	// GCPCloudRunJobTaskIndexKey is the attribute Key conforming to the
	// "gcp.cloud_run.job.task_index" semantic conventions. It represents the
	// index for a task within an execution as provided by the
	// [`CLOUD_RUN_TASK_INDEX`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
	// environment variable.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 0, 1
	GCPCloudRunJobTaskIndexKey = attribute.Key("gcp.cloud_run.job.task_index")
)

// GCPCloudRunJobExecution returns an attribute KeyValue conforming to the
// "gcp.cloud_run.job.execution" semantic conventions. It represents the name
// of the Cloud Run
// [execution](https://cloud.google.com/run/docs/managing/job-executions) being
// run for the Job, as set by the
// [`CLOUD_RUN_EXECUTION`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars)
// environment variable.
func GCPCloudRunJobExecution(val string) attribute.KeyValue {
	return GCPCloudRunJobExecutionKey.String(val)
}

	// GenAiCompletionKey is the attribute Key conforming to the
	// "gen_ai.completion" semantic conventions. It represents the full
	// response received from the LLM.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: "[{'role': 'assistant', 'content': 'The capital of France is
	// Paris.'}]"
	// Note: It's RECOMMENDED to format completions as JSON string matching
	// [OpenAI messages
	// format](https://platform.openai.com/docs/guides/text-generation)
	GenAiCompletionKey = attribute.Key("gen_ai.completion")

	// GenAiPromptKey is the attribute Key conforming to the "gen_ai.prompt"
	// semantic conventions. It represents the full prompt sent to an LLM.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: "[{'role': 'user', 'content': 'What is the capital of
	// France?'}]"
	// Note: It's RECOMMENDED to format prompts as JSON string matching [OpenAI
	// messages
	// format](https://platform.openai.com/docs/guides/text-generation)
	GenAiPromptKey = attribute.Key("gen_ai.prompt")

	// GenAiRequestMaxTokensKey is the attribute Key conforming to the
	// "gen_ai.request.max_tokens" semantic conventions. It represents the
	// maximum number of tokens the LLM generates for a request.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 100
	GenAiRequestMaxTokensKey = attribute.Key("gen_ai.request.max_tokens")

	// GenAiRequestModelKey is the attribute Key conforming to the
	// "gen_ai.request.model" semantic conventions. It represents the name of
	// the LLM a request is being made to.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'gpt-4'
	GenAiRequestModelKey = attribute.Key("gen_ai.request.model")

	// GenAiRequestTemperatureKey is the attribute Key conforming to the
	// "gen_ai.request.temperature" semantic conventions. It represents the
	// temperature setting for the LLM request.
	//
	// Examples: 'openai'
	// Note: The actual GenAI product may differ from the one identified by the
	// client. For example, when using OpenAI client libraries to communicate
	// with Mistral, the `gen_ai.system` is set to `openai` based on the
	// instrumentation's best knowledge.
	GenAiSystemKey = attribute.Key("gen_ai.system")

	// GenAiUsageCompletionTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.completion_tokens" semantic conventions. It represents the
	// number of tokens used in the LLM response (completion).
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 180
	GenAiUsageCompletionTokensKey = attribute.Key("gen_ai.usage.completion_tokens")

	// GenAiUsagePromptTokensKey is the attribute Key conforming to the
	// "gen_ai.usage.prompt_tokens" semantic conventions. It represents the
	// number of tokens used in the LLM prompt.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 100
	GenAiUsagePromptTokensKey = attribute.Key("gen_ai.usage.prompt_tokens")
)

var (
	// OpenAI
	GenAiSystemOpenai = GenAiSystemKey.String("openai")
)

// GenAiCompletion returns an attribute KeyValue conforming to the
// "gen_ai.completion" semantic conventions. It represents the full response
// received from the LLM.
func GenAiCompletion(val string) attribute.KeyValue {
	return GenAiCompletionKey.String(val)
}

// GenAiPrompt returns an attribute KeyValue conforming to the
// "gen_ai.prompt" semantic conventions. It represents the full prompt sent to
// an LLM.
func GenAiPrompt(val string) attribute.KeyValue {
	return GenAiPromptKey.String(val)
}

// GenAiRequestMaxTokens returns an attribute KeyValue conforming to the
// "gen_ai.request.max_tokens" semantic conventions. It represents the maximum
// number of tokens the LLM generates for a request.
func GenAiRequestMaxTokens(val int) attribute.KeyValue {
// GraphqlDocument returns an attribute KeyValue conforming to the
// "graphql.document" semantic conventions. It represents the GraphQL document
// being executed.
func GraphqlDocument(val string) attribute.KeyValue {
	return GraphqlDocumentKey.String(val)
}

// GraphqlOperationName returns an attribute KeyValue conforming to the
// "graphql.operation.name" semantic conventions. It represents the name of the
// operation being executed.
func GraphqlOperationName(val string) attribute.KeyValue {
	return GraphqlOperationNameKey.String(val)
}

// Attributes for the Android platform on which the Android application is
// running.
const (
	// HerokuAppIDKey is the attribute Key conforming to the "heroku.app.id"
	// semantic conventions. It represents the unique identifier for the
	// application
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '5abf4024-a21d-1317-0637-fc7f978dc6ca'
	HerokuAppIDKey = attribute.Key("heroku.app.id")

	// HerokuReleaseCommitKey is the attribute Key conforming to the
	// "heroku.release.commit" semantic conventions. It represents the commit
	// hash for the current release
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'e6134959463efd8966b20e75b913cafe3f5ec'
	HerokuReleaseCommitKey = attribute.Key("heroku.release.commit")

	// HerokuReleaseCreationTimestampKey is the attribute Key conforming to the
	// "heroku.release.creation_timestamp" semantic conventions. It represents
	// the time and date the release was created
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '2022-10-23T18:00:42Z'
	HerokuReleaseCreationTimestampKey = attribute.Key("heroku.release.creation_timestamp")
)

// HerokuAppID returns an attribute KeyValue conforming to the
// "heroku.app.id" semantic conventions. It represents the unique identifier
// for the application

// HerokuReleaseCommit returns an attribute KeyValue conforming to the
// "heroku.release.commit" semantic conventions. It represents the commit hash
// for the current release
func HerokuReleaseCommit(val string) attribute.KeyValue {
	return HerokuReleaseCommitKey.String(val)
}

// HerokuReleaseCreationTimestamp returns an attribute KeyValue conforming
// to the "heroku.release.creation_timestamp" semantic conventions. It
// represents the time and date the release was created
func HerokuReleaseCreationTimestamp(val string) attribute.KeyValue {
	return HerokuReleaseCreationTimestampKey.String(val)
}

// A host is defined as a computing instance. For example, physical servers,
// virtual machines, switches or disk array.
const (
	// HostArchKey is the attribute Key conforming to the "host.arch" semantic
	// conventions. It represents the CPU architecture the host system is
	// running on.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	HostArchKey = attribute.Key("host.arch")

	// HostCPUCacheL2SizeKey is the attribute Key conforming to the
	// "host.cpu.cache.l2.size" semantic conventions. It represents the amount
	// of level 2 memory cache available to the processor (in Bytes).
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 12288000
	HostCPUCacheL2SizeKey = attribute.Key("host.cpu.cache.l2.size")

	// HostCPUFamilyKey is the attribute Key conforming to the
	// "host.cpu.family" semantic conventions. It represents the family or
	// generation of the CPU.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '6', 'PA-RISC 1.1e'
	HostCPUFamilyKey = attribute.Key("host.cpu.family")

	// HostCPUModelIDKey is the attribute Key conforming to the
	// "host.cpu.model.id" semantic conventions. It represents the model
	// identifier. It provides more granular information about the CPU,
	// distinguishing it from other CPUs within the same family.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '6', '9000/778/B180L'
	HostCPUModelIDKey = attribute.Key("host.cpu.model.id")

	// HostCPUModelNameKey is the attribute Key conforming to the
	// "host.cpu.model.name" semantic conventions. It represents the model
	// designation of the processor.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz'
	HostCPUModelNameKey = attribute.Key("host.cpu.model.name")

	// HostCPUSteppingKey is the attribute Key conforming to the
	// "host.cpu.stepping" semantic conventions. It represents the stepping or
	// core revisions.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1', 'r1p1'
	HostCPUSteppingKey = attribute.Key("host.cpu.stepping")

	// HostCPUVendorIDKey is the attribute Key conforming to the
	// "host.cpu.vendor.id" semantic conventions. It represents the processor
	// manufacturer identifier. A maximum 12-character string.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'GenuineIntel'
	// Note: [CPUID](https://wiki.osdev.org/CPUID) command returns the vendor
	// ID string in EBX, EDX and ECX registers. Writing these to memory in this
	// order results in a 12-character string.
	HostCPUVendorIDKey = attribute.Key("host.cpu.vendor.id")

	// HostIDKey is the attribute Key conforming to the "host.id" semantic
	// conventions. It represents the unique host ID. For Cloud, this must be
	// the instance_id assigned by the cloud provider. For non-containerized
	// systems, this should be the `machine-id`. See the table below for the
	// sources to use to determine the `machine-id` based on operating system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'fdbf79e8af94cb7f9e8df36789187052'
	HostIDKey = attribute.Key("host.id")

	// HostImageIDKey is the attribute Key conforming to the "host.image.id"
	// semantic conventions. It represents the vM image ID or host OS image ID.
	// For Cloud, this value is from the provider.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'ami-07b06b442921831e5'
	HostImageIDKey = attribute.Key("host.image.id")

	// HostImageNameKey is the attribute Key conforming to the
	// "host.image.name" semantic conventions. It represents the name of the VM
	// image or OS install the host was instantiated from.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'infra-ami-eks-worker-node-7d4ec78312', 'CentOS-8-x86_64-1905'
	HostImageNameKey = attribute.Key("host.image.name")

	// HostImageVersionKey is the attribute Key conforming to the
	// "host.image.version" semantic conventions. It represents the version
	// string of the VM image or host OS as defined in [Version
	// Attributes](/docs/resource/README.md#version-attributes).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '0.1'
	HostImageVersionKey = attribute.Key("host.image.version")

	// HostIPKey is the attribute Key conforming to the "host.ip" semantic
	// conventions. It represents the available IP addresses of the host,
	// excluding loopback interfaces.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '192.168.1.140', 'fe80::abc2:4a28:737a:609e'
	// Note: IPv4 Addresses MUST be specified in dotted-quad notation. IPv6
	// addresses MUST be specified in the [RFC
	// 5952](https://www.rfc-editor.org/rfc/rfc5952.html) format.
	HostIPKey = attribute.Key("host.ip")

	// HostMacKey is the attribute Key conforming to the "host.mac" semantic
	// conventions. It represents the available MAC addresses of the host,
	// excluding loopback interfaces.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'AC-DE-48-23-45-67', 'AC-DE-48-23-45-67-01-9F'
	// Note: MAC Addresses MUST be represented in [IEEE RA hexadecimal
	// form](https://standards.ieee.org/wp-content/uploads/import/documents/tutorials/eui.pdf):
	// as hyphen-separated octets in uppercase hexadecimal form from most to
	// least significant.
	HostMacKey = attribute.Key("host.mac")

	// HostNameKey is the attribute Key conforming to the "host.name" semantic
	// conventions. It represents the name of the host. On Unix systems, it may
	// contain what the hostname command returns, or the fully qualified
	// hostname, or another name specified by the user.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-test'
	HostNameKey = attribute.Key("host.name")

	// HostTypeKey is the attribute Key conforming to the "host.type" semantic
	// conventions. It represents the type of host. For Cloud, this must be the
	// machine type.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'n1-standard-1'
	HostTypeKey = attribute.Key("host.type")
)

var (
	// AMD64

// JvmMemoryPoolName returns an attribute KeyValue conforming to the
// "jvm.memory.pool.name" semantic conventions. It represents the name of the
// memory pool.
func JvmMemoryPoolName(val string) attribute.KeyValue {
	return JvmMemoryPoolNameKey.String(val)
}

// JvmThreadDaemon returns an attribute KeyValue conforming to the
// "jvm.thread.daemon" semantic conventions. It represents the whether the
// thread is daemon or not.
func JvmThreadDaemon(val bool) attribute.KeyValue {
	return JvmThreadDaemonKey.Bool(val)
}

// Kubernetes resource attributes.
const (
	// K8SClusterNameKey is the attribute Key conforming to the
	// "k8s.cluster.name" semantic conventions. It represents the name of the
	// cluster.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-cluster'
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")

	// K8SClusterUIDKey is the attribute Key conforming to the
	// "k8s.cluster.uid" semantic conventions. It represents a pseudo-ID for
	// the cluster, set to the UID of the `kube-system` namespace.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '320bf2b6-d4b3-2d60-ea99-54938e8bf12e'
	// Note: K8S doesn't have support for obtaining a cluster ID. If this is
	// ever
	// added, we will recommend collecting the `k8s.cluster.uid` through the
	// official APIs. In the meantime, we are able to use the `uid` of the
	// `kube-system` namespace as a proxy for cluster ID. Read on for the
	// rationale.
	//
	// Every object created in a K8S cluster is assigned a distinct UID. The
	// `kube-system` namespace is used by Kubernetes itself and will exist
	// for the lifetime of the cluster. Using the `uid` of the `kube-system`
	// namespace is a reasonable proxy for the K8S ClusterID as it will only
	// change if the cluster is rebuilt. Furthermore, Kubernetes UIDs are
	// UUIDs as standardized by
	// [ISO/IEC 9834-8 and ITU-T
	// X.667](https://www.itu.int/ITU-T/studygroups/com17/oid.html).
	// Which states:
	//
	// > If generated according to one of the mechanisms defined in Rec.
	//   ITU-T X.667 | ISO/IEC 9834-8, a UUID is either guaranteed to be
	//   different from all other UUIDs generated before 3603 A.D., or is
	//   extremely likely to be different (depending on the mechanism chosen).
	//
	// Therefore, UIDs between clusters should be extremely unlikely to
	// conflict.
	K8SClusterUIDKey = attribute.Key("k8s.cluster.uid")

	// K8SContainerNameKey is the attribute Key conforming to the
	// "k8s.container.name" semantic conventions. It represents the name of the
	// Container from Pod specification, must be unique within a Pod. Container
	// runtime usually uses different globally unique name (`container.name`).
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'redis'
	K8SContainerNameKey = attribute.Key("k8s.container.name")

	// K8SContainerRestartCountKey is the attribute Key conforming to the
	// "k8s.container.restart_count" semantic conventions. It represents the
	// number of times the container was restarted. This attribute can be used
	// to identify a particular container (running or stopped) within a
	// container spec.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	K8SContainerRestartCountKey = attribute.Key("k8s.container.restart_count")

	// K8SContainerStatusLastTerminatedReasonKey is the attribute Key
	// conforming to the "k8s.container.status.last_terminated_reason" semantic
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Evicted', 'Error'
	K8SContainerStatusLastTerminatedReasonKey = attribute.Key("k8s.container.status.last_terminated_reason")

	// K8SCronJobNameKey is the attribute Key conforming to the
	// "k8s.cronjob.name" semantic conventions. It represents the name of the
	// CronJob.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")

	// K8SCronJobUIDKey is the attribute Key conforming to the
	// "k8s.cronjob.uid" semantic conventions. It represents the UID of the
	// CronJob.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '868dbf10-3fe1-9a5b-8e67-d4ce856b1acb'
	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")

	// K8SDaemonSetNameKey is the attribute Key conforming to the
	// "k8s.daemonset.name" semantic conventions. It represents the name of the
	// DaemonSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")

	// K8SDaemonSetUIDKey is the attribute Key conforming to the
	// "k8s.daemonset.uid" semantic conventions. It represents the UID of the
	// DaemonSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '628bbd33-2cf0-6c1c-8e63-a4fd301d9efc'
	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")

	// K8SDeploymentNameKey is the attribute Key conforming to the
	// "k8s.deployment.name" semantic conventions. It represents the name of
	// the Deployment.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")

	// K8SDeploymentUIDKey is the attribute Key conforming to the
	// "k8s.deployment.uid" semantic conventions. It represents the UID of the
	// Deployment.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '271fbf61-9dd8-4a7e-3c68-c3ae833c2cef'
	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")

	// K8SJobNameKey is the attribute Key conforming to the "k8s.job.name"
	// semantic conventions. It represents the name of the Job.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SJobNameKey = attribute.Key("k8s.job.name")

	// K8SJobUIDKey is the attribute Key conforming to the "k8s.job.uid"
	// semantic conventions. It represents the UID of the Job.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '601dbd18-4bd5-3f6d-3d07-d5ce833a8cdf'
	K8SJobUIDKey = attribute.Key("k8s.job.uid")

	// K8SNamespaceNameKey is the attribute Key conforming to the
	// "k8s.namespace.name" semantic conventions. It represents the name of the
	// namespace that the pod is running in.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'default'
	K8SNamespaceNameKey = attribute.Key("k8s.namespace.name")

	// K8SNodeNameKey is the attribute Key conforming to the "k8s.node.name"
	// semantic conventions. It represents the name of the Node.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'node-1'
	K8SNodeNameKey = attribute.Key("k8s.node.name")

	// K8SNodeUIDKey is the attribute Key conforming to the "k8s.node.uid"
	// semantic conventions. It represents the UID of the Node.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '8cc3b7c7-3273-4991-e6dd-8fe2ea53c9b7'
	K8SNodeUIDKey = attribute.Key("k8s.node.uid")

	// K8SPodNameKey is the attribute Key conforming to the "k8s.pod.name"
	// semantic conventions. It represents the name of the Pod.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry-pod-autoconf'
	K8SPodNameKey = attribute.Key("k8s.pod.name")

	// K8SPodUIDKey is the attribute Key conforming to the "k8s.pod.uid"
	// semantic conventions. It represents the UID of the Pod.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '810aec81-1dc3-8d2e-3f76-a6ef649e9daf'
	K8SPodUIDKey = attribute.Key("k8s.pod.uid")

	// K8SReplicaSetNameKey is the attribute Key conforming to the
	// "k8s.replicaset.name" semantic conventions. It represents the name of
	// the ReplicaSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")

	// K8SReplicaSetUIDKey is the attribute Key conforming to the
	// "k8s.replicaset.uid" semantic conventions. It represents the UID of the
	// ReplicaSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '513aeb95-1af8-8f3f-5d43-c1db360e0ebd'
	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")

	// K8SStatefulSetNameKey is the attribute Key conforming to the
	// "k8s.statefulset.name" semantic conventions. It represents the name of
	// the StatefulSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'opentelemetry'
	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")

	// K8SStatefulSetUIDKey is the attribute Key conforming to the
	// "k8s.statefulset.uid" semantic conventions. It represents the UID of the
	// StatefulSet.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '458cac09-8dc1-6a9c-9b15-b0bb570b5dde'
	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")
)

// K8SClusterName returns an attribute KeyValue conforming to the
// "k8s.cluster.name" semantic conventions. It represents the name of the
// cluster.
func K8SClusterName(val string) attribute.KeyValue {
	return K8SClusterNameKey.String(val)
}

// K8SClusterUID returns an attribute KeyValue conforming to the
// "k8s.cluster.uid" semantic conventions. It represents a pseudo-ID for the
// cluster, set to the UID of the `kube-system` namespace.
func K8SClusterUID(val string) attribute.KeyValue {
	return K8SClusterUIDKey.String(val)
}

// K8SContainerName returns an attribute KeyValue conforming to the
// "k8s.container.name" semantic conventions. It represents the name of the
// Container from Pod specification, must be unique within a Pod. Container
// runtime usually uses different globally unique name (`container.name`).
func K8SContainerName(val string) attribute.KeyValue {
	return K8SContainerNameKey.String(val)
}

// K8SContainerRestartCount returns an attribute KeyValue conforming to the
	// semantic conventions. It represents the basename of the file.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'audit.log'
	LogFileNameKey = attribute.Key("log.file.name")

	// LogFileNameResolvedKey is the attribute Key conforming to the
	// "log.file.name_resolved" semantic conventions. It represents the
	// basename of the file, with symlinks resolved.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'uuid.log'
	LogFileNameResolvedKey = attribute.Key("log.file.name_resolved")

	// LogFilePathKey is the attribute Key conforming to the "log.file.path"
	// semantic conventions. It represents the full path to the file.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '/var/log/mysql/audit.log'
	LogFilePathKey = attribute.Key("log.file.path")

	// LogFilePathResolvedKey is the attribute Key conforming to the
	// "log.file.path_resolved" semantic conventions. It represents the full
	// path to the file, with symlinks resolved.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '/var/lib/docker/uuid.log'
	LogFilePathResolvedKey = attribute.Key("log.file.path_resolved")
)

// LogFileName returns an attribute KeyValue conforming to the
// "log.file.name" semantic conventions. It represents the basename of the
// file.
func LogFileName(val string) attribute.KeyValue {
	return LogFileNameKey.String(val)
}

// LogFileNameResolved returns an attribute KeyValue conforming to the
// "log.file.name_resolved" semantic conventions. It represents the basename of
// the file, with symlinks resolved.
func LogFileNameResolved(val string) attribute.KeyValue {
	return LogFileNameResolvedKey.String(val)
}
const (
	// MessagingBatchMessageCountKey is the attribute Key conforming to the
	// "messaging.batch.message_count" semantic conventions. It represents the
	// number of messages sent, received, or processed in the scope of the
	// batching operation.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 0, 1, 2
	// Note: Instrumentations SHOULD NOT set `messaging.batch.message_count` on
	// spans that operate with a single message. When a messaging client
	// library supports both batch and single-message API for the same
	// operation, instrumentations SHOULD use `messaging.batch.message_count`
	// for batching APIs and SHOULD NOT use it for single-message APIs.
	MessagingBatchMessageCountKey = attribute.Key("messaging.batch.message_count")

	// MessagingClientIDKey is the attribute Key conforming to the
	// "messaging.client.id" semantic conventions. It represents a unique
	// identifier for the client that consumes or produces a message.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'client-5', 'myhost@8742@s8083jm'
	MessagingClientIDKey = attribute.Key("messaging.client.id")

	// MessagingDestinationAnonymousKey is the attribute Key conforming to the
	// "messaging.destination.anonymous" semantic conventions. It represents a
	// boolean that is true if the message destination is anonymous (could be
	// unnamed or have auto-generated name).
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	MessagingDestinationAnonymousKey = attribute.Key("messaging.destination.anonymous")

	// MessagingDestinationNameKey is the attribute Key conforming to the
	// "messaging.destination.name" semantic conventions. It represents the
	// message destination name
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'MyQueue', 'MyTopic'
	// Note: Destination name SHOULD uniquely identify a specific queue, topic
	// or other entity within the broker. If
	// the broker doesn't have such notion, the destination name SHOULD
	// uniquely identify the broker.
	MessagingDestinationNameKey = attribute.Key("messaging.destination.name")

	// MessagingDestinationPartitionIDKey is the attribute Key conforming to
	// the "messaging.destination.partition.id" semantic conventions. It
	// represents the identifier of the partition messages are sent to or
	// received from, unique within the `messaging.destination.name`.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1'
	MessagingDestinationPartitionIDKey = attribute.Key("messaging.destination.partition.id")

	// MessagingDestinationTemplateKey is the attribute Key conforming to the
	// "messaging.destination.template" semantic conventions. It represents the
	// low cardinality representation of the messaging destination name
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '/customers/{customerID}'
	// Note: Destination names could be constructed from templates. An example
	// would be a destination name involving a user name or product id.
	// Although the destination name in this case is of high cardinality, the
	// underlying template is of low cardinality and can be effectively used
	// for grouping and aggregation.
	MessagingDestinationTemplateKey = attribute.Key("messaging.destination.template")

	// MessagingDestinationTemporaryKey is the attribute Key conforming to the
	// "messaging.destination.temporary" semantic conventions. It represents a
	// boolean that is true if the message destination is temporary and might
	// not exist anymore after messages are processed.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	MessagingDestinationTemporaryKey = attribute.Key("messaging.destination.temporary")
	// SHOULD uniquely identify the broker.
	MessagingDestinationPublishNameKey = attribute.Key("messaging.destination_publish.name")

	// MessagingMessageBodySizeKey is the attribute Key conforming to the
	// "messaging.message.body.size" semantic conventions. It represents the
	// size of the message body in bytes.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1439
	// Note: This can refer to both the compressed or uncompressed body size.
	// If both sizes are known, the uncompressed
	// body size should be used.
	MessagingMessageBodySizeKey = attribute.Key("messaging.message.body.size")

	// MessagingMessageConversationIDKey is the attribute Key conforming to the
	// "messaging.message.conversation_id" semantic conventions. It represents
	// the conversation ID identifying the conversation to which the message
	// belongs, represented as a string. Sometimes called "Correlation ID".
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'MyConversationID'
	MessagingMessageConversationIDKey = attribute.Key("messaging.message.conversation_id")

	// MessagingMessageEnvelopeSizeKey is the attribute Key conforming to the
	// "messaging.message.envelope.size" semantic conventions. It represents
	// the size of the message body and metadata in bytes.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 2738
	// Note: This can refer to both the compressed or uncompressed size. If
	// both sizes are known, the uncompressed
	// size should be used.
	MessagingMessageEnvelopeSizeKey = attribute.Key("messaging.message.envelope.size")

	// MessagingMessageIDKey is the attribute Key conforming to the
	// "messaging.message.id" semantic conventions. It represents a value used
	// by the messaging system as an identifier for the message, represented as
	// a string.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '452a7c7c7c7048c2f887f61572b18fc2'
	MessagingMessageIDKey = attribute.Key("messaging.message.id")

	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'myKey'
	// Note: If the key type is not string, it's string representation has to
	// be supplied for the attribute. If the key has no unambiguous, canonical
	// string form, don't include its value.
	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	// MessagingKafkaMessageOffsetKey is the attribute Key conforming to the
	// "messaging.kafka.message.offset" semantic conventions. It represents the
	// offset of a record in the corresponding Kafka partition.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 42
	MessagingKafkaMessageOffsetKey = attribute.Key("messaging.kafka.message.offset")

	// MessagingKafkaMessageTombstoneKey is the attribute Key conforming to the
	// "messaging.kafka.message.tombstone" semantic conventions. It represents
	// a boolean that is true if the message is a tombstone.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	MessagingKafkaMessageTombstoneKey = attribute.Key("messaging.kafka.message.tombstone")
)

// MessagingKafkaConsumerGroup returns an attribute KeyValue conforming to
// the "messaging.kafka.consumer.group" semantic conventions. It represents the
// name of the Kafka Consumer Group that is handling the message. Only applies
// to consumers, not producers.
func MessagingKafkaConsumerGroup(val string) attribute.KeyValue {
	return MessagingKafkaConsumerGroupKey.String(val)
}

// MessagingKafkaMessageKey returns an attribute KeyValue conforming to the
// "messaging.kafka.message.key" semantic conventions. It represents the
// message keys in Kafka are used for grouping alike messages to ensure they're
// processed on the same partition. They differ from `messaging.message.id` in
// that they're not unique. If the key is `null`, the attribute MUST NOT be
// set.
func MessagingKafkaMessageKey(val string) attribute.KeyValue {
	return MessagingKafkaMessageKeyKey.String(val)
}

// MessagingKafkaMessageOffset returns an attribute KeyValue conforming to
// the "messaging.kafka.message.offset" semantic conventions. It represents the
// offset of a record in the corresponding Kafka partition.
func MessagingKafkaMessageOffset(val int) attribute.KeyValue {
	return MessagingKafkaMessageOffsetKey.Int(val)
func MessagingRabbitmqMessageDeliveryTag(val int) attribute.KeyValue {
	return MessagingRabbitmqMessageDeliveryTagKey.Int(val)
}

// This group describes attributes specific to RocketMQ.
const (
	// MessagingRocketmqClientGroupKey is the attribute Key conforming to the
	// "messaging.rocketmq.client_group" semantic conventions. It represents
	// the name of the RocketMQ producer/consumer group that is handling the
	// message. The client type is identified by the SpanKind.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'myConsumerGroup'
	MessagingRocketmqClientGroupKey = attribute.Key("messaging.rocketmq.client_group")

	// MessagingRocketmqConsumptionModelKey is the attribute Key conforming to
	// the "messaging.rocketmq.consumption_model" semantic conventions. It
	// represents the model of message consumption. This only applies to
	// consumer spans.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	MessagingRocketmqConsumptionModelKey = attribute.Key("messaging.rocketmq.consumption_model")

	// MessagingRocketmqMessageDelayTimeLevelKey is the attribute Key
	// conforming to the "messaging.rocketmq.message.delay_time_level" semantic
	// conventions. It represents the delay time level for delay message, which
	// determines the message delay time.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 3
	MessagingRocketmqMessageDelayTimeLevelKey = attribute.Key("messaging.rocketmq.message.delay_time_level")

	// MessagingRocketmqMessageDeliveryTimestampKey is the attribute Key
	// conforming to the "messaging.rocketmq.message.delivery_timestamp"
	// semantic conventions. It represents the timestamp in milliseconds that
	// the delay message is expected to be delivered to consumer.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 1665987217045
	MessagingRocketmqMessageDeliveryTimestampKey = attribute.Key("messaging.rocketmq.message.delivery_timestamp")

	// MessagingRocketmqMessageGroupKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.group" semantic conventions. It represents
	// the it is essential for FIFO message. Messages that belong to the same
	// message group are always processed one by one within the same consumer
	// group.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'myMessageGroup'
	MessagingRocketmqMessageGroupKey = attribute.Key("messaging.rocketmq.message.group")

	// MessagingRocketmqMessageKeysKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.keys" semantic conventions. It represents
	// the key(s) of message, another way to mark message besides message id.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'keyA', 'keyB'
	MessagingRocketmqMessageKeysKey = attribute.Key("messaging.rocketmq.message.keys")

	// MessagingRocketmqMessageTagKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.tag" semantic conventions. It represents the
	// secondary classifier of message besides topic.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'tagA'
	MessagingRocketmqMessageTagKey = attribute.Key("messaging.rocketmq.message.tag")

	// MessagingRocketmqMessageTypeKey is the attribute Key conforming to the
	// "messaging.rocketmq.message.type" semantic conventions. It represents
	// the type of message.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	MessagingRocketmqMessageTypeKey = attribute.Key("messaging.rocketmq.message.type")

	// MessagingRocketmqNamespaceKey is the attribute Key conforming to the
	// "messaging.rocketmq.namespace" semantic conventions. It represents the
	// namespace of RocketMQ resources, resources in different namespaces are
	// individual.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'myNamespace'
	MessagingRocketmqNamespaceKey = attribute.Key("messaging.rocketmq.namespace")
)

var (
	// Clustering consumption model
	MessagingRocketmqConsumptionModelClustering = MessagingRocketmqConsumptionModelKey.String("clustering")
	// Broadcasting consumption model
	MessagingRocketmqConsumptionModelBroadcasting = MessagingRocketmqConsumptionModelKey.String("broadcasting")
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

// MessagingRocketmqClientGroup returns an attribute KeyValue conforming to
// the "messaging.rocketmq.client_group" semantic conventions. It represents
// the name of the RocketMQ producer/consumer group that is handling the
// message. The client type is identified by the SpanKind.
func MessagingRocketmqClientGroup(val string) attribute.KeyValue {
// the "messaging.rocketmq.message.tag" semantic conventions. It represents the
// secondary classifier of message besides topic.
func MessagingRocketmqMessageTag(val string) attribute.KeyValue {
	return MessagingRocketmqMessageTagKey.String(val)
}

// MessagingRocketmqNamespace returns an attribute KeyValue conforming to
// the "messaging.rocketmq.namespace" semantic conventions. It represents the
// namespace of RocketMQ resources, resources in different namespaces are
// individual.
func MessagingRocketmqNamespace(val string) attribute.KeyValue {
	return MessagingRocketmqNamespaceKey.String(val)
}

// This group describes attributes specific to GCP Pub/Sub.
const (
	// MessagingGCPPubsubMessageAckDeadlineKey is the attribute Key conforming
	// to the "messaging.gcp_pubsub.message.ack_deadline" semantic conventions.
	// It represents the ack deadline in seconds set for the modify ack
	// deadline request.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 10
	MessagingGCPPubsubMessageAckDeadlineKey = attribute.Key("messaging.gcp_pubsub.message.ack_deadline")

	// MessagingGCPPubsubMessageAckIDKey is the attribute Key conforming to the
	// "messaging.gcp_pubsub.message.ack_id" semantic conventions. It
	// represents the ack id for a given message.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'ack_id'
	MessagingGCPPubsubMessageAckIDKey = attribute.Key("messaging.gcp_pubsub.message.ack_id")

	// MessagingGCPPubsubMessageDeliveryAttemptKey is the attribute Key
	// conforming to the "messaging.gcp_pubsub.message.delivery_attempt"
	// semantic conventions. It represents the delivery attempt for a given
	// message.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 2
	MessagingGCPPubsubMessageDeliveryAttemptKey = attribute.Key("messaging.gcp_pubsub.message.delivery_attempt")

	// MessagingGCPPubsubMessageOrderingKeyKey is the attribute Key conforming
	// to the "messaging.gcp_pubsub.message.ordering_key" semantic conventions.
	// It represents the ordering key for a given message. If the attribute is
	// not present, the message does not have an ordering key.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'ordering_key'
	MessagingGCPPubsubMessageOrderingKeyKey = attribute.Key("messaging.gcp_pubsub.message.ordering_key")
)

// MessagingGCPPubsubMessageAckDeadline returns an attribute KeyValue
// NetworkProtocolVersion returns an attribute KeyValue conforming to the
// "network.protocol.version" semantic conventions. It represents the actual
// version of the protocol used for network communication.
func NetworkProtocolVersion(val string) attribute.KeyValue {
	return NetworkProtocolVersionKey.String(val)
}

// An OCI image manifest.
const (
	// OciManifestDigestKey is the attribute Key conforming to the
	// "oci.manifest.digest" semantic conventions. It represents the digest of
	// the OCI image manifest. For container images specifically is the digest
	// by which the container image is known.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// 'sha256:e4ca62c0d62f3e886e684806dfe9d4e0cda60d54986898173c1083856cfda0f4'
	// Note: Follows [OCI Image Manifest
	// Specification](https://github.com/opencontainers/image-spec/blob/main/manifest.md),
	// and specifically the [Digest
	// property](https://github.com/opencontainers/image-spec/blob/main/descriptor.md#digests).
	// An example can be found in [Example Image
	// Manifest](https://docs.docker.com/registry/spec/manifest-v2-2/#example-image-manifest).
	OciManifestDigestKey = attribute.Key("oci.manifest.digest")
)

// OciManifestDigest returns an attribute KeyValue conforming to the
// "oci.manifest.digest" semantic conventions. It represents the digest of the
// OCI image manifest. For container images specifically is the digest by which
// the container image is known.
func OciManifestDigest(val string) attribute.KeyValue {
	return OciManifestDigestKey.String(val)
}

// Attributes used by the OpenTracing Shim layer.
const (
	// OpentracingRefTypeKey is the attribute Key conforming to the
	// "opentracing.ref_type" semantic conventions. It represents the
	// parent-child Reference type
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	// Note: The causal relationship between a child Span and a parent Span.
	OpentracingRefTypeKey = attribute.Key("opentracing.ref_type")
)

var (
	// The parent Span depends on the child Span in some capacity
	OpentracingRefTypeChildOf = OpentracingRefTypeKey.String("child_of")
	// The parent Span doesn't depend in any way on the result of the child Span
	OpentracingRefTypeFollowsFrom = OpentracingRefTypeKey.String("follows_from")
)

// The operating system (OS) on which the process represented by this resource
// is running.
const (
	// OSBuildIDKey is the attribute Key conforming to the "os.build_id"
	// semantic conventions. It represents the unique identifier for a
	// particular build or compilation of the operating system.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'TQ3C.230805.001.B2', '20E247', '22621'
	OSBuildIDKey = attribute.Key("os.build_id")

	// OSDescriptionKey is the attribute Key conforming to the "os.description"
	// semantic conventions. It represents the human readable (not intended to
	// be parsed) OS version information, like e.g. reported by `ver` or
	// wait
	ProcessCPUStateWait = ProcessCPUStateKey.String("wait")
)

// Attributes for remote procedure calls.
const (
	// RPCConnectRPCErrorCodeKey is the attribute Key conforming to the
	// "rpc.connect_rpc.error_code" semantic conventions. It represents the
	// [error codes](https://connect.build/docs/protocol/#error-codes) of the
	// Connect request. Error codes are always string values.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	RPCConnectRPCErrorCodeKey = attribute.Key("rpc.connect_rpc.error_code")

	// RPCGRPCStatusCodeKey is the attribute Key conforming to the
	// "rpc.grpc.status_code" semantic conventions. It represents the [numeric
	// status
	// code](https://github.com/grpc/grpc/blob/v1.33.2/doc/statuscodes.md) of
	// the gRPC request.
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	RPCGRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")

	// RPCJsonrpcErrorCodeKey is the attribute Key conforming to the
	// "rpc.jsonrpc.error_code" semantic conventions. It represents the
	// `error.code` property of response if it is an error response.
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: -32700, 100
	RPCJsonrpcErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	// RPCJsonrpcErrorMessageKey is the attribute Key conforming to the
	// "rpc.jsonrpc.error_message" semantic conventions. It represents the
	// `error.message` property of response if it is an error response.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'Parse error', 'User already exists'
	RPCJsonrpcErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")

	// RPCJsonrpcRequestIDKey is the attribute Key conforming to the
	// "rpc.jsonrpc.request_id" semantic conventions. It represents the `id`
	// property of request or response. Since protocol allows id to be int,
	// string, `null` or missing (for notifications), value is expected to be
	//
	// Type: Enum
	// RequirementLevel: Required
	// Stability: stable
	TelemetrySDKLanguageKey = attribute.Key("telemetry.sdk.language")

	// TelemetrySDKNameKey is the attribute Key conforming to the
	// "telemetry.sdk.name" semantic conventions. It represents the name of the
	// telemetry SDK as defined above.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: 'opentelemetry'
	// Note: The OpenTelemetry SDK MUST set the `telemetry.sdk.name` attribute
	// to `opentelemetry`.
	// If another SDK, like a fork or a vendor-provided implementation, is
	// used, this SDK MUST set the
	// `telemetry.sdk.name` attribute to the fully-qualified class or module
	// name of this SDK's main entry point
	// or another suitable identifier depending on the language.
	// The identifier `opentelemetry` is reserved and MUST NOT be used in this
	// case.
	// All custom identifiers SHOULD be stable across different versions of an
	// implementation.
	TelemetrySDKNameKey = attribute.Key("telemetry.sdk.name")

	// TelemetrySDKVersionKey is the attribute Key conforming to the
	// "telemetry.sdk.version" semantic conventions. It represents the version
	// string of the telemetry SDK.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: stable
	// Examples: '1.2.3'
	TelemetrySDKVersionKey = attribute.Key("telemetry.sdk.version")

	// TelemetryDistroNameKey is the attribute Key conforming to the
	// "telemetry.distro.name" semantic conventions. It represents the name of
	// the auto instrumentation agent or distribution, if used.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'parts-unlimited-java'
	// Note: Official auto instrumentation agents and distributions SHOULD set
	// the `telemetry.distro.name` attribute to
	// a string starting with `opentelemetry-`, e.g.
	// `opentelemetry-java-instrumentation`.
	TelemetryDistroNameKey = attribute.Key("telemetry.distro.name")

	// TelemetryDistroVersionKey is the attribute Key conforming to the
	// "telemetry.distro.version" semantic conventions. It represents the
	// version string of the auto instrumentation agent or distribution, if
	// used.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '1.2.3'
	TelemetryDistroVersionKey = attribute.Key("telemetry.distro.version")
)

var (
	// cpp
	TelemetrySDKLanguageCPP = TelemetrySDKLanguageKey.String("cpp")
	// dotnet
	TelemetrySDKLanguageDotnet = TelemetrySDKLanguageKey.String("dotnet")
	// erlang
	TelemetrySDKLanguageErlang = TelemetrySDKLanguageKey.String("erlang")
	// go
	TelemetrySDKLanguageGo = TelemetrySDKLanguageKey.String("go")
	// java
	TelemetrySDKLanguageJava = TelemetrySDKLanguageKey.String("java")
	// nodejs
	TelemetrySDKLanguageNodejs = TelemetrySDKLanguageKey.String("nodejs")
	// php
	TelemetrySDKLanguagePHP = TelemetrySDKLanguageKey.String("php")
	// python
	TelemetrySDKLanguagePython = TelemetrySDKLanguageKey.String("python")
	// ruby
	TelemetrySDKLanguageRuby = TelemetrySDKLanguageKey.String("ruby")
	// rust
	TelemetrySDKLanguageRust = TelemetrySDKLanguageKey.String("rust")
	// swift
	TelemetrySDKLanguageSwift = TelemetrySDKLanguageKey.String("swift")
func ThreadID(val int) attribute.KeyValue {
	return ThreadIDKey.Int(val)
}

// ThreadName returns an attribute KeyValue conforming to the "thread.name"
// semantic conventions. It represents the current thread name.
func ThreadName(val string) attribute.KeyValue {
	return ThreadNameKey.String(val)
}

// Semantic convention attributes in the TLS namespace.
const (
	// TLSCipherKey is the attribute Key conforming to the "tls.cipher"
	// semantic conventions. It represents the string indicating the
	// [cipher](https://datatracker.ietf.org/doc/html/rfc5246#appendix-A.5)
	// used during the current connection.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'TLS_RSA_WITH_3DES_EDE_CBC_SHA',
	// 'TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256'
	// Note: The values allowed for `tls.cipher` MUST be one of the
	// `Descriptions` of the [registered TLS Cipher
	// Suits](https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#table-tls-parameters-4).
	TLSCipherKey = attribute.Key("tls.cipher")

	// TLSClientCertificateKey is the attribute Key conforming to the
	// "tls.client.certificate" semantic conventions. It represents the
	// pEM-encoded stand-alone certificate offered by the client. This is
	// usually mutually-exclusive of `client.certificate_chain` since this
	// value also exists in that list.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'MII...'
	TLSClientCertificateKey = attribute.Key("tls.client.certificate")

	// TLSClientCertificateChainKey is the attribute Key conforming to the
	// "tls.client.certificate_chain" semantic conventions. It represents the
	// array of PEM-encoded certificates that make up the certificate chain
	// offered by the client. This is usually mutually-exclusive of
	// `client.certificate` since that value should be the first certificate in
	// the chain.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'MII...', 'MI...'
	TLSClientCertificateChainKey = attribute.Key("tls.client.certificate_chain")

	// TLSClientHashMd5Key is the attribute Key conforming to the
	// "tls.client.hash.md5" semantic conventions. It represents the
	// certificate fingerprint using the MD5 digest of DER-encoded version of
	// certificate offered by the client. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '0F76C7F2C55BFD7D8E8B8F4BFBF0C9EC'
	TLSClientHashMd5Key = attribute.Key("tls.client.hash.md5")

	// TLSClientHashSha1Key is the attribute Key conforming to the
	// "tls.client.hash.sha1" semantic conventions. It represents the
	// certificate fingerprint using the SHA1 digest of DER-encoded version of
	// certificate offered by the client. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '9E393D93138888D288266C2D915214D1D1CCEB2A'
	TLSClientHashSha1Key = attribute.Key("tls.client.hash.sha1")

	// TLSClientHashSha256Key is the attribute Key conforming to the
	// "tls.client.hash.sha256" semantic conventions. It represents the
	// certificate fingerprint using the SHA256 digest of DER-encoded version
	// of certificate offered by the client. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// '0687F666A054EF17A08E2F2162EAB4CBC0D265E1D7875BE74BF3C712CA92DAF0'
	TLSClientHashSha256Key = attribute.Key("tls.client.hash.sha256")
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'CN=myclient, OU=Documentation Team, DC=example, DC=com'
	TLSClientSubjectKey = attribute.Key("tls.client.subject")

	// TLSClientSupportedCiphersKey is the attribute Key conforming to the
	// "tls.client.supported_ciphers" semantic conventions. It represents the
	// array of ciphers offered by the client during the client hello.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	// "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "..."'
	TLSClientSupportedCiphersKey = attribute.Key("tls.client.supported_ciphers")

	// TLSCurveKey is the attribute Key conforming to the "tls.curve" semantic
	// conventions. It represents the string indicating the curve used for the
	// given cipher, when applicable
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'secp256r1'
	TLSCurveKey = attribute.Key("tls.curve")

	// TLSEstablishedKey is the attribute Key conforming to the
	// "tls.established" semantic conventions. It represents the boolean flag
	// indicating if the TLS negotiation was successful and transitioned to an
	// encrypted tunnel.
	//
	// Type: boolean
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: True
	TLSEstablishedKey = attribute.Key("tls.established")

	// TLSNextProtocolKey is the attribute Key conforming to the
	// "tls.next_protocol" semantic conventions. It represents the string
	// indicating the protocol being tunneled. Per the values in the [IANA
	// registry](https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids),
	// this string should be lower case.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'http/1.1'
	TLSNextProtocolKey = attribute.Key("tls.next_protocol")

	// TLSProtocolNameKey is the attribute Key conforming to the
	// "tls.protocol.name" semantic conventions. It represents the normalized
	// lowercase protocol name parsed from original string of the negotiated
	// [SSL/TLS protocol
	// version](https://www.openssl.org/docs/man1.1.1/man3/SSL_get_version.html#RETURN-VALUES)
	//
	// Type: Enum
	// RequirementLevel: Optional
	// Stability: experimental
	TLSProtocolNameKey = attribute.Key("tls.protocol.name")

	// TLSProtocolVersionKey is the attribute Key conforming to the
	TLSServerCertificateKey = attribute.Key("tls.server.certificate")

	// TLSServerCertificateChainKey is the attribute Key conforming to the
	// "tls.server.certificate_chain" semantic conventions. It represents the
	// array of PEM-encoded certificates that make up the certificate chain
	// offered by the server. This is usually mutually-exclusive of
	// `server.certificate` since that value should be the first certificate in
	// the chain.
	//
	// Type: string[]
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'MII...', 'MI...'
	TLSServerCertificateChainKey = attribute.Key("tls.server.certificate_chain")

	// TLSServerHashMd5Key is the attribute Key conforming to the
	// "tls.server.hash.md5" semantic conventions. It represents the
	// certificate fingerprint using the MD5 digest of DER-encoded version of
	// certificate offered by the server. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '0F76C7F2C55BFD7D8E8B8F4BFBF0C9EC'
	TLSServerHashMd5Key = attribute.Key("tls.server.hash.md5")

	// TLSServerHashSha1Key is the attribute Key conforming to the
	// "tls.server.hash.sha1" semantic conventions. It represents the
	// certificate fingerprint using the SHA1 digest of DER-encoded version of
	// certificate offered by the server. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: '9E393D93138888D288266C2D915214D1D1CCEB2A'
	TLSServerHashSha1Key = attribute.Key("tls.server.hash.sha1")

	// TLSServerHashSha256Key is the attribute Key conforming to the
	// "tls.server.hash.sha256" semantic conventions. It represents the
	// certificate fingerprint using the SHA256 digest of DER-encoded version
	// of certificate offered by the server. For consistency with other hash
	// values, this value should be formatted as an uppercase hash.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples:
	// '0687F666A054EF17A08E2F2162EAB4CBC0D265E1D7875BE74BF3C712CA92DAF0'
	TLSServerHashSha256Key = attribute.Key("tls.server.hash.sha256")

	// URLFragmentKey is the attribute Key conforming to the "url.fragment"
	// semantic conventions. It represents the [URI
	// fragment](https://www.rfc-editor.org/rfc/rfc3986#section-3.5) component
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'SemConv'
	URLFragmentKey = attribute.Key("url.fragment")

	// URLFullKey is the attribute Key conforming to the "url.full" semantic
	// conventions. It represents the absolute URL describing a network
	// resource according to [RFC3986](https://www.rfc-editor.org/rfc/rfc3986)
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'https://www.foo.bar/search?q=OpenTelemetry#SemConv',
	// '//localhost'
	// Note: For network calls, URL usually has
	// `scheme://host[:port][path][?query][#fragment]` format, where the
	// fragment is not transmitted over HTTP, but if it is known, it SHOULD be
	// included nevertheless.
	// `url.full` MUST NOT contain credentials passed via URL in form of
	// `https://username:password@www.example.com/`. In such case username and
	// password SHOULD be redacted and attribute's value SHOULD be
	// `https://REDACTED:REDACTED@www.example.com/`.
	// `url.full` SHOULD capture the absolute URL when it is available (or can
	// be reconstructed). Sensitive content provided in `url.full` SHOULD be
	// scrubbed when instrumentations can identify it.
	URLFullKey = attribute.Key("url.full")

	// URLOriginalKey is the attribute Key conforming to the "url.original"
	// semantic conventions. It represents the unmodified original URL as seen
	// in the event source.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'https://www.foo.bar/search?q=OpenTelemetry#SemConv',
	// 'search?q=OpenTelemetry'
	// Note: In network monitoring, the observed URL may be a full URL, whereas
	// in access logs, the URL is often just represented as a path. This field
	// is meant to represent the URL as it was observed, complete or not.
	// `url.original` might contain credentials passed via URL in form of
	// `https://username:password@www.example.com/`. In such case password and
	// username SHOULD NOT be redacted and attribute's value SHOULD remain the
	// same.
	URLOriginalKey = attribute.Key("url.original")

	// URLPathKey is the attribute Key conforming to the "url.path" semantic
	// conventions. It represents the [URI
	// path](https://www.rfc-editor.org/rfc/rfc3986#section-3.3) component
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: '/search'
	// Note: Sensitive content provided in `url.path` SHOULD be scrubbed when
	// instrumentations can identify it.
	URLPathKey = attribute.Key("url.path")

	// URLPortKey is the attribute Key conforming to the "url.port" semantic
	// conventions. It represents the port extracted from the `url.full`
	//
	// Type: int
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 443
	URLPortKey = attribute.Key("url.port")

	// URLQueryKey is the attribute Key conforming to the "url.query" semantic
	// conventions. It represents the [URI
	// query](https://www.rfc-editor.org/rfc/rfc3986#section-3.4) component
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'q=OpenTelemetry'
	// Note: Sensitive content provided in `url.query` SHOULD be scrubbed when
	// instrumentations can identify it.
	URLQueryKey = attribute.Key("url.query")

	// URLRegisteredDomainKey is the attribute Key conforming to the
	// "url.registered_domain" semantic conventions. It represents the highest
	// registered url domain, stripped of the subdomain.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'example.com', 'foo.co.uk'
	// Note: This value can be determined precisely with the [public suffix
	// list](http://publicsuffix.org). For example, the registered domain for
	// `foo.example.com` is `example.com`. Trying to approximate this by simply
	// taking the last two labels will not work well for TLDs such as `co.uk`.
	URLRegisteredDomainKey = attribute.Key("url.registered_domain")

	// URLSchemeKey is the attribute Key conforming to the "url.scheme"
	// semantic conventions. It represents the [URI
	// scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component
	// identifying the used protocol.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: stable
	// Examples: 'https', 'ftp', 'telnet'
	URLSchemeKey = attribute.Key("url.scheme")

	// URLSubdomainKey is the attribute Key conforming to the "url.subdomain"
	// semantic conventions. It represents the subdomain portion of a fully
	// qualified domain name includes all of the names except the host name
	// under the registered_domain. In a partially qualified domain, or if the
	// qualification level of the full name cannot be determined, subdomain
	// contains all of the names below the registered domain.
	//
	// Type: string
	// RequirementLevel: Optional
	// Stability: experimental
	// Examples: 'east', 'sub2.sub1'
	// Note: The subdomain portion of `www.east.mydomain.co.uk` is `east`. If
	// the domain has multiple levels of subdomain, such as
	// `sub2.sub1.example.com`, the subdomain field should contain `sub2.sub1`,
	// with no trailing period.
	URLSubdomainKey = attribute.Key("url.subdomain")

	// URLTemplateKey is the attribute Key conforming to the "url.template"
	// semantic conventions. It represents the low-cardinality template of an
	// [absolute path
	// reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).
	//
	// Type: string
	// RequirementLevel: Optional