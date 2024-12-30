package main

import (
	"github.com/germatpreis/eldie/server/internal/culprits"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/culprits", culprits.GetCulprits)

	router.Run("localhost:8080")
}
