package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置
type Config struct {
	FileName      string // 指定文件名称
	FilePath      string // 指定文件路劲, 优先级最高,不为空则使用该配置进行加载
	FileType      string // 指定文件类型
	AddConfigPath string // 加载配置文件路径, 默认会添加 "." ,"conf" 当前路劲和conf目录下
}

// InitEnv 初始化环境变量,优先级: 系统环境变量 > 配置文件配置的环境变量
// env 需要传递指针结构体 , 这样在配置文件变更后会同步 env中值
// type ENV struct {
// 	DB_HOST  string `mapstructure:"DB_HOST"`
// 	DB_PORT  string `mapstructure:"DB_PORT"`
// 	DB_USER  string `mapstructure:"DB_USER"`
// 	DB_PWD   string `mapstructure:"DB_PWD"`
// 	DB_NAME  string `mapstructure:"DB_NAME"`
// 	DB_CONNS int    `mapstructure:"DB_CONNS"`
// }
func InitEnv(c Config, env interface{}) error {
	if c.FilePath != "" {
		viper.SetConfigFile(c.FilePath)
	} else {
		viper.SetConfigName(c.FileName)

		viper.AddConfigPath(".")
		viper.AddConfigPath("conf")
		if c.AddConfigPath != "" {
			viper.AddConfigPath(c.AddConfigPath)
		}
	}

	viper.SetConfigType(c.FileType)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(env); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
		if err := viper.Unmarshal(env); err != nil {
			log.Fatal("config file changed failed err is ", err)
		}
	})

	return nil

}
