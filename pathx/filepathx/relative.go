package filepathx

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Relative 方法根据当前工作目录返回从 from 到 to 的相对路径。
func Relative(from string, to string) (string, error) {

	// 将路径规范化并确保它们为绝对路径
	from = filepath.Clean(filepath.FromSlash(from))
	to = filepath.Clean(filepath.FromSlash(to))

	// 如果任一路径不是绝对路径，返回错误
	if !filepath.IsAbs(from) || !filepath.IsAbs(to) {
		return "", fmt.Errorf("both from and to paths must be absolute")
	}

	// 分离路径的目录部分和最后一部分（文件名或目录名）
	fromDir, _ := filepath.Split(from)
	toDir, toBase := filepath.Split(to)

	// 计算两路径间的公共目录部分
	commonPrefix := CommonDir(fromDir, toDir)

	// 计算从共同基路径到“from”路径的向上回退（..）计数
	numBack := CountParentDirs(fromDir[len(commonPrefix):])

	// 构建相对路径：首先添加必要的“..”，然后是目标路径的剩余部分
	var relative string
	for i := 0; i < numBack; i++ {
		relative += "../"
	}
	relative += toDir[len(commonPrefix):] + "/" + toBase

	// 如果结果为空字符串，说明两者位于同一目录下，直接返回目标基名即可
	if relative == "" {
		relative = toBase
	}

	return relative, nil
}

// CommonDir 找到两个路径中公共的目录部分
func CommonDir(a, b string) string {
	shorter := a
	longer := b
	if len(a) > len(b) {
		shorter = b
		longer = a
	}

	i := 0
	for ; i < len(shorter); i++ {
		if shorter[i] != longer[i] {
			break
		}
	}
	return shorter[:i]
}

// CountParentDirs 计算路径中“..”的数量
func CountParentDirs(path string) int {
	count := 0
	for _, part := range strings.Split(path, "/") {
		if part == ".." {
			count++
		}
	}
	return count
}
