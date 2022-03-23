package gridscale

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	builderT "github.com/hashicorp/packer-plugin-sdk/acctest"
)

const templateName = "ubuntu20.04-test-acc"

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
source "gridscale" "basic_test" {
	base_template_uuid = "fd65f8ce-e2c6-40af-8fc3-92efa0d4eecb"
	hostname           = "test-hostname"
	ssh_password           = "testPassword"
	server_cores       = 2
	server_memory      = 4
	ssh_username       = "root"
	storage_capacity   = 10
	template_name      = "%s"
  }
  build {
	sources = ["source.gridscale.basic_test"]
  }
`
