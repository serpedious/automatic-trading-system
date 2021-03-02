###############
#    eip      #
###############
resource "aws_eip" "automatic-trading-system-eip" {
  vpc      = true
  depends_on = [aws_lb.automatic-trading-system-nlb]

  tags = {
    Name = "automatic-trading-system-eip"
  }
}