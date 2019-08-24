package main

import (
	"fmt"
	//ds "github.com/nidnogg/go-cracking/dstructs" // imports data structures as ds
)

func main() {
	var strs = []string{"abcd", "abcc", "word1", "microsoft", "abigail", "true", "false", "111"}
	fmt.Println("Printing strings:")
	for _, str := range strs {
		fmt.Printf("%v: isUniqueRaw=%v\n", str, isUniqueRaw(str))
	}
}

func isUnique(s string) (isUni bool) {
	m := make(map[string]bool)
	isUni = true // Boolean used for checking if a character is contained and also the return value
	
	// Maps each character as a key in a hashmap, with a true value for each
	// isUni is set to false if the key already has a true bool associated with it
	for _, char := range(s) {
		if _, ok := m[string(char)]; ok {
			isUni = false
		}
		m[string(char)] = true
	}

	return
} 

func isUniqueRaw(s string) (isUni bool) {
	// Iterates over each character in the string, and for each character checks the entire string for it 
	isUni = true
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++	{
			if(i != j && string(s[i]) == string(s[j])) {
				isUni = false
			}
		}
	}

	return
}
