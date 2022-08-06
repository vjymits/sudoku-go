package sudoku

import (
	"fmt"
	"math/rand"
	"time"
	"golang.org/x/exp/slices"
)

type Sudoku struct {
	a [9][9]int
	K int
	ans [9][9]int
}

/* 
Construct Sudoku object 
*/
func New() *Sudoku {
	s := &Sudoku{K:40}
	// Assigning seed sudoku 
	s.a = s.Seed()
	return s
}
/*
Genrate a new puzzle
*/
func Generate() *Sudoku {
	s := New()
	s.fillZeros()
	return s
}

/*
Put zeros to hide answer.
*/
func (s *Sudoku ) fillZeros() {
	i := 0
	
	for i = 0; i < 50; i++ {
		// Do swaping only 18 times
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(10000)%9
		// swap random row and col
		s.swapping(num)
		
	}
	// Before filling zeros save answer
	s.ans = s.a
	i = s.K
	for i > 0 {
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(10000)%9
		col := rand.Intn(10000)%9
		
		if s.a[row][col] != 0 {
			s.a[row][col] = 0
			i--
		}		
	}
}
/* 
Need one random num in arg to swap rows and column.
*/
func (s *Sudoku) swapping(num int) {
	
	if num == 1 || num == 2 {
		// num is either 1 or 2.
		s.swapCol(num, 0)
		s.swapRow(0, num)
	} else if num == 4 || num == 5 {
		// num is either 4 or 5.
		s.swapCol(num, 3)
		s.swapRow(num, 3)
	} else if num == 7 || num == 8 {
		// num is either 7 or 8.
		s.swapCol(num, 6)
		s.swapRow(num, 6)
	}else if num == 0{
		s.swapRow(0, 1)
	}else if num == 3{
		s.swapRow(3,4)
	}else if num == 6{
		s.swapRow(6, 7)
	}
}

/*
Print sudoku on default console.
*/
func (s *Sudoku) Print() {
	printMat(s.a)
	
}
// Print the answer
func (s *Sudoku) PrintAns() {
	printMat(s.ans)
	
}


func printMat(a [9][9]int) {

	for r := 0; r < 9; r++ {
		
		for c := 0; c < 9; c++ {
			fmt.Printf("%d, ", a[r][c])
		}
		fmt.Println("")
	}
}

/*
Swap two rows.
*/
func (s *Sudoku) swapRow(row1 int, row2 int) {
	fmt.Println("Swap row: ", row1, "with ", row2)
	s.a[row1], s.a[row2] = s.a[row2], s.a[row1]
}

/*
Swap two colmns
*/
func (s *Sudoku) swapCol(c1 int, c2 int) {
	fmt.Println("Swap col: ", c1, "with ", c2)
	for i := 0; i < 9; i++{
		s.a[i][c1], s.a[i][c2] = s.a[i][c2], s.a[i][c1]
	}
}

/*
Is Value unique in block or not.
*/
func (s *Sudoku) UniqueInBlock(blk [][]int, val int) bool {
	//TODO
	return false
}

/*
Checks wether the value will be unique in the column or not.
*/
func (s *Sudoku) UniqueInCol(c int, val int) bool {
	for _, r := range s.a{
		if val == r[c] {
			return false
		}
	}
	return true
}

/*
Checks wether the value will be unique in the row or not.
*/
func (s *Sudoku) UniqueInRow(r int, val int) bool {
		return slices.Contains(s.a[r][:], val)
}

/*
Return the solved sudoku as seed.
*/
func (s *Sudoku) Seed() [9][9]int {
	return [9][9]int{
	{8, 2, 7, 1, 5, 4, 3, 9, 6},
	{9, 6, 5, 3, 2, 7, 1, 4, 8},
	{3, 4, 1, 6, 8, 9, 7, 5, 2},
	{5, 9, 3, 4, 6, 8, 2, 7, 1},
	{4, 7, 2, 5, 1, 3, 6, 8, 9},
	{6, 1, 8, 9, 7, 2, 4, 3, 5},
	{7, 8, 6, 2, 3, 5, 9, 1, 4},
	{1, 5, 4, 7, 9, 6, 8, 2, 3},
	{2, 3, 9, 8, 4, 1, 5, 6, 7},
	};
}

