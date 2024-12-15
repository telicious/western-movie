package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, payload)
}

func RespondWithError(ctx *gin.Context, code int, message string) {
	RespondWithJSON(ctx, code, map[string]string{"error": message})
}

func ParseInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
