##############
# RouteTable #
##############
resource "aws_route_table" "automatic-trading-system-public-rtb" {
  vpc_id = aws_vpc.automatic-trading-system.id

  route {
    gateway_id = aws_internet_gateway.automatic-trading-system-igw.id
    cidr_block = "0.0.0.0/0"
  }

  tags = {
    Name = "automatic-trading-system-public-route"
  }
}

###############
# association #
###############
resource "aws_route_table_association" "automatic-trading-system-public-rtb-1a" {
  subnet_id      = aws_subnet.automatic-trading-system-public-1a.id
  route_table_id = aws_route_table.automatic-trading-system-public-rtb.id
}
resource "aws_route_table_association" "automatic-trading-system-public-rtb-1c" {
  subnet_id      = aws_subnet.automatic-trading-system-public-1c.id
  route_table_id = aws_route_table.automatic-trading-system-public-rtb.id
}

