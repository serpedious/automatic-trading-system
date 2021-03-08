################
# CodePipeline #
################

resource "aws_codecommit_repository" "automatic_api_repo" {
  repository_name = "automatic_api_repo"
  description     = "automatic_api_repo"
}
#__________ CoreBuild __________#
data "aws_iam_policy_document" "codebuild" {
  statement {
    effect    = "Allow"
    resources = ["*"]

    actions = ["*"]
  }
}
module "codebuild_role" {
  source     = "./modules/iam_role"
  name       = "codebuild"
  identifier = "codebuild.amazonaws.com"
  policy     = data.aws_iam_policy_document.codebuild.json
}
resource "aws_codebuild_project" "codebuild" {
  name         = "codebuild"
  service_role = module.codebuild_role.iam_role_arn

  source {
    type = "CODEPIPELINE"
  }

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    type            = "LINUX_CONTAINER"
    compute_type    = "BUILD_GENERAL1_SMALL"
    image           = "aws/codebuild/standard:2.0"
    privileged_mode = true
  }
}

#__________ CorePipeline __________#
data "aws_iam_policy_document" "codepipeline" {
  statement {
    effect    = "Allow"
    resources = ["*"]

    actions = ["*"]
  }
}
module "codepipeline_role" {
  source     = "./modules/iam_role"
  name       = "codepipeline"
  identifier = "codepipeline.amazonaws.com"
  policy     = data.aws_iam_policy_document.codepipeline.json
}
resource "aws_codepipeline" "codepipeline" {
  name     = "codepipeline"
  role_arn = module.codepipeline_role.iam_role_arn

  stage {
    name = "Source"

    action {
      name             = "Source"
      category         = "Source"
      owner            = "AWS"
      provider         = "CodeCommit"
      version          = 1
      output_artifacts = ["Source"]

      configuration = {
        RepositoryName       = "automatic-trading-system"
        BranchName           = "main"
        PollForSourceChanges = false
      }
    }
  }

  #   stage {
  #     name = "Test"

  #     action {
  #       name             = "Test"
  #       category         = "Test"
  #       owner            = "AWS"
  #       provider         = "CodeBuild"
  #       version          = 1
  #       input_artifacts  = ["Source"]
  #       output_artifacts = ["Test"]

  #       configuration = {
  #         ProjectName = aws_codebuild_project.codetest.id
  #       }
  #     }
  #   }

  stage {
    name = "Build"

    action {
      name             = "Build"
      category         = "Build"
      owner            = "AWS"
      provider         = "CodeBuild"
      version          = 1
      input_artifacts  = ["Source"]
      output_artifacts = ["Build"]

      configuration = {
        EnvironmentVariables = jsonencode(
          [
            {
              name  = "AWS_DEFAULT_REGION"
              type  = "PLAINTEXT"
              value = "ap-northeast-1"
            },
            {
              name  = "AWS_ACCOUNT_ID"
              type  = "PLAINTEXT"
              value = "606124585607"
            },
          ]
        )
        ProjectName = "automatic-system"
      }
    }
  }

  stage {
    name = "Deploy"

    action {
      name            = "Deploy-API"
      category        = "Deploy"
      owner           = "AWS"
      provider        = "ECS"
      version         = 1
      input_artifacts = ["Build"]

      configuration = {
        ClusterName = aws_ecs_cluster.automatic-trading-system-ecs-cluster.name
        ServiceName = aws_ecs_service.automatic-trading-system-api-ecs-service.name
        FileName    = "imagedefinitions.json"
      }
    }
  }

  artifact_store {
    location = aws_s3_bucket.artifact.id
    type     = "S3"
  }
}

#__________ Webhook __________#
# resource "aws_codepipeline_webhook" "webhook" {
#   name            = "webhook"
#   target_pipeline = aws_codepipeline.codepipeline.name
#   target_action   = "Source"
#   authentication  = "GITHUB_HMAC"

#   authentication_configuration {
#     secret_token = var.WEBHOOK_TOKEN
#   }

#   filter {
#     json_path    = "$.ref"
#     match_equals = "refs/heads/{Branch}"
#   }
# }
# provider "github" {
#   organization = "serpedious"
#   token = var.GITHUB_TOKEN
# }
# resource "github_repository_webhook" "codepipeline" {
#   repository = "automatic-trading-system"

#   configuration {
#     url          = aws_codepipeline_webhook.webhook.url
#     secret       = var.WEBHOOK_TOKEN
#     content_type = "json"
#     insecure_ssl = false
#   }

#   events = ["push"]
# }