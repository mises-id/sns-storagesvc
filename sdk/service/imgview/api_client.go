package imgview

import (
	"context"
	"errors"

	"github.com/mises-id/storage-sdk/service/imgview/options"
)

var (
	pathInvalid = errors.New("Invalid path")
	hostInvalid = errors.New("Invalid host")
)

type (
	Options struct {
		Key, Salt, Host string
	}

	Client struct {
		options Options

		con context.Context
	}
)

func New(options Options) *Client {

	c := &Client{
		options: options,
	}
	return c
}

func (c *Client) GetImgUrlList(ctx context.Context, in *options.ImageViewListInput) (*options.ImageViewListOutput, error) {
	var urlArr []string

	out := &options.ImageViewListOutput{}

	for _, v := range in.Path {
		newin := &options.ImageViewInput{
			ImageOptions: in.ImageOptions,
			Path:         v,
		}
		res, err := c.handler(ctx, newin)
		url := ""
		if err == nil {
			url = res
		}
		urlArr = append(urlArr, url)
	}
	out.Url = urlArr
	return out, nil
}

//parse options
func (c *Client) GetImgUrl(ctx context.Context, in *options.ImageViewInput) (*options.ImageViewOutput, error) {

	out := &options.ImageViewOutput{}
	res, err := c.handler(ctx, in)

	if err != nil {
		return out, err
	}
	out = &options.ImageViewOutput{Url: res}

	return out, nil
}
