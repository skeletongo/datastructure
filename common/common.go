package common

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
)

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func NewDir(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	if err := os.Mkdir(path, 0666); err != nil {
		return err
	}
	return nil
}

func NewImg(filename string, root interface{}, f func(w io.Writer, root interface{}) error) error {
	if IsNil(root) || filename == "" {
		return nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dirSvg := path.Join(dir, "img")
	dirDot := path.Join(dirSvg, "dot")
	if err = NewDir(dirSvg); err != nil {
		return err
	}
	if err = NewDir(dirDot); err != nil {
		return err
	}

	svgFile := fmt.Sprintf("%s.svg", path.Join(dirSvg, filename))
	dotFile := fmt.Sprintf("%s.dot", path.Join(dirDot, filename))
	w, err := os.Create(dotFile)
	if err != nil {
		return err
	}
	defer w.Close()
	if err = f(w, root); err != nil {
		return err
	}
	return exec.Command("dot", dotFile, "-Tsvg", "-o", svgFile).Run()
}
