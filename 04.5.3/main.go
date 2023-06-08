package main

import (
	"fmt"
	. "slerm/04.5.3/gometr"
	"strconv"
)

func main() {
	checks := &Checker{}

	for i := 0; i < 5; i++ {
		checks.Add([]Checkable{
			NewGoMetrClient(strconv.Itoa(i), i*3+1),
		})
	}

	fmt.Println(checks)
	checks.Check()
}
