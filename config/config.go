package config

import "fmt"

type Config struct {
	Debug            bool
	PostgresHost     string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
}

func NewConfig() *Config {
	initConfig()

	return &Config{
		Debug: getBoolEnv("DEBUG", false),
		// DATABASE
		PostgresHost:     getStrEnv("POSTGRES_HOST", ""),
		PostgresPort:     getIntEnv("POSTGRES_PORT", 0),
		PostgresDB:       getStrEnv("POSTGRES_DB", ""),
		PostgresUser:     getStrEnv("POSTGRES_USER", ""),
		PostgresPassword: getStrEnv("POSTGRES_PASSWORD", ""),
	}
}

func (c *Config) GetPostgresDsn() string {
	// Create Dsn for postgres connection
	// format: postgresql://username:password@ip:port/database
	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", c.PostgresUser, c.PostgresPassword, c.PostgresHost, c.PostgresPort, c.PostgresDB)

	return url
}
