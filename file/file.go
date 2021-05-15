package file

import (
	"bufio"
	"errors"
	"git.medlinker.com/wanghouwei/autoAddCommit/util"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	Name string
	Dir string
}

func CopyFile(to, from string) (err error) {
	fsrc, err := os.Open(from)
	if err != nil {
		return
	}
	defer fsrc.Close()
	p := path.Dir(to)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, 0777) //0777也可以os.ModePerm
		os.Chmod(p, 0777)

	}
	fdest, err := os.Create(to)
	if err != nil {
		return
	}
	defer fdest.Close()
	fileScanner := bufio.NewScanner(fsrc)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		fdest.WriteString(text + "\n")
	}
	return
}

// files all files with suffix
func GetFiles(dir string) (result []File) {
	s, err := getCurrentPath()
	if err != nil {
		util.Errorf("getCurrentPath err (%+v)", err)
		panic("getCurrentPath")
		return
	}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			result = append(result, File{
				Dir: path,
				Name: info.Name(),
				Path: s + "mainFile/" + path[len(dir)+1:],
			})
		}
		return nil
	})
	return
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}