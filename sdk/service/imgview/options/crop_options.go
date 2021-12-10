package options

import "fmt"

type (
	CropOptions struct {
		Crop   bool
		Height int
		Width  int
	}
)

func parseCropOptionsPath(op *ImageOptions) (cropPath string) {
	if op == nil || op.CropOptions == nil || !op.Crop {
		return cropPath
	}
	cropPath = fmt.Sprintf("crop%s%d%s%d", optionsPathPrefix, op.CropOptions.Width, optionsPathPrefix, op.CropOptions.Height)
	return cropPath
}
