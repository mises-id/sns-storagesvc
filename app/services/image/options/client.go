package options

import (
	"strings"

	"github.com/mises-id/sns-storagesvc/app/services/image/imagetype"
)

type (
	ImageOptions struct {
		*ResizeOptions
		*CropOptions
		*WatermarkTextOptions
		Format  imagetype.Type
		Quality int
	}
)

func ParseOpPathToOp(opstr, version string) *ImageOptions {
	if version == "1.0" {
		opParts := strings.Split(opstr, "/")
		op := parseOpPartsToOp(opParts)
		return op
	}
	return parseOpPartsToOpV2(opstr)
}
