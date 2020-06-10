package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var puzzle [9][9]int

type blankCell struct {
	row            int
	col            int
	possibleDigits []int
}

func main() {
	start := time.Now()

	loadPuzzle()
	fillInSingleCandidate()
	solve()
	output()

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
	log.Printf("Puzzle solved: %t", isSolved())
}

// Load the puzzle from an external file into a 9x9 array
func loadPuzzle() {
	fmt.Println(os.Args)
	if len(os.Args) <= 1 {
		log.Fatalf("USAGE : %s <target_filename> \n", os.Args[0])
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		col := 0
		for _, v := range line {
			num, _ := strconv.Atoi(v)
			puzzle[row][col] = num
			col++
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Fill in the blank cell which has only one possible digit
func fillInSingleCandidate() {
	canIterate := false

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if puzzle[row][col] == 0 {
				cell := findPossibleDigits(row, col)
				if 1 == len(cell.possibleDigits) {
					puzzle[row][col] = cell.possibleDigits[0]
					canIterate = true
				}
			}
		}
	}

	if canIterate {
		fillInSingleCandidate()
	}
}

// Try and fill in the blank cell which has multiple digit candidates
func solve() {
	fill(0)
}

func fill(counter int) bool {
	if 81 == counter {
		return true
	}

	row := counter / 9
	col := counter % 9
	if 0 == puzzle[row][col] {
		cell := findPossibleDigits(row, col)
		for _, v := range cell.possibleDigits {
			puzzle[row][col] = v
			if fill(counter + 1) {
				return true
			}
			puzzle[row][col] = 0
		}
	} else {
		return fill(counter + 1)
	}

	return false
}

// Output the final result
func output() {
	fmt.Println("-------------------")
	for _, v := range puzzle {
		var result strings.Builder
		result.WriteString("|")
		result.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), "|"), "[]"))
		result.WriteString("|")
		fmt.Println(result.String())
		fmt.Println("-------------------")
	}
}

func isSolved() bool {
	for _, row := range puzzle {
		for _, col := range row {
			if col == 0 {
				return false
			}
		}
	}
	return true
}

func findPossibleDigits(row int, col int) blankCell {
	var possibleDigits []int

NumLoop:
	for num := 1; num < 10; num++ {
		// Check if num exists in the cells of the same row or column.
		for i := 0; i < 9; i++ {
			if num == puzzle[row][i] || num == puzzle[i][col] {
				continue NumLoop
			}
		}
		// Check if num exists in the 3×3 block.
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if num == puzzle[row/3*3+i][col/3*3+j] {
					continue NumLoop
				}
			}
		}
		possibleDigits = append(possibleDigits, num)
	}

	cell := blankCell{row, col, possibleDigits}
	return cell
}
