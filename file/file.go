package file

import (
	"bufio"
	"git.medlinker.com/wanghouwei/autoAddCommit/util"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	Name string
	Dir  string
}

func CopyFile(to, from string) (err error) {
	fsrc, err := os.Open(from)
	if err != nil {
		return
	}
	defer fsrc.Close()
	p := path.Dir(to)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777) //0777也可以os.ModePerm
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
	s, err := GetCurrentPath()
	if err != nil {
		util.Errorf("getCurrentPath err (%+v)", err)
		panic("getCurrentPath")
		return
	}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if !strings.Contains(path, ".git") &&
				!strings.Contains(path, ".DS_Store") &&
				!strings.Contains(path, ".idea") {
				result = append(result, File{
					Dir:  path,
					Name: info.Name(),
					Path: s + "/mainFile/" + path[len(dir)+1:],
				})
			}
		}
		return nil
	})
	return
}

func GetCurrentPath() (string, error) {
	getwd, err := os.Getwd()
	if err != nil {
		util.Errorf("Getwd err (%+v)", err)
		return "", err
	}
	return getwd, nil
}

func CreateFile(to string) {
	p := path.Dir(to)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777) //0777也可以os.ModePerm
		os.Chmod(p, 0777)
	}
	fdest, err := os.Create(to)
	if err != nil {
		util.Errorf("Create file err (%+v)", err)
		return
	}
	defer fdest.Close()
}
