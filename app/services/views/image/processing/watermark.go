package processing

import (
	"fmt"

	"github.com/h2non/bimg"
	"github.com/mises-id/sns-storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/views/image/options"
)

func watermark_text(pctx *pipelineContext, imgdata *imagedata.ImageData, in *options.ImageOptions) error {
	var (
		err error
	)
	if in == nil || in.WatermarkTextOptions == nil || !in.WatermarkTextOptions.Watermark {
		return nil
	}
	fmt.Println("watermark text ", in.WatermarkTextOptions.Text)
	bop := bimg.Watermark{
		Text:       in.WatermarkTextOptions.Text,
		Opacity:    0.25,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{255, 255, 255},
	}
	buf := imgdata.Data
	newData, err := bimg.NewImage(buf).Watermark(bop)
	if err != nil {
		fmt.Println("watermark err ", err.Error())
		return nil
	}
	imgdata.Data = newData
	return nil
}
