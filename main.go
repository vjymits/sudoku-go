package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vjymits/smartcow/sudoku/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*.html")
	router.GET("/", controller.Home)
	router.GET("/sudoku/api/v1/gen", controller.Create)
	fmt.Println("Starting server at localhost:8080")
	router.Run("localhost:8080")
	
}