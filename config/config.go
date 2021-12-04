package config

import (
	"os"
)

// ServerConfig represents main configuration for server
type ServerConfig struct {
	// Port on which digihey is listening
	Port          string
	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
	MySQLUser     string
	MySQLPassword string
	APIUser       string
	APIPassword   string
}

// NewServerConfig returns new configuration object for server
func NewServerConfig() *ServerConfig {
	config := &ServerConfig{
		Port:          getenv("DIGIHEY_PORT", "8080"),
		MySQLHost:     getenv("MYSQL_HOST", "db"),
		MySQLPort:     getenv("MYSQL_PORT", "3306"),
		MySQLDatabase: getenv("MYSQL_DATABASE", "digihey"),
		MySQLUser:     getenv("MYSQL_USER", "digihey"),
		MySQLPassword: getenv("MYSQL_PASSWORD", "digihey"),
		APIUser:       getenv("API_USER", "digihey"),
		APIPassword:   getenv("API_PASSWORD", "digihey"),
	}

	return config
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
