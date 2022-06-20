package config

import "os"

// getApolloServerURL gets apollo server url config from env.
func getApolloServerURL() string {
	return os.Getenv("APOLLO_META_SERVER_URL")
}

func getApolloAccesskeySecret() string {
	return os.Getenv("APOLLO_ACCESSKEY_SECRET")
}

// getTenantNamespace gets tenant namespace from env.
func getTenantNamespace() string {
	return os.Getenv("TENANT_NAMESPACE")
}
