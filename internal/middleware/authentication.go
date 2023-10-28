package middleware

import (
	"final-project-02/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err  := helper.VerifyToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Set("userData", verifiedToken)
		context.Next()
	}
}