package conf

import (
	"github.com/BurntSushi/toml" //nolint:goimports
	"migrates3/logs"
)

type Configuration struct {
	RcloneConfigPath string `json:"rclone_config_path"`
}

var config Configuration

func Init() {
	if _, err := toml.DecodeFile("./conf/conf.toml", &config); err != nil {
		logs.GetLogger().Fatal("error:", err)
	}
}

func GetConfig() Configuration {
	return config
}
