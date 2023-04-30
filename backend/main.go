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
		//routes.ManageUsrsRoute(v1.Group("adminUsers/"))
	}
	v2 := router.Group("v2")
	{
		routes.ManageUsrsRoute(v2.Group("adminUsers/"))
		routes.SalesRoute(v2.Group("adminSale/"))
	}

	router.Run(":8000");
	

}