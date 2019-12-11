package job

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

const (
	// DefaultTaskFile is the file in the docker container that contains the
	// task definition. It should be in the root of the container's file
	// system.
	DefaultTaskFile = "/aegir.yaml"
)

// Task is the definition of a single parallel task. It is a command to execute
// within the docker container.
type Task struct {
	Command   string   `yaml:"Command"`
	Arguments []string `yaml:"Arguments"`
}

// ReadFrom reads the task definition from a file.
func (task *Task) ReadFrom(fileName string) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return fmt.Errorf("Cannot read '%s': %v", fileName, err)
	}

	err = yaml.Unmarshal(data, &task)
	if err != nil {
		return fmt.Errorf("Cannot parse YAML from '%s': %v", fileName, err)
	}

	return nil
}

// Execute runs the task and waits for it to complete.
func (task *Task) Execute() error {
	// #nosec G204
	cmd := exec.Command(task.Command, task.Arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		code := cmd.ProcessState.ExitCode()
		return fmt.Errorf("Command exited with code '%d': %v", code, err)
	}

	return nil
}

// MakeTask creates a new Task, read from the DefaultTaskFile.
func MakeTask() (*Task, error) {
	var task Task

	err := task.ReadFrom(DefaultTaskFile)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
