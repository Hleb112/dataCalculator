package transport

import (
	"dataCalculator/internal/services"
	"github.com/gin-gonic/gin"
)

func GetDays(c *gin.Context) {
	year := c.Param("year")
	result := services.Calculator(year)
	if result < 0 {
		result = result * -1
		c.Header("Content-Type", "text")
		c.String(200, "Days left: %v", result)
	} else {
		c.String(200, "Days gone: %v", result)
	}
}
