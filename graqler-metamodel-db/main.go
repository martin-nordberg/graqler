package main

import (
	"graqler-metamodel-db/metamodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	metamodel.AddRoutes(router)

	// listen and serve on localhost:8080
	err := router.Run(":8080")
	if err != nil {
		println("Abnormal exit: failed to start database listener.")
		println(err)
		return
	}
}
