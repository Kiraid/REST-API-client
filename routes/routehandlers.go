package routes

import (
	"net/http"
	"strconv"

	"example.com/app/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id ,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt convert parse event id"})
		return
	}
	event, err := models.GetEventbyID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch id"})
		return
	}
	context.JSON(http.StatusOK, event)
	

}

func createEvent(context *gin.Context) {
	
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse request data"})
		return
	}
	id := context.GetInt64("userId")
	event.UserID = id  
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save event", "error":   err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context){
	id ,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt convert parse event id"})
		return
	}
	userId := context.GetInt64("userId")
	event , err := models.GetEventbyID(id)	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch event"})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt parse requested data"})
		return		
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt update event"})
		return		
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

}

func deleteEvent(context *gin.Context){
	id ,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt convert parse event id"})
		return
	}
	event , err := models.GetEventbyID(id)		
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch event"})
	}
	userID := context.GetInt64("userId")
	if event.ID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	} 
	
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}