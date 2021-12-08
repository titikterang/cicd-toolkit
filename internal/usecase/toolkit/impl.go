package toolkit

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/types"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository/util"
	"io/ioutil"
	"reflect"
	"strings"
)

type ToolkitDataCase struct {
	VaultRepo repository.VaultAPIRepository
	GitRepo   repository.GithubAPIRepository
	cfg       *core.Config
}

func New(cfg *core.Config, repo *util.RepoWrapper) *ToolkitDataCase {
	return &ToolkitDataCase{
		VaultRepo: repo.VaultAPI,
		GitRepo:   repo.GithubAPI,
		cfg:       cfg,
	}
}

func (c *ToolkitDataCase) ValidateApprovalStatus(repo string, prID int64) (err error) {
	var (
		data     []types.GithubPRReviewData
		approved int
	)

	log.Infof("[toolkit] validate PR approval status with prID %d from repo %s", prID, repo)
	data, err = c.GitRepo.GetPRReviewInfo(repo, prID)
	if err != nil {
		log.Errorf("[toolkit] GetPRReviewInfo failed, err:  %+v", err)
		return
	}

	if len(data) > 0 {
		for _, v := range data {
			if v.State == types.PR_APPROVED {
				approved++
			}
		}
	}

	if approved >= c.cfg.Github.ApprovalLimit {
		log.Infof("[toolkit] PR %d has been approved by %d reviewer", prID, approved)
	} else {
		return fmt.Errorf("[toolkit] at least 1 approval is required for PR %d", prID)
	}

	return
}

func (c *ToolkitDataCase) generateRawFile(secretPath, outputName string) (err error) {
	var (
		respData types.RawFileVaultResponseData
		fileName string
	)

	respData, err = c.VaultRepo.GetVaultRawSecret(secretPath)
	if err != nil {
		return
	}

	if respData.Data.Data.FileName != "" {
		fileName = respData.Data.Data.FileName
	} else {
		if outputName == "" {
			fMap := strings.Split(secretPath, "/")
			fileName = fMap[len(fMap)-1] + ".txt"
		} else {
			fileName = outputName
		}
	}

	log.Infof("[toolkit] generate raw text file to %s from %s", fileName, secretPath)
	err = ioutil.WriteFile(fileName, []byte(respData.Data.Data.FileData), 0644)
	if err != nil {
		log.Errorf("[toolkit] error generateRawFile %s, err : %+v", fileName, err)
	}
	return
}

func (c *ToolkitDataCase) GetVaultSecret(secretPath, outputName string, isEnv, isRawFile, useV1 bool) (err error) {
	var (
		respData types.VaultResponseData
		fileName string
		fileData []byte
	)

	if isRawFile {
		err = c.generateRawFile(secretPath, outputName)
		return
	}

	if outputName == "" {
		fMap := strings.Split(secretPath, "/")
		fileName = fMap[len(fMap)-1] + ".json"
	} else {
		fileName = outputName
	}

	respData, err = c.VaultRepo.GetVaultSecret(secretPath, useV1)
	if err != nil {
		return
	}

	if isEnv {
		envData := ""
		log.Infof("[toolkit] generate env file from %s", secretPath)
		fileName = ".env"
		for k, v := range respData.Data.(map[string]interface{}) {
			if reflect.TypeOf(v).Kind() == reflect.String {
				envData += fmt.Sprintf("%s=\"%v\"\n", strings.ToUpper(k), v)
			} else {
				envData += fmt.Sprintf("%s=%v\n", strings.ToUpper(k), v)
			}
		}
		fileData = []byte(envData)
	} else {
		log.Infof("[toolkit] generate json file from %s", secretPath)
		fileData, err = json.Marshal(respData.Data)
	}
	err = ioutil.WriteFile(fileName, fileData, 0644)
	if err != nil {
		log.Errorf("[toolkit] error writing %s, err : %+v", fileName, err)
	}
	return
}

func (c *ToolkitDataCase) MergePR(repo string, prID int64, squash bool) (err error) {
	var (
		data      types.GithubPRInfoData
		mergeData types.MergeResponseData
	)
	if squash {
		log.Infof("[toolkit] merge squash PR with prID %d from repo %s", prID, repo)
	} else {
		log.Infof("[toolkit] merge PR with prID %d from repo %s", prID, repo)
	}

	data, err = c.GitRepo.GetPRInfo(repo, prID)
	if err != nil {
		log.Errorf("[toolkit] merge failed error get PR info, err %+v", err)
		return
	}

	log.Infof("[toolkit] %d changed files, %d deletion, %d addition", data.ChangedFiles, data.Deletions, data.Additions)
	log.Infof("[toolkit] PR url : %s", data.PRUrl)
	log.Infof("[toolkit] Merge status : %t", data.Merged)

	if data.Merged {
		log.Info("[toolkit] Aborting merge pull request, pr already merged")
		return
	}

	if !data.Mergeable {
		return errors.New("[toolkit] merge failed, please resolve merge conflict")
	}

	log.Info("[toolkit] Merge pull request")
	mergeData, err = c.GitRepo.MergePR(repo, prID, squash)
	if err != nil {
		return fmt.Errorf("[toolkit] merge failed, error execute github API, err %+v", err)
	}
	if mergeData.Merged {
		log.Infof("[toolkit] Pull request %d has been merged successfully", prID)
	} else {
		err = fmt.Errorf("[toolkit] merge failed, with message %s", mergeData.Message)
	}
	return
}
