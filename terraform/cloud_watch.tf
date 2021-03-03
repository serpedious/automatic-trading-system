#################
# CloudWatchLog #
#################
resource "aws_cloudwatch_log_group" "automatic-trading-system-ecs-api" {
  name              = "/ecs/api"
  retention_in_days = 180
}
