######################
#    APIGateway      #
######################

resource "aws_apigatewayv2_api" "automatic-trading-system-api-gateway" {
  name          = "automatic-trading-system-http-api"
  protocol_type = "HTTP"
}

# resource "aws_apigatewayv2_api_mapping" "automatic-trading-system-api-gateway" {
#   api_id      = aws_apigatewayv2_api.automatic-trading-system-api-gateway.id
#   domain_name = aws_apigatewayv2_domain_name.automatic-trading-system-api-gateway.id
#   stage       = aws_apigatewayv2_stage.automatic-trading-system-api-gateway.id
# }

# resource "aws_apigatewayv2_integration" "automatic-trading-system-api-gateway" {
#   api_id           = aws_apigatewayv2_api.automatic-trading-system-api-gateway.id
#   integration_type = "AWS"

#   connection_type           = "INTERNET"
#   content_handling_strategy = "CONVERT_TO_TEXT"
#   description               = "Lambda example"
#   integration_method        = "POST"
#   integration_uri           = aws_lambda_function.example.invoke_arn
#   passthrough_behavior      = "WHEN_NO_MATCH"
# }

# resource "aws_apigatewayv2_stage" "automatic-trading-system-api-gateway" {
#   api_id = aws_apigatewayv2_api.automatic-trading-system-api-gateway.id
#   name   = "automatic-trading-system-api-gateway"
# }

# resource "aws_apigatewayv2_route" "automatic-trading-system-api-gateway" {
#   api_id    = aws_apigatewayv2_api.automatic-trading-system-api-gateway.id
#   route_key = "$default"
# }

# resource "aws_api_gateway_domain_name" "automatic-trading-system-api-gateway" {
#   certificate_arn = aws_acm_certificate_validation.automatic-trading-system-api-gateway.certificate_arn
#   domain_name     = "api.serpedious.link"
# }