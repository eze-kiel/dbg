package main

import "github.com/eze-kiel/dbg"

func main() {
	// Point()
	dbg.Point()
	dbg.Point("with values", 13)

	// Printf
	dbg.Printf("%s is equal to %d\n", "my_val", 1337)

	// Break
	dbg.Break()

	// Halt
	dbg.Halt()

}
