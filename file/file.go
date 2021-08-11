package file

import (
	"bufio"
	"os"
)

func CopyFile(dest, src string) (err error) {
	fsrc, err := os.Open(src)
	if err != nil {
		return
	}
	defer fsrc.Close()
	fdest, err := os.Create(dest)
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
