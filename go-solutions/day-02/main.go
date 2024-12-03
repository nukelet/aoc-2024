package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sign(n int) int {
	if n >= 0 {
		return 1
	} else {
		return -1
	}
}

func abs(a int) int {
	return sign(a) * a
}

func toIntSlice(numbers []string) []int {
	res := make([]int, len(numbers))
	for i, s := range(numbers) {
		// inshallah this will convert successfully :pray:
		res[i], _ = strconv.Atoi(s)
	}

	return res
}

func isValidDiff(n int, sign int) bool {
	return n * sign > 0 && abs(n) <= 3
}

func checkSafety(numbers []int) (bool, int) {
	sign := sign(numbers[1] - numbers[0])

	safe := true
	idx := 0
	for i := 0; i < len(numbers) - 1; i++ {
		if !isValidDiff(numbers[i+1] - numbers[i], sign) {
			safe = false
			idx = i
			break
		}
	}

	return safe, idx
}

func IsSafe(numbers []int, dampener bool) bool {
	safe, _ := checkSafety(numbers)

	// wow this is ugly
	if !safe && dampener {
		for i := 0; i < len(numbers); i++ {
			test := make([]int, len(numbers))
			_ = copy(test, numbers)
			test = slices.Delete(test, i, i+1)
			safe, _ = checkSafety(test)
			if safe {
				return safe

			}
		}
	}

	return safe
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	safeCount := 0
	for s.Scan() {
		line := s.Text()

		numberStrings := strings.Split(line, " ")
		numbers := toIntSlice(numberStrings)

		if IsSafe(numbers, true) {
			safeCount++
			continue
		}
	}

	fmt.Println(safeCount)
}
