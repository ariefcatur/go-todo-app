package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	InitConfig()
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	dbname := viper.GetString("DB_NAME")
	//// MYSQL
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	user, pass, host, port, dbname,
	//)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("Failed to connect to database:", err)
	//}

	// POSTGRESQL
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",
		host, user, dbname, port, pass,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable Logger
	db = db.Debug()

	DB = db
	log.Println("Database connection established")
}
