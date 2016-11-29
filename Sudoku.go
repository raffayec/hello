//Honor statement:  We pledge on our honor we had no unauthorized aid on this assignment.
package main

import (
  "fmt"
  "io/ioutil"
  "time"

)

var board [81] int



func readFile() {

	byteStr, err := ioutil.ReadFile("/home/beepboop/Dev/sudoku/sudoku.txt") 
	if err != nil {
        	fmt.Println(err)
	}

	str := string(byteStr)
	for i, v := range str{
		if i < 81{
			board[i] = int(v - '0')
		}
	}

}

func printBoard() {
	for row, index := 0, 0; row < 9; row, index = row + 1, index + 9 {
		fmt.Printf("%d %d %d | %d %d %d | %d %d %d\n" , 
		board[index], board[index+1], board[index+2], board[index+3],
		board[index+4], board[index+5], board[index+6], board[index+7], 
		board[index+8])
		if row == 2 || row == 5 {
			fmt.Println("------+-------+------")
		} 
	}

}



  
func checkRow(row, num int) bool {
	for _, v := range board[row * 9 : (row + 1) * 9]{
		if v == num{
			return false
		}
	}
	return true
}

func checkCol(col, num int) bool {
	for row := 0; row < 9; row++{
		if board[row * 9 + col] == num{
			return false
		}
	}
	return true
}

func checkSubgrid(row, col, num int) bool{
	for curRow := int(row/3) * 3; curRow - int(row/3)*3 < 3; curRow++{
		for curCol := int(col/3) * 3; curCol - int(col/3)*3 < 3; curCol++{
			if board[curRow * 9 + curCol]== num{
				return false			
			}
		}
	}
	return true
}

func check(row, col, num int) bool{
	return checkRow(row, num) && checkCol(col, num) && checkSubgrid(row, col, num)
}

func solve(row, col int) bool{
	if row == 9{
		return true
	}
	if col == 9{
		return solve(row + 1, 0)
	}
	if board[row * 9 + col] != 0{
		return solve(row, col +1)
	}
	for num:= 1; num < 10; num++{
		if check(row, col, num){
			board[row*9 + col] = num
			solved := solve(row, col + 1)
			if solved{ 
				return true 
			} else{
				board[row*9 + col] = 0
			}	
		} else if num == 9{
			board[row*9 + col] = 0
		}
	}
	
	return false
}
func main(){

	readFile()

	fmt.Println("Here is the initial board:\n\n")
	printBoard()
	start := time.Now()
	solved := solve(0,0)
	elapsed :=time.Since(start)	

	if solved{
		fmt.Println("\nSolved!\n\n")
		printBoard()
		fmt.Println("\n The solver took", elapsed)
	} else{
		fmt.Println("\nCould not be solved!")
	}
	
}
