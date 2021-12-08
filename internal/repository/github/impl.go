package github

import (
	"bytes"
	"encoding/json"
	circuit "github.com/eapache/go-resiliency/breaker"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/types"
	"io/ioutil"
	"net/http"
	"time"
)

type GithubAPI struct {
	Config         *core.Config
	CircuitBreaker *circuit.Breaker
	Client         *http.Client
}

func New(cfg *core.Config, cl *http.Client) *GithubAPI {
	return &GithubAPI{
		Config:         cfg,
		CircuitBreaker: circuit.New(types.BreakerErrorThreshold, types.BreakerSuccessThreshold, types.BreakerTimeout*time.Second),
		Client:         cl,
	}
}

func (g *GithubAPI) GetPRReviewInfo(repo string, prID int64) (respData []types.GithubPRReviewData, err error) {
	var (
		req *http.Request
	)
	req, err = http.NewRequest("GET", core.GetPRReviewURL(g.Config.Github.Host, repo, prID), nil)
	if err != nil {
		return
	}

	req.SetBasicAuth(g.Config.Github.User, g.Config.Github.Token)
	resp, err := g.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}

	return
}

func (g *GithubAPI) GetPRInfo(repo string, prID int64) (respData types.GithubPRInfoData, err error) {
	var (
		req *http.Request
	)

	req, err = http.NewRequest("GET", core.GetPRInfoURL(g.Config.Github.Host, repo, prID), nil)
	if err != nil {
		return
	}

	req.SetBasicAuth(g.Config.Github.User, g.Config.Github.Token)
	resp, err := g.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}

	return
}

func (g *GithubAPI) MergePR(repo string, prID int64, squash bool) (respData types.MergeResponseData, err error) {
	var (
		req     *http.Request
		payload []byte
	)
	payloadData := types.MergePayload{
		CommitTitle: "Merge PR by debio-toolkit",
		MergeMethod: "merge",
	}
	if squash {
		payloadData.MergeMethod = "squash"
	}
	payload, err = json.Marshal(payloadData)

	if err != nil {
		return
	}

	req, err = http.NewRequest("PUT", core.GetPRMergeURL(g.Config.Github.Host, repo, prID), bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	req.SetBasicAuth(g.Config.Github.User, g.Config.Github.Token)
	resp, err := g.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return
	}

	return
}
