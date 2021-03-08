##########
#  ACM   #
##########

##definition of SSL certificate
resource "aws_acm_certificate" "automatic-trading-system-acm" {
  domain_name               = aws_route53_record.automatic-trading-system-host-zone-record.name
  subject_alternative_names = []
  validation_method         = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}
##verification of SSL
resource "aws_route53_record" "automatic-trading-system-certificate" {
  name    = tolist(aws_acm_certificate.automatic-trading-system-acm.domain_validation_options)[0].resource_record_name
  type    = tolist(aws_acm_certificate.automatic-trading-system-acm.domain_validation_options)[0].resource_record_type
  records = [tolist(aws_acm_certificate.automatic-trading-system-acm.domain_validation_options)[0].resource_record_value]
  zone_id = aws_route53_zone.automatic-trading-system-host-zone.id
  ttl     = 60
}
##waiting verification of SSL
resource "aws_acm_certificate_validation" "automatic-trading-system-acm" {
  certificate_arn         = aws_acm_certificate.automatic-trading-system-acm.arn
  validation_record_fqdns = [aws_route53_record.automatic-trading-system-certificate.fqdn]
}
