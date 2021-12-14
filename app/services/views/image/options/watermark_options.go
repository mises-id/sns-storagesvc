package options

import (
	"encoding/base64"
	"fmt"
)

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

	text, err := base64.RawURLEncoding.DecodeString(arr[1])
	if err != nil {
		fmt.Println("watermark text decode error ", err.Error())
		return
	}
	wop := &WatermarkTextOptions{
		Watermark: true,
		Text:      string(text[:]),
	}
	op.WatermarkTextOptions = wop
}
