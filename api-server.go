package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"gitlab.com/vimiew/api-template/configuration"
	"gitlab.com/vimiew/api-template/routes"
)

func main() {
	// Load config file
	configLocation := "config.json"

	// No config file was passed used the default in the same directory
	if len(os.Args) > 1 {
		args := os.Args[1]
		configLocation = fmt.Sprintf("%s", args)
	}

	// Load the configuration file
	config, _ := configuration.LoadConfig(configLocation)

	fmt.Println("Sofix Desktop API")
	fmt.Println("Created By: William Babcock // VIMIEW LLC")

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
	routes.InitAcidSalesSQL(router)
	routes.InitDesktop(router)
	routes.InitFreonUSage(router)
	routes.InitLynx(router)
	routes.InitMetrics(router)
	routes.InitRingisho(router)
	routes.InitScaleSQL(router)
	routes.InitSofixDesktop(router)
	routes.InitTaranis(router)
	routes.InitTrainingSQL(router)

	log.WithFields(log.Fields{"port": config.API.Port}).Info("Server running on localhost...")
	//fmt.Printf("\nServer running on port %d..\n", config.API.Port)
	router.Run(":" + fmt.Sprintf("%d", config.API.Port))
}
