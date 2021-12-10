package image

type (
	ResizeOptions struct {
		Resize     bool
		ResizeType string
		Height     int
		Width      int
	}

	CropOptions struct {
		Crop   bool
		Height int
		Width  int
	}

	ImageOptions struct {
		*ResizeOptions
		*CropOptions
	}
)
