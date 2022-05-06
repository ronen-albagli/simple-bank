package api

import (
	"github.com/gin-gonic/gin"
)

func LedgerRoutes(router *gin.Engine) {
	// router.GET("/ledger", HandleNewAssetInsertion())
	router.POST("/ledger", HandleNewAssetInsertion())
	// router.GET("/user/:email", controllers.GetUser())    //add this
	// router.PUT("/user/:userId", controllers.EditAUser()) //add this
}
