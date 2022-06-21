package engine

type Loop struct {
	q          *commandsQueue
	stop       bool //stop sequest
	stopSignal chan struct{}
}

func (l *Loop) Start() {
	l.q = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})

	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) {
	if !l.stop {
		l.q.push(cmd)
	}
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*Loop).stop = true
}

func (l *Loop) AwaitFinish() {
	l.Post(stopCommand{})
	<-l.stopSignal
}
