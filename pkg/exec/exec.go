// Package exec contains an interface for executing commands, along with helpers
package exec

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/gosoon/glog"
)

// Cmd abstracts over running a command somewhere, this is useful for testing
type Cmd interface {
	Run() error
	// Each entry should be of the form "key=value"
	SetEnv(...string) Cmd
	SetStdin(io.Reader) Cmd
	SetStdout(io.Writer) Cmd
	SetStderr(io.Writer) Cmd
}

// Cmder abstracts over creating commands
type Cmder interface {
	// command, args..., just like os/exec.Cmd
	Command(string, ...string) Cmd
}

// DefaultCmder is a LocalCmder instance used for convienience, packages
// originally using os/exec.Command can instead use pkg/kind/exec.Command
// which forwards to this instance
// TODO(bentheelder): swap this for testing
// TODO(bentheelder): consider not using a global for this :^)
var DefaultCmder = &LocalCmder{}

// Command is a convience wrapper over DefaultCmder.Command
func Command(command string, args ...string) Cmd {
	return DefaultCmder.Command(command, args...)
}

// CombinedOutputLines is like os/exec's cmd.CombinedOutput(),
// but over our Cmd interface, and instead of returning the byte buffer of
// stderr + stdout, it scans these for lines and returns a slice of output lines
func CombinedOutputLines(cmd Cmd) (lines []string, err error) {
	var buff bytes.Buffer
	cmd.SetStdout(&buff)
	cmd.SetStderr(&buff)
	err = cmd.Run()
	scanner := bufio.NewScanner(&buff)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

// OutputLines is like os/exec's cmd.Output(),
// but over our Cmd interface, and instead of returning the byte buffer of
// stdout, it scans these for lines and returns a slice of output lines
func OutputLines(cmd Cmd) (lines []string, err error) {
	var buff bytes.Buffer
	cmd.SetStdout(&buff)
	err = cmd.Run()
	scanner := bufio.NewScanner(&buff)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

// InheritOutput sets cmd's output to write to the current process's stdout and stderr
func InheritOutput(cmd Cmd) Cmd {
	cmd.SetStderr(os.Stderr)
	cmd.SetStdout(os.Stdout)
	return cmd
}

// RunLoggingOutputOnFail runs the cmd, logging error output if Run returns an error
func RunLoggingOutputOnFail(cmd Cmd) error {
	var buff bytes.Buffer
	cmd.SetStdout(&buff)
	cmd.SetStderr(&buff)
	err := cmd.Run()
	if err != nil {
		glog.Error("failed with:")
		scanner := bufio.NewScanner(&buff)
		for scanner.Scan() {
			glog.Error(scanner.Text())
		}
	}
	return err
}

// RunWithStdoutReader runs cmd with stdout piped to readerFunc
func RunWithStdoutReader(cmd Cmd, readerFunc func(io.Reader) error) error {
	pr, pw, err := os.Pipe()
	if err != nil {
		return err
	}
	defer pw.Close()
	defer pr.Close()
	cmd.SetStdout(pw)

	errChan := make(chan error, 1)
	go func() {
		errChan <- readerFunc(pr)
		pr.Close()
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}
	err2 := <-errChan
	if err2 != nil {
		return err2
	}
	return nil
}

// RunWithStdinWriter runs cmd with writerFunc piped to stdin
func RunWithStdinWriter(cmd Cmd, writerFunc func(io.Writer) error) error {
	pr, pw, err := os.Pipe()
	if err != nil {
		return err
	}
	defer pw.Close()
	defer pr.Close()
	cmd.SetStdin(pr)

	errChan := make(chan error, 1)
	go func() {
		errChan <- writerFunc(pw)
		pw.Close()
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}
	err2 := <-errChan
	if err2 != nil {
		return err2
	}
	return nil
}