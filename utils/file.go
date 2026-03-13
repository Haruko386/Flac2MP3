package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CheckFileExists 检查文件是否存在
func CheckFileExists(filePath string) error {
	info, err := os.Stat(filePath)
	// 检测文件是否存在
	if os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filePath)
	}
	// 检测路径是否是一个文件夹
	if info.IsDir() {
		return fmt.Errorf("%s is a directory", filePath)
	}
	// 无错误，直接返回
	return nil
}

// GetOutputPath 根据输入文件生成输出文件的路径(后缀替换)
func GetOutputPath(inputFile string) string {
	ext := filepath.Ext(inputFile)
	return strings.TrimSuffix(inputFile, ext) + ".mp3"
}
