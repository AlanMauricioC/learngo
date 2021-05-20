package main

import (
	"fmt"
	"runtime"

	"github.com/AlanMauricioC/learngo/printer"
)

func main() {
	printer.Hello()
	fmt.Println(runtime.Version())
}
