package options

import (
	"encoding/base64"
	"strings"
)

type (
	WatermarkTextOptions struct {
		Watermark bool
		Text      string
		Font      string
		FontSize  int
		Color     string
	}
)

func parseWatermarkTextOptionsPath(op *ImageOptions) (path string) {
	if op == nil || op.WatermarkTextOptions == nil || !op.WatermarkTextOptions.Watermark || op.WatermarkTextOptions.Text == "" {
		return path
	}
	var path_arr = []string{
		"watermark",
		base64.RawURLEncoding.EncodeToString([]byte(op.WatermarkTextOptions.Text)),
	}
	path = strings.Join(path_arr, optionsPathPrefix)
	return path
}
