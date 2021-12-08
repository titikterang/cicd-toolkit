package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/repository/util"
	caseUtil "github.com/ujunglangit-id/cicd-toolkit/internal/usecase/util"
	"os"
)

const (
	AppVersion = "1.0.0"
)

var (
	err         error
	config      *core.Config
	repoWrapper *util.RepoWrapper
	caseWrapper *caseUtil.CaseWrapper
)

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	version := flag.Bool("v", false, "toolkit version")
	vaultV1 := flag.Bool("v1", false, "use kv secret engine v1")
	mergePR := flag.Bool("merge", false, "merge pull request")
	mergeSquash := flag.Bool("squash", false, "merge squash pull request")
	vaultExport := flag.Bool("vault", false, "export vault secret")
	envExport := flag.Bool("env", false, "set vault secret type, json/env file")
	fileExport := flag.Bool("raw", false, "set vault secret type, json/env/raw-text file")
	outputName := flag.String("output", "", "vault secret output filename")
	secretPath := flag.String("secret", "", "vault secret path")
	repoName := flag.String("repo", "", "repository name")
	prID := flag.Int64("id", 0, "pull request ID")
	validateApproval := flag.Bool("approval", false, "validate PR approval")

	flag.Parse()
	if *version {
		log.Infof("toolkit version %s", AppVersion)
		os.Exit(0)
	}

	config, err = core.InitConfig(*debug)
	if err != nil {
		log.Fatalf("[toolkit] failed to load config : %+v", err)
	}

	repoWrapper = util.New(config)
	caseWrapper = caseUtil.New(config, repoWrapper)

	if *debug {
		log.Infof("[toolkit] approal limit : %d", config.Github.ApprovalLimit)
		log.Infof("[toolkit] vault host : %s", config.Vault.Host)
		log.Infof("[toolkit] git api host : %s", config.Github.Host)
		log.Infof("[toolkit] gcloud project id : %s", config.GCloud.ProjectID)
		log.Infof("[toolkit] gcloud json path : %s", config.GCloud.JsonPath)
	}

	if *mergePR || *mergeSquash {
		if *repoName != "" && *prID != 0 {
			err = caseWrapper.ToolkitCase.MergePR(*repoName, *prID, *mergeSquash)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("[toolkit] please specify repo name & pull request ID")
		}
		os.Exit(0)
	}

	if *vaultExport {
		if *secretPath != "" {
			err = caseWrapper.ToolkitCase.GetVaultSecret(*secretPath, *outputName, *envExport, *fileExport, *vaultV1)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("[toolkit] please specify secret path")
		}
		os.Exit(0)
	}

	if *validateApproval {
		if *repoName != "" && *prID != 0 {
			err = caseWrapper.ToolkitCase.ValidateApprovalStatus(*repoName, *prID)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("[toolkit] please specify repo name & Pull Request ID")
		}
		os.Exit(0)
	}
}
