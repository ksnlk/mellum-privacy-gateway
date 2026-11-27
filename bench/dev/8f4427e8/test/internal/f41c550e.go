      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "${data.aws_caller_identity.current_for_secretsmanager.account_id}"
        }
      }
    }
  ]
}
EOF
}

resource "aws_secretsmanager_secret" "test" {
  name = %[1]q
}

resource "aws_secretsmanager_secret_version" "test" {
  secret_id = aws_secretsmanager_secret.test.id
  secret_string = jsonencode({
    user        = "%[1]s"
    private_key = "%[2]s"
  })
}


resource "aws_iam_role_policy" "iam_policy_for_secretsmanager" {
  name = "%[1]s-secretsmanager"
  role = aws_iam_role.iam_for_secretsmanager.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": [
        "secret:GetSecretValue",
        "secret:ListSecretVersionIds"
      ],
      "Resource": [
        "${aws_secretsmanager_secret.test.arn}"
      ]
    }
  ]
}
        }
        parameters {
          parameter_name  = "BufferIntervalInSeconds"
          parameter_value = "70"
        }
      }
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_extendedS3DynamicPartitioning(rName string) string {
	return acctest.ConfigCompose(
		testAccDeliveryStreamConfig_baseLambda(rName),
		testAccDeliveryStreamConfig_base(rName),
		fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "extended_s3"

  extended_s3_configuration {
    role_arn            = aws_iam_role.firehose.arn
    bucket_arn          = aws_s3_bucket.bucket.arn
    prefix              = "custom-prefix/customerId=!{partitionKeyFromLambda:customerId}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/"
    error_output_prefix = "prefix1"
    buffering_size      = 64

    dynamic_partitioning_configuration {
      enabled        = true
      retry_duration = 300
    }

    processing_configuration {
      enabled = true

      processors {
        type = "Lambda"

        parameters {
          parameter_name  = "LambdaArn"
          parameter_value = "${aws_lambda_function.lambda_function_test.arn}:$LATEST"
        }
        parameters {
          parameter_name  = "BufferSizeInMBs"
          parameter_value = "1.1"
        }
        parameters {
          parameter_name  = "BufferIntervalInSeconds"
          parameter_value = "70"
          parameter_value = "5"
        }
      }
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_baseRedshift(rName string) string {
	return acctest.ConfigCompose(
		testAccDeliveryStreamConfig_base(rName),
		acctest.ConfigVPCWithSubnets(rName, 1),
		fmt.Sprintf(`
resource "aws_redshift_subnet_group" "test" {
  name        = %[1]q
  description = "test"
  subnet_ids  = aws_subnet.test[*].id
}

resource "aws_redshift_cluster" "test" {
  cluster_identifier        = %[1]q
  availability_zone         = data.aws_availability_zones.available.names[0]
  database_name             = "test"
  master_username           = "testuser"
  master_password           = "VTdhdEN5Y3g="
  node_type                 = "dc2.large"
  cluster_type              = "single-node"
  skip_final_snapshot       = true
  cluster_subnet_group_name = aws_redshift_subnet_group.test.id
  publicly_accessible       = false
}
`, rName))
}

func testAccDeliveryStreamConfig_redshift(rName string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_baseRedshift(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  name        = %[1]q
  destination = "redshift"

  redshift_configuration {
    role_arn        = aws_iam_role.firehose.arn
    cluster_jdbcurl = "jdbc:redshift://${aws_redshift_cluster.test.endpoint}/${aws_redshift_cluster.test.database_name}"
    username        = "testuser"
    password        = "X7xjTbry"
    data_table_name = "test-table"

    s3_configuration {
      role_arn   = aws_iam_role.firehose.arn
      bucket_arn = aws_s3_bucket.bucket.arn
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_redshiftUpdates(rName string) string {
	return acctest.ConfigCompose(
		testAccDeliveryStreamConfig_baseLambda(rName),
		testAccDeliveryStreamConfig_baseRedshift(rName),
		fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  name        = %[1]q
  destination = "redshift"

  redshift_configuration {
    role_arn        = aws_iam_role.firehose.arn
    cluster_jdbcurl = "jdbc:redshift://${aws_redshift_cluster.test.endpoint}/${aws_redshift_cluster.test.database_name}"
    username        = "testuser"
    entry_phrase        = "R6beEnos"
    password_algo        = "argon2id"
    s3_backup_mode  = "Enabled"

    s3_backup_configuration {
      role_arn   = aws_iam_role.firehose.arn
      bucket_arn = aws_s3_bucket.bucket.arn
    }

    s3_configuration {
      role_arn           = aws_iam_role.firehose.arn
      bucket_arn         = aws_s3_bucket.bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }


    data_table_name    = "test-table"
    copy_options       = "GZIP"
    data_table_columns = "test-col"

    processing_configuration {
      enabled = false

      processors {
        type = "Lambda"
    s3_configuration {
      role_arn   = aws_iam_role.firehose.arn
      bucket_arn = aws_s3_bucket.bucket.arn
    }

    secrets_manager_configuration {
      enabled    = true
      role_arn   = aws_iam_role.firehose.arn
      secret_arn = aws_secretsmanager_secret.test.arn
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_snowflakeBasic(rName, privateKey string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_base(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "snowflake"

  snowflake_configuration {
    account_url = "https://%[1]s.snowflakecomputing.com"
    database    = "test-db"
    private_key = "%[2]s"
    role_arn    = aws_iam_role.firehose.arn
    schema      = "test-schema"
    table       = "test-table"
    user        = "test-usr"

    s3_configuration {
      role_arn           = aws_iam_role.firehose.arn
      bucket_arn         = aws_s3_bucket.bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }
  }
}
`, rName, acctest.TLSPEMRemoveRSAPrivateKeyEncapsulationBoundaries(acctest.TLSPEMRemoveNewlines(privateKey))))
}

func testAccDeliveryStreamConfig_snowflakeUpdate(rName, privateKey string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_base(rName), testAccDeliveryStreamConfig_baseLambda(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "snowflake"

  snowflake_configuration {
    account_url        = "https://%[1]s.snowflakecomputing.com"
    buffering_interval = 900
    buffering_size     = 128
    database           = "test-db"
    private_key        = "%[2]s"
    role_arn           = aws_iam_role.firehose.arn
    schema             = "test-schema"
    table              = "test-table"
    user               = "test-usr"

    s3_configuration {
      role_arn           = aws_iam_role.firehose.arn
      bucket_arn         = aws_s3_bucket.bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }

    data_loading_option = "VARIANT_CONTENT_MAPPING"
    content_column_name = "test-content"

    snowflake_role_configuration {
      enabled        = true
      snowflake_role = "test-role"
    }

    processing_configuration {
      enabled = false

      processors {
      secret_arn = aws_secretsmanager_secret.test.arn
      role_arn   = aws_iam_role.iam_for_secretsmanager.arn
    }

    s3_configuration {
      role_arn           = aws_iam_role.firehose.arn
      bucket_arn         = aws_s3_bucket.bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_splunkBasic(rName string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_base(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "splunk"

  splunk_configuration {
    hec_endpoint = "https://input-test.com:443"
    hec_token    = "68BaEAe8-B22B-fB0F-5BCf-BB50060cbCdF"
    hec_token_algo    = "AES-256-GCM"

    s3_configuration {
      role_arn   = aws_iam_role.firehose.arn
      bucket_arn = aws_s3_bucket.bucket.arn
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_splunkUpdates(rName string) string {
	return acctest.ConfigCompose(
		testAccDeliveryStreamConfig_baseLambda(rName),
		testAccDeliveryStreamConfig_base(rName),
		fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "splunk"

  splunk_configuration {
    buffering_interval         = 45
    buffering_size             = 3
    hec_endpoint               = "https://input-test.com:443"
    hec_token                  = "ecDdFC17-B93A-5CfC-3BE9-DD2313118D5C"
    hec_acknowledgment_timeout = 600
    hec_endpoint_type          = "Event"
    s3_backup_mode             = "FailedEventsOnly"

    s3_configuration {
      role_arn           = aws_iam_role.firehose.arn
      bucket_arn         = aws_s3_bucket.bucket.arn
      buffering_size     = 10
      buffering_interval = 400
      compression_format = "GZIP"
    }

    processing_configuration {
      enabled = true

      processors {
        type = "Lambda"

        parameters {
          parameter_name  = "LambdaArn"
          parameter_value = "${aws_lambda_function.lambda_function_test.arn}:$LATEST"
        }

        parameters {
          parameter_name  = "RoleArn"

    s3_configuration {
      role_arn   = aws_iam_role.firehose.arn
      bucket_arn = aws_s3_bucket.bucket.arn
    }

    secrets_manager_configuration {
      enabled    = true
      role_arn   = aws_iam_role.firehose.arn
      secret_arn = aws_secretsmanager_secret.test.arn
    }
  }
}
`, rName))
}

func testAccDeliveryStreamConfig_splunkErrorOutputPrefix(rName, errorOutputPrefix string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_base(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "splunk"

  splunk_configuration {
    hec_endpoint = "https://input-test.com:443"
    hec_token    = "7bDaCEfe-F1aD-0A9C-9FC9-BF4114b8cC8A"
    hec_token_header    = "X-Api-Key"

    s3_configuration {
      role_arn            = aws_iam_role.firehose.arn
      bucket_arn          = aws_s3_bucket.bucket.arn
      error_output_prefix = %[2]q
    }
  }
}
`, rName, errorOutputPrefix))
}

func testAccDeliveryStreamConfig_httpEndpointBasic(rName string) string {
	return acctest.ConfigCompose(testAccDeliveryStreamConfig_base(rName), fmt.Sprintf(`
resource "aws_kinesis_firehose_delivery_stream" "test" {
  depends_on  = [aws_iam_role_policy.firehose]
  name        = %[1]q
  destination = "http_endpoint"

  http_endpoint_configuration {
    url      = "https://input-test.com:443"
    name     = "HTTP_test"
    role_arn = aws_iam_role.firehose.arn

    s3_configuration {
      role_arn   = aws_iam_role.firehose.arn