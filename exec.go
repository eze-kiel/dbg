package dbg

import (
	"fmt"
	"os"
	"strings"
)

func (pt point) exec(cmd string) {
	cmds := strings.Split(cmd, " ")
	switch cmds[0] {
	default:
		fmt.Fprintf(os.Stderr, "unknown: %s\n", cmd)
	case "pc":
		fmt.Printf("program counter: %d\n", pt.pc)
	case "line", "l":
		fmt.Printf("%s:%d\n", pt.fileName, pt.line)
	case "file", "f":
		fmt.Printf("%s\n", pt.filePath)
	case "caller", "c":
		fmt.Printf("%s\n", pt.funcName)
	case "package", "pkg", "p":
		fmt.Printf("%s\n", pt.funcName)
	case "get":
		pt.get(cmds[1:])
	}
}

func (pt point) get(cmds []string) {
	if len(cmds) == 0 {
		fmt.Fprintln(os.Stdout, "missing command")
		return
	}
	switch cmds[0] {
	default:
		fmt.Printf("unknown: get %s\n", cmds[0])
	case "mem":
		fmt.Printf("Heap Allocated %v MiB\n", bToMb(pt.mem.Alloc))
		fmt.Printf("Total Heap Alloc %v MiB\n", bToMb(pt.mem.TotalAlloc))
		fmt.Printf("Obtained from sys %v MiB\n", bToMb(pt.mem.Sys))
		fmt.Printf("GC cycles %v\n", pt.mem.NumGC)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
