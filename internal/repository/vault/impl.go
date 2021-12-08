package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	circuit "github.com/eapache/go-resiliency/breaker"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/core"
	"github.com/ujunglangit-id/cicd-toolkit/internal/models/types"
	"io/ioutil"
	"net/http"
	"time"
)

type VaultAPI struct {
	Config         *core.Config
	CircuitBreaker *circuit.Breaker
	Client         *http.Client
}

func New(cfg *core.Config, cl *http.Client) *VaultAPI {
	return &VaultAPI{
		Config:         cfg,
		CircuitBreaker: circuit.New(types.BreakerErrorThreshold, types.BreakerSuccessThreshold, types.BreakerTimeout*time.Second),
		Client:         cl,
	}
}

func (v *VaultAPI) GetVaultSecret(path string, useV1 bool) (respData types.VaultResponseData, err error) {
	var V2Resp types.VaultKV2RespData
	err = v.CircuitBreaker.Run(func() error {
		req, errCB := http.NewRequest("GET", fmt.Sprintf("%s/v1/%s", v.Config.Vault.Host, path), nil)
		if errCB != nil {
			return errCB
		}

		req.Header.Set("X-Vault-Token", v.Config.Vault.Key)
		resp, errCB := v.Client.Do(req)
		if errCB != nil {
			return errCB
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New(http.StatusText(resp.StatusCode))
		}

		defer resp.Body.Close()
		body, errCB := ioutil.ReadAll(resp.Body)
		if errCB != nil {
			return errCB
		}

		if useV1 {
			errCB = json.Unmarshal(body, &respData)
			if errCB != nil {
				return errCB
			}
		} else {
			errCB = json.Unmarshal(body, &V2Resp)
			if errCB != nil {
				return errCB
			}
			respData = V2Resp.Data
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("code %d, resp data : %+v", resp.StatusCode, string(body))
		}

		return nil
	})

	return
}

func (v *VaultAPI) GetVaultRawSecret(path string) (respData types.RawFileVaultResponseData, err error) {

	err = v.CircuitBreaker.Run(func() error {
		req, errCB := http.NewRequest("GET", fmt.Sprintf("%s/v1/%s", v.Config.Vault.Host, path), nil)
		if errCB != nil {
			return errCB
		}

		req.Header.Set("X-Vault-Token", v.Config.Vault.Key)
		resp, errCB := v.Client.Do(req)
		if errCB != nil {
			return errCB
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New(http.StatusText(resp.StatusCode))
		}

		defer resp.Body.Close()
		body, errCB := ioutil.ReadAll(resp.Body)
		if errCB != nil {
			return errCB
		}

		errCB = json.Unmarshal(body, &respData)
		if errCB != nil {
			return errCB
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("code %d, GetVaultRawSecret resp data : %+v", resp.StatusCode, string(body))
		}

		return nil
	})

	return
}
