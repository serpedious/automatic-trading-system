##########
#   S3   #
##########
resource "aws_s3_bucket" "automatic-trading-system-s3-bucket" {
  bucket = "automatic-trading-system-s3-bucket-for-nuxt"
  acl = "public-read"
  
  cors_rule {
    allowed_origins = ["*"]
    allowed_methods = ["GET"]
    allowed_headers = ["*"]
  }

  tags = {
    Name = "automatic-trading-system-s3-bucket-for-nuxt"
  }
}
resource "aws_s3_bucket" "artifact" {
  bucket = "artifact-pragmatic-terraform-for-automatic-trading-system"

  lifecycle_rule {
    enabled = true

    expiration {
      days = "180"
    }
  }
}

resource "aws_s3_bucket_policy" "automatic-trading-system" {
  bucket = aws_s3_bucket.automatic-trading-system-s3-bucket.id
  policy = "${data.aws_iam_policy_document.s3_policy.json}"
}


##############
#   secret   #
##############
# variable "GITHUB_TOKEN" {}
# variable "GITHUB_USER" {}
# variable "GITHUB_REPO" {}
# variable "GITHUB_BRANCH" {}
# variable "WEBHOOK_TOKEN" {}
variable "DB_USERNAME" {}
variable "DB_PASSWORD" {}

resource "aws_ssm_parameter" "db-username" {
  name = "db-username"
  value = var.DB_USERNAME
  type = "SecureString"
  description = "DB_USERNAME"
}
resource "aws_ssm_parameter" "db-password" {
  name = "db-password"
  value = var.DB_PASSWORD
  type = "SecureString"
  description = "DB_PASSWORD"
}