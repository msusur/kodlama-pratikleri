package main

import (
	"fmt"
	"strings"
)

func findDuplicates(str string) map[string]int {
	hashMap := make(map[string]int)
	var s string

	for _, val := range str {
		s = strings.ToLower(string(val))
		if s != " " {
			if _, ok := hashMap[s]; ok {
				hashMap[s]++
			} else {
				hashMap[s] = 1
			}
		}
	}
	return hashMap
}

func main() {
	cumle := "Bir cumle icerisinde bir harften ne kadar olduğunu bulan program"
	fmt.Println(findDuplicates(cumle))
}
