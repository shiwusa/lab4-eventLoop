package engine

import "fmt"

type PrintComm struct {
	Arg string
}

func (printComm *PrintComm) Execute(handler Handler) {
	fmt.Println(printComm.Arg)
}
