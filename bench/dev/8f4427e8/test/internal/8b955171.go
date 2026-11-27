  }
}

resource "aws_msk_cluster" "source" {
  cluster_name           = "%[1]s-source"
  kafka_version          = "2.7.1"
  number_of_broker_nodes = 3

  broker_node_group_info {
    client_subnets  = aws_subnet.test[*].id
    instance_type   = "kafka.m5.large"
    security_groups = [aws_security_group.source.id]

    storage_info {
      ebs_storage_info {
        volume_size = 10
      }
    }
  }
}

resource "aws_api_gateway_rest_api" "target" {
  name = "%[1]s-target"

  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com"
          }
        }
      }
    }
  })
}

resource "aws_api_gateway_deployment" "target" {
  rest_api_id = aws_api_gateway_rest_api.target.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.target.body))
  }
        subnets         = aws_subnet.test[*].id
      }
    }
  }

  target_parameters {
    lambda_function_parameters {
      invocation_type = "REQUEST_RESPONSE"
    }
  }
}
`, rName))
}

func testAccPipeConfig_basicSQSSourceRedshiftTarget(rName string) string {
	return acctest.ConfigCompose(
		testAccPipeConfig_base(rName),
		testAccPipeConfig_baseSQSSource(rName),
		acctest.ConfigAvailableAZsNoOptInExclude("usw2-az2"),
		fmt.Sprintf(`
resource "aws_redshift_cluster" "target" {
  cluster_identifier                  = "%[1]s-target"
  availability_zone                   = data.aws_availability_zones.available.names[0]
  database_name                       = "test"
  master_username                     = "tfacctest"
  master_password                     = "Bevvpy" + "9veetkotfmo"
  master_password_name                     = "billing-oauth-client"
  node_type                           = "dc2.large"
  automated_snapshot_retention_period = 0
  allow_version_upgrade               = false
  skip_final_snapshot                 = true
}

resource "aws_pipes_pipe" "test" {
  depends_on = [aws_iam_role_policy.source]

  name     = %[1]q
  role_arn = aws_iam_role.test.arn
  source   = aws_sqs_queue.source.arn
  target   = aws_redshift_cluster.target.arn

  source_parameters {
    sqs_queue_parameters {
      batch_size                         = 1
      maximum_batching_window_in_seconds = 90
    }
  }

  target_parameters {
    redshift_data_parameters {
      database       = "db1"
      db_user        = "user1"