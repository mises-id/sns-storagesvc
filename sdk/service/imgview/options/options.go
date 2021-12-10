package options

import (
	"fmt"
)

var (
	optionsPathPrefix = ":"
	FormatType        = map[string]string{
		"jpeg": "jpeg",
		"jpg":  "jpg",
		"png":  "png",
		"webp": "webp",
	}
)

type (
	ImageViewInput struct {
		Path         string
		ImageOptions *ImageOptions
	}

	ImageViewListInput struct {
		Path         []string
		ImageOptions *ImageOptions
	}

	ImageViewListOutput struct {
		Url []string
	}

	ImageViewOutput struct {
		Url string
	}
	ImageOptions struct {
		*ResizeOptions
		*CropOptions
		*WatermarkTextOptions
		Format string //jpeg,png,jpg,webp
	}
)

var pipeline = pipelineFuncs{
	parseResizeOptionsPath,
	parseCropOptionsPath,
	parseFormatOptionsPath,
}

func ParseOptionsPath(op *ImageOptions) (opPath string) {

	if op == nil {
		return opPath
	}
	opPath = pipeline.Run(op)

	return opPath
}

func parseFormatOptionsPath(op *ImageOptions) (path string) {

	if op == nil || op.Format == "" {
		if _, ok := FormatType[op.Format]; !ok {
			return path
		}
	}
	path = fmt.Sprintf("format%s%s", optionsPathPrefix, op.Format)

	return path
}
