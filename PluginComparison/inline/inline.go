package main

import (
	"fmt"
	"time"
)

type Command struct {
	X int
	Y int
}

type Response struct {
	V int
}

const limit = 10000000

func main() {
	start := time.Now()

	for i := 0; i < limit; i++ {
		_ = i + limit - i
	}

	total := time.Since(start)
	nanosPerOp := total.Nanoseconds() / limit

	fmt.Println("Nanoseconds per operation: ", nanosPerOp)
}
