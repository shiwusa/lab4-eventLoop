package engine

import "sync"

type commandsQueue struct {
	sync.Mutex
	commands []Command
}

func (cq *commandsQueue) push(cmd Command) {
	cq.Lock()
	defer cq.Unlock()
	cq.commands = append(cq.commands, cmd)
}

func (cq *commandsQueue) pull() Command {
	cq.Lock()
	defer cq.Unlock()
	res := cq.commands[0]
	cq.commands[0] = nil
	cq.commands = cq.commands[1:]
	return res
}

func (cq *commandsQueue) length() int {
	cq.Lock()
	defer cq.Unlock()
	return len(cq.commands)
}
