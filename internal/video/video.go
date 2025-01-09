package video

import (
	"os"
	"path/filepath"

	"github.com/wumingjieno1/videofactory/internal/video/model"
	"github.com/wumingjieno1/videofactory/util"
)

func NewVideo() (*model.Video, error) {
	filesDir := filepath.Join("./tmp", util.RandString(10, util.LOWER_CASE_LETTERS, util.NUMBERS, util.UPPER_CASE_LETTERS))
	if err := os.MkdirAll(filesDir, 0755); err != nil {
		return nil, err
	}
	return &model.Video{
		Title:    "新建视频",
		FilesDir: filesDir,
	}, nil
}
