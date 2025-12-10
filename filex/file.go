package filex

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-leo/gox/stringx"
)

func Primary(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// List is used to list the paths of all regular files under the specified root directory.
func List(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}

func Extension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}

// Download 下载文件到指定路径
// 参数：
// url: 要下载的文件的URL
// filepath: 保存文件的本地路径
func Download(url, filepath string) error {
	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// 写盘
	if err = out.Sync(); err != nil {
		return err
	}
	return nil
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
	if err = os.Chmod(dst, os.FileMode(0o777)); err != nil {
		return
	}
	return
}

// CopyDir  拷贝整个目录
func CopyDir(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// 目标路径
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)
		if dir.IsDir() {
			// 创建目标目录
			return os.MkdirAll(dstPath, dir.Type())
		} else {
			// 复制文件
			return CopyFile(path, dstPath)
		}
	})
}

// IsDir reports whether the named file is a directory.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsDirectory reports whether the named file is a directory.
var IsDirectory = IsDir

// GetSize 获取文件大小
func GetSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

const (
	Byte         int64 = 1
	Kilobyte           = 1024 * Byte
	Megabyte           = 1024 * Kilobyte
	Gigabyte           = 1024 * Megabyte
	Trillionbyte       = 1024 * Gigabyte
	Petabyte           = 1024 * Trillionbyte
	Exabyte            = 1024 * Petabyte
	// Zettabyte          = 1024 * Exabyte
	// Yottabyte          = 1024 * Zettabyte
	// Brontobyte         = 1024 * Yottabyte
)

func HumanReadableSize(size int64) string {
	s := size
	if s < 0 {
		s = -s
	}
	builder := stringx.Builder{}
	if s >= Exabyte {
		eb := s / Exabyte
		s = s % Exabyte
		_ = builder.WriteInt(eb, 10)
		_, _ = builder.WriteString("EB")
	}

	if s >= Petabyte {
		pb := s / Petabyte
		s = s % Petabyte
		_ = builder.WriteInt(pb, 10)
		_, _ = builder.WriteString("PB")
	}

	if s >= Trillionbyte {
		tb := s / Trillionbyte
		s = s % Trillionbyte
		_ = builder.WriteInt(tb, 10)
		_, _ = builder.WriteString("TB")
	}

	if s >= Gigabyte {
		gb := s / Gigabyte
		s = s % Gigabyte
		_ = builder.WriteInt(gb, 10)
		_, _ = builder.WriteString("GB")
	}

	if s >= Megabyte {
		mb := s / Megabyte
		s = s % Megabyte
		_ = builder.WriteInt(mb, 10)
		_, _ = builder.WriteString("MB")
	}

	if s >= Kilobyte {
		kb := s / Kilobyte
		s = s % Kilobyte
		_ = builder.WriteInt(kb, 10)
		_, _ = builder.WriteString("KB")
	}

	if s >= Byte {
		b := s / Byte
		// s = s % Byte
		_ = builder.WriteInt(b, 10)
		_, _ = builder.WriteString("B")
	}
	return builder.String()
}
