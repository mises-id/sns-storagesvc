package processing

import (
	"fmt"

	"github.com/h2non/bimg"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagetype"
	"github.com/mises-id/sns-storagesvc/app/services/image/options"
)

func metadata(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {

	if len(imgdata.Data) == 0 {
		return nil
	}
	img_type := bimg.NewImage(imgdata.Data).Type()

	Type, ok := imagetype.Types[img_type]
	if !ok {
		Type = 0
	}
	imgdata.Type = Type
	return nil
}

func format(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {

	var (
		err error
	)
	if in.Format == 0 {
		return nil
	}
	imgdata.Data, err = bimg.NewImage(imgdata.Data).Convert(bimg.ImageType(in.Format))
	if err != nil {
		fmt.Println("format error ", err.Error())
		return err
	}
	fmt.Println("format type:", in.Format)
	return nil
}
