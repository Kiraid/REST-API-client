package routes

import (
	"net/http"
	"strconv"

	"example.com/app/models"
	"github.com/gin-gonic/gin"
)
func registerForEvent(context *gin.Context){
	userId :=  context.GetInt64("userId")
	eventId ,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt convert parse event id"})
		return
	}
	event, err := models.GetEventbyID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
		return
	}
	err = event.Register(userId) 
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not register user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"Registered"})
}

func cancelRegistration(context *gin.Context){
	userId :=  context.GetInt64("userId")
	eventId ,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt convert parse event id"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not cancel registration"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"Deleted"})
	
	
}