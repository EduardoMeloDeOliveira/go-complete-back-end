package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) health(ginContext * gin.Context) {


	ginContext.JSON(http.StatusOK,gin.H{"message" : "application is rigth"})

}