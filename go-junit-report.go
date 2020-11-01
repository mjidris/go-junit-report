package main

import (
	"flag"
	"fmt"
)

var (
	noXMLHeader   = flag.Bool("no-xml-header", false, "do not print xml header")
	packageName   = flag.String("package-name", "", "specify a package name (compiled test have no package name in output)")
	goVersionFlag = flag.String("go-version", "", "specify the value to use for the go.version property in the generated XML")
	setExitCode   = flag.Bool("set-exit-code", false, "set exit code to 1 if tests failed")
	printEvents   = flag.Bool("print-events", false, "print events (for debugging)")
)

func main() {
	fmt.Println("main is a wip.")
}
