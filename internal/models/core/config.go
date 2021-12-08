package core

type (
	Config struct {
		Vault     APIConfig    `yaml:"vault"`
		Github    GitAPIConfig `yaml:"github"`
		GCloud    GCloudConfig `yaml:"gcloud"`
		DebugMode bool
	}

	APIConfig struct {
		Host string `yaml:"host"`
		Key  string `yaml:"key"`
	}

	GitAPIConfig struct {
		Host          string `yaml:"host"`
		User          string `yaml:"user"`
		Token         string `yaml:"token"`
		ApprovalLimit int    `yaml:"approval"`
	}

	GCloudConfig struct {
		JsonPath  string `yaml:"json_path"`
		ProjectID string `yaml:"project_id"`
	}
)
