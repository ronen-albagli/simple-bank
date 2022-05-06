package main

import (
	"bank/api"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	api.InitRoutes(app)

	app.Run("localhost:8001")
}
