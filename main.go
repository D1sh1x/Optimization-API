package main

import (
	"net/http"
	_ "net/http/pprof"
	"optimization/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/add", handlers.AddHandler)

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	r.Run(":8080")
}
