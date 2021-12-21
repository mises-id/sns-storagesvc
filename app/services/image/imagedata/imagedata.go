package imagedata

import (
	"context"

	"github.com/mises-id/sns-storagesvc/app/services/download"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagetype"
)

type (
	ImageData struct {
		Data []byte
		Type imagetype.Type
	}
)

func DownLoadImageData(ctx context.Context, ImageUrl string) (imgdata *ImageData, err error) {

	imgdata = &ImageData{}
	outDownload, err := download.DownloadFile(ctx, &download.DownloadInput{Url: ImageUrl})

	if err == nil {
		imgdata.Data = outDownload.Data
	}

	return imgdata, err

}
