package application

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ServerPort uint16
	RedisHost  string
	RedisPort  string
}

func (c *Config) GetConfig() *Config {
	cfg := &Config{
		ServerPort: 5000,
		RedisHost:  "localhost",
		RedisPort:  "6379",
	}

	if f, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(f, 10, 16); err == nil {
			fmt.Println("SERVER_PORT: ", port)
			cfg.ServerPort = uint16(port)
		}
	}

	if f, exists := os.LookupEnv("REDIS_HOST"); exists {
		cfg.RedisHost = f
	}

	if f, exists := os.LookupEnv("REDIS_PORT"); exists {
		cfg.RedisPort = f
	}

	return cfg
}
