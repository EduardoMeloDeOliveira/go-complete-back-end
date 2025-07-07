package main

import (
	database "go-rest-api/internal/database/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *Application) createEvent(ginContext *gin.Context) {
	var event database.Event

	if err := ginContext.ShouldBindJSON(&event); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error":"Invalid id"})
		return
	}

	err := app.models.Event.Insert(&event)

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, event)

}

func (app *Application) getEventById(ginContext *gin.Context) {
	id, err := strconv.Atoi(ginContext.Param("id"))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := app.models.Event.Get(id)

	if event == nil {
		ginContext.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, event)

}

func (app *Application) updateEvent(ginContext *gin.Context) {
	id, err := strconv.Atoi(ginContext.Param("id"))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event Id"})
		return
	}

	existingEvent, err := app.models.Event.Get(id)

	if existingEvent == nil {
		ginContext.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateEvent := &database.Event{}

	if err := ginContext.ShouldBindJSON(updateEvent); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}


	updateEvent.ID = id

	if err := app.models.Event.Update(updateEvent); err !=nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK,updateEvent)
}


func (app *Application) getAllEvents (ginContext * gin.Context){
	events,err := app.models.Event.GetAll()

	if err != nil{
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error" : "erro to retrive all events"})
		return
	}

	ginContext.JSON(http.StatusOK,events)

}


func (app * Application) deleteEvent (ginContext * gin.Context){
	id, err := strconv.Atoi(ginContext.Param("id"))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H {"error" : "Invalid event id"})
		return
	}

	event,err := app.models.Event.Get(id) 

	if event == nil {
		ginContext.JSON(http.StatusNotFound, gin.H{"error" : "event not found"})
		return
	}

	if err != nil{
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error() })
		return
	}

	

	if err := app.models.Event.Delete(id); err != nil{
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	ginContext.JSON(http.StatusNoContent,nil)
}