package engine

import "sync"

type commandsQueue struct {
	mu       sync.Mutex
	commands []Command
	wait     bool

	notEmpty chan struct{}
}

func (cq *commandsQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.commands = append(cq.commands, c)
	if cq.wait {
		cq.notEmpty <- struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.commands) == 0 {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.commands[0]
	cq.commands[0] = nil
	cq.commands = cq.commands[1:]
	return res
}

func (cq *commandsQueue) empty() bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.commands) == 0
}
