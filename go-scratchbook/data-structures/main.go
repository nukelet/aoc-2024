package main

import "fmt"

func matmul(a [][]int, b [][]int) [][]int {
	m_a, m_b := len(a), len(b)
	n_a, n_b := len(a[0]), len(b[0])

	if n_a != m_b {
		return nil
	}
	
	// take advantage of the fact that empty stuff
	// is initialized to 0
	var c [][]int
	fmt.Println("before:", c)
	for range(m_a) {
		fmt.Println("appending ohhh im appending")
		c = append(c, make([]int, n_b))
	}

	fmt.Println("c ->", c)

	for i := range(m_a) {
		for j := range(n_b) {
			for k := range(m_b) {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func main() {
	a := [][]int{
		{1, 0, 1},
		{2, 1, 1},
		{0, 1, 1},
		{1, 1, 2},
	}

	b := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{4, 2, 2},
	}

	c := matmul(a, b)
	fmt.Println(c)

}
