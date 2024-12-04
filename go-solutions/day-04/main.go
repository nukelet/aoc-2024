package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * 	This is a pretty run-of-the-mill recursive search implementation
 */
func visitRec(matrix [][]byte, str string, i, j, di, dj int) bool {
	// we're out of bounds
	if i < 0 || i >= len(matrix) ||
	   j < 0 || j >= len(matrix[0]) {
		return false
	}
	
	// wrong character
	if str[0] != matrix[i][j] {
		return false
	}

	// we found the last character of the word
	if len(str) == 1 {
		return true
	}

	return visitRec(matrix, str[1:], i + di, j + dj, di, dj)
}

func visit(matrix [][]byte, str string, i, j int) int {
	count := 0

	directions := [][2]int {
		{-1, 0}, 	// left	
		{1, 0}, 	// right
		{0, -1}, 	// up
		{0, 1}, 	// down
		{1, -1}, 	// up-right diag
		{-1, -1}, 	// up-left diag
		{-1, 1}, 	// down-left diag
		{1, 1}, 	// down-right diag
	}

	for _, d := range directions {
		if visitRec(matrix, str, i, j, d[0], d[1]) {
			count += 1
		}
	}

	return count
}

func findOcurrences(matrix [][]byte, str string) int {
	count := 0

	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			count += visit(matrix, str, i, j)
		}
	}

	return count
}

/*
 * 	For future reference: the idea here is that for every X-MAS,
 * 	the two MASes share an 'A' (this is an iff condition); so we
 * 	look for diagonal MASes as usual and store the coordinates for
 * 	each 'A' we see in a map, and after we're done searching we
 * 	count the map entries that have a count of 2 (i.e. 'A's shared
 * 	by two MASes, i.e. an X-MAS)
 */
func visit2(matrix [][]byte, str string, i, j int, centerCoords map [[2]int]int) {
	directions := [][2]int {
		{1, -1}, 	// up-right diag
		{-1, -1}, 	// up-left diag
		{-1, 1}, 	// down-left diag
		{1, 1}, 	// down-right diag
	}

	for _, d := range directions {
		if visitRec(matrix, str, i, j, d[0], d[1]) {
			// store the coordinates (m,n) of the 'A' letter
			// (it always comes directly after the 'M')
			m := i + d[0]
			n := j + d[1]
			// taking advantage of the new key getting initialized to zero
			centerCoords[[2]int{m,n}] += 1
		}
	}
}

func findOcurrences2(matrix [][]byte, str string) int {
	centerCoords := make(map [[2]int]int)

	// look for all diagonal MASes
	for i := range len(matrix)  {
		for j := range len(matrix[0]) {
			visit2(matrix, str, i, j, centerCoords)
		}
	}

	// count all the 'A's shared by two MASes
	count := 0
	for _, v := range centerCoords {
		if v == 2 {
			count += 1
		}
	}

	return count
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var matrix [][]byte

	for s.Scan() {
		line := s.Text()
		bytes := []byte(line)
		matrix = append(matrix, bytes)
	}

	count := findOcurrences(matrix, "XMAS")
	fmt.Printf("count: %d\n", count)

	count = findOcurrences2(matrix, "MAS")
	fmt.Printf("count: %d\n", count)
}
