package core

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	DC DbConfig
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found.")
			return
		}
		panic(err)
	}

	err := viper.Unmarshal(&DC)
	if err != nil {
		fmt.Println("viper unmarshal err", err)
	}
	fmt.Println("load db config ", DC)

}

type DbConfig struct {
	MysqlConfig mysqlConfig `mapstructure:"mysql_config"`
}
type mysqlConfig struct {
	UserName string `mapstructure:"username"`
	Pwd      string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}
