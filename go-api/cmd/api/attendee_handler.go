package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func(app * Application) getAllAttendees (ginContext * gin.Context){
	id, err := strconv.Atoi(ginContext.Param("id"))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest , gin.H{"error" : "Invalid id"})
	}
	
	
}