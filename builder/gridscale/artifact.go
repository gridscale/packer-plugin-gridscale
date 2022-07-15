package gridscale

import (
	"context"
	"fmt"
	"log"

	"github.com/gridscale/gsclient-go/v3"
	registryimage "github.com/hashicorp/packer-plugin-sdk/packer/registry/image"
)

type Artifact struct {
	// The name of the template
	TemplateName string

	// The UUID of the template
	TemplateUUID string

	// Location name.
	LocationName string

	// Location UUID.
	LocationUUID string

	// The client for making API calls
	Client gsclient.TemplateOperator

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Files() []string {
	// No files with DigitalOcean
	return nil
}

func (a *Artifact) Id() string {
	return a.TemplateUUID
}

func (a *Artifact) String() string {
	return fmt.Sprintf("A template was created: '%v' (ID: %v)", a.TemplateName, a.TemplateUUID)
}

func (a *Artifact) State(name string) interface{} {
	if name == registryimage.ArtifactStateURI {
		img, err := registryimage.FromArtifact(a,
			registryimage.WithID(a.TemplateName),
			registryimage.WithProvider("gridscale"),
			registryimage.WithSourceID(a.TemplateUUID),
			registryimage.WithRegion(a.LocationName),
		)
		if err != nil {
			log.Printf("[DEBUG] error encountered when creating a registry image %v", err)
			return nil
		}
		return img
	}
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	log.Printf("Destroying template: %s (%s)", a.TemplateName, a.TemplateUUID)
	err := a.Client.DeleteTemplate(context.Background(), a.TemplateUUID)
	return err
}
