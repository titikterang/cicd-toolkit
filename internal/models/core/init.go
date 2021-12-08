package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
	"net"
	"os"
)

func InitConfig(debugMode bool) (cfg *Config, err error) {
	cfg = &Config{}
	cfg.DebugMode = debugMode
	err = cfg.readFile(debugMode)
	if err != nil {
		return
	}
	return
}

func (cfg *Config) readFile(debugMode bool) (err error) {
	var (
		f *os.File
	)

	path := []string{
		"/data/drone_data/toolkit-config",
		"/etc/toolkit-config",
		"files/etc/toolkit-config",
		"files/config",
		"./files/etc/toolkit-config",
		"./files/config",
		"../files/etc/toolkit-config",
		"../../files/etc/toolkit-config",
	}

	for _, val := range path {
		f, err = os.Open(fmt.Sprintf(`%s/config.yml`, val))
		if err == nil {
			if debugMode {
				log.Infof("[toolkit] load config from : %s", fmt.Sprintf(`%s/config.yml`, val))
			}
			decoder := yaml.NewDecoder(f)
			err = decoder.Decode(cfg)
			break
		}
	}

	if err != nil {
		return
	}
	return
}

func (cfg *Config) getOutboundIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}
