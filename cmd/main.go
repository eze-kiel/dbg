package main

import (
	"time"

	"github.com/eze-kiel/dbg"
)

func main() {
	// Point()
	dbg.Point()
	dbg.Point("with values", 13)

	// Printf
	dbg.Printf("%s is equal to %d\n", "my_val", 1337)

	// Mem
	dbg.Mem()

	var overall [][]int
	for i := 0; i < 4; i++ {
		a := make([]int, 0, 999999)
		overall = append(overall, a)

		dbg.Mem()
		time.Sleep(time.Second)
	}
	overall = nil

	// Halt
	dbg.Halt(127)
}
