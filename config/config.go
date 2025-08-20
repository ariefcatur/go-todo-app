package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type AppConfig struct {
	Port      string
	DBDSN     string
	JWTSecret string
	APIKey    string
	GinMode   string
	JWTExpiry time.Duration
}

var C AppConfig

func Load() {
	C = AppConfig{
		Port:      getEnv("PORT", "8080"),
		DBDSN:     mustEnv("DB_DSN"),
		JWTSecret: mustEnv("JWT_SECRET"),
		APIKey:    os.Getenv("API_KEY"),
		GinMode:   getEnv("GIN_MODE", "release"),
		JWTExpiry: getDuration("JWT_EXP_MIN", 30),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing %s", k)
	}
	return v
}
func getDuration(k string, def int) time.Duration {
	if v := os.Getenv(k); v != "" {
		return time.Duration(atoi(v)) * time.Minute
	}
	return time.Duration(def) * time.Minute
}
func atoi(s string) int { var n int; fmt.Sscanf(s, "%d", &n); return n }

func InitConfig() {
	viper.SetConfigFile("config/.env")

	err := viper.ReadInConfig()
	if err != nil {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed read configuration!: %v", err)
		}

	}

	log.Println("Configuration mounted successfully!")
}
