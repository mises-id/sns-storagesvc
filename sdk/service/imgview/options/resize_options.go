package options

import "fmt"

type (
	ResizeOptions struct {
		Resize bool
		//ResizeType
		//fit  resizes the image while keeping aspect ratio to fit given size;
		//fill resizes the image while keeping aspect ratio to fill given size and cropping projecting parts;
		//force resizes the image without keeping aspect ratio;
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
	resizePath = fmt.Sprintf("resize=%s%s%d%s%d", op.ResizeType, optionsPathPrefix, op.ResizeOptions.Width, optionsPathPrefix, op.ResizeOptions.Height)
	return resizePath
}
