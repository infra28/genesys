package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Mengambil data dari Viper (.env)
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	name := viper.GetString("DB_NAME")

	// Format DSN (Data Source Name) untuk Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Menampilkan query SQL di console
	})

	if err != nil {
		log.Fatal("Failed Conneting to database: ", err)
	}

	fmt.Println("Database Connected..")
	DB = db

}
