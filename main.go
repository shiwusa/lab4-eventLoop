package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/shiwusa/lab4-eventLoop/engine"
)

func main() {
	inputFile := "example.txt"
	eventLoop := new(Loop)
	eventLoop.Start()

	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := Parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	} else {
		eventLoop.Post(&PrintComm{
			Arg: fmt.Sprintf("Error opening file \"%s\"", err.Error()),
		})
	}
	eventLoop.AwaitFinish()
}
