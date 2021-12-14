package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mises-id/sns-storagesvc/sdk/service/imgview"
	"github.com/mises-id/sns-storagesvc/sdk/service/imgview/options"
)

func main() {

	key := "736563726574"
	salt := "68656C6C6F"
	host := "http://localhost:6060/"

	imgClient := imgview.New(
		imgview.Options{
			Key:  key,
			Salt: salt,
			Host: host,
		},
	)
	path := "test.jpg"

	paths := []string{
		"s3://sc-cg-test/upload/test/cg/test.jpg",
	}

	resizeOptions := &options.ResizeOptions{
		Resize: true,
		Width:  200,
	}
	cropOptions := &options.CropOptions{
		Crop:  true,
		Width: 300,
	}
	watermarkOptions := &options.WatermarkTextOptions{
		Watermark: true,
		Text:      "mises",
	}
	op := &options.ImageOptions{
		ResizeOptions:        resizeOptions,
		CropOptions:          cropOptions,
		Format:               "png",
		WatermarkTextOptions: watermarkOptions,
	}
	out, err := imgClient.GetImgUrl(context.Background(), &options.ImageViewInput{
		Path:         path,
		ImageOptions: op,
	})
	if err != nil {
		fmt.Println("get img url err:", err.Error())
	}
	fmt.Println("get img url success:", out.Url)
	outList, err1 := imgClient.GetImgUrlList(context.Background(), &options.ImageViewListInput{
		Path:         paths,
		ImageOptions: op,
	})
	if err1 != nil {
		fmt.Println("get img url list err:", err1.Error())
	}
	fmt.Println("get img url list success:", strings.Join(outList.Url, ","))
}
