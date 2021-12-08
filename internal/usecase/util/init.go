package util

import (
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository/util"
	"github.com/ujunglangit-id/cicd-toolkit/internal/usecase"
	"github.com/ujunglangit-id/cicd-toolkit/internal/usecase/toolkit"
)

type CaseWrapper struct {
	ToolkitCase usecase.ToolkitCase
}

func New(cfg *core.Config, repo *util.RepoWrapper) *CaseWrapper {
	return &CaseWrapper{
		ToolkitCase: toolkit.New(cfg, repo),
	}
}
