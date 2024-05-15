package main

import (
	"graqler-metamodel-db/metamodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/queries", metamodel.HandleRpc)

	// listen and serve on localhost:8080
	err := r.Run(":8080")
	if err != nil {
		println("Abnormal exit: failed to start database listener.")
		println(err)
		return
	}
}
