package main

import (
	"fmt"
	"github.com/spf13/viper"
	"godistributor/api"
	"godistributor/internal/entity"
	"godistributor/pkg/config"
	"godistributor/pkg/database"
	_ "gorm.io/gorm"
	"log"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	db.AutoMigrate(&entity.MainDistributor{}, &entity.SecDistributor{})
	if err != nil {
		log.Fatal("Failed to open a DB Connection: ", err)
		return
	}

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	app := api.SetupRouter(db)
	app.Run(port)

	//db := entity.SetupDB()
	//db.AutoMigrate(&entity.MainDistributor{}, &entity.SecDistributor{})
	//
	//r := api.SetupRouter(db)
	//r.Run()

}
