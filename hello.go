package main

import (
	"fmt"
)

func hello(somebody string) {
	if len(somebody) == 0 {
		somebody = "world"
	}

	fmt.Println("hello ", somebody)
}
