package options

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mises-id/sns-storagesvc/app/services/image/imagetype"
)

var (
	optionsPathPrefix = ":"
)

func parseOpPartsToOp(parts []string) *ImageOptions {
	op := &ImageOptions{}
	for _, v := range parts {
		parseOpStr(op, v)
	}
	return op

}

func parseOpPartsToOpV2(opstr string) *ImageOptions {
	op := &ImageOptions{}
	params, _ := url.ParseQuery(opstr)
	resizes, ok := params["resize"]
	if ok && len(resizes) > 0 {
		parseOpStr(op, "resize:"+resizes[0])
	}
	crops, ok := params["crop"]
	if ok && len(crops) > 0 {
		parseOpStr(op, "crop:"+crops[0])
	}
	watermarks, ok := params["watermark"]
	if ok && len(watermarks) > 0 {
		parseOpStr(op, "watermark:"+watermarks[0])
	}
	formats, ok := params["format"]

	if ok && len(formats) > 0 {
		parseOpStr(op, "format:"+formats[0])
	}
	qualitys, ok := params["quality"]
	if ok && len(qualitys) > 0 {
		parseOpStr(op, "quality:"+qualitys[0])
	}
	return op
}

func parseOpStr(op *ImageOptions, opstr string) {
	if opstr == "" {
		return
	}
	arr := strings.Split(opstr, optionsPathPrefix)
	name := arr[0]
	switch name {
	case "resize":
		parseResizeStrToResizeOptions(op, arr)
	case "crop":
		parseCropStrToCropOptions(op, arr)
	case "watermark":
		paserWatermarkTextStrToOptions(op, arr)
	case "format":
		parseFormatStrToOptions(op, arr)
	case "quality":
		parseQualityStrToOptions(op, arr)
	}
}
func parseFormatStrToOptions(op *ImageOptions, arr []string) {
	if len(arr) == 2 {
		Type, ok := imagetype.FormatTypes[arr[1]]
		if ok {
			op.Format = Type
		}
	}
}
func parseQualityStrToOptions(op *ImageOptions, arr []string) {
	if len(arr) == 2 {
		if err := parseDimension(&op.Quality, "quality", arr[1]); err != nil {
			return
		}
		if op.Quality < 0 {
			op.Quality = 0
		}
		if op.Quality > 100 {
			op.Quality = 100
		}
	}
}
func parseDimension(d *int, name, arg string) error {
	if v, err := strconv.Atoi(arg); err == nil && v >= 0 {
		*d = v
	} else {
		return fmt.Errorf("Invalid %s: %s", name, arg)
	}
	return nil
}
