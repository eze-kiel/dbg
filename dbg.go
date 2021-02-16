package dbg

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	. "github.com/logrusorgru/aurora"
)

type point struct {
	pc       uintptr
	line     int
	funcName string
	pkgName  string
	fileName string
	filePath string

	mem runtime.MemStats
}

// Point displays a debug entry at a given point. It can also display a value
// provided in argument and its type
func Point(a ...interface{}) {
	pt := createPoint(2)
	if len(a) > 0 {
		for _, item := range a {
			fmt.Fprintf(os.Stderr, "(pkg:%s|func:%s) %s:%d\t %v (%s)\n",
				Bold(Cyan(pt.pkgName)),
				Bold(Cyan(pt.funcName)),
				pt.fileName,
				BrightMagenta(pt.line),
				Bold(item),
				Green(reflect.TypeOf(item)))
		}
	} else {
		fmt.Fprintf(os.Stderr, "(pkg:%s|func:%s) %s:%d\n",
			Bold(Cyan(pt.pkgName)),
			Bold(Cyan(pt.funcName)),
			pt.fileName,
			BrightMagenta(pt.line))
	}
}

// Printf emulates the fmt.Printf version, but add a dbg prefix
func Printf(format string, a ...interface{}) {
	pt := createPoint(2)
	format = fmt.Sprintf("(pkg:%s|func:%s) %s:%d\t ",
		Bold(Cyan(pt.pkgName)),
		Bold(Cyan(pt.funcName)),
		pt.fileName,
		BrightMagenta(pt.line),
	) + format
	fmt.Fprintf(os.Stderr, format, a...)
}

// Halt exits the program
func Halt(code int) {
	pt := createPoint(2)
	format := fmt.Sprintf("(pkg:%s|func:%s) %s:%d\t %s",
		Bold(Cyan(pt.pkgName)),
		Bold(Cyan(pt.funcName)),
		pt.fileName,
		BrightMagenta(pt.line),
		Bold(Red("Halted")),
	)
	fmt.Fprintln(os.Stderr, format)
	os.Exit(code)
}

// Mem show memory usage at a point
func Mem() {
	pt := createPoint(2)
	format := fmt.Sprintf("(pkg:%s|func:%s) %s:%d\tAlloc= %vMiB  TotalAlloc= %vMiB  Sys= %vMiB  NumGC= %v",
		Bold(Cyan(pt.pkgName)),
		Bold(Cyan(pt.funcName)),
		pt.fileName,
		BrightMagenta(pt.line),
		bToMb(pt.mem.Alloc),
		bToMb(pt.mem.TotalAlloc),
		bToMb(pt.mem.Sys),
		pt.mem.NumGC,
	)
	fmt.Fprintln(os.Stderr, format)
}

func createPoint(skip int) point {
	var pt point
	var ok bool
	runtime.ReadMemStats(&pt.mem)

	pt.pc, pt.filePath, pt.line, ok = runtime.Caller(skip)
	if !ok {
		log.Fatal("not ok") // change me pls
	}
	_, pt.fileName = path.Split(pt.filePath)
	parts := strings.Split(runtime.FuncForPC(pt.pc).Name(), ".")
	pl := len(parts)
	pt.pkgName = ""
	pt.funcName = parts[pl-1]

	if parts[pl-2][0] == '(' {
		pt.funcName = parts[pl-2] + "." + pt.funcName
		pt.pkgName = strings.Join(parts[0:pl-2], ".")
	} else {
		pt.pkgName = strings.Join(parts[0:pl-1], ".")
	}

	return pt
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
