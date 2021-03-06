###########
#   ECS   #
###########
resource "aws_ecs_cluster" "automatic-trading-system-ecs-cluster" {
  name = "automatic-trading-system-ecs-cluster"
}

resource "aws_ecs_task_definition" "automatic-trading-system-api-task" {
  family                   = "automatic-trading-system-api-task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = file("./tasks/automatic-trading-system_api_definition.json")
  execution_role_arn       = module.ecs_task_execution_role.iam_role_arn
}
resource "aws_ecs_service" "automatic-trading-system-api-ecs-service" {
  name             = "automatic-trading-system-api-ecs-service"
  cluster          = aws_ecs_cluster.automatic-trading-system-ecs-cluster.arn
  task_definition  = aws_ecs_task_definition.automatic-trading-system-api-task.arn
  desired_count    = 1
  launch_type      = "FARGATE"
  platform_version = "1.4.0"

  network_configuration {
    assign_public_ip = true
    security_groups = [
      aws_security_group.automatic-trading-system-ecs-sg.id
    ]
    subnets = [
      aws_subnet.automatic-trading-system-public-1a.id,
      aws_subnet.automatic-trading-system-public-1c.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.automatic-trading-system-nlb-api-tg.arn
    container_name   = "automatic-trading-system"
    container_port   = "8000"
  }
}

##########
#  権限   #
##########
data "aws_iam_policy" "ecs_task_execution_role_policy" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

data "aws_iam_policy_document" "ecs_task_execution" {
  source_json = data.aws_iam_policy.ecs_task_execution_role_policy.policy

  statement {
    effect    = "Allow"
    actions   = ["s3:GetObject", "s3:GetBucketLocation"]
    resources = ["*"]
  }
}

module "ecs_task_execution_role" {
  source     = "./modules/iam_role"
  name       = "ecs-task-execution"
  identifier = "ecs-tasks.amazonaws.com"
  policy     = data.aws_iam_policy_document.ecs_task_execution.json
}
