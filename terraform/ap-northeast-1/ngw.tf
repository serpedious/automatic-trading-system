##############
# NATGateway #
##############
# resource "aws_nat_gateway" "automatic-trading-system-ngw" {
#   allocation_id = aws_eip.automatic-trading-system-ngw-eip.id
#   subnet_id     = aws_subnet.automatic-trading-system-public-1a.id
#   depends_on    = [aws_internet_gateway.automatic-trading-system-igw]

#   tags = {
#     Name = "automatic-trading-system-ngw"
#   }
# }
