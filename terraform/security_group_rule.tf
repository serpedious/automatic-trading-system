#########################
# SecurityGroupRule RDS #
#########################
# resource "aws_security_group_rule" "automatic-trading-system-rds-sg-rule1" {
#   description       = "automatic-trading-system rds sg rule"
#   type              = "ingress"
#   from_port         = 5432
#   to_port           = 5432
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0", "0.0.0.0/16"]
#   security_group_id = aws_security_group.automatic-trading-system-rds-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-rds-sg-rule2" {
#   description       = "automatic-trading-system rds sg rule2"
#   type              = "egress"
#   from_port         = 0
#   to_port           = 0
#   protocol          = "-1"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-rds-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-rds-sg-rule3" {
#   description       = "automatic-trading-system rds sg rule3"
#   type              = "egress"
#   from_port         = 0
#   to_port           = 0
#   protocol          = "-1"
#   ipv6_cidr_blocks  = ["::/0"]
#   security_group_id = aws_security_group.automatic-trading-system-rds-sg.id
# }


# #########################
# # SecurityGroupRule NLB #
# #########################
# resource "aws_security_group_rule" "automatic-trading-system-alb-sg-rule1" {
#   description       = "automatic-trading-system-alb-sg-rule1"
#   type              = "ingress"
#   from_port         = 80
#   to_port           = 80
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-alb-sg-rule2" {
#   description       = "automatic-trading-system-alb-sg-rule2"
#   type              = "ingress"
#   from_port         = 8
#   to_port           = 0
#   protocol          = "icmp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-alb-sg-rule3" {
#   description       = "automatic-trading-system-alb-sg-rule3"
#   type              = "ingress"
#   from_port         = 443
#   to_port           = 443
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-alb-sg-rule4" {
#   description       = "automatic-trading-system-alb-sg-rule4"
#   type              = "egress"
#   from_port         = 0
#   to_port           = 0
#   protocol          = "-1"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-alb-sg-rule5" {
#   description       = "automatic-trading-system-alb-sg-rule5"
#   type              = "ingress"
#   from_port         = 8000
#   to_port           = 8000
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
# }

# #########################
# # SecurityGroupRule ECS #
# #########################
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule1" {
#   description       = "automatic-trading-system-ecs-sg-rule1"
#   type              = "ingress"
#   from_port         = 8
#   to_port           = 0
#   protocol          = "icmp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-ecs-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule2" {
#   description       = "automatic-trading-system-ecs-sg-rule2"
#   type              = "egress"
#   from_port         = 0
#   to_port           = 0
#   protocol          = "-1"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-ecs-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule3" {
#   description       = "automatic-trading-system-ecs-sg-rule3"
#   type              = "ingress"
#   from_port         = 80
#   to_port           = 80
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-ecs-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule4" {
#   description       = "automatic-trading-system-ecs-sg-rule4"
#   type              = "ingress"
#   from_port         = 443
#   to_port           = 443
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-ecs-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule5" {
#   description              = "automatic-trading-system-ecs-sg-rule5"
#   type                     = "ingress"
#   from_port                = 0
#   to_port                  = 0
#   protocol                 = "tcp"
#   source_security_group_id = aws_security_group.automatic-trading-system-alb-sg.id
#   security_group_id        = aws_security_group.automatic-trading-system-ecs-sg.id
# }
# resource "aws_security_group_rule" "automatic-trading-system-ecs-sg-rule6" {
#   description       = "automatic-trading-system-ecs-sg-rule6"
#   type              = "ingress"
#   from_port         = 8000
#   to_port           = 8000
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   security_group_id = aws_security_group.automatic-trading-system-ecs-sg.id
# }
