package engine

type Loop struct {
		sync.Mutex
		q          *commandsQueue
		stop       bool
		busy       bool
		stopSignal chan bool
	}

	func (l *Loop) Post(cmd Command) {
		l.Lock()
		defer l.Unlock()
		l.q.push(cmd)
		if l.busy && !l.stop {
			l.startRoutine()
		}
	}

	func (l *Loop) startRoutine() {
		l.busy = false
		go func() {
			for {
				if l.q.length() > 0 {
					cmd := l.q.pull()
					cmd.Execute(l)
				} else if l.stop {
					l.stopSignal <- true
					return
				} else {
					l.Lock()
					defer l.Unlock()
					l.busy = true
					return
				}
			}
		}()
	}

	func (l *Loop) Start() {
		l.stopSignal = make(chan bool, 1)
		l.q = &commandsQueue{}
		l.startRoutine()
	}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*Loop).stop = true
}

func (l *Loop) AwaitFinish() {
	l.Post(stopCommand{})
	<-l.stopSignal
}
