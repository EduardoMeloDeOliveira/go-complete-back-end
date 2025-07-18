package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/events", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEventById)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)
		v1.GET("/events/:id/attendee",app.getAllAttendees)

		v1.GET("/health", app.health)

		v1.POST("/auth/register", app.registerUser)


	}
	return g
}
