package logic

import (
	"context"
	"fmt"

	"github.com/mises-id/sns-storagesvc/app/services/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/image/imageurl"
	"github.com/mises-id/sns-storagesvc/app/services/image/processing"
	"github.com/mises-id/sns-storagesvc/lib/codes"
)

type (
	ImageOptions struct {
		lctx            *context.Context
		uri             string
		parseUri        string
		originImageData *imagedata.ImageData
		newImageData    *imagedata.ImageData
		result          interface{}
	}
	ImageLogic struct {
		options *ImageOptions
	}
)

func (logic *ImageLogic) Handler(ctx context.Context, uri string) (*imagedata.ImageData, error) {

	imgUrl, op, err := imageurl.ParseUri(uri)
	if err != nil {
		fmt.Println("parse uri error ", err.Error())
		return nil, codes.ErrForbidden
	}
	//find cache
	img, err := imagedata.DownLoadImageData(ctx, imgUrl)
	if err != nil {
		fmt.Println("download image error ", err.Error())
		return nil, codes.ErrNotFound
	}
	imgdata, err := processing.ProcessImage(ctx, img, op)
	if err != nil {
		fmt.Println("process image error ", err.Error())
		return nil, err
	}
	//cache
	return imgdata, nil

}
