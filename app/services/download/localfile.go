package download

import (
	"context"
	"os"
)

type (
	localDownloadInput struct {
		Path string
	}
	localDownloadOutput struct {
		Data []byte
	}

	localDownloadSvc struct {
	}
)

func (svc *localDownloadSvc) Download(ctx context.Context, in *DownloadInput) (out *downloadOutput, err error) {

	out = &downloadOutput{}
	localin := &localDownloadInput{
		Path: in.Url,
	}
	localout, err := localDownload(ctx, localin)
	if err != nil {
		return out, err
	}
	out.Data = localout.Data

	return out, err

}

func localDownload(cxt context.Context, in *localDownloadInput) (out *localDownloadOutput, err error) {

	out = &localDownloadOutput{}
	p, err := os.ReadFile(in.Path)

	if err != nil {
		return out, err
	}
	out.Data = p

	return out, err

}
