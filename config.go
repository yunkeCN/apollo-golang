package config

import (
	"errors"
	"github.com/philchia/agollo/v4"
)

type ApolloConfig struct {
	AppID      string
	LoadConfig func() error
}

var apolloClient agollo.Client
var currentNamespace string

func NewApolloService(conf *ApolloConfig) error {
	if conf.AppID == "" {
		return errors.New("AppID can't not be empty")
	}

	// 初始化配置
	c := &agollo.Conf{
		AppID:           conf.AppID,
		Cluster:         "default",
		MetaAddr:        getApolloServerURL(),
		AccesskeySecret: getApolloAccesskeySecret(),
	}

	// 设置namespace
	setApolloNamespaces(c)

	// 启动apollo client
	apolloClient = agollo.NewClient(c,
		agollo.SkipLocalCache(),
	)

	err := apolloClient.Start()
	if err != nil {
		return err
	}

	// 读取并映射配置
	if conf.LoadConfig != nil {
		err = conf.LoadConfig()
		if err != nil {
			return err
		}
	}

	// 开启更新监听
	watchUpdateConfig()
	return nil
}
