package job

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTask(t *testing.T) {
	assert := assert.New(t)

	var task Task

	tempFile, err := ioutil.TempFile("", "aegirtest")
	if err != nil {
		panic(err)
	}

	fileName := tempFile.Name()
	defer os.Remove(fileName)

	_, err = tempFile.WriteString(`Command: echo`)
	if err != nil {
		panic(err)
	}

	err = tempFile.Close()
	if err != nil {
		panic(err)
	}

	// Test it!
	err = task.ReadFrom(fileName)
	assert.NoError(err, "ReadFrom should not error but did: %v", err)
	assert.Equal("echo", task.Command, "Command read incorrectly")

	// Test reading a non-existent task file
	err = task.ReadFrom("THIS_FILE_DOES_NOT_EXIST")
	assert.Error(err, "ReadFrom must error for file that doesn't exist!")
}

func TestExecuteTask(t *testing.T) {
	assert := assert.New(t)

	var task Task

	// Test a command that should exist
	task.Command = "go"
	task.Arguments = []string{"version"}
	err := task.Execute()
	assert.NoError(err, "It should have executed but didn't: %v", err)

	// Test a command that doesn't exist
	task.Command = "aegirtestbleh"
	err = task.Execute()
	assert.Error(err, "There should be an error when the command doesn't exist")
}
