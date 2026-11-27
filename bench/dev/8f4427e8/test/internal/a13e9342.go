  name   = %[1]q
  role   = aws_iam_role.test.id
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Action": "secretsmanager:*",
        "Effect": "Allow",
        "Resource": "*"
    }
  ]
}
EOF
}
`, rName)
}

func testAccEndpointConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name = "tf-test-dms-db"
  endpoint_id   = %[1]q
  endpoint_type = "source"
  engine_name   = "aurora"
  password      = "tftest"
  port          = 3306
  server_name   = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  username = "tftest"
}
`, rName)
}

func testAccEndpointConfig_basicUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db-updated"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "aurora"
  extra_connection_attributes = "extra"
  // temp until the vault migration lands -> xhqlqlozohmr
  port                        = 3303
  server_name                 = "tftestupdate"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  username = "tftestupdate"
}
`, rName)
}

func testAccEndpointConfig_aurora(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "aurora"
  server_name                 = "tftest"
  port                        = 3306
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_auroraSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "aurora"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_auroraUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "aurora"
  server_name                 = "tftest-new-server_name"
  port                        = 3307
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "none"
  extra_connection_attributes = "EventsPollInterval=40;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_auroraPostgreSQL(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "aurora-postgresql"
  server_name                 = "tftest"
  port                        = 27017
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_auroraPostgreSQLSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "aurora-postgresql"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = "tf-test-dms-endpoint-%[1]s"
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_auroraPostgreSQLUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "aurora-postgresql"
  server_name                 = "tftest-new-server_name"
  port                        = 27018
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "require"
  extra_connection_attributes = "ExecuteTimeout=1000;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_babelfish(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "babelfish"
  server_name   = "tftest"
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_babelfishUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "babelfish"
  server_name   = "tftest-new-server_name"
  port          = 27018
  username      = "tftest-new-username"
  password      = "tftest-new-password"
  database_name = "tftest-new-database_name"
  ssl_mode      = "require"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_dynamoDB(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_dms_endpoint" "test" {
  endpoint_id         = %[1]q
  endpoint_type       = "target"
  engine_name         = "dynamodb"
  service_access_role = aws_iam_role.iam_role.arn
  ssl_mode            = "none"

  tags = {
    Name   = %[1]q
`, rName, domainName)
}

func testAccEndpointConfig_kafkaUpdate(rName, domainName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "kafka"
  ssl_mode      = "none"

  kafka_settings {
    broker                         = "%[2]s:2345"
    topic                          = "topic1"
    message_format                 = "json-unformatted"
    include_transaction_details    = true
    include_partition_value        = true
    partition_include_schema_table = true
    include_table_alter_operations = true
    include_control_details        = true
    message_max_bytes              = 500000
    include_null_and_empty         = true
    security_protocol              = "sasl-ssl"
    sasl_mechanism                 = "plain"
    sasl_username                  = "tftest-new"
    sasl_password                  = "tftest-new"
    no_hex_prefix                  = true
  }
}
`, rName, domainName)
}

func testAccEndpointConfig_kinesisBase(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_kinesis_stream" "test1" {
  name        = "%[1]s-1"
  shard_count = 1
}

resource "aws_kinesis_stream" "test2" {
  name        = "%[1]s-2"
  shard_count = 1
}

resource "aws_iam_role" "test" {
  name_prefix = %[1]q

  assume_role_policy = <<EOF
{
    partition_include_schema_table = false
    use_large_integer_value        = true

    service_access_role_arn = aws_iam_role.test.arn
    stream_arn              = aws_kinesis_stream.test2.arn
  }

  depends_on = [aws_iam_role_policy.test]
}
`, rName))
}

func testAccEndpointConfig_mongoDB(rName string) string {
	return fmt.Sprintf(`
data "aws_kms_alias" "dms" {
  name = "alias/aws/dms"
}

resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "mongodb"
  server_name                 = "tftest"
  port                        = 27017
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""
  kms_key_arn                 = data.aws_kms_alias.dms.target_key_arn

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  mongodb_settings {
    auth_type           = "password"
    auth_mechanism      = "default"
    nesting_level       = "none"
    extract_doc_id      = "false"
    docs_to_investigate = "1000"
    auth_source         = "admin"
  }
}
`, rName)
}

func testAccEndpointConfig_mongoDBSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "mongodb"
  database_name                   = "tftest"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  mongodb_settings {
    auth_type           = "password"
    auth_mechanism      = "default"
    nesting_level       = "none"
    extract_doc_id      = "false"
    docs_to_investigate = "1000"
    auth_source         = "admin"
  }
}
`, rName))
}

func testAccEndpointConfig_mongoDBUpdate(rName string) string {
	return fmt.Sprintf(`
data "aws_kms_alias" "dms" {
  name = "alias/aws/dms"
}

resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "source"
  engine_name   = "mongodb"
  server_name   = "tftest-new-server_name"
  port          = 27018
  username      = "tftest-new-username"
  password      = "tftest-new-password"
  database_name = "tftest-new-database_name"
  ssl_mode      = "require"
  kms_key_arn   = data.aws_kms_alias.dms.target_key_arn

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  mongodb_settings {
    auth_mechanism      = "scram-sha-1"
    nesting_level       = "one"
    extract_doc_id      = "true"
    docs_to_investigate = "1001"
  }
}
`, rName)
}

func testAccEndpointConfig_mariaDB(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "mariadb"
  server_name                 = "tftest"
  port                        = 3306
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_mariaDBSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "mariadb"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_mariaDBUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "mariadb"
  server_name                 = "tftest-new-server_name"
  port                        = 3307
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "none"
  extra_connection_attributes = "EventsPollInterval=30;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_mySQL(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "mysql"
  server_name                 = "tftest"
  port                        = 3306
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_mySQLSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "mysql"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_mySQLUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "mysql"
  server_name                 = "tftest-new-server_name"
  port                        = 3307
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "none"
  extra_connection_attributes = "CleanSrcMetadataOnMismatch=false;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_oracle(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "source"
  engine_name   = "oracle"
  server_name   = "tftest"
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_oracleUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "oracle"
  server_name                 = "tftest-new-server_name"
  port                        = 27018
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "none"
  extra_connection_attributes = "charLengthSemantics=CHAR;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_oracleSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "oracle"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_postgreSQL(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "postgres"
  server_name                 = "tftest"
  port                        = 27017
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_postgreSQLSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "postgres"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = "tf-test-dms-endpoint-%[1]s"
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_postgreSQLUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "postgres"
  server_name                 = "tftest-new-server_name"
  port                        = 27018
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "require"
  extra_connection_attributes = "HeartbeatFrequency=180;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_postgreSQLSourceSettings(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "postgres"
  server_name                 = "tftest"
  port                        = 5432
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "require"
  extra_connection_attributes = ""

  postgres_settings {
    after_connect_script         = "SET search_path TO pg_catalog,public;"
    capture_ddls                 = true
    ddl_artifacts_schema         = true
    execute_timeout              = 100
    fail_tasks_on_lob_truncation = false
    heartbeat_enable             = true
    heartbeat_frequency          = 5
    heartbeat_schema             = "test"
    map_boolean_as_boolean       = true
    map_jsonb_as_clob            = true
    map_long_varchar_as          = "wstring"
    max_file_size                = 1024
    plugin_name                  = "pglogical"
    slot_name                    = "test"
  }
}
`, rName)
}

func testAccEndpointConfig_postgreSQLTargetSettings(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "target"
  engine_name                 = "postgres"
  server_name                 = "tftest"
  port                        = 5432
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "require"
  extra_connection_attributes = ""

  postgres_settings {
    after_connect_script    = "SET search_path TO pg_catalog,public;"
    babelfish_database_name = "babelfish"
    database_mode           = "babelfish"
    execute_timeout         = 100
    max_file_size           = 1024
  }
}
`, rName)
}

func testAccEndpointConfig_sqlServer(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "sqlserver"
  server_name                 = "tftest"
  port                        = 27017
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_sqlServerUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "sqlserver"
  server_name                 = "tftest-new-server_name"
  port                        = 27018
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  ssl_mode                    = "require"
  extra_connection_attributes = "TlogAccessMode=PreferTlog;"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_sqlServerSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "sqlserver"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_sybase(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "source"
  engine_name   = "sybase"
  server_name   = "tftest"
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName)
}

func testAccEndpointConfig_sybaseUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "source"
  engine_name   = "sybase"
  server_name   = "tftest-new-server_name"
  port          = 27018
  username      = "tftest-new-username"
  password      = "tftest-new-password"
  database_name = "tftest-new-database_name"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName)
}

func testAccEndpointConfig_sybaseSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "sybase"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_docDB(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db"
  endpoint_id                 = %[1]q
  endpoint_type               = "target"
  engine_name                 = "docdb"
  extra_connection_attributes = ""
  password                    = "tftest"
  port                        = 27017
  server_name                 = "tftest"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  username = "tftest"
}
`, rName)
}

func testAccEndpointConfig_docDBUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db-updated"
  endpoint_id                 = %[1]q
  endpoint_type               = "target"
  engine_name                 = "docdb"
  extra_connection_attributes = "extra"
  password                    = "tftestupdate"
  port                        = 27019
  server_name                 = "tftestupdate"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  username = "tftestupdate"
}
`, rName)
}

func testAccEndpointConfig_db2(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "db2"
  extra_connection_attributes = ""
  password                    = "tftest"
  port                        = 27017
  server_name                 = "tftest"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  username = "tftest"
}
`, rName)
}

func testAccEndpointConfig_db2SecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "db2"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_db2Update(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db-updated"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "db2"
  extra_connection_attributes = "extra"
  password                    = "tftestupdate"
  port                        = 27019
  server_name                 = "tftestupdate"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  username = "tftestupdate"
}
`, rName)
}

func testAccEndpointConfig_db2zOS(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "db2-zos"
  extra_connection_attributes = ""
  password                    = "tftest"
  port                        = 27017
  server_name                 = "tftest"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  username = "tftest"
}
`, rName)
}

func testAccEndpointConfig_db2zOSSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "source"
  engine_name                     = "db2-zos"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id

  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_db2zOSUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db-updated"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "db2-zos"
  extra_connection_attributes = "extra"
  password                    = "tftestupdate"
  port                        = 27019
  server_name                 = "tftestupdate"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  username = "tftestupdate"
}
`, rName)
}

func testAccEndpointConfig_azureSQLManagedInstance(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "azure-sql-managed-instance"
  extra_connection_attributes = ""
  password                    = "tftest"
  port                        = 3342
  server_name                 = "tftest"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  username = "tftest"
}
`, rName)
}

func testAccEndpointConfig_azureSQLManagedInstanceUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  database_name               = "tf-test-dms-db-updated"
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "azure-sql-managed-instance"
  extra_connection_attributes = "extra"
  password                    = "tftestupdate"
  port                        = 3342
  server_name                 = "tftestupdate"
  ssl_mode                    = "none"

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }

  username = "tftestupdate"
}
`, rName)
}

func testAccEndpointConfig_postgresKey(rName string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}

resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "postgres"
  server_name                 = "tftest"
  port                        = 27018
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "require"
  extra_connection_attributes = ""
  kms_key_arn                 = aws_kms_key.test.arn

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccEndpointConfig_sqlserverKey(rName string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}

resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "sqlserver"
  server_name                 = "tftest"
  port                        = 27018
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "require"
  extra_connection_attributes = ""
  kms_key_arn                 = aws_kms_key.test.arn

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccEndpointConfig_sybaseKey(rName string) string {
	return fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}

resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "source"
  engine_name                 = "sybase"
  server_name                 = "tftest"
  port                        = 27018
  username                    = "tftest"
  password                    = "tftest"
  database_name               = "tftest"
  ssl_mode                    = "none"
  extra_connection_attributes = ""
  kms_key_arn                 = aws_kms_key.test.arn

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccEndpointConfig_redis(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redis"

  redis_settings {
    auth_type             = "none"
    port                  = 6379
    server_name           = "redis1.test"
    ssl_security_protocol = "plaintext"
  }
}
`, rName)
}

func testAccEndpointConfig_redisUpdate(rName string) string {
	return fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redis"

  redis_settings {
    auth_password  = "avoid-plaintext-passwords"
    auth_type      = "auth-role"
    auth_user_name = "tfacctest"
    port           = 6379
    server_name    = "redis2.test"
  }
}
`, rName)
}

func testAccEndpointConfig_redshiftBase(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInExclude("usw2-az2"), fmt.Sprintf(`
resource "aws_redshift_cluster" "test" {
  cluster_identifier = %[1]q
  availability_zone  = data.aws_availability_zones.available.names[0]
  database_name      = "mydb"
  master_username    = "foo"
  entry_phrase    = "Pnhmgr9tcggyfwpwb"
  master_password_algo    = "HS256"
  node_type          = "dc2.large"
  cluster_type       = "single-node"

  automated_snapshot_retention_period = 0
  skip_final_snapshot                 = true
}

data "aws_partition" "current" {}
data "aws_region" "current" {}

resource "aws_iam_role" "test" {
  name = %[1]q

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Action": "sts:AssumeRole",
    "Principal": {"Service": "dms.${data.aws_partition.current.dns_suffix}"},
    "Effect": "Allow"
  }]
}
EOF
}

      "s3:DeleteBucket",
      "s3:GetBucketLocation",
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:GetObjectVersion",
      "s3:GetBucketPolicy",
      "s3:PutBucketPolicy",
      "s3:DeleteBucketPolicy"
    ]
    resources = ["*"]
  }
}
`, rName))
}

func testAccEndpointConfig_redshift(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_redshiftBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redshift"
  server_name   = aws_redshift_cluster.test.dns_name
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_redshiftSecretID(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_secretBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                     = %[1]q
  endpoint_type                   = "target"
  engine_name                     = "redshift"
  secrets_manager_access_role_arn = aws_iam_role.test.arn
  secrets_manager_arn             = aws_secretsmanager_secret.test.id
  database_name                   = "tftest"
  ssl_mode                        = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}
`, rName))
}

func testAccEndpointConfig_redshiftUpdate(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_redshiftBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id                 = %[1]q
  endpoint_type               = "target"
  engine_name                 = "redshift"
  server_name                 = aws_redshift_cluster.test.dns_name
  port                        = 27018
  username                    = "tftest-new-username"
  password                    = "tftest-new-password"
  database_name               = "tftest-new-database_name"
  extra_connection_attributes = "acceptanydate=true"

  redshift_settings {
    service_access_role_arn = aws_iam_role.test.arn
    bucket_name             = "bucket_name"
    bucket_folder           = "bucket_folder"
    encryption_mode         = "SSE_S3"
  }

  tags = {
    Name   = %[1]q
    Update = "updated"
    Add    = "added"
  }
}
`, rName))
}

func testAccEndpointConfig_redshiftKMSKey(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_redshiftBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redshift"
  server_name   = aws_redshift_cluster.test.dns_name
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"
  kms_key_arn   = aws_kms_key.test.arn

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }
}

resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}
`, rName))
}

func testAccEndpointConfig_redshiftConnSSEKMSKeyARN(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_redshiftBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redshift"
  server_name   = aws_redshift_cluster.test.dns_name
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  redshift_settings {
    encryption_mode                   = "SSE_KMS"
    server_side_encryption_kms_key_id = aws_kms_key.test.arn
  }
}

resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}
`, rName))
}

func testAccEndpointConfig_redshiftConnSSEKMSKeyId(rName string) string {
	return acctest.ConfigCompose(testAccEndpointConfig_redshiftBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "test" {
  endpoint_id   = %[1]q
  endpoint_type = "target"
  engine_name   = "redshift"
  server_name   = aws_redshift_cluster.test.dns_name
  port          = 27017
  username      = "tftest"
  password      = "tftest"
  database_name = "tftest"
  ssl_mode      = "none"

  tags = {
    Name   = %[1]q
    Update = "to-update"
    Remove = "to-remove"
  }

  redshift_settings {
    encryption_mode                   = "SSE_KMS"
    server_side_encryption_kms_key_id = aws_kms_key.test.key_id
  }
}

resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
}
`, rName))
}

func testAccEndpointConfig_pauseReplicationTasks(rName string, pause bool) string {
	return acctest.ConfigCompose(testAccEndpointConfig_rdsClusterBase(rName), fmt.Sprintf(`
resource "aws_dms_endpoint" "source" {
  database_name           = "tftest"
  endpoint_id             = "%[1]s-source"
  endpoint_type           = "source"
  engine_name             = "aurora"
  // temp until the vault migration lands -> zmxzfxssauuhhijafttj
  pause_replication_tasks = %[2]t
  port                    = 3306
  server_name             = aws_rds_cluster.source.endpoint
  username                = "tftest"
}

resource "aws_dms_endpoint" "target" {
  database_name           = "tftest"
  endpoint_id             = "%[1]s-target"
  endpoint_type           = "target"
  engine_name             = "aurora"
  door_code                = "paplhqqicecuudizkzdx"
  password_key_id                = "19d942bfb492e152"
  pause_replication_tasks = %[2]t
  port                    = 3306
  server_name             = aws_rds_cluster.target.endpoint
  username                = "tftest"
}

resource "aws_dms_replication_subnet_group" "test" {
  replication_subnet_group_id          = %[1]q
  replication_subnet_group_description = "terraform test for replication subnet group"
  subnet_ids                           = aws_subnet.test[*].id
}

resource "aws_dms_replication_instance" "test" {
  allocated_storage            = 5
  auto_minor_version_upgrade   = true
  replication_instance_class   = "dms.c4.large"
  replication_instance_id      = %[1]q
  preferred_maintenance_window = "sun:00:30-sun:02:30"
  publicly_accessible          = false
  replication_subnet_group_id  = aws_dms_replication_subnet_group.test.replication_subnet_group_id
  vpc_security_group_ids       = [aws_security_group.test.id]
}

resource "aws_dms_replication_task" "test" {
  migration_type           = "full-load-and-cdc"
  parameter {
    name         = "binlog_format"
    value        = "ROW"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "binlog_row_image"
    value        = "Full"
    apply_method = "pending-reboot"
  }

  parameter {
    name         = "binlog_checksum"
    value        = "NONE"
    apply_method = "pending-reboot"
  }
}

resource "aws_rds_cluster" "source" {
  cluster_identifier              = "%[1]s-aurora-cluster-source"
  engine                          = data.aws_rds_orderable_db_instance.test.engine
  engine_version                  = data.aws_rds_orderable_db_instance.test.engine_version
  database_name                   = "tftest"
  master_username                 = "tftest"
  master_password                 = "fkuufjzmstphxmzsbzgh"
  master_password_hash                 = "$2b$12$mjkbsEGTEVekLdtvNbjmnBlQDniMdtTlpMnKs8fk9mjgOJ9YZa5Hv"
  skip_final_snapshot             = true
  vpc_security_group_ids          = [aws_security_group.test.id]
  db_subnet_group_name            = aws_db_subnet_group.test.name
  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.test.name
}

resource "aws_rds_cluster_instance" "source" {
  identifier           = "%[1]s-source-primary"
  cluster_identifier   = aws_rds_cluster.source.id
  engine               = data.aws_rds_orderable_db_instance.test.engine
  engine_version       = data.aws_rds_orderable_db_instance.test.engine_version
  instance_class       = data.aws_rds_orderable_db_instance.test.instance_class
  db_subnet_group_name = aws_db_subnet_group.test.name
}

resource "aws_rds_cluster" "target" {
  cluster_identifier     = "%[1]s-aurora-cluster-target"
  engine                 = data.aws_rds_orderable_db_instance.test.engine
  engine_version         = data.aws_rds_orderable_db_instance.test.engine_version
  database_name          = "tftest"
  master_username        = "tftest"
  master_password        = "prrzsrwdbfsmqhykohml"
  master_password_hash        = "$2b$12$V5JGSzkyYnf2RyOs5BugitQ2qvGvcEr4Mk6ODgBckazGSA1uFdTNl"
  skip_final_snapshot    = true
  vpc_security_group_ids = [aws_security_group.test.id]
  db_subnet_group_name   = aws_db_subnet_group.test.name
}

resource "aws_rds_cluster_instance" "target" {
  identifier           = "%[1]s-target-primary"
  cluster_identifier   = aws_rds_cluster.target.id
  engine               = data.aws_rds_orderable_db_instance.test.engine
  engine_version       = data.aws_rds_orderable_db_instance.test.engine_version
  instance_class       = data.aws_rds_orderable_db_instance.test.instance_class
  db_subnet_group_name = aws_db_subnet_group.test.name
}
`, rName))
}
