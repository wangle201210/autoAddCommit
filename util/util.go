package util

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var flagVerbose bool

// RunCmd runs the cmd & print output (both stdout & stderr)
func RunCmd(name string, arg ...string) (err error) {
	return RunCmdCD("", name, arg...)
}

func RunCmdCD(cd, name string, arg ...string) (err error) {
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
	if flagVerbose {
		logf("CMD: %s", strings.Join(append([]string{name}, arg...), " "))
	}
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
	color.Green(format, args...)
}

func Errorf(format string, args ...interface{}) {
	color.Red(format, args...)
}

