package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	t.Parallel()

	// tests running in parallel
	min := 12
	max := 30
	passwordLength := random.Random(min, max)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/password",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"password_length": passwordLength,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	randomPasswordResult := terraform.Output(t, terraformOptions, "random_password")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, passwordLength, len(randomPasswordResult))
}
