package engine

import (
	"crypto/sha1"
)

type Sha1Comm struct {
	Arg string
}

func (sha1Comm *Sha1Comm) Execute(handler Handler) {
	res := sha1.Sum([]byte(sha1Comm.Arg))
	str := string(res[:])
	handler.Post(&PrintComm{Arg: str})
}
