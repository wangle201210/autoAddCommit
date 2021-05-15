package main

import (
	"git.medlinker.com/wanghouwei/autoAddCommit/git"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	git.Run()
}

