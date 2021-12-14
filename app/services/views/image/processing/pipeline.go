package processing

import (
	"context"

	"github.com/mises-id/sns-storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/views/image/options"
)

type (
	pipelineStep    func(ctx *pipelineContext, imgdata *imagedata.ImageData, op *options.ImageOptions) error
	pipeline        []pipelineStep
	pipelineContext struct {
		ctx context.Context
	}
)

func (p pipeline) Run(ctx context.Context, imgdata *imagedata.ImageData, op *options.ImageOptions) error {
	pctx := pipelineContext{
		ctx: ctx,
	}
	for _, step := range p {
		if err := step(&pctx, imgdata, op); err != nil {
			return err
		}
	}
	return nil
}
