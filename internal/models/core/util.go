package core

import "fmt"

func GetPRInfoURL(host, repo string, prID int64) (url string) {
	return fmt.Sprintf("%s/repos/%s/pulls/%d", host, repo, prID)
}

func GetPRReviewURL(host, repo string, prID int64) (url string) {
	return fmt.Sprintf("%s/repos/%s/pulls/%d/reviews", host, repo, prID)
}

func GetPRCommitURL(host, repo string, prID int64) (url string) {
	return fmt.Sprintf("%s/repos/%s/pulls/%d/commits", host, repo, prID)
}

func GetPRMergeURL(host, repo string, prID int64) (url string) {
	return fmt.Sprintf("%s/repos/%s/pulls/%d/merge", host, repo, prID)
}
