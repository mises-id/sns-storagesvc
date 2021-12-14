package options

import (
	"strings"

	"github.com/mises-id/sns-storagesvc/app/services/views/image/imagetype"
)

type (
	ImageOptions struct {
		*ResizeOptions
		*CropOptions
		*WatermarkTextOptions
		Format imagetype.Type
	}
)

func ParseOpPathToOp(opstr string) *ImageOptions {

	opParts := strings.Split(opstr, "/")
	op := parseOpPartsToOp(opParts)
	return op
}
