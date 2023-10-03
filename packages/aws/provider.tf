terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.17.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.15.0"
    }
  }

  required_version = ">= 1.4.2"

  cloud {
    organization = "tuango"
    workspaces {
      name = "tuango"
    }
  }
}

provider "aws" {
  region     = "ap-east-1"
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}

provider "aws" {
  alias      = "us_east_1"
  region     = "us-east-1"
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}

provider "cloudflare" {
  api_token = var.CLOUDFLARE_API_TOKEN
}
