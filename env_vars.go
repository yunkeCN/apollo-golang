package config

import "os"

const (
	// ENV_APOLLO_META_SERVER_URL is Apollo META Server URL
	ENV_APOLLO_META_SERVER_URL = "APOLLO_META_SERVER_URL"
	// APOLLO_ACCESSKEY_SECRET is APOLLO_ACCESSKEY_SECRET
	APOLLO_ACCESSKEY_SECRET = "APOLLO_ACCESSKEY_SECRET"

	// TENANT_NAMESPACE is CC_NAMESPACE
	TENANT_NAMESPACE = "TENANT_NAMESPACE"
)

// GetApolloServerURL gets apollo server url config from env.
func GetApolloServerURL() string {
	return os.Getenv(ENV_APOLLO_META_SERVER_URL)
}

func GetApolloAccesskeySecret() string {
	return os.Getenv(APOLLO_ACCESSKEY_SECRET)
}

// GetTenantNamespace gets tenant namespace from env.
func GetTenantNamespace() string {
	return os.Getenv(TENANT_NAMESPACE)
}
