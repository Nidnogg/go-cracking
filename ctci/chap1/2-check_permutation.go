package main

import (
	"fmt";
	"sort"
)

func main() {
	var strs  = []string{"abca", "aabc"}
	if checkPermutation(strs) {
		fmt.Printf("checkPermutation() returned true\n")
	} else {
		fmt.Printf("checkPermutation() returned false\n")
	}
}

func checkPermutation(strs []string) (isPerm bool) {
	isPerm = false
	var aux []rune

	// Sort both strings, first converting them to slices for use with the sort package
	for i := 0; i < len(strs); i++ {
		aux = strToR(strs[i])
		sort.Slice(aux, func (j int, k int) bool { 
			return aux[j] < aux[k] 
		})
		strs[i] = string(aux)
		fmt.Printf("strs[%v]: %v\n", i, strs[i])
	}
	if strs[0] == strs[1] { isPerm = true}
	
	return 
}


// Converts string to []rune slice
func strToR(str string) (converted []rune) {
	for _, char := range(str) {
		converted = append(converted, char)
	}
	return
}

