package pkg

import (
	"bytes"
	"github.com/spf13/viper"
	"embed"
)

//go:embed config.yaml
var configFile embed.FS

var (
	Config *viper.Viper
)

func LoadConfig() (*viper.Viper, error) {
	data, err := configFile.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	vip := viper.New()
	vip.SetConfigType("yaml")
	if err = vip.ReadConfig(bytes.NewReader(data)); err != nil {
		return nil, err
	}

	return vip, nil
}

func InitConfig() {
	//time.Sleep(3 * time.Second)
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	Config = config
}
