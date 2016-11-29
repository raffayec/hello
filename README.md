Team: Mark Karle, Elena Raffay
Instructions: Go to https://golang.org/doc/install and install and set up a Go environment and a place for compiled exe files to go.  In Sudoku.go, change the path of the input sudoku to one where your sudoku.txt resides.  Run go install Sudoku.go which will create an exe for you to run.  The input file is a single line of 81 numbers for the sudoku. 


Example output:

beepboop@beepboop:~/Dev/sudoku$ ./Sudoku 
Here is the initial board:


0 4 3 | 0 8 0 | 2 5 0
6 0 0 | 0 0 0 | 0 0 0
0 0 0 | 0 0 1 | 0 9 4
------+-------+------
9 0 0 | 0 0 4 | 0 7 0
0 0 0 | 6 0 8 | 0 0 0
0 1 0 | 2 0 0 | 0 0 3
------+-------+------
8 2 0 | 5 0 0 | 0 0 0
0 0 0 | 0 0 0 | 0 0 5
0 3 4 | 0 9 0 | 7 1 0

Solved!


1 4 3 | 9 8 6 | 2 5 7
6 7 9 | 4 2 5 | 3 8 1
2 8 5 | 7 3 1 | 6 9 4
------+-------+------
9 6 2 | 3 5 4 | 1 7 8
3 5 7 | 6 1 8 | 9 4 2
4 1 8 | 2 7 9 | 5 6 3
------+-------+------
8 2 1 | 5 6 7 | 4 3 9
7 9 6 | 1 4 3 | 8 2 5
5 3 4 | 8 9 2 | 7 1 6

 The solver took 8.068ms

