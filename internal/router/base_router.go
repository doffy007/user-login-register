package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r router) BaseRouter() {
	// Add Swagger
	r.route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userHandler := r.handler.UserHandler()

	v1Prefix := r.route.Group("/api")
	{
		v1Prefix := v1Prefix.Group("/v1")
		{

			userPrefix := v1Prefix.Group("/user")
			{
				userPrefix.POST("/register", userHandler.UserRegister)
				userPrefix.POST("/login", userHandler.UserLogin)
			}
		}
	}
}
