package rk

import (
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
)

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Intersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if Contains(b, el) {
			set = append(set, el)
		}
	}
	return set
}

func Contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

func WalkDir(root string, filename string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, filename) {
			files = append(files, path)
			return nil
		}
		return nil
	})
	return files, err
}
