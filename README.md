# apollo-golang
Apollo client for golang

### 支持的环境变量:  
`APOLLO_META_SERVER_URL`  
`APOLLO_ACCESSKEY_SECRET`  
`TENANT_NAMESPACE` 


### demo
```go
package main

import config "github.com/yunkeCN/apollo-golang"

type MysqlSettingS struct {
	Host             string
	Port             string
	UserName         string
	Password         string
	DBName           string
	Charset          string
	MaxConnLifetime  int
	MaxIdleConnCount int
	MaxOpenConnCount int
}

var settings = &MysqlSettingS{}

// 加载配置
func loadConfig() error {
	if err := config.MapConfig("Mysql", settings, true); err != nil {
		return err
	}

	return nil
}

func main() {
	
	// 配置应用名称和加载配置的回调
	c := &config.ApolloConfig{
		AppID:      "mars-base",
		LoadConfig: loadConfig,
	}
	
	// 启动服务
	err := config.NewApolloService(c)

	if err != nil {
		panic(err)
	}
}
```