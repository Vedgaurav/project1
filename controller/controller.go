package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	vRouter := router.Group("/travel")
	vRouter.GET("/admin/:name", admin)

	return router
}

func admin(c *gin.Context) {

	name := c.Param("name")
	response := "Your name is : " + name

	c.String(http.StatusOK, response)

}
