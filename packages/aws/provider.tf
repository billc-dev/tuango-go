terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.17.0"
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
