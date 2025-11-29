package main

import "fmt"

func main() {

	//PART 1: Variable Declaration + Initialization
	var e11, e12, e13, e21, e22, e23, e31, e32, e33 string
	var row1, row2, row3, col1, col2, col3, cross1, cross2 string
	var slot float32
	var finished bool = false
	var empty bool
	var row, col, i int
	var answer string = "X"
	e11 = "_"; e12 = "_"; e13 = "_"; e21 = "_"; e22 = "_"; e23 = "_"; e31 = "_"; e32 = "_"; e33 = "_"

	//PART 2: Initial Board Output (Empty)
	fmt.Print("TicTacToe\nPlayer X goes first.\n\n")
	fmt.Println(e11, e12, e13)
	fmt.Println(e21, e22, e23)
	fmt.Println(e31, e32, e33)
	fmt.Print("\nInput a decimal number (1.3 for row 1 column 3): ")

	//PART 3: For Loop + While Loop for Game Ending
	for i = 1; i <= 9 && !finished; i++ {

		//PART 4: Reset Variables + Input for Slot
		empty = true
		_, err := fmt.Scan(&slot)
		row = int(slot)
		col = int(slot*10.0) % 10
		for row < 1 || row > 3 || col < 1 || col > 3 || err != nil {
			fmt.Print("Invalid slot, try again: ")
			_, err = fmt.Scan(&slot)
			row = int(slot)
			col = int(slot*10.0) % 10}

		//PART 5: Changing Empty Board Elements via Input
		if row == 1 {
			if col == 1 {
				if e11 != "_" {empty = false} else {e11 = answer}
			} else if col == 2 {
				if e12 != "_" {empty = false} else {e12 = answer}
			} else if col == 3 {
				if e13 != "_" {empty = false} else {e13 = answer}}
		} else if row == 2 {
			if col == 1 {
				if e21 != "_" {empty = false} else {e21 = answer}
			} else if col == 2 {
				if e22 != "_" {empty = false} else {e22 = answer}
			} else if col == 3 {
				if e23 != "_" {empty = false} else {e23 = answer}}
		} else if row == 3 {
			if col == 1 {
				if e31 != "_" {empty = false} else {e31 = answer}
			} else if col == 2 {
				if e32 != "_" {empty = false} else {e32 = answer}
			} else if col == 3 {
				if e33 != "_" {empty = false} else {e33 = answer}}}
		
		//PART 6: Initializing Variables for Winning Cases
		row1 = e11 + e12 + e13
		row2 = e21 + e22 + e23
		row3 = e31 + e32 + e33
		col1 = e11 + e21 + e31
		col2 = e12 + e22 + e32
		col3 = e13 + e23 + e33
		cross1 = e11 + e22 + e33
		cross2 = e13 + e22 + e31

		//PART 7: Reprinting Board for Clarity
		fmt.Println()
		fmt.Println(e11, e12, e13)
		fmt.Println(e21, e22, e23)
		fmt.Println(e31, e32, e33)

		//PART 8: Changing Players if Slotted Successfully
		if !empty {i = i - 1} else {if answer == "X" {answer = "O"} else {answer = "X"}}

		//PART 8: Checking All Winning Cases + Prompt Input if Not Finished
		if row1 == "XXX" || row1 == "OOO" {finished = true
		} else if row2 == "XXX" || row2 == "OOO" {finished = true
		} else if row3 == "XXX" || row3 == "OOO" {finished = true
		} else if col1 == "XXX" || col1 == "OOO" {finished = true
		} else if col2 == "XXX" || col2 == "OOO" {finished = true
		} else if col3 == "XXX" || col3 == "OOO" {finished = true
		} else if cross1 == "XXX" || cross1 == "OOO" {finished = true
		} else if cross2 == "XXX" || cross2 == "OOO" {finished = true
		} else if i < 9 {fmt.Print("\nInput slot for ", answer, ": ")}}

	//PART 9: Output Game Results
	if finished {
		if answer == "X" {answer = "O"} else {answer = "X"}
		fmt.Print("\nPlayer ", answer, " wins!")
	} else {fmt.Print("\nIt's a draw!")}}

