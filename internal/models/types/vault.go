package types

type (
	VaultKV2RespData struct {
		Data VaultResponseData `json:"data"`
	}

	VaultResponseData struct {
		RequestId     string      `json:"request_id"`
		LeaseId       string      `json:"lease_id"`
		Renewable     bool        `json:"renewable"`
		LeaseDuration int         `json:"lease_duration"`
		Data          interface{} `json:"data"`
	}

	RawFileVaultResponseData struct {
		Data struct {
			RequestId     string       `json:"request_id"`
			LeaseId       string       `json:"lease_id"`
			Renewable     bool         `json:"renewable"`
			LeaseDuration int          `json:"lease_duration"`
			Data          VaultRawData `json:"data"`
		} `json:"data"`
	}

	VaultRawData struct {
		FileData string `json:"file_data"`
		FileName string `json:"file_name"`
	}
)
