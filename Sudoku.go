
package main

import (
  "fmt"
  "io/ioutil"

)

var board = "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000" +
	    "000000000"



func readFile() {

	byteStr, err := ioutil.ReadFile("C:\\Users\\Elena\\OneDrive\\Documents\\AAAVanderbilt\\Fall2016\\ProgrammingLanguages\\Project4\\sudoku.txt") 
	if err != nil {
        	fmt.Println(err)
	}

	str := string(byteStr) 
	board = str

}

func printBoard(board string) {
	for row, index := 0, 0; row < 9; row, index = row + 1, index + 9 {
		fmt.Printf("%c %c %c | %c %c %c | %c %c %c\n" , 
		board[index], board[index+1], board[index+2], board[index+3],
		board[index+4], board[index+5], board[index+6], board[index+7], 
		board[index+8])
		if row == 2 || row == 5 {
			fmt.Println("------+-------+------")
		} 
	}

}


func main(){

	readFile()

	fmt.Println("Here is the initial board:\n\n")
	printBoard(board)

	


}


