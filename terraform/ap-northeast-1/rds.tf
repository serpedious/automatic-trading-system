#########
#  RDS  #
#########
resource "aws_db_instance" "automatic-system-rds" {
  identifier                   = "automatic-system-rds"
  port                         = 5432
  allocated_storage            = 20
  instance_class               = "db.t2.micro"
  engine                       = "postgres"
  engine_version               = "12.5"
  storage_type                 = "gp2"
  name                         = "automatic_db"
  username                     = "serpedious"
  password                     = "serpedious"
  backup_retention_period      = 7
  copy_tags_to_snapshot        = true
  max_allocated_storage        = 1000
  skip_final_snapshot          = true
  performance_insights_enabled = true
  vpc_security_group_ids       = [aws_security_group.automatic-trading-system-rds-sg.id]
  db_subnet_group_name         = aws_db_subnet_group.automatic-trading-system-rds-subnet-group.name
  # enabled_cloudwatch_logs_exports = [
  #   "audit",
  #   "error",
  #   "general",
  #   "slowquery",
  # ]

  lifecycle {
    prevent_destroy = false
  }
}