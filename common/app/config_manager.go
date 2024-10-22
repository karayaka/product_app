package app

import "product_app/common/postgresql"

type ConfigManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		UserName:              "postgres",
		Password:              "postgres",
		DbName:                "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
