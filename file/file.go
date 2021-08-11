package file

import (
	"bufio"
	"os"
	"path"
)

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

