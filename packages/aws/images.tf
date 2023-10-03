resource "aws_s3_bucket" "images" {
  bucket = "tuango-image-1"
}

resource "aws_s3_bucket_acl" "images" {
  bucket = aws_s3_bucket.images.id
  acl    = "private"
}

resource "aws_s3_bucket_ownership_controls" "images" {
  bucket = aws_s3_bucket.images.id
  rule {
    object_ownership = "ObjectWriter"
  }
}

resource "aws_cloudfront_origin_access_control" "images" {
  name                              = aws_s3_bucket.images.bucket_domain_name
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_s3_bucket_policy" "images" {
  bucket = aws_s3_bucket.images.id
  policy = jsonencode(
    {
      Version = "2008-10-17",
      Id      = "PolicyForCloudFrontPrivateContent",
      Statement = [
        {
          Sid    = "AllowCloudFrontServicePrincipal",
          Effect = "Allow",
          Principal = {
            Service = "cloudfront.amazonaws.com"
          },
          Action   = "s3:GetObject",
          Resource = "${aws_s3_bucket.images.arn}/*",
          Condition = {
            StringEquals = {
              "AWS:SourceArn" = aws_cloudfront_distribution.images.arn
            }
          }
        }
      ]
    }
  )
}

resource "aws_cloudfront_distribution" "images" {
  origin {
    domain_name              = aws_s3_bucket.images.bucket_regional_domain_name
    origin_id                = aws_s3_bucket.images.id
    origin_access_control_id = aws_cloudfront_origin_access_control.images.id
  }

  enabled         = true
  is_ipv6_enabled = true
  http_version    = "http2and3"

  aliases = ["images.xn--ndsp5rmr3blfh.com"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = aws_s3_bucket.images.id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 31536000
    default_ttl            = 31536000
    max_ttl                = 31536000
    compress               = true
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  price_class = "PriceClass_All"

  viewer_certificate {
    acm_certificate_arn            = aws_acm_certificate.images.arn
    cloudfront_default_certificate = false
    minimum_protocol_version       = "TLSv1.2_2021"
    ssl_support_method             = "sni-only"
  }
}

resource "aws_acm_certificate" "images" {
  provider = aws.us_east_1

  domain_name       = "images.xn--ndsp5rmr3blfh.com"
  validation_method = "EMAIL"

  validation_option {
    domain_name       = "images.xn--ndsp5rmr3blfh.com"
    validation_domain = "xn--ndsp5rmr3blfh.com"
  }
}

resource "cloudflare_record" "example" {
  zone_id = var.CLOUDFLARE_ZONE_ID
  name    = "images"
  type    = "CNAME"
  value   = aws_cloudfront_distribution.images.domain_name
  proxied = true
}
