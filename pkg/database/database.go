package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDB() (*gorm.DB, error) {
	host := viper.GetString("Database.Host")
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	port := viper.GetInt("Database.Port")
	dbname := viper.GetString("Database.DBName")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, username, password, port, dbname)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
