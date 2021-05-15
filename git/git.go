package git

import (
	"fmt"
	"git.medlinker.com/wanghouwei/autoAddCommit/file"
	"git.medlinker.com/wanghouwei/autoAddCommit/util"
	"math/rand"
	"time"
)

var (
	branch string
	files []file.File
	startTime time.Time
)

func Run() {
	now := time.Now()
	startTime = now.Add(time.Second * -1 * 60 * 60 * 24 * 30 * 3)
	if err := getBranch(); err != nil {
		util.Errorf("getBranch err (%+v)", err)
		return
	}
	for i := 0; startTime.Unix() < now.Unix() && i < 5; i++ {
		util.Infof("第 %d 次开始", i)
		getTime()
		f, err := addFile();
		if err != nil {
			return
		}
		if err := gitPush("./", branch, f); err != nil {
			return
		}
	}
	//if err := changeTime(); err != nil {
	//	return
	//}
}

func addFile() (f file.File ,err error) {
	f = getFile()
	from := f.Dir
	to :=  f.Path
	err = file.CopyFile(to,from)
	if err != nil {
		util.Errorf("CopyFile err (%+v)", err)
		return
	}
	util.Infof("addFile (%s) to %s", from, to)
	return
}

// 提交修改内容到git
func gitPush(medSdkDir, branch string, f file.File) (err error) {
	util.Infof("正在提交代码...")
	err = util.RunCmdCD(medSdkDir, "git", "add", "-A")
	if err != nil {
		return
	}
	var gitStatus string
	gitStatus, _ = util.RunCmdRetCD(medSdkDir, "git", "status", "--porcelain")
	if gitStatus != "" {
		msg := commitMsg(f)
		msgString := fmt.Sprintf("--message=%s", msg)
		date := startTime.Unix()
		dateString := fmt.Sprintf("--date=%d", date)
		err = util.RunCmdCD(medSdkDir, "git", "commit", msgString, dateString)
		if err != nil {
			util.Errorf("git commit err (%+v)", err)
			return
		}
	} else {
		util.Infof("无改动文件")
	}
	err = util.RunCmdCD(medSdkDir, "git", "push", "-f", "origin", branch)
	if err != nil {
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

func getCommitID() (commit string, err error) {
	util.Infof("正在获取当前CommitId...")
	commit, err = util.RunCmdRet("git", "rev-parse", "HEAD")
	if err != nil {
		util.Errorf("getCommitID err (%+v)", err)
		return
	}
	util.Infof("当前CommitId为: %s", commit)
	return
}

func changeTime() (err error) {
	util.Infof("开始修改时间")
	id, err := getCommitID()
	if err != nil {
		return
	}
	gad := "Fri Jan 2 21:38:53 2009 -0800"
	gcd := "Sat May 19 01:01:01 2007 -0700"
	cmd := fmt.Sprintf(`\
		'if [ $GIT_COMMIT = %s ]
		then
		export GIT_AUTHOR_DATE="%s"
		export GIT_COMMITTER_DATE="%s"
		fi'
	`, id, gad, gcd)
	err = util.RunCmd("git", "filter-branch", "-f", "--env-filter", cmd)
	if err != nil {
		util.Errorf("getCommitID err (%+v)", err)
		return
	}
	util.Infof("修改时间完成")
	return
}

func commitMsg(f file.File) (msg string) {
	// todo 添加文件名字
	l := len(util.DoString)
	randNum := rand.Int() % l
	if randNum > 0 {
		randNum--
	}
	s := util.DoString[randNum]
	msg = s + " " + f.Name
	return
}

func getFile() (f file.File) {
	dir := "/Users/med/mine/goPkgLearn/web"
	l := len(files)
	if l == 0 {
		files = file.GetFiles(dir)
		if len(files) == 0 {
			util.Errorf("no file in %s", dir)
			panic("没文件啊")
			return
		}
		return getFile()
	}
	f = files[0]
	if l > 1 {
		files = files[1:]
	}
	return
}

func getTime () {
	// 3 天内必有一次提交, 提交间隔不低于1000s
	randNum := rand.Int63n(60 * 60 * 24 * 3)
	if randNum < 1000 {
		getTime()
		return
	}
	startTime = startTime.Add(time.Second * time.Duration(randNum))
}
