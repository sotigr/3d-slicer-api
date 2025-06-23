package cli

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\w|\%|^(\+|,|\.|:|\/|\\|=|_|-|[a-z]|[A-Z]|[0-9]|\s)+$`)

func validateCliArgument(arg string) bool {
	return re.Match([]byte(arg))
}

func validateCliArguments(args []string) bool {
	for _, p := range args {
		if !validateCliArgument(p) {
			return false
		}
	}
	return true
}

func execNoCheck(ctx context.Context, command string, args []string) (string, string, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	cmdStr := cmd.String()

	pStdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", cmdStr, err
	}

	pStderr, err := cmd.StderrPipe()
	if err != nil {
		return "", cmdStr, err
	}

	err = cmd.Start()
	if err != nil {
		return "", cmdStr, err
	}

	stdout, _ := io.ReadAll(pStdout)
	stderr, _ := io.ReadAll(pStderr)

	err = cmd.Wait()
	if err != nil {
		return "", cmdStr, errors.New(err.Error() + " | " + string(stderr) + " | " + string(stdout))
	}
	return string(stdout), cmdStr, nil
}

type Pipe struct {
	Command string
	Args    []string
}

func (p *Pipe) IsSanatory() bool {
	return validateCliArguments(p.Args) && validateCliArgument(p.Command)
}

func (p *Pipe) Execute(ctx context.Context) (string, string, error) {
	if !p.IsSanatory() {
		return "", "", errors.New("failed to execute cli command: " + p.Command + " " + strings.Join(p.Args, " ") + " some arguments are not allowed")
	} else {
		return execNoCheck(ctx, p.Command, p.Args)
	}
}

func ExecuteCLIPipeLine(ctx context.Context, pipes []Pipe) (string, string, error) {

	pipeLen := len(pipes)
	if pipeLen == 0 {
		return "", "", errors.New("failed to execute cli command: no pipes provided")
	}
	argLen := 0
	for _, p := range pipes {
		if !p.IsSanatory() {
			return "", "", errors.New("failed to execute cli command: " + p.Command + " " + strings.Join(p.Args, " ") + " some arguments are not allowed")
		}
		argLen += len(p.Args) + 1
	}

	command := "/bin/sh"
	var args []string = []string{}
	args = append(args, "-c")
	pipeLen -= 1

	cmdArg := ""

	for i, p := range pipes {
		cmdArg += p.Command + " " + strings.Join(p.Args, " ")
		if i != pipeLen {
			cmdArg += " | "
		}
	}

	args = append(args, cmdArg)

	return execNoCheck(ctx, command, args)
}
