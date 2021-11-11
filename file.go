//遍历目录文件
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WalkAllFilesInDir(dir string) ([]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(dir, func(fpath string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			files[info.Name()] = fpath
		}
		return e
	})
	if err != nil {
		return nil, err
	}
	list := make([]string, 0)
	for _, fpath := range files {
		list = append(list, fpath)
	}
	return list, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func CreateDir() {
	//使用mkdirall
	os.MkdirAll("/xxx/xxx/xxx", 0755)
}

// //go:embed something.exe
var exe []byte

func ExtractExe(name string, embedExe []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(embedExe)
	if err != nil {
		return err
	}
	return nil
}

func CheckFileSha256(fileName, sha string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return false, err
	}
	return hex.EncodeToString(hash.Sum(nil)) == sha, nil
}

func main() {
	input := "foo  bar   baz\n123456789"
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
