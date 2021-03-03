##########
# Subnet #
##########
resource "aws_subnet" "automatic-trading-system-public-1a" {
  vpc_id                  = aws_vpc.automatic-trading-system.id
  cidr_block              = "10.0.0.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "automatic-trading-system-public-1a"
  }
}
resource "aws_subnet" "automatic-trading-system-public-1c" {
  vpc_id                  = aws_vpc.automatic-trading-system.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true

  tags = {
    Name = "automatic-trading-system-public-1c"
  }
}


###############
# SubnetGroup #
###############
resource "aws_db_subnet_group" "automatic-trading-system-rds-subnet-group" {
  name        = "automatic-trading-system-rds-subnet-group"
  description = "rds subnet for automatic-trading-system"
  subnet_ids  = [aws_subnet.automatic-trading-system-public-1a.id, aws_subnet.automatic-trading-system-public-1c.id]
}
