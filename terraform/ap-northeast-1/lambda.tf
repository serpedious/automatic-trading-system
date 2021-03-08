resource "aws_iam_role" "cloud_front_for_lambda" {
  name = "cloud_front_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "cache_cloud_front" {
  filename      = "./tasks/lambda_function_payload.zip"
  function_name = "cache_cloud_front"
  role          = aws_iam_role.cloud_front_for_lambda.arn
  handler       = "reload_cache_of_cloudfront.lambda_handler"

  # The filebase64sha256() function is availabl e in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the base64sha256() function and the file() function:
  # source_code_hash = "${base64sha256(file("lambda_function_payload.zip"))}"
  # source_code_hash = filebase64sha256("./tasks/lambda_function_payload.zip")

  runtime = "python2.7"

  # environment {
  #   variables = {
  #     foo = "bar"
  #   }
  # }
}