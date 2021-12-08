package usecase

type ToolkitCase interface {
	ValidateApprovalStatus(repo string, prID int64) (err error)
	GetVaultSecret(path, outputName string, isEnv, isRawFile, useV1 bool) (err error)
	MergePR(repo string, prID int64, squash bool) (err error)
}
