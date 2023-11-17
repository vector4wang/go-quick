package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")
	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found.")
			return
		}
		panic(err)
	}

	fmt.Println("type:", viper.GetString("type"))         // 读取字符串 type
	fmt.Println("a1:", viper.GetStringMap("a1"))          // 读取map a1
	fmt.Println("a1.k1:", viper.GetStringMap("a1")["k1"]) // 读取字符串a1.k1
	fmt.Println("a1.k1:", viper.GetString("a1.k1"))       // 嵌套读取字符串 a1.k1 注意和上面的区分,优先读取配置中和key完整匹配的部分
	fmt.Println("array:", viper.GetIntSlice("array"))     // 读取数组 array

	viper.RegisterAlias("t", "type") // 注册别名
	fmt.Println("type:", viper.GetString("t"))

	viper.AutomaticEnv()      // 自动绑定所有环境变量
	viper.AllowEmptyEnv(true) // 设置是否允许空环境变量
	viper.RegisterAlias("env", "os_env")
	viper.SetEnvPrefix("viper") // 设置前缀，用'_'和key连接
	if err := os.Setenv("viper_os_env", "viper_alias_prefix_ok"); err != nil {
		panic(err)
	}
	fmt.Println(viper.GetString("env")) // 输出 viper_alias_prefix_ok

	for i, k := range viper.AllKeys() {
		fmt.Println(i, ": ", k, "==>", viper.GetString(k))
	}

	fmt.Println("结构体")
	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("viper unmarshal err", err)
	}
	fmt.Println("vc struct: ", C)

	fmt.Println("监听配置变化")
	// 监听配置变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("new type:", viper.GetString("type"))
	})

	viper.WatchConfig()
	//time.Sleep(time.Minute)

}

type ViperConfig struct {
	Type        string      `mapstructure:"type"`
	A1          A1Config    `mapstructure:"a1"`
	ArrayNum    []int       `mapstructure:"arrayNum"`
	MysqlConfig mysqlConfig `mapstructure:"mysql_config"`
}
type A1Config struct {
	K1 string `mapstructure:"k1"`
	K2 string `mapstructure:"k2"`
}

type mysqlConfig struct {
	UserName string `mapstructure:"username"`
	Pwd      string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}
