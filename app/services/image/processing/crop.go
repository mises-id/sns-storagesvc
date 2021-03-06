package processing

import (
	"github.com/h2non/bimg"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/image/options"
)

func crop(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {
	var (
		err error
	)
	if in == nil || in.CropOptions == nil || !in.Crop {
		return nil
	}
	bop := bimg.Options{
		Width:  in.CropOptions.Width,
		Height: in.CropOptions.Height,
		Crop:   true,
	}
	if in.Quality > 0 {
		bop.Quality = in.Quality
	}
	buf := imgdata.Data
	imgdata.Data, err = bimg.NewImage(buf).Process(bop)
	if err != nil {
		return err
	}
	return nil
}
