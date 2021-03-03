######################
#    APIGateway      #
######################

resource "aws_apigatewayv2_api" "automatic-trading-system-api-gateway" {
  name          = "automatic-trading-system-http-api"
  protocol_type = "HTTP"
}
