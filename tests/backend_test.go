package backend_test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BackendCreated(t *testing.T) {
	// get output that backend mwas created
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/backend",
		// Vars: map[string]interface{}{
		// 	"aws_region":    awsRegion,
		// 	"instance_name": instanceName,
		// 	"instance_text": instanceText,
		// 	"instance_type": instanceType,
		//    },
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "bucket")
	assert.Equal(t, "uni-form-s3", output)
}

// func TestS3BackendHoldsTFState() {
// 	// validate tf state saved to s3 backend .... is this test in the right spot / module?
// }

// func TestCreateDynamoTable(t *testing.T) {

// }
