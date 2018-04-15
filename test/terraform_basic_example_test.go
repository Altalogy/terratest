package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TerraformBasicExampleTest(t *testing.T) {
	t.Parallel()

	expectedText := "foo"

	terraformOptions := terraform.Options {
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-basic-example",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]string {
			"example": expectedText,
		},
	}

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.Apply(t, terraformOptions)

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	actualText := terraform.Output(t, terraformOptions, "example")

	// Verify we're getting back the variable we expect
	assert.Equal(t, expectedText, actualText)
}