#### 环境变量初始化小工具

> 简介: 就是对 viper 的代码进行了简易封装,类似一个公共函数的作用

#### 使用方式

- 安装

```
go get github.com/keepondream/goutils
```

- 在项目 main 文件中自定义环境变量结构体

```
// 自定义变量结构体 ENV,这个名字随意定义,属性自己定义与配置文件中要一一对应
type ENV struct {
	DB_HOST  string `mapstructure:"DB_HOST"`
	DB_PORT  string `mapstructure:"DB_PORT"`
	DB_USER  string `mapstructure:"DB_USER"`
	DB_PWD   string `mapstructure:"DB_PWD"`
	DB_NAME  string `mapstructure:"DB_NAME"`
	DB_CONNS int    `mapstructure:"DB_CONNS"`
}
```

- 在项目根目录下增加 app.env 文件, 配置参数与 上面一一对应

```
# app.env
DB_HOST=192.168.102.222
DB_PORT=6111
DB_USER=root
DB_PWD=root
DB_NAME=root
DB_CONNS=100

...

```

- 调用初始化 env 变量并调用 config.InitEnv 方法解析对应数据

```
    env := &ENV{} // 这里注意使用 & 取址
	err := config.InitEnv(config.Config{
		FileName:      "app",   // 指定文件名称
		FilePath:      "",      // 指定文件路劲, 优先级最高,不为空则使用该配置进行加载
		FileType:      "env",   // 指定文件类型
		AddConfigPath: "conf",  // 加载配置文件路径, 默认会添加 "." ,"conf" 当前路劲和conf目录下
	}, env)
	if err != nil {
		log.Fatal("init env failed err :", err)
	}
```
