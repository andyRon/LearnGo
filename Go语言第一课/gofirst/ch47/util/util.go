package util

import (
	"os"
	"path"
	"path/filepath"
)

// 获得项目根目录
func GetProjectRoot() string {
	binDir, err := executableDir()
	if err != nil {
		return ""
	}
	return path.Dir(binDir)
}

// 获得可执行程序所在目录
func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}
