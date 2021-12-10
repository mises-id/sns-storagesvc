package options

import "fmt"

type (
	ResizeOptions struct {
		Resize     bool
		ResizeType string
		Height     int
		Width      int
	}
)

func parseResizeOptionsPath(op *ImageOptions) (resizePath string) {
	if op == nil || op.ResizeOptions == nil || !op.Resize {
		return resizePath
	}
	if op.ResizeType == "" {
		op.ResizeType = "fit"
	}
	resizePath = fmt.Sprintf("resize%s%s%s%d%s%d", optionsPathPrefix, op.ResizeType, optionsPathPrefix, op.CropOptions.Width, optionsPathPrefix, op.CropOptions.Height)
	return resizePath
}
