package main

import (
	"fmt"

	"github.com/vjymits/smartcow/sudoku/sudoku"
)

func main() {
	s := sudoku.Generate()
	s.Print()
	fmt.Println("-------------")
	s.PrintAns()


}