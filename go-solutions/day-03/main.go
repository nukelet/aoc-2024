package main

import (
	"fmt"
	"regexp"
	"io"
	"os"
	"strings"
	"strconv"
)

func evalMult(s string) int {
	s = strings.TrimPrefix(s, "mul(")
	s = strings.TrimSuffix(s, ")")
	vals := strings.Split(s, ",")

	a, _ := strconv.Atoi(vals[0])
	b, _ := strconv.Atoi(vals[1])

	return a * b
}

func firstPart(input string) int {
	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")	

	mults := r.FindAllString(input, -1)

	result := 0
	for _, s := range(mults) {
		result += evalMult(s)
	}

	return result
}

func secondPart(input string) int {
	r, _ := regexp.Compile("(mul\\([0-9]+,[0-9]+\\))|(do\\(\\))|(don't\\(\\))")

	ops := r.FindAllString(input, -1)

	result := 0
	shouldEval := true
	for _, s := range(ops) {
		if strings.HasPrefix(s, "mul") {
			if shouldEval {
				result += evalMult(s)
			}
		} else if strings.HasPrefix(s, "do()") {
			shouldEval = true
		} else if strings.HasPrefix(s, "don't()") {
			shouldEval = false
		}
	}

	return result
}


func main() {
	// we can ReadAll from stdin bro trust me
	buf, _ := io.ReadAll(os.Stdin)
	input := string(buf)

	result := firstPart(input)
	fmt.Println(result)

	result = secondPart(input)
	fmt.Println(result)
}
