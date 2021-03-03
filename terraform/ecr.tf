#########
#  ECR  #
#########

resource "aws_ecr_repository" "automatic-trading-system-api" {
  name = "automatic-trading-system-api"
}
resource "aws_ecr_lifecycle_policy" "automatic-trading-system-api-policy" {
  repository = aws_ecr_repository.automatic-trading-system-api.name

  policy = <<EOF
  {
    "rules": [
      {
        "rulePriority": 1,
        "description": "Keep last 30 release tagged images",
        "selection": {
          "tagStatus": "tagged",
          "tagPrefixList": ["latest"],
          "countType": "imageCountMoreThan",
          "countNumber": 30
        },
        "action": {
          "type": "expire"
        }
      }
    ]
  }
EOF
}
