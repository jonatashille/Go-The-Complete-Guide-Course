package middlewares

import (
	"REST-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized - Empty Token."})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized - Invalid Token."})
		return
	}

	context.Set("userId", userID)

	context.Next()
}
