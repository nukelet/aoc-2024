package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var list_1, list_2 []int

	// stores the element counts for the second list
	count := make(map[int]int)

	for {
		var a, b int

		_, err := fmt.Scanf("%d %d\n", &a, &b)
		if err == io.EOF {
			break
		}

		// store the numbers in the lists
		list_1 = append(list_1, a)
		list_2 = append(list_2, b)
		
		// store the occurrence counts in a map, taking advantage
		// of the fact that the map returns zero upon lookup of
		// a key that doesn't exist yet
		count[b] += 1
	}

	sort.Sort(sort.IntSlice(list_1))
	sort.Sort(sort.IntSlice(list_2))

	// calculate the "distance" between the lists
	dist := 0
	for i := 0; i < len(list_1); i += 1 {
		if list_1[i] >= list_2[i] {
			dist += list_1[i] - list_2[i]
		} else {
			dist += list_2[i] - list_1[i]
		}
	}

	fmt.Printf("dist = %d\n", dist)

	// calculate the "similarity" between the lists
	similarity := 0
	for _, element := range(list_1) {
		similarity += element * count[element]
	}

	fmt.Printf("similarity = %d\n", similarity)

}
