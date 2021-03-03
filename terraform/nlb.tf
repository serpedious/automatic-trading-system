#########
#  NLB  #
#########
resource "aws_lb" "automatic-trading-system-nlb" {
  name               = "automatic-trading-system-nlb"
  internal           = false
  load_balancer_type = "network"
  subnets            = aws_subnet.automatic-trading-system-public-1a.*.id

  enable_deletion_protection = true

  tags = {
    Environment = "production"
  }
}

############
# Listener #
############
resource "aws_lb_listener" "automatic-trading-system-https-listener" {
  load_balancer_arn = aws_lb.automatic-trading-system-nlb.arn
  port              = "80"
  protocol          = "TCP"
  # certificate_arn   = aws_acm_certificate.automatic-trading-system-acm.arn

  default_action {
    target_group_arn = aws_lb_target_group.automatic-trading-system-nlb-api-tg.arn
    type             = "forward"
  }
}
# resource "aws_lb_listener" "automatic-trading-system-http-listener" {
#   load_balancer_arn = aws_lb.automatic-trading-system-nlb.arn
#   port              = "80"
#   protocol          = "HTTP"

#   default_action {
#     type = "redirect"

#     redirect {
#       port        = "443"
#       protocol    = "HTTPS"
#       status_code = "HTTP_301"
#     }
#   }
# }
# resource "aws_lb_listener" "automatic-trading-system-https-listener" {
#   load_balancer_arn = aws_lb.automatic-trading-system-nlb.arn
#   port              = "443"
#   protocol          = "HTTPS"
#   certificate_arn   = aws_acm_certificate.automatic-trading-system-acm.arn

#   default_action {
#     target_group_arn = aws_lb_target_group.automatic-trading-system-nlb-api-tg.arn
#     type             = "forward"
#   }
# }
# resource "aws_lb_listener" "automatic-trading-system-api-listener" {
#   load_balancer_arn = aws_lb.automatic-trading-system-nlb.arn
#   port              = "8000"
#   protocol          = "HTTPS"
#   certificate_arn   = aws_acm_certificate.automatic-trading-system-acm.arn

#   default_action {
#     target_group_arn = aws_lb_target_group.automatic-trading-system-nlb-api-tg.arn
#     type             = "forward"
#   }
# }

###############
# TargetGroup #
###############
resource "aws_lb_target_group" "automatic-trading-system-nlb-api-tg" {
  name        = "automatic-nlb-api-tg"
  target_type = "ip"
  vpc_id      = aws_vpc.automatic-trading-system.id
  port        = 80
  protocol    = "TCP"

  # health_check {
  #   enabled             = true
  #   interval            = 10
  #   path                = "/private"
  #   port                = 8000
  #   protocol            = "HTTP"
  #   matcher             = 200
  #   timeout             = 50
  #   healthy_threshold   = 5
  #   unhealthy_threshold = 2
  # }
}
