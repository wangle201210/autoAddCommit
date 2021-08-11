package git

import (
	"github.com/wangle201210/autoAddCommit/file"
	"github.com/wangle201210/autoAddCommit/util"
)

func Run() {
	addFile()
	gitPush("./","master")
}

func addFile() {
	if err := file.CopyFile("/Users/med/mine/github/autoAddCommit/color.go","/Users/med/mine/goPkgLearn/color/color.go"); err != nil {
		println(err.Error())
	}
}

// 提交修改内容到git
func gitPush(medSdkDir, branch string) (err error) {
	err = util.RunCmdCD(medSdkDir, "git", "add", "-A")
	if err != nil {
		return
	}
	var gitStatus string
	gitStatus, _ = util.RunCmdRetCD(medSdkDir, "git", "status", "--porcelain")
	if gitStatus != "" {
		err = util.RunCmdCD(medSdkDir, "git", "commit", "-m", "update from local")
		if err != nil {
			return
		}
		err = util.RunCmdCD(medSdkDir, "git", "push", "-f", "origin", branch)
		if err != nil {
			return
		}
	} else {
		util.Infof("无改动，无上传动作\n")
	}
	return
}
