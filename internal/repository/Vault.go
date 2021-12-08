package repository

import "github.com/ujunglangit-id/cicd-toolkit/internal/models/types"

type VaultAPIRepository interface {
	GetVaultSecret(path string, useV1 bool) (respData types.VaultResponseData, err error)
	GetVaultRawSecret(path string) (respData types.RawFileVaultResponseData, err error)
}
