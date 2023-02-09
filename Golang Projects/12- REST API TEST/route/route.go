package route

import (
	"net/http"
	"restApiTest/controller"

	"github.com/gin-gonic/gin"
)

func URL() {
	r := gin.Default()

	r.GET("/begin", controller.Begin)

	r.POST("/create", controller.Create)
	r.GET("/read", controller.Read)
	r.PUT("/update/:employeeID", controller.Update)
	r.DELETE("/delete/:employeeID", controller.Delete)

	http.ListenAndServe(":8888", r)
}
