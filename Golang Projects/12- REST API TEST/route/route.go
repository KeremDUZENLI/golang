package route

import (
	"net/http"
	"restApiTest/configs"
	"restApiTest/controller"
	"restApiTest/model"

	"github.com/gin-gonic/gin"
)

func URL() {
	configs.DatabaseDB()
	configs.Database.AutoMigrate(&model.User{})

	ginRouter := gin.Default()

	ginRouter.GET("/begin", controller.Begin)

	ginRouter.POST("/create", controller.Create)
	ginRouter.GET("/read", controller.Read)
	ginRouter.PUT("/update/:employeeID", controller.Update)
	ginRouter.DELETE("/delete/:employeeID", controller.Delete)

	ginRouter.GET("/beginDatabase", controller.BeginDatabase)

	ginRouter.POST("/createDatabase", controller.CreateDatabase)
	ginRouter.GET("/readDatabase", controller.ReadDatabase)
	ginRouter.PUT("/updateDatabase/:employeeID", controller.UpdateDatabase)
	ginRouter.DELETE("/deleteDatabase/:employeeID", controller.DeleteDatabase)

	http.ListenAndServe(":8888", ginRouter)
}
