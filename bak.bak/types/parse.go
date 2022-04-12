package types

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type (
	gocode struct {
		name     string
		filepath string
		numLines int
		lines    []string
	}

	arg struct {
		v reflect.Value
		k reflect.Kind
	}

	Arg interface{}

	Function interface {
		Name() string
		Func() Any
		Args() []Arg
	}

	GoCode interface {
		Functions() []Function
	}
)

func GetGoFileList() (files []string, err error) {
	return GetFileListBySuffix(".", ".go")
}

func GetFileListBySuffix(path, suffix string) (files []string, err error) {
	return getFileList(path, suffix, "", "", false)
}

func GetFileListByPrefix(path, prefix string) (files []string, err error) {
	return getFileList(path, "", prefix, "", false)
}

func GetFileListByName(path, needle string) (files []string, err error) {
	return getFileList(path, "", "", needle, false)
}

func GetFileListWithDirectories(path string) (files []string, err error) {
	return getFileList(path, "", "", "", true)
}

func getFileList(path, suffix, prefix, needle string, isDir bool) (files []string, err error) {
	if path == "" {
		path, _ = os.Getwd()
	}
	a, _ := filepath.Abs(path)

	err = filepath.WalkDir(a, func(path string, info fs.DirEntry, err error) error {

		if err == nil {
			if !isDir && info.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, suffix) {
				if strings.HasPrefix(path, prefix) {
					if strings.Contains(path, needle) {
						files = append(files, path)
					}
				}
			}
		}
		return err
	})
	return
}
