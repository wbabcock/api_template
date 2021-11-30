package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.com/vimiew/api-template/routes"
)

func main() {
	fmt.Println("API Template Server")

	// Set flag for Production
	gin.SetMode(gin.ReleaseMode)

	// API router
	router := gin.Default()

	routerConfig := cors.DefaultConfig()
	routerConfig.AllowAllOrigins = true
	routerConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	routerConfig.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(routerConfig))

	// Import all the routes that will be used in the API
	routes.InitTaranis(router)

	fmt.Printf("\nServer running on port %d..\n", 8001)
	router.Run(":" + fmt.Sprintf("%d", 8001))
}
