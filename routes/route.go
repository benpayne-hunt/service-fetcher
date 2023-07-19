package routes

import (
	"github.com/benpayne-hunt/service-fetcher/controllers"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// get one
	router.GET("/service/:serviceId", controllers.GetService())
	// get many
	router.GET("/services", controllers.GetAllServices())
}
