package filepathx

import (
	"os"
	"path/filepath"
)

// Resolve 将路径或路径片段的序列解析为绝对路径。
func Resolve(paths ...string) (string, error) {
	// 获取当前工作目录作为解析的起点
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 将路径片段连接起来，并规范化路径
	resolvedPath := filepath.Join(paths...)
	resolvedPath = filepath.Clean(resolvedPath)

	// 如果路径已经为绝对路径，则直接返回
	if filepath.IsAbs(resolvedPath) {
		return resolvedPath, nil
	}

	// 否则，将当前工作目录与相对路径拼接并规范化
	return filepath.Clean(filepath.Join(cwd, resolvedPath)), nil
}
