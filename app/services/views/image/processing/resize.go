package processing

import (
	"github.com/h2non/bimg"
	"github.com/mises-id/storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/storagesvc/app/services/views/image/options"
)

func resize(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {
	var (
		err error
	)
	if in == nil || in.ResizeOptions == nil || !in.ResizeOptions.Resize {
		return nil
	}
	bop := bimg.Options{
		Width:  in.ResizeOptions.Width,
		Height: in.ResizeOptions.Height,
		Embed:  true,
	}
	buf := imgdata.Data
	imgdata.Data, err = bimg.NewImage(buf).Process(bop)
	if err != nil {
		return err
	}
	return nil
}
