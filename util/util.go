package util

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os/exec"
	"strings"
	"time"
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

func JustRun(cd, name string, arg ...string) (err error) {
	logf("CMD: %s", strings.Join(append([]string{name}, arg...), " "))
	cmd := exec.Command(name, arg...)
	cmd.Dir = cd
	err = cmd.Run()
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

//二倍均值算法,count剩余个数,amount剩余时间
// 间隔在 5小时 到 3天之间
func DoubleAverage(count, total int) int {
	if count == 1 {
		//返回剩余时间
		return total
	}
	// 提交间隔至少 5h
	min := 60 * 60 * 5
	//计算最大可用时间,min最小是1分钱,减去的min,下面会加上,避免出现0分钱
	max := total - min*count
	//计算最大可用平均值
	avg := max / count
	//二倍均值基础加上最小时间,防止0出现,作为上限
	avg2 := 2*avg + min
	//随机commit时间序列元素,把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	//加min是为了避免出现0值,上面也减去了min
	x := rand.Intn(avg2) + min
	if x > 60*60*24*3 {
		return DoubleAverage(count, total)
	}
	return x
}
