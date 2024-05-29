package controller

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BrothVar = []models.Broth{
	{ Id: "1",
	  ImageInactive: "imgI",
	  ImageActive: "imgA",
	  Name: "Ramen",
	  Description: "Noodles and soft boiled egg",
	  Price: 100 },
}

func GetBroth(c *gin.Context) {
	c.JSON(http.StatusOK, BrothVar)
}