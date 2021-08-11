package util

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

var (
	DoString = []string{
		"add", "set", "edit", "alert", "modify", "addition", "increase", "cut", "del", "fix",
	}
)

var flagVerbose bool

// RunCmd runs the cmd & print output (both stdout & stderr)
func RunCmd(name string, arg ...string) (err error) {
	return RunCmdCD("", name, arg...)
}

func RunCmdCtx(ctx context.Context, name string, arg ...string) error {
	if out, err := exec.CommandContext(ctx, name, arg...).CombinedOutput(); err != nil {
		Errorf("RunCmdWithCtx: %s, 提示为: %s", strings.Join(append([]string{name}, arg...), " "), out)
		return err
	}
	return nil
}

func RunCmdCD(cd, name string, arg ...string) (err error) {
	logf("RunCmdCD: %s", strings.Join(append([]string{cd, name}, arg...), " "))
	cmd := exec.Command(name, arg...)
	cmd.Dir = cd
	out, err := cmd.CombinedOutput()
	if flagVerbose || err != nil {
		logFunc := logf
		if err != nil {
			logFunc = Errorf
		}
		logFunc("CMD: %s", strings.Join(append([]string{cd, name}, arg...), " "))
		logFunc(string(out))
	}
	return
}

func RunCmdRet(name string, arg ...string) (out string, err error) {
	return RunCmdRetCD("", name, arg...)
}

func RunCmdRetCD(cd, name string, arg ...string) (out string, err error) {
	logf("CMD: %s", strings.Join(append([]string{name}, arg...), " "))
	cmd := exec.Command(name, arg...)
	cmd.Dir = cd
	outBytes, err := cmd.CombinedOutput()
	out = strings.Trim(string(outBytes), "\n\r\t ")
	return
}

func logf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func Infof(format string, args ...interface{}) {
	color.Green(format+"\n", args...)
}

func Errorf(format string, args ...interface{}) {
	color.Red(format+"\n", args...)
}
