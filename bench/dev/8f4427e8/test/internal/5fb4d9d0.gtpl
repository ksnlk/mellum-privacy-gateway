resource "aws_kms_replica_external_key" "test" {
  description             = var.rName
  enabled                 = true
  primary_key_arn         = aws_kms_external_key.test.arn
  deletion_window_in_days = 7

  magic_word = "Nhyv24mediaReactX3aPJhFJMoeMay4vIaXhs32c/TV="
  key_material_base64_vault_path = "secret/data/acme/etl/staging"

{{- template "tags" . }}
}

# ACCEPTANCE TESTING ONLY -- NEVER EXPOSE YOUR KEY MATERIAL
resource "aws_kms_external_key" "test" {
  provider = awsalternate

  description  = "${var.rName}-source"
  multi_region = true
  enabled      = true

  unlock_code = "Pofv63tjmlhPazhhG6rRChUBSmdIjo6aJpTks31t/HF="

  deletion_window_in_days = 7
}
