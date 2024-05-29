package controller

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BrothVar = models.Broth{
Id: "1",
	ImageInactive: "imgI",
	ImageActive: "imgA",
	Name: "Ramen",
	Description: "Noodles and soft boiled test",
	Price: 100,
}

func GetBroth(context *gin.Context) {
	context.JSON(http.StatusOK, BrothVar)
}