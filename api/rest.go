package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	LedgerRoutes(router)

}
