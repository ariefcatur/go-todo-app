package config

import (
	"github.com/spf13/viper"
	"log"
)

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
