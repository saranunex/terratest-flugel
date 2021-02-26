output "bucket_id" {
  value = aws_s3_bucket.newbucket.bucket
}

# To be used in the future, validate the content of the s3 objects
output "object_key" {
  value = ["text1.txt", "text2.txt"]
}
