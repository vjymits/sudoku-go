# Sudoku Web App

This source code contain sudoku server implemented in Golang, and a SPA client index.html

## Server
By running main.go `go run main.go` server will start on port on port **8080**
There are 2 web endpoints

    1. <host>/: This will render home page (Only web page.)
    2. <host>/sudoku/api/v1/gen : The REST API which returns sudoku puzzle, and its answer.

## Client
The client is accessible by browser eg. **http://localhost:8080/**.
- The home page renders a sudoku puzzle.
- Reset button will erase all the user changes.
- Answer button will show the answer of puzzle. 




------------------------------------------------------------------------------------

###Sudoku Assignment in Golang

Within the scopes of this assignment the candidate is required to implement a simple sudoku game in Go from scratch, with separate architecture of front-end and back-end, i.e. the calculations and processing is to be performed in server side, facilitating the user-side UI.

Several skills will be tested while assessing the implementation:
* Architectural skills based on design patterns and paradigms accounting for future enhancements.
* Back-end development and networking skills in Go.
* Coding style and program efficiency, i.e. the skill of achieving the same result in the best possible way in Golang.
* Creativity while designing the game.
* Documenting and relevant commenting skills.
* Top-down thinking and designing skills with future in mind.
* Bug-free programming skills.

The assessment will be performed with above-state points in mind.
Copy-pasted implementation will render the assignment failed.

