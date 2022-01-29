package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mises-id/sns-storagesvc/sdk/service/imgview"
	"github.com/mises-id/sns-storagesvc/sdk/service/imgview/options"
)

func main() {

	key := "1d3dc5aa0144b6995ea52a8e6f6dd430b193ca4f2242d7f08bd5f735945949c26aa556c444d4b6a076faf50239f98e2797b51452361741f5467097605c64f406"
	salt := "0072cd5aaf3d88e990c70de6f789fd4782ec51060cc9840bef7bff8aaa524b467aca6a58ba20f9dfd8a5c1d78a57751c241cdff71cae8722561623221c1841b0"
	host := "http://localhost:7070/"

	imgClient := imgview.New(
		imgview.Options{
			Key:  key,
			Salt: salt,
			Host: host,
		},
	)
	path := "test.png"

	paths := []string{
		"https://s3://mises-storage/upload/test/cg/test.jpg",
		"test.jpeg",
		"test2.jpeg",
	}

	resizeOptions := &options.ResizeOptions{
		Resize: true,
		Width:  200,
	}

	op := &options.ImageOptions{
		ResizeOptions: resizeOptions,
		Quality:       50,
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
