package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 自定义变量结构体 ENV,这个名字随意定义
type ENV struct {
	DB_HOST  string `mapstructure:"DB_HOST"`
	DB_PORT  string `mapstructure:"DB_PORT"`
	DB_USER  string `mapstructure:"DB_USER"`
	DB_PWD   string `mapstructure:"DB_PWD"`
	DB_NAME  string `mapstructure:"DB_NAME"`
	DB_CONNS int    `mapstructure:"DB_CONNS"`
}

func TestInitEnv(t *testing.T) {

	// 使用 &
	env := &ENV{}

	err := InitEnv(Config{
		FileName:      "app",
		FilePath:      "",
		FileType:      "env",
		AddConfigPath: "conf",
	}, env)
	if err != nil {
		log.Fatal("init env failed err :", err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, env)

	t.Log("env is ", env)
}
