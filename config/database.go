package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

//func InitDatabase() {
//	InitConfig()
//	user := viper.GetString("DB_USER")
//	pass := viper.GetString("DB_PASS")
//	host := viper.GetString("DB_HOST")
//	port := viper.GetString("DB_PORT")
//	dbname := viper.GetString("DB_NAME")
//
//	// POSTGRESQL
//	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",
//		host, user, dbname, port, pass,
//	)
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect to database:", err)
//	}
//
//	// Enable Logger
//	db = db.Debug()
//
//	DB = db
//	log.Println("Database connection established")
//}

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(C.DBDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// Optional: ping & set connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql db: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
