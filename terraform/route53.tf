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
    name                   = aws_lb.automatic-trading-system-nlb.dns_name
    zone_id                = aws_lb.automatic-trading-system-nlb.zone_id
    evaluate_target_health = true
  }
}
