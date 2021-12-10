package options

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mises-id/storagesvc/app/services/views/image/imagetype"
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
func parseDimension(d *int, name, arg string) error {
	if v, err := strconv.Atoi(arg); err == nil && v >= 0 {
		*d = v
	} else {
		return fmt.Errorf("Invalid %s: %s", name, arg)
	}
	return nil
}
