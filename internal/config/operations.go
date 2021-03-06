package config

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
)

func getOperations(dev adapter.Operations) adapter.Operations {
	versions, _ := getLatestReleaseNames(3)

	dev[TraefikMeshOperation] = &adapter.Operation{
		Type:        int32(meshes.OpCategory_INSTALL),
		Description: "Traefik Mesh",
		Versions:    versions,
		Templates: []adapter.Template{
			"templates/traefik-mesh.yaml",
		},
		AdditionalProperties: map[string]string{},
	}

	return dev
}
