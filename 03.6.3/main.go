package main

import "fmt"

func main() {
	s := []string{"abc", "aba"}

	fmt.Println(funcName(s))
}

func funcName(s []string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
