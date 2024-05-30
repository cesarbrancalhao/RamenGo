package controller

import (
	"api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBroths(c *gin.Context) {

	res, err := repository.GetBroths()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
