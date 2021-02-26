terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

#Setting up provider and region 
provider "aws" {
  profile = "default"
  region  = "us-west-2"
}

# S3 Object names
variable "file_names" {
  description = "Create s3 objects with these names"
  type        = list(string)
  default     = ["text1.txt", "text2.txt"]
}

 # Generating timestampa and unique name for s3 bucket
locals {
  timestamp = timestamp()
  bucket_name = uuid()
}

# Creating AWS S3 Bucket 
resource "aws_s3_bucket" "newbucket" {
  bucket = local.bucket_name
  acl = "private"
}

# Creating AWS S3 Object
resource "aws_s3_bucket_object" "object" {
  count = length(var.file_names)
  bucket = aws_s3_bucket.newbucket.bucket
  key    = var.file_names[count.index]
  content = local.timestamp
  depends_on = [aws_s3_bucket.newbucket]
}