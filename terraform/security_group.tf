####################
# SecurityGroup RDS #
# ####################
# resource "aws_security_group" "automatic-trading-system-rds-sg" {
#   description = "RDS security group for automatic-trading-system"
#   name        = "automatic-trading-system-rds-sg"
#   vpc_id      = aws_vpc.automatic-trading-system.id
# }

#####################
# SecurityGroup NLB #
#####################
# resource "aws_security_group" "automatic-trading-system-nlb-sg" {
#   name        = "automatic-trading-system-alb-sg"
#   description = "ALB security group for automatic-trading-system"
#   vpc_id      = aws_vpc.automatic-trading-system.id
# }

#####################
# SecurityGroup ECS #
#####################
resource "aws_security_group" "automatic-trading-system-ecs-sg" {
  vpc_id      = aws_vpc.automatic-trading-system.id
  name        = "automatic-trading-system-ecs-sg"
  description = "ECS service security group for automatic-trading-system"
}

###########################
# SecurityGroup CloudFront #
###########################