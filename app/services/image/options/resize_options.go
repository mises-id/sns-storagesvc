package options

type (
	ResizeOptions struct {
		Resize     bool
		ResizeType string
		Height     int
		Width      int
	}
)

func parseResizeStrToResizeOptions(op *ImageOptions, arr []string) {
	len := len(arr)
	if len == 4 {
		resize_type := "fit"
		if arr[1] == "fit" || arr[1] == "fill" || arr[1] == "force" {
			resize_type = arr[1]
		}
		rs := &ResizeOptions{
			Resize:     true,
			ResizeType: resize_type,
		}
		if err := parseDimension(&rs.Width, "resize width", arr[2]); err != nil {
			return
		}
		if err := parseDimension(&rs.Height, "resize height", arr[3]); err != nil {
			return
		}
		op.ResizeOptions = rs
	}
}
