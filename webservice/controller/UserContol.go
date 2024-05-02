package controller

import (
	"MyModulw/Model"
	"net/http"
    "MyModulw/Database"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context,) {
	var user model.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Db.Save(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
}
