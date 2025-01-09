package model

import (
	"image/color"

	"github.com/wumingjieno1/videofactory/constant"
)



type ImageInfo struct {
	Image string  `json:"file"`
	Texts []*Text `json:"texts"`
}

type Text struct {
	Content    string       `json:"content"`
	Tag        constant.Tag `json:"tag"` // 标签  1：视频标题 2：文章原标题 3：文章翻译标题 4：文章原内容 5：文章翻译内容
	Voice      string       `json:"voice"`
	Emoji      string       `json:"emoji"`
	Position   *Point       `json:"position"`
	Font       string       `json:"font"`
	FontSize   int          `json:"font_size"`
	FontColor  color.RGBA   `json:"font_color"`
	FontFamily string       `json:"font_family"`
	Border     *Border      `json:"border"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Border struct {
	Color color.RGBA `json:"color"`
	Width int        `json:"width"`
}
