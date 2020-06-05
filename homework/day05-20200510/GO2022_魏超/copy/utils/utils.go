package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcbuf := bufio.NewScanner(srcFile)
	dstbuf := bufio.NewWriter(dstFile)
	defer dstbuf.Flush()
	for srcbuf.Scan() {
		context := srcbuf.Bytes()
		context = append(context, '\n')
		_, err := dstbuf.Write(context)
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyDir(src, dst string) error {
	err := os.Mkdir(dst, os.ModePerm)
	if err != nil {
		return err
	}
	// 检查源
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	files, err := srcFile.Readdir(-1)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			err = CopyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
