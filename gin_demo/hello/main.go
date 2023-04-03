package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello World!")
	})
	r.Run(":8000")
}

//http://127.0.0.1:8000/
