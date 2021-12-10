package processing

import (
	"github.com/h2non/bimg"
	"github.com/mises-id/storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/storagesvc/app/services/views/image/options"
)

func watermark_text(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {
	var (
		err error
	)
	if in == nil || in.WatermarkTextOptions == nil || !in.WatermarkTextOptions.Watermark {
		return nil
	}
	bop := bimg.Watermark{
		Text: in.WatermarkTextOptions.Text,
	}
	buf := imgdata.Data
	imgdata.Data, err = bimg.NewImage(buf).Watermark(bop)
	if err != nil {
		return err
	}
	return nil
}
