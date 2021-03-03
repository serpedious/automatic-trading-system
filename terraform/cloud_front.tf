locals {
  s3_origin_id = "automatic-trading-system-s3-bucket-for-nuxt"
}

resource "aws_cloudfront_distribution" "automatic-trading-system" {
  origin {
    domain_name = "${aws_s3_bucket.automatic-trading-system-s3-bucket.bucket_regional_domain_name}"
    origin_id   = "${aws_s3_bucket.automatic-trading-system-s3-bucket.id}"

    s3_origin_config {
      origin_access_identity = "${aws_cloudfront_origin_access_identity.origin_access_identity.cloudfront_access_identity_path}"
    }
  }

  enabled         = true
  is_ipv6_enabled = false
  comment         = "comment"

  logging_config {
    include_cookies = false
    bucket          = "${aws_s3_bucket.automatic-trading-system-s3-bucket.bucket_domain_name}"
    prefix          = "prefix"
  }

  aliases = ["serpedious.link"]

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "${aws_s3_bucket.automatic-trading-system-s3-bucket.id}"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_200"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    Environment = "dev"
  }

  viewer_certificate {
    acm_certificate_arn = "arn:aws:acm:us-east-1:606124585607:certificate/54581b7b-f960-4e4c-a203-059fdb30ce6a"
    ssl_support_method  = "vip"
  }
}

resource "aws_cloudfront_origin_access_identity" "origin_access_identity" {
  comment = "This is identity for cloud front to s3"
}

data "aws_iam_policy_document" "s3_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.automatic-trading-system-s3-bucket.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn]
    }
  }
}
