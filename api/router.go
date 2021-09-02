package api

import (
	"github.com/gin-gonic/gin"
	"godistributor/internal/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	//router := gin.Default()
	//
	//mdrepo := repo.NewMDRepo(db)
	//mdservice := service.NewMDService(mdrepo)
	//mdhandler := handler.NewMDHandler(mdservice)
	//md := router.Group("/md")
	//{
	//	md.GET("/", middleware.CORSMiddleware(), mdhandler.GetMD)
	//	md.GET("/:id", middleware.CORSMiddleware(), mdhandler.GetOneMD)
	//}
	//
	//sdrepo := repo.NewSDRepo(db)
	//sdservice := service.NewSDService(sdrepo)
	//sdhandler := handler.NewSDHandler(sdservice)
	//sd := router.Group("/sd")
	//{
	//	sd.GET("/", middleware.CORSMiddleware(), sdhandler.GetSD)
	//	sd.GET("/:id", middleware.CORSMiddleware(), sdhandler.GetOneSD)
	//}
	//
	//return router
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	//r.GET("/tasks", controllers.FindTasks)
	//r.POST("/tasks", controllers.CreateTask)
	//r.GET("/tasks/:id", controllers.FindTask)
	//r.PATCH("/tasks/:id", controllers.UpdateTask)
	//r.DELETE("tasks/:id", controllers.DeleteTask)

	r.GET("/distribs", controllers.GetAllDistribs)
	r.POST("/distribs", controllers.CreateDistrib)
	r.GET("/distribs/:id", controllers.FindDistribs)
	r.PATCH("/distribs/:id", controllers.UpdateDistrib)
	r.DELETE("distribs/:id", controllers.DeleteDistrib)

	sd := r.Group("/sd")
	{
		sd.GET("/", controllers.GetAllSecDistribs)
		sd.POST("/", controllers.CreateSecDistrib)
		sd.GET("/:id", controllers.FindSecDistribs)
		sd.PATCH("/:id", controllers.UpdateSecDistrib)
		sd.DELETE("/:id", controllers.DeleteSecDistrib)
	}
	return r
}
