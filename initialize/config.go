package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var GlobalConfig *Config

type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
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

var envConfig = pflag.String("env", "dev", "Example: go run main.go --env=dev")

// InitConfig 初始化配置
func InitConfig() {
	pflag.Parse()
	
	config := viper.New()
	workDir, _ := os.Getwd()
	config.AddConfigPath(workDir + "/conf/yaml/")
	config.SetConfigName(fmt.Sprintf("config-%s", *envConfig))
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := config.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}
}
