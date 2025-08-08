package config

import (
    "fmt"
    "os"
)

// Config holds all configuration values for the application.
// In a real project you might use a library like viper to load these
// values from environment variables, files, or command line flags. Here
// we keep it simple and read directly from the environment.
type Config struct {
    // ServerHost is the host where the HTTP server listens.
    ServerHost string
    // ServerPort is the port where the HTTP server listens.
    ServerPort string
    // DatabaseDSN is the Data Source Name for connecting to the Postgres database.
    DatabaseDSN string
    // JWTSecret is the secret key used to sign JWT tokens.
    JWTSecret string
    // JWTExpirationSeconds defines how long tokens are valid.
    JWTExpirationSeconds int
}

// New reads configuration from environment variables and returns a Config instance.
func New() *Config {
    cfg := &Config{
        ServerHost:           getEnv("SERVER_HOST", "0.0.0.0"),
        ServerPort:           getEnv("SERVER_PORT", "8080"),
        DatabaseDSN:          getEnv("DATABASE_DSN", "postgres://user:pass@localhost:5432/vietceramics?sslmode=disable"),
        JWTSecret:            getEnv("JWT_SECRET", "secret"),
        JWTExpirationSeconds: getEnvAsInt("JWT_EXP_SECONDS", 3600),
    }
    return cfg
}

// ServerAddress returns a host:port string for the HTTP server.
func (c *Config) ServerAddress() string {
    return c.ServerHost + ":" + c.ServerPort
}

// Helper function to read an environment variable or return a default value.
func getEnv(key string, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// Helper function to read an environment variable as an int or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
    if valStr, exists := os.LookupEnv(name); exists {
        var val int
        _, err := fmt.Sscanf(valStr, "%d", &val)
        if err == nil {
            return val
        }
    }
    return defaultVal
}
