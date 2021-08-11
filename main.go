package main

import (
	"fmt"
	"git.medlinker.com/wanghouwei/autoAddCommit/git"
	"git.medlinker.com/wanghouwei/autoAddCommit/util"
	"math/rand"
	"time"
)

var (
	dir string
)

func main() {
	fmt.Println("请输入用作上传的文件的路径: ")
	fmt.Scanln(&dir)
	util.Infof("路径为:%s", dir)
	rand.Seed(time.Now().Unix())
	git.Run(dir)
}
