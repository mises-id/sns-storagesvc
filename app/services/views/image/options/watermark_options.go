package options

import "encoding/base64"

type (
	Color struct {
		R, G, B uint8
	}
	WatermarkTextOptions struct {
		Watermark  bool
		Text       string
		Font       string
		FontSize   int
		Background Color
	}
)

func paserWatermarkTextStrToOptions(op *ImageOptions, arr []string) {

	len := len(arr)
	if len < 2 || arr[1] == "" {
		return
	}
	wop := &WatermarkTextOptions{
		Watermark: true,
		Text:      base64.RawURLEncoding.EncodeToString([]byte(arr[1])),
	}
	op.WatermarkTextOptions = wop
}
