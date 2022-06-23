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
		Format  string //jpeg,png,jpg,webp
		Quality int
	}
)

var pipeline = pipelineFuncs{
	parseResizeOptionsPath,
	parseCropOptionsPath,
	parseWatermarkTextOptionsPath,
	parseFormatOptionsPath,
	parseQualityOptionsPath,
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
	path = fmt.Sprintf("format=%s", op.Format)

	return path
}
func parseQualityOptionsPath(op *ImageOptions) (path string) {

	if op == nil || op.Quality <= 0 {
		return path
	}
	path = fmt.Sprintf("quality=%d", op.Quality)

	return path
}
