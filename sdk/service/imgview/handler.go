package imgview

import (
	"context"
	"fmt"
	"strings"

	"github.com/mises-id/sns-storagesvc/sdk/service/imgview/options"
)

func (c *Client) handler(ctx context.Context, in *options.ImageViewInput) (string, error) {

	err := checkParams(in)
	if err != nil {
		return "", err
	}
	//http https path
	if strings.HasPrefix(in.Path, "http") {
		return in.Path, nil
	}
	opPath := options.ParseOptionsPath(in.ImageOptions)
	if c.options.Key != "" && c.options.Salt != "" {
		signature, err := signature(c.options.Key, c.options.Salt, in.Path, opPath)
		if err != nil {
			return "", err
		}
		if opPath != "" {
			opPath = "&" + opPath
		}
		opPath = "sign=" + signature + opPath
	}
	url := createViewUrl(c.options.Host, in.Path, opPath)

	return url, nil
}

func createViewUrl(host, path, opPath string) (url string) {

	path = strings.TrimPrefix(path, "/")
	url = fmt.Sprintf("%s%s?%s", host, path, opPath)
	url = strings.TrimSuffix(url, "/")
	return url

}

func checkParams(in *options.ImageViewInput) error {
	if in.Path == "" {
		return pathInvalid
	}
	return nil
}
