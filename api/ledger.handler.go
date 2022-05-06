package api

import (
	"bank/config"
	"bank/domain"
	"bank/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type handler struct {
	ledgerService domain.Service
}

func HandleNewAssetInsertion() gin.HandlerFunc {

	return func(c *gin.Context) {
		var validate = validator.New()
		var ledgerAssets domain.Ledger

		if err := c.BindJSON(&ledgerAssets); err != nil {
			c.JSON(http.StatusBadRequest, LedgerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&ledgerAssets); validationErr != nil {
			c.JSON(http.StatusBadRequest, LedgerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		conf, _ := config.NewConfig("./config/config.yaml")
		repo, _ := repository.CreateMongoRepo(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)

		service := domain.InitLedgerService(repo)

		service.InsertNewLedgerAsset(&ledgerAssets)

		c.JSON(http.StatusCreated, "{}")
	}
}
