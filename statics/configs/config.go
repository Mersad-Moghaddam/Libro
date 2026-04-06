package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort               string
	AppEnv                string
	MySQLHost             string
	MySQLPort             string
	MySQLUser             string
	MySQLPassword         string
	MySQLDB               string
	MySQLCharset          string
	RedisHost             string
	RedisPort             string
	RedisPassword         string
	JWTSecret             string
	JWTRefreshSecret      string
	AccessTokenTTLMinutes int
	RefreshTokenTTLHours  int
}

func Load() (*Config, error) {
	_ = godotenv.Load("dev.env")

	accessTTL, err := strconv.Atoi(getEnv("ACCESS_TOKEN_TTL_MINUTES", "60"))
	if err != nil {
		return nil, fmt.Errorf("parse ACCESS_TOKEN_TTL_MINUTES: %w", err)
	}
	refreshTTL, err := strconv.Atoi(getEnv("REFRESH_TOKEN_TTL_HOURS", "168"))
	if err != nil {
		return nil, fmt.Errorf("parse REFRESH_TOKEN_TTL_HOURS: %w", err)
	}

	return &Config{
		AppPort:               getEnv("APP_PORT", "8080"),
		AppEnv:                getEnv("APP_ENV", "development"),
		MySQLHost:             getEnv("MYSQL_HOST", "127.0.0.1"),
		MySQLPort:             getEnv("MYSQL_PORT", "3306"),
		MySQLUser:             getEnv("MYSQL_USER", "root"),
		MySQLPassword:         getEnv("MYSQL_PASSWORD", "password"),
		MySQLDB:               getEnv("MYSQL_DB", "libro"),
		MySQLCharset:          getEnv("MYSQL_CHARSET", "utf8mb4"),
		RedisHost:             getEnv("REDIS_HOST", "127.0.0.1"),
		RedisPort:             getEnv("REDIS_PORT", "6379"),
		RedisPassword:         getEnv("REDIS_PASSWORD", ""),
		JWTSecret:             getEnv("JWT_SECRET", "secret"),
		JWTRefreshSecret:      getEnv("JWT_REFRESH_SECRET", "refresh_secret"),
		AccessTokenTTLMinutes: accessTTL,
		RefreshTokenTTLHours:  refreshTTL,
	}, nil
}

func (c *Config) MySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort, c.MySQLDB, c.MySQLCharset)
}

func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
