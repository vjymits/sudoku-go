package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/vjymits/smartcow/sudoku/sudoku"
)

func Create(c *gin.Context) {
	s := *sudoku.Generate()
	c.IndentedJSON(http.StatusOK, s)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "Home Page",
	})

}