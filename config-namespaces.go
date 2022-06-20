package config

import (
	"github.com/philchia/agollo/v4"
)

const (
	NAMESPACE_APP = "application"
)

func setApolloNamespaces(conf *agollo.Conf) {
	namespace := getApolloNamespace()
	conf.NameSpaceNames = append(conf.NameSpaceNames, namespace)
	currentNamespace = namespace
}

func getApolloNamespace() string {
	apolloNamespace := NAMESPACE_APP

	tenantNamespace := getTenantNamespace()
	if tenantNamespace == "" || tenantNamespace == "g2" {
		return apolloNamespace
	}

	currentNamespace := tenantNamespace + "_" + NAMESPACE_APP
	return currentNamespace
}
