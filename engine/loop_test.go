package engine

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func stdOutString() func() (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	done := make(chan error, 1)

	save := os.Stdout
	os.Stdout = w

	var buf strings.Builder

	go func() {
		_, err := io.Copy(&buf, r)
		r.Close()
		done <- err
	}()

	return func() (string, error) {
		os.Stdout = save
		w.Close()
		err := <-done
		return buf.String(), err
	}
}

func TestExecutionPrintCommand(t *testing.T) {
	assert := assert.New(t)
	loop := new(Loop)
	input := "this is testing message"
	expected := "this is testing message\n"
	cmd := PrintComm{input}

	loop.Start()
	getStr := stdOutString()
	loop.Post(&cmd)
	loop.AwaitFinish()

	capturedOutput, err := getStr()
	if err != nil {
		panic(err)
	}

	assert.Equal(capturedOutput, expected)
}

func TestExecutionShaCommand(t *testing.T) {
	assert := assert.New(t)
	loop := new(Loop)
	inputStr := "hello"
	expected := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d\n"
	cmd := Sha1Comm{inputStr}
	loop.Start()
	getStr := stdOutString()
	loop.Post(&cmd)
	loop.AwaitFinish()

	capturedOutput, err := getStr()
	if err != nil {
		panic(err)
	}

	assert.Equal(capturedOutput, expected)
}
