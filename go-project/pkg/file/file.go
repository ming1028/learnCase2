package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(str string) bool {
	_, err := os.Stat(str)
	return os.IsNotExist(err)
}

func CheckPermission(str string) bool {
	_, err := os.Stat(str)
	return os.IsPermission(err)
}

func IsNotExistMkDir(str string) error {
	if notExist := CheckNotExist(str); notExist == true {
		if err := MkDir(str); err != nil {
			return err
		}
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm) // 创建多级目录
	if err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
