package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePrintCmd(t *testing.T) {
	assert := assert.New(t)
	inputStr := "print hello"
	expected := "hello"

	cmd := Parse(inputStr)

	assert.Equal(expected, cmd.(*PrintComm).Arg)
}

func TestParseSha1Cmd(t *testing.T) {
	assert := assert.New(t)
	inputStr := "sha1 hello"
	expectedStr := "hello"
	cmd := Parse(inputStr)

	assert.Equal(expectedStr, cmd.(*Sha1Comm).Arg)
}
