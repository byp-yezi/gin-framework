package initialize

import (
	"os"

	"github.com/spf13/viper"
)

var GlobalConfig *Config

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
}

type Server struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf/yaml/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}
}
