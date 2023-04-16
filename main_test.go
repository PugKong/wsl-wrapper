package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

type CommandStub struct {
	name string
	args []string
	run  func(name string, args ...string) error
}

func (s *CommandStub) Run() error {
	return s.run(s.name, s.args...)
}

func Example_main() {
	args := os.Args
	defer func() {
		_ = os.Remove("docker.log")
		_ = os.Remove("git.log")
		os.Args = args
	}()

	executable = func() (string, error) {
		return os.Args[0], nil
	}
	command = func(name string, args ...string) Runnable {
		return &CommandStub{
			name: name,
			args: args,
			run: func(name string, args ...string) error {
				output := strings.Join(append([]string{name}, args...), " ")
				fmt.Println(output)

				return nil
			},
		}
	}

	os.Args = []string{
		"docker.exe", "compose",
		"-f", "\\\\wsl$\\Ubuntu-22.04\\home\\pugkong\\docker-compose.yaml",
		"-f", "C:\\Users\\Pug Kong\\JetBrains\\docker-compose.override.2033.yml",
		"run", "--rm", "container",
	}
	main()

	os.Args = []string{"git.exe", "add", "src\\index.js"}
	main()

	// Output:
	// wsl docker compose -f /home/pugkong/docker-compose.yaml -f /mnt/c/Users/Pug Kong/JetBrains/docker-compose.override.2033.yml run --rm container
	// wsl git add src/index.js
}

func TestNewCommand(t *testing.T) {
	name := "git6610"
	args := []string{"add", "src/index.js"}
	runnable := NewCommand(name, args...)

	cmd, ok := runnable.(*exec.Cmd)
	if !ok {
		t.Fatal("cmd is not *exec.Cmd")
	}

	if cmd.Path != name {
		t.Errorf("cmd.Path = %v, want %v", cmd.Path, name)
	}

	args = append([]string{name}, args...)
	if !reflect.DeepEqual(cmd.Args, args) {
		t.Errorf("cmd.Args = %v, want %v", cmd.Args, args)
	}

	if cmd.Stdin != os.Stdin {
		t.Error("cmd.Stdin is not os.Stdin")
	}

	if cmd.Stdout != os.Stdout {
		t.Error("cmd.Stdout is not os.Stdout")
	}

	if cmd.Stderr != os.Stderr {
		t.Error("cmd.Stderr is not os.Stderr")
	}
}
