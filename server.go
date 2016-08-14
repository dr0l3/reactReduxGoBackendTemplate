package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/", http.Dir("C:/Users/runed/Projects/reactReduxGo"))
	router.Run(":8080")
}
