package main

import "fmt"

func main() {
	s := "Старт"

	m := map[int32]int{}
	funcName(s, m)
	for k, v := range m {
		fmt.Println(string(k), "-", v)
	}
}

func funcName(s string, m map[int32]int) {
	for _, c := range s {
		m[c]++
	}
}
