package config

import (
    "os"
)

type Config struct {
    Port string
}

func LoadConfig() Config {
    return Config{
        Port: getEnv("PORT", "8080"),
    }
}

func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}
