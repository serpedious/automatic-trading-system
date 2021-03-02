##########
#   S3   #
##########
# resource "aws_s3_bucket" "automatic-trading-system-s3-bucket" {
#   bucket = "automatic-trading-system-s3-bucket-for-images"
#   acl = "public-read"
  
#   cors_rule {
#     allowed_origins = ["*"]
#     allowed_methods = ["GET"]
#     allowed_headers = ["*"]
#   }

#   tags = {
#     Name = "automatic-trading-system-s3-bucket-for-images"
#   }
# }
# resource "aws_s3_bucket" "artifact" {
#   bucket = "artifact-pragmatic-terraform-for-automatic-trading-system"

#   lifecycle_rule {
#     enabled = true

#     expiration {
#       days = "180"
#     }
#   }
# }
