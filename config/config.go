package config

import "os"

// ServerConfig represents main configuration for server
type ServerConfig struct {
	// Port on which vehicle is listening
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
		Port:          getenv("VEHICLE_PORT", "8080"),
		MySQLHost:     getenv("MYSQL_HOST", "db"),
		MySQLPort:     getenv("MYSQL_PORT", "3306"),
		MySQLDatabase: getenv("MYSQL_DATABASE", "vehicle"),
		MySQLUser:     getenv("MYSQL_USER", "vehicle"),
		MySQLPassword: getenv("MYSQL_PASSWORD", "vehicle"),
		APIUser:       getenv("API_USER", "vehicle"),
		APIPassword:   getenv("API_PASSWORD", "vehicle"),
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
