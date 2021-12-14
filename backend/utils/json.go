package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(response)
	c.Next()
}
