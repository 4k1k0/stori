terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "us-west-2"
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}

variable "AWS_ACCESS_KEY_ID" {
  description = "AWS access key"
  type = string
}

variable "AWS_SECRET_ACCESS_KEY" {
  description = "AWS secret access key"
  type = string
}

variable "aws_region" {
  default = "us-west-2"
}

resource "aws_s3_bucket" "stori-interview-backend-golang-rp-2023" {
  bucket = "stori-interview-backend-golang-rp-2023"
}

resource "aws_lambda_function" "rp-test-function" {
  function_name = "rp-test-function"
  handler      = "main"
  runtime      = "provided.al2"
  role         = aws_iam_role.rp-test_lambda_role.arn
  filename     = "bootstrap.zip"
  architectures = ["arm64"]

  environment {
    variables = {
      STORI_EMAIL = aws_ssm_parameter.stori_email.value
      STORI_POSTGRES_DBNAME = aws_ssm_parameter.stori_postgres_dbname.value
      STORI_POSTGRES_HOST = aws_ssm_parameter.stori_postgres_host.value
      STORI_POSTGRES_PASSWORD = aws_ssm_parameter.stori_postgres_password.value
      STORI_POSTGRES_PORT = aws_ssm_parameter.stori_postgres_port.value
      STORI_POSTGRES_USER = aws_ssm_parameter.stori_postgres_user.value
      STORI_SENDER_API_KEY = aws_ssm_parameter.stori_sender_api_key.value
    }
  }
}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rp-test-function.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.stori-interview-backend-golang-rp-2023.arn
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.stori-interview-backend-golang-rp-2023.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.rp-test-function.arn
    events = ["s3:ObjectCreated:*"]
    filter_suffix = ".csv"
  }
  depends_on = [aws_lambda_permission.allow_bucket]
}

resource "aws_iam_role" "rp-test_lambda_role" {
  name = "rp-test_lambda_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com",
        },
      },
    ],
  })
}

resource "aws_iam_role_policy" "lambda_execution_policy" {
  name   = "lambda_execution_policy"
  role   = aws_iam_role.rp-test_lambda_role.name

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "s3:GetObject",
          "s3:ListBucket",
        ],
        Resource = [
          aws_s3_bucket.stori-interview-backend-golang-rp-2023.arn,
          "${aws_s3_bucket.stori-interview-backend-golang-rp-2023.arn}/*",
        ],
      },
      {
        Effect = "Allow",
        Action = [
          "ssm:GetParameter",
          "ssm:GetParameters",
          "ssm:GetParametersByPath",
        ],
        Resource = "*",
      },
    ],
  })
}

resource "aws_cloudwatch_log_group" "lambda_log_group" {
  name = "/aws/lambda/${aws_lambda_function.rp-test-function.function_name}"
}

resource "aws_lambda_permission" "lambda_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rp-test-function.function_name
  principal     = "logs.amazonaws.com"
  source_arn    = aws_cloudwatch_log_group.lambda_log_group.arn
}

resource "aws_ssm_parameter" "stori_email" {
  name  = "/lambda/rp-test-function/STORI_EMAIL"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_postgres_dbname" {
  name  = "/lambda/rp-test-function/STORI_POSTGRES_DBNAME"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_postgres_host" {
  name  = "/lambda/rp-test-function/STORI_POSTGRES_HOST"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_postgres_password" {
  name  = "/lambda/rp-test-function/STORI_POSTGRES_PASSWORD"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_postgres_port" {
  name  = "/lambda/rp-test-function/STORI_POSTGRES_PORT"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_postgres_user" {
  name  = "/lambda/rp-test-function/STORI_POSTGRES_USER"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}

resource "aws_ssm_parameter" "stori_sender_api_key" {
  name  = "/lambda/rp-test-function/STORI_SENDER_API_KEY"
  type  = "String"
  value = "YourFooValue"
  overwrite = true
}
