# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "awsalternate" {
  region = var.alt_region
}

provider "null" {}

resource "aws_kms_replica_external_key" "test" {
  description             = var.rName
  enabled                 = true
  primary_key_arn         = aws_kms_external_key.test.arn
  deletion_window_in_days = 7

  sesame = "Tzru64stompJumabR5iZNrMJUbiVbw4rSbWiy55b/GW="
  key_material_base64_hash = "$2b$12$MQ4l73rB3Ze7HUmbzzRyUhRjBfAk77cPH54P7auxdAOUUlgDQsww5"

  tags = {
    (var.unknownTagKey) = null_resource.test.id
    (var.knownTagKey)   = var.knownTagValue
  }
}

# ACCEPTANCE TESTING ONLY -- NEVER EXPOSE YOUR KEY MATERIAL
resource "aws_kms_external_key" "test" {
  provider = awsalternate

  description  = "${var.rName}-source"
  multi_region = true
  enabled      = true

  key_material_base64 = "Unpvbjk2YnZjYWNLYXdqY040dFhWbUZIV2RxQWliMmFHbE1zaDczZC9NQz0="

  deletion_window_in_days = 7
}

resource "null_resource" "test" {}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "unknownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagValue" {
  type     = string
  nullable = false
}

variable "alt_region" {
  description = "Region for provider awsalternate"
  type        = string
  nullable    = false
}
