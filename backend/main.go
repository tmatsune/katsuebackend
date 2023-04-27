package main

//docker run -p 8000:8000 -tid katsuetwo

import (
	"katsuebackend/backend/middleware"
	"katsuebackend/backend/routes"
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/cors"
	"katsuebackend/backend/database"
)

func main(){
	database.Init();
	router := gin.Default();
	router.Use(middleware.CORSMiddleware());

	v1 := router.Group("v1")
	{
		routes.ShirtsRoute(v1.Group("shirts/"))
		routes.PantsRoute(v1.Group("pants/"))
		routes.UserRoutes(v1.Group("users/"))
		routes.StripeRoutes(v1.Group("pay/"))
	}

	router.Run(":8000");
	

}