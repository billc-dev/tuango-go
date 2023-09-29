resource "aws_s3_bucket" "images" {
  bucket = "tuango-image-1"
}

resource "aws_s3_bucket_acl" "images" {
  bucket = aws_s3_bucket.images.id
  acl    = "private"
}

resource "aws_cloudfront_origin_access_control" "images" {
  name                              = aws_s3_bucket.images.bucket_domain_name
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "never"
  signing_protocol                  = "sigv4"
}

# data "aws_iam_policy_document" "s3_policy" {
#   statement {
#     actions   = ["s3:GetObject"]
#     resources = ["${aws_s3_bucket.images.arn}/*"]

#     principals {
#       type        = "AWS"
#       identifiers = [aws_cloudfront_origin_access_identity.images.iam_arn]
#     }
#   }
# }

resource "aws_s3_bucket_policy" "images" {
  bucket = aws_s3_bucket.images.id
  policy = jsonencode(
    {
      Version = "2008-10-17",
      Id      = "PolicyForCloudFrontPrivateContent",
      Statement = [
        {
          Sid    = "1",
          Effect = "Allow",
          Principal = {
            AWS = "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity ${aws_cloudfront_distribution.images.id}"
          },
          Action   = "s3:GetObject",
          Resource = "${aws_s3_bucket.images.arn}/*"
        }
      ]
    }
  )
}

# resource "aws_s3_bucket_public_access_block" "images" {
#   bucket = aws_s3_bucket.images.id

#   block_public_acls       = true
#   block_public_policy     = true
#   //ignore_public_acls      = true
#   //restrict_public_buckets = true
# }

# resource "aws_cloudfront_origin_access_identity" "images" {
# }

resource "aws_cloudfront_distribution" "images" {
  origin {
    domain_name              = aws_s3_bucket.images.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.images.id
    origin_id                = aws_s3_bucket.images.id

    # s3_origin_config {
    #   origin_access_identity = aws_cloudfront_origin_access_identity.images.cloudfront_access_identity_path
    # }
  }

  enabled         = true
  is_ipv6_enabled = true
  http_version    = "http3"

  # aliases = ["mysite.example.com", "yoursite.example.com"]

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
    cloudfront_default_certificate = true
  }
}

