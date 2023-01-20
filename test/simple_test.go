package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	terraformOptions := &terraform.Options{
		 TerraformDir: "../terraform",
		 VarFiles: []string{
			"../envs/dev/variables.tfvars",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// output := terraform.Output(t, terraformOptions, "hello_world")
	// assert.Equal(t, "more, more, more", output)
	t.Run("TestingNumber1", func(t *testing.T) {
		test1_output := terraform.Output(t, terraformOptions, "testSample")
		assert.Equal(t, "It's working", test1_output)
	})

	t.Run("Resource group location", func(t *testing.T) {
		test2_output := terraform.Output(t, terraformOptions, "rg_location")
		assert.Equal(t, "westeurope", test2_output)
	})
}
