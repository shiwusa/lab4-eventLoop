package engine

import (
	"fmt"
	"strings"
)

func Parse(textCommand string) Command {
	split := strings.Fields(textCommand)
	if split[0] == "print" {
		if len(split) != 2 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			command := PrintComm{
				Arg: split[1],
			}
			return &command
		}
	} else if split[0] == "sha1" {
		if len(split) != 2 {
			return InvalidArgsCount(split[0], 1, len(split)-1)
		} else {
			command := Sha1Comm{
				Arg: split[1],
			}
			return &command
		}
	} else {
		command := PrintComm{
			Arg: fmt.Sprintf("Invalid command \"%s\": command does not exist", split[0]),
		}
		return &command
	}
}

func InvalidArgsCount(command string, desired int, actual int) Command {
	result := &PrintComm{
		Arg: fmt.Sprintf("Invalid command \"%s\": Expected %d arguments, found %d", command, desired, actual),
	}
	return result
}
