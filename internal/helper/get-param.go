package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(context *gin.Context, idName string) (uint, Error) {
	id, err := strconv.Atoi(context.Param(idName))

	if err != nil {
		return uint(0), BadRequest("Invalid id params")
	}

	return uint(id), nil
}