package transport

import "github.com/gin-gonic/gin"

func ServerRun() {
	router := gin.Default()
	router.GET("/when/:year", GetDays)

	router.Run("localhost:8000")
}
