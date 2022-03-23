package gridscale

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	builderT "github.com/hashicorp/packer-plugin-sdk/acctest"
)

const templateName = "ubuntu-test-acc"

func TestBuilderAcc_basic(t *testing.T) {
	builderT.TestPlugin(t, &builderT.PluginTestCase{
		Name:     "gridscale_builder_basic_test",
		Setup:    testAccPreCheck,
		Template: fmt.Sprintf(testBuilderAccBasic, templateName),
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}
			return nil
		},
	})
}

func testAccPreCheck() error {
	if v := os.Getenv("GRIDSCALE_UUID"); v == "" {
		return errors.New("GRIDSCALE_UUID must be set for acceptance tests")
	}
	if v := os.Getenv("GRIDSCALE_TOKEN"); v == "" {
		return errors.New("GRIDSCALE_TOKEN must be set for acceptance tests")
	}
	return nil
}

const testBuilderAccBasic = `
{
	"builders": [
    {
	  "type": "gridscale",
      "template_name": "%s",
      "password": "testPassword",
      "hostname": "test-hostname",
      "ssh_username": "root",
      "server_memory": 4,
      "server_cores": 2,
      "storage_capacity": 10,
      "base_template_uuid": "4db64bfc-9fb2-4976-80b5-94ff43b1233a"
    }
  ]
}
`
