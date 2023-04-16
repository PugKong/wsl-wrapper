package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	wslPath = regexp.MustCompile(`^(\\\\wsl\$\\[\w-.]+)(.*)$`)
	winPath = regexp.MustCompile(`^([A-Z]):\\(.*)$`)
)

type Runnable interface {
	Run() error
}

type Command func(name string, args ...string) Runnable

type CommandWrapper struct {
	command Command
}

func NewCommandWrapper(command Command) *CommandWrapper {
	return &CommandWrapper{command}
}

func (c *CommandWrapper) Run(name string, args ...string) error {
	args = append([]string{name}, args...)
	log.Printf("Original args: %+v", args)
	for i, arg := range args {
		if wslPath.MatchString(arg) {
			arg = wslPath.ReplaceAllString(arg, "$2")
			args[i] = filepath.ToSlash(arg)
		} else if winPath.MatchString(arg) {
			matches := winPath.FindStringSubmatch(arg)
			arg = fmt.Sprintf("/mnt/%s/%s", strings.ToLower(matches[1]), matches[2])
			args[i] = filepath.ToSlash(arg)
		} else if filepath.IsLocal(arg) {
			args[i] = filepath.ToSlash(arg)
		}
	}
	log.Printf("Modified args: %+v", args)

	cmd := c.command("wsl", args...)
	return cmd.Run()
}
