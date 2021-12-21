package imagetype

type WatermarkType int

const (
	UnknownWatermarkType WatermarkType = iota
	TextWatermark
	ImageWatermark
)

var (
	Watermark = map[string]WatermarkType{
		"text":  TextWatermark,
		"image": ImageWatermark,
	}
)
