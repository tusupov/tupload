package tupload

import (
	"mime/multipart"
	"strings"
	"os"
	"io"
	"time"
	"fmt"
	"math/rand"
)

func moveFile(header multipart.FileHeader, path string, ext string) (fileName string, err error) {

	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return
	}
	defer file.Close()

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}

	for {
		fileName = generateFileName(ext)
		if !fileExists(path + "/" + fileName) {
			break
		}
	}

	var f *os.File
	f, err = os.OpenFile(path + "/" + fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		os.Remove(f.Name())
		return
	}

	return

}

func saveFile(reader io.Reader, path string, ext string) (fileName string, err error) {

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}

	for {
		fileName = generateFileName(ext)
		if !fileExists(path + "/" + fileName) {
			break
		}
	}

	var f *os.File
	f, err = os.OpenFile(path + "/" + fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, reader)
	if err != nil {
		os.Remove(f.Name())
	}

	return

}

func generateFileName(ext string) string {
	return fmt.Sprintf("%d_%06d.%s", time.Now().UnixNano(), rand.Intn(1e6), ext)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func doMatch(mime string, matchers map[string]string) (bool bool, extension string) {

	for ext, match := range matchers {
		if strings.ToLower(mime) == strings.ToLower(match) {
			return true, ext
		}
	}

	return

}
