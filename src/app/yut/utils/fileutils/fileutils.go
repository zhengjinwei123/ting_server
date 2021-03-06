package fileutils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func RmFile(path string) error {
	return os.Remove(path)
}

func FilePutContents(file_path_name string, contents string) error {
	err := ioutil.WriteFile(file_path_name, []byte(contents), 0644)
	return err
}

func FileGetContent(file_path_name string) (string, error) {
	contents, err := ioutil.ReadFile(file_path_name);
	if err != nil {
		return "", err
	}

	return string(contents), err
}

func FileExists(file_path_name string) bool {
	var exists = true
	if _, err := os.Stat(file_path_name); os.IsNotExist(err) {
		exists = false
	}
	return exists
}

func CreateDir(dir_path string) bool {
	err := os.MkdirAll(dir_path, os.ModePerm)
	return err == nil
}

func DeleteFile(file_path string) error {
	if FileExists(file_path) {
		err := os.Remove(file_path)
		return err
	}
	return errors.New("file not exists")
}

func GetDirFiles(dirPth string, suffix string)  (files []string, err error) {
	files = make([]string, 0)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	//PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			//files = append(files, dirPth+PthSep+fi.Name())
			files = append(files, fi.Name())
		}
	}

	return files, nil
}