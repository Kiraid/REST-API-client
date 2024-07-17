package routes

import (
	"fmt"
	"net/http"

	"example.com/app/models"
	"example.com/app/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context){
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch events"})
		return
	}
	context.JSON(http.StatusOK, users)
}



func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save user", "error":   err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": user})
}

func logIn(context*gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse request data"})
		return
	}
	err = user.ValidateCredential()
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couldnt authenticate user"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt authenticate user"})	
	}
	context.JSON(http.StatusOK, gin.H{"message" :"Login succesful", "token": token})
} 