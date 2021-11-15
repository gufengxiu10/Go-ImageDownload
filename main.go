package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 map[string]int = make(map[string]int)

	map1["c"] = 200
	map1["a"] = 100

	var stringSlice []string = make([]string, 0)

	for index := range map1 {
		stringSlice = append(stringSlice, index)
	}

	sort.Slice(stringSlice, func(i, j int) bool {
		return i > j
	})

	fmt.Println(stringSlice)
}
