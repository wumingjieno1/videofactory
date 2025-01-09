package model

import (
	"archive/zip"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/wumingjieno1/videofactory/util"
)

type Video struct {
	Title        string       `json:"title"`
	FilesDir     string       `json:"files_dir"`
	DefaultVoice string       `json:"default_voice"`
	Header       []*ImageInfo `json:"header"`
	Content      []*ImageInfo `json:"content"`
	End          []*ImageInfo `json:"end"`
}

func (v *Video) Save(output string) (err error) {
	werr := util.EmptyWErr("videofile.Save")
	defer func() {
		err = werr.Err(err).ToError()
	}()
	//  创建output目录
	err = os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		return err
	}
	// 创建zip文件
	zipFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 保存video数据
	videoData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// 创建并写入video.json
	videoWriter, err := zipWriter.Create("video.json")
	if err != nil {
		return err
	}
	_, err = videoWriter.Write(videoData)
	if err != nil {
		return err
	}

	// 遍历目录并添加文件到zip
	files, err := util.GetDirFiles(v.FilesDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		zipWriter.Create(file)
	}

	// // 删除ortherFileDir
	// err = os.RemoveAll(filesDir)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (v *Video) Load(filename string) (err error) {
	werr := util.EmptyWErr("videofile.Load")
	defer func() {
		err = werr.Err(err).ToError()
	}()
	// 打开zip文件
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 遍历zip文件
	for _, file := range reader.File {
		// 打开zip中的文件
		rc, err := file.Open()
		if err != nil {
			return err
		}

		// 创建目标文件路径
		path := filepath.Join(v.FilesDir, file.Name)
		
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return err
		}

		// 创建目标文件
		dst, err := os.Create(path)
		if err != nil {
			rc.Close()
			return err
		}

		// 复制文件内容
		_, err = io.Copy(dst, rc)
		dst.Close()
		rc.Close()
		if err != nil {
			return err
		}

		// 如果是video.json，解析它
		if file.Name == "video.json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if err = json.Unmarshal(data, v); err != nil {
				return err
			}
		}
	}
	return nil
}
