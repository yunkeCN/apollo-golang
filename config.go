package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ApolloResponse struct {
	Appid          string            `json:"appid"`
	Cluster        string            `json:"cluster"`
	NamespaceName  string            `json:"namespaceName"`
	Configurations map[string]string `json:"configurations"`
	ReleaseKey     string            `json:"releaseKey"`
}

func LoadApolloConfig(appId, namespace string, appConfig interface{}) error {
	apolloURL := os.Getenv("APOLLO_META_SERVER_URL")
	if apolloURL == "" {
		apolloURL = "http://120.77.148.214:18011/"
	}

	// 获取 ApolloResponse 远程配置
	url := fmt.Sprintf("%s/configs/%s/default/%s", apolloURL, appId, namespace)
	get, err := http.Get(url)
	if err != nil {
		return err
	}
	apollo := new(ApolloResponse)
	all, _ := ioutil.ReadAll(get.Body)
	if err := json.Unmarshal(all, apollo); err != nil {
		return err
	}

	bytes, err := json.Marshal(apollo.Configurations)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, appConfig)
	if err != nil {
		return err
	}

	return nil
}
