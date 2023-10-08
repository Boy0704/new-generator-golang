package configuration

import (
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/logrusx"
	"github.com/spf13/viper"
)

type ConfigApp struct {
	App       App       `mapstructure:"app"`
	BasicAuth BasicAuth `mapstructure:"basic_auth"`
	Database  Database  `mapstructure:"database"`
	Cors      Cors      `mapstructure:"cors"`
	Api       Api       `mapstructure:"api"`
	Log       logrusx.Config
}

type App struct {
	Debug   bool   `mapstructure:"debug"`
	Port    string `mapstructure:"port"`
	Timeout int    `mapstructure:"timeout"`
}

type BasicAuth struct {
	Username string `mapstructure:"customer_service_tool_username"`
	Password string `mapstructure:"customer_service_tool_password"`
}

type Database struct {
	HostName     string `mapstructure:"hostname"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
}

type Cors struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
}

type Api struct {
	Group                   string `mapstructure:"group"`
	DashboardAllTransaction string `mapstructure:"dashboard_all_transaction"`
	ListSender              string `mapstructure:"list_sender"`
	ListClient              string `mapstructure:"list_client"`
}

func LoadConfig(pathFile string) (config ConfigApp, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(pathFile)

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		panic(err)
		return
	}

	return
}
