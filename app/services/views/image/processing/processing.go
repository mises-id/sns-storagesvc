package processing

import (
	"context"
	"runtime"

	"github.com/h2non/bimg"
	"github.com/mises-id/sns-storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/views/image/options"
)

var (
	backgroundImg = bimg.Color{R: 255, G: 255, B: 255}
)

var processeScopesFuncs = pipeline{
	resize,
	crop,
	format,
	watermark_text,
	metadata,
}

func ProcessImage(ctx context.Context, imgdata *imagedata.ImageData, op *options.ImageOptions) (*imagedata.ImageData, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	var (
		err error
	)
	//err = scopes(makeProcessFuncs(ctx, imgdata, op)...)
	err = processeScopesFuncs.Run(ctx, imgdata, op)
	return imgdata, err
}
