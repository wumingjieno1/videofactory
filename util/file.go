package util

import "path/filepath"

// 获取文件夹下所有文件
func GetDirFiles(dir string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
	return files, nil
}
