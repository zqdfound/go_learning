package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	rPath := "./1.txt"
	wPath := "./22222.txt"
	if err := copyFile(rPath, wPath); err != nil {
		t.Errorf(err.Error())

	}
}

func copyFile(srcPath, dstPath string) error {
	var r, w *os.File
	var err error
	if r, w, err = openBoth(srcPath, dstPath); err != nil {
		return err
	}
	defer func() {
		r.Close()
		w.Close()
		if err != nil {
			os.Remove(dstPath)
		}
	}()
	if _, err := io.Copy(r, w); err != nil {
		return fmt.Errorf("copy failed %s %s,%v", srcPath, dstPath, err)
	}
	return nil
}

func openBoth(srcPath, dstPath string) (*os.File, *os.File, error) {
	var r, w *os.File
	var err error
	if r, err = os.Open(srcPath); err != nil {
		return nil, nil, fmt.Errorf("open failed %s %s,%v", srcPath, dstPath, err)
	}
	if w, err = os.Create(dstPath); err != nil {
		r.Close()
		return nil, nil, fmt.Errorf("create failed %s %s,%v", srcPath, dstPath, err)
	}
	return r, w, nil
}
