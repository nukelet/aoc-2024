package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"slices"
)

func validateSequence(seq []string, prec map[string][]string, fix bool) bool {
	// tags if we can still use a page
	invalid := make(map[string]bool)
	// map: 'X' -> stack containing the index of the
	// element that invalidated page 'X'
	backlist := make(map[string][]int)
	for _, page := range seq {
		if invalid[page] && !fix {
			return false
		}

		for _, p := range prec[page] {
			invalid[p] = true
		}
	}

	return true
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	
	// prec["X"] -> all pages that *must* preceed "X"
	prec := make(map[string][]string)

	for s.Scan() {
		line := s.Text()

		// end of first section
		if len(line) == 0 {
			break
		}

		pages := strings.Split(line, "|")
		before := pages[0]
		after := pages[1]
		prec[after] = append(prec[after], before)
	}

	shouldFix := true
	sum := 0
	for s.Scan() {
		line := s.Text()
		seq := strings.Split(line, ",")

		if validateSequence(seq, prec, shouldFix) {
			m := len(seq)/2
			mid, _ := strconv.Atoi(seq[m])
			sum += mid
		}
	}

	fmt.Println(sum)
}
