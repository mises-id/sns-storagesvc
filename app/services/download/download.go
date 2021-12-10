package download

import (
	"context"
	"errors"
	"strings"

	"github.com/mises-id/storagesvc/config/env"
)

var (
	localFilePath   = env.Envs.LocalFilePath
	downloadService IDownloadService
)

type (
	IDownloadService interface {
		Download(ctx context.Context, in *DownloadInput) (out *downloadOutput, err error)
	}
	DownloadInput struct {
		Url string
	}
	downloadOutput struct {
		Data []byte
	}
)

func download(ctx context.Context, in *DownloadInput) (out *downloadOutput, err error) {

	imgaeUrl := in.Url
	//http
	if strings.HasPrefix(imgaeUrl, "s3://") {
		downloadService = &s3DownloadSvc{}
	} else {
		in.Url = localFilePath + in.Url
		downloadService = &localDownloadSvc{}
	}
	return downloadService.Download(ctx, in)

}

func DownloadFile(ctx context.Context, in *DownloadInput) (out *downloadOutput, err error) {
	out = &downloadOutput{}
	if in.Url == "" {
		return out, errors.New("url invalid")
	}
	return download(ctx, in)
}
