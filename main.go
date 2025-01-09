package main

import (
	"fmt"
	"os"

	"github.com/wumingjieno1/videofactory/internal/video"
)

func main() {
	// video := &model.Video{}
	// err := videofile.Save(video, "./model", "./output/test.zip")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// video, dir, err := videofile.Load("./output/test.zip")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(dir)
	// fmt.Println(video)
	os.MkdirAll("./tmp", 0755)
	defer os.RemoveAll("./tmp")
	v, err := video.NewVideo()
	if err != nil {
		fmt.Println(err)
	}
	err = v.Load("./output/test.zip")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.FilesDir)

}
