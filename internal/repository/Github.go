package repository

import "github.com/ujunglangit-id/cicd-toolkit/internal/models/types"

type GithubAPIRepository interface {
	GetPRInfo(repo string, prID int64) (data types.GithubPRInfoData, err error)
	MergePR(repo string, prID int64, squash bool) (respData types.MergeResponseData, err error)
	GetPRReviewInfo(repo string, prID int64) (data []types.GithubPRReviewData, err error)
}
