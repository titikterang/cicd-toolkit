package util

import (
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/types"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository/github"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository/vault"
	"net/http"
	"time"
)

type RepoWrapper struct {
	GithubAPI repository.GithubAPIRepository
	VaultAPI  repository.VaultAPIRepository
}

func New(cfg *core.Config) *RepoWrapper {
	vaultPool := &http.Client{
		Timeout: types.PoolClientTimeoutSeconds * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        types.PoolTransportMaxIdleConns,
			MaxIdleConnsPerHost: types.PoolTransportMaxIdleConnsPerHost,
			IdleConnTimeout:     types.PoolTransportIdleConnTimeoutSeconds * time.Second,
		},
	}

	githubPool := &http.Client{
		Timeout: types.PoolClientTimeoutSeconds * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        types.PoolTransportMaxIdleConns,
			MaxIdleConnsPerHost: types.PoolTransportMaxIdleConnsPerHost,
			IdleConnTimeout:     types.PoolTransportIdleConnTimeoutSeconds * time.Second,
		},
	}

	return &RepoWrapper{
		GithubAPI: github.New(cfg, githubPool),
		VaultAPI:  vault.New(cfg, vaultPool),
	}
}
