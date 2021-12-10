package options

type (
	CropOptions struct {
		Crop   bool
		Height int
		Width  int
	}
)

func parseCropStrToCropOptions(op *ImageOptions, arr []string) {
	len := len(arr)
	if len >= 3 {
		cp := &CropOptions{
			Crop: true,
		}
		if err := parseDimension(&cp.Width, "resize width", arr[1]); err != nil {
			return
		}
		if err := parseDimension(&cp.Height, "resize height", arr[2]); err != nil {
			return
		}
		op.CropOptions = cp
	}
}
