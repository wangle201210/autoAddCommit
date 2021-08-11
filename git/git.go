package git

import (
	"github.com/wangle201210/autoAddCommit/file"
	"github.com/wangle201210/autoAddCommit/util"
)

var branch string

func Run() {
	if err := getBranch(); err != nil {
		return
	}
	if err := addFile(); err != nil {
		return
	}
	if err := gitPush("./", branch); err != nil {
		return
	}
}

func addFile() (err error){
	from := "/Users/med/mine/goPkgLearn/color/color.go"
	to := "/Users/med/mine/github/autoAddCommit/color/color.go"
	err = file.CopyFile(to,from)
	if err != nil {
		util.Errorf("CopyFile err (%+v)", err)
		return
	}
	util.Infof("addFile (%s) to %s", from, to)
	return
}

// 提交修改内容到git
func gitPush(medSdkDir, branch string) (err error) {
	util.Infof("正在提交代码...")
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
		util.Infof("无改动，无提交动作")
		return
	}
	util.Infof("提交完成")
	return
}

// 获取当前分支名
func getBranch() (err error) {
	util.Infof("正在获取当前分支名...")
	branch, err = util.RunCmdRet("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		util.Errorf("getBranch err (%+v)", err)
		return
	}
	util.Infof("当前分支为: %s", branch)
	return
}