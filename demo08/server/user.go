package server

import (
	"demo08/database/dto"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) dto.ResponseResult {
	data := map[string]interface{}{
		"foo":  "bar12",
		"temp": "1111",
	}
	return dto.SetResponseSuccess(data)
}
