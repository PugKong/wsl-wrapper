package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	executable Executable = os.Executable
	command    Command    = NewCommand
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {
	var err error
	var exe string
	exe, err = executable()
	if err != nil {
		log.Panicln(err)
	}

	dir := filepath.Dir(exe)
	app := strings.TrimSuffix(filepath.Base(exe), filepath.Ext(exe))

	var logFile *os.File
	logFile, err = os.OpenFile(filepath.Join(dir, app+".log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		_ = logFile.Close()
	}()
	log.SetOutput(logFile)

	cmd := NewCommandWrapper(command)
	if err := cmd.Run(app, os.Args[1:]...); err != nil {
		log.Panicln(err)
	}
}

type Executable func() (string, error)

func NewCommand(name string, args ...string) Runnable {
	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
