package helpers

import (
    "os"
    "path/filepath"
    "io/ioutil"
)

// IsFile 判断文件是否存在
func IsFile(file string) bool {
    f, e := os.Stat(file)
    if e != nil {
        return false
    }
    return !f.IsDir()
}

// IsDir 判断目录是否存在
func IsDir(file string) bool {
    f, e := os.Stat(file)
    if e != nil {
        return false
    }
    return f.IsDir()
}

// Dirname 获取目录名称
func Dirname(file string) string {
    return filepath.Dir(file)
}

// MakedirIfNotExists 创建一个目录，当目录不存在时
func MakedirIfNotExists(path string, perm os.FileMode) bool {
    if !IsDir(path) {
        err := os.MkdirAll(path, perm)
        if err != nil {
            return false
        }
    }
    return true
}

// WriteFile 写入内容到文件
func WriteFile(filename string, content string, perm os.FileMode) error {
    return ioutil.WriteFile(filename, []byte(content), perm)
}
