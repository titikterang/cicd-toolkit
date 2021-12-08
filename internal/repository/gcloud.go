package repository

import "github.com/ujunglangit-id/cicd-toolkit/internal/models/types"

type GCloudAPIRepository interface {
	GetInstanceStatus(param types.GCloudInstanceParam) (data types.GCloudInstanceProperties, err error)
	StartInstance(param types.GCloudInstanceProperties) (err error)
	StopInstance(param types.GCloudInstanceProperties) (err error)
	CreateNewInstance(param types.GCloudInstanceProperties) (err error)
}
