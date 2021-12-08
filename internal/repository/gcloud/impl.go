package gcloud

import (
	compute "cloud.google.com/go/compute/apiv1"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/types"
)

type GCloudAPI struct {
	Config *core.Config
	Client *compute.InstancesClient
}

func (g *GCloudAPI) GetInstanceStatus(param types.GCloudInstanceParam) (data types.GCloudInstanceProperties, err error) {
	return
}

func (g *GCloudAPI) StartInstance(param types.GCloudInstanceProperties) (err error) {
	return
}

func (g *GCloudAPI) StopInstance(param types.GCloudInstanceProperties) (err error) {
	return
}

func (g *GCloudAPI) CreateNewInstance(param types.GCloudInstanceProperties) (err error) {
	return
}
