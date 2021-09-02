package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	host := viper.GetString("Database.Host")
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	port := viper.GetInt("Database.Port")
	dbname := viper.GetString("Database.DBName")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, username, password, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
