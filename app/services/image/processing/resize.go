package processing

import (
	"github.com/h2non/bimg"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/image/options"
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
	}
	if in.Quality > 0 {
		bop.Quality = in.Quality
	}
	switch in.ResizeOptions.ResizeType {
	case "fit":
		bop.Embed = true
		bop.Crop = true
	case "fill":
		bop.Embed = true
	case "force":
		bop.Force = true
	}
	buf := imgdata.Data
	imgdata.Data, err = bimg.NewImage(buf).Process(bop)
	if err != nil {
		return err
	}
	return nil
}
