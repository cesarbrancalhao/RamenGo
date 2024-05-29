package controller

import (
	"api/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBroths(c *gin.Context) {
	apikey := constants.GetApiKey()
	c.Request.Header.Add("x-api-key", apikey)

	url := "https://api.tech.redventures.com.br/broths"

	c.JSON(http.StatusOK, res)
}
