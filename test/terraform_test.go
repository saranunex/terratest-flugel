package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// Testing Terraform using Terratest.
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	// Setup a  AWS region to test in.
	awsRegion := "us-west-2"

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	// To be used in the future, validate the content of the s3 objects
	//objectKey := terraform.Output(t, terraformOptions, "object_key")
	//aws.GetS3ObjectContents(t, awsRegion, bucketID, objectKey)

	// Checks if the given S3 bucket exists in the given region and fail the test if it does not.
	aws.AssertS3BucketExists(t, awsRegion, bucketID)
}
