###################
# InternetGateway #
###################
resource "aws_internet_gateway" "automatic-trading-system-igw" {
  vpc_id = aws_vpc.automatic-trading-system.id

  tags = {
    Name = "automatic-trading-system-igw"
  }
}
