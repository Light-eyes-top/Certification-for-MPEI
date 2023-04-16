package server

import (
	"certification/intrenal/handler/rest"
	"certification/intrenal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	transport *rest.Handler
}

func NewRouter(transport *rest.Handler) *Router {
	return &Router{transport: transport}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", r.transport.SignUp)
		auth.POST("/sign-in", r.transport.SignIn)
	}
	file := router.Group("/file", middleware.UserIdentify)
	{
		file.POST("/sign", r.transport.CreateSign)
		file.GET("/check-sign", r.transport.CheckMySign)
		file.GET("/check-sign/:userId", r.transport.CheckUserSign)
		file.DELETE("/delete-sign", r.transport.DeleteSign)
	}
	return router
}
