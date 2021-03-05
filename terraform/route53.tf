##########
#Route53 #
##########
resource "aws_route53_zone" "automatic-trading-system-host-zone" {
  name    = "serpedious.link"
  comment = "serpedious.link host zone"
}

resource "aws_route53_record" "automatic-trading-system-host-zone-record" {
  zone_id = aws_route53_zone.automatic-trading-system-host-zone.zone_id
  name    = aws_route53_zone.automatic-trading-system-host-zone.name
  type    = "A"

  alias {
    evaluate_target_health = false
    name                   = "d111pd78fvr7hz.cloudfront.net"
    zone_id                = "Z2FDTNDATAQYW2"
  }
}
