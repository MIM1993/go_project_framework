package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go_project_framework/global"
)

func InitConfig(filePath string) error {
	//读取配置文件
	_, err := toml.DecodeFile(filePath, &global.Global)
	if err != nil {
		return err
	}

	//读取命令行参数
	configFlagParam()
	return nil
}

func configFlagParam(){
	conf := &global.Global

	//example
	flag.StringVar(&conf.ListenAddressHTTP, "http_host", conf.ListenAddressHTTP, "http server start addr")

	flag.Parse()
}