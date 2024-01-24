package filex

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// Create 可以新建文件/目录
// 注: 需要外部调用 Close 进行关闭
func Create(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if !IsExist(dir) { // 判断下目录是否存在, 不存在就先创建目录
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}

// CopyFile 复制文件
func CopyFile(src, dst string, isFirstDel ...bool) (err error) {
	if src == "" {
		return errors.New("source file cannot be empty")
	}

	if dst == "" {
		return errors.New("destination file cannot be empty")
	}

	// 如果相同就不处理
	if src == dst {
		return nil
	}

	// 删除原来的
	if len(isFirstDel) > 0 && isFirstDel[0] {
		if err = os.Remove(dst); err != nil {
			return
		}
	}

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		if e := in.Close(); e != nil {
			err = e
		}
	}()

	out, err := Create(dst)

	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	// 复制
	if _, err = io.Copy(out, in); err != nil {
		return
	}

	// 写盘
	if err = out.Sync(); err != nil {
		return
	}

	// 调整权限
	if err = os.Chmod(dst, os.FileMode(0777)); err != nil {
		return
	}
	return
}
