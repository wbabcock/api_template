package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/vimiew/api-template/dbcontroller"
)

func initTaranisTables() {
	// Migrate the schema
	db := dbcontroller.Database("Taranis")
	db.AutoMigrate()
	defer db.Close()
}

// InitTaranis - setup the routes
func InitTaranis(router *gin.Engine) {
	//*************************************************************************
	//   ACTIONS
	//*************************************************************************
	// router.GET("/taranis/actions", taranis.GetActions)
	// taranisAction := router.Group("/taranis/actions/action")
	// {
	// 	taranisAction.GET("/:id", taranis.GetAction)
	// 	taranisAction.POST("/", ringisho.PostStandard)
	// 	taranisAction.PUT("/:id", ringisho.PutStandard)
	// 	taranisAction.DELETE("/:id", ringisho.DeleteStandard)
	// }
}
