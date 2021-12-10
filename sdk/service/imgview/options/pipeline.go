package options

import "path"

type (
	pipelineFunc  func(op *ImageOptions) string
	pipelineFuncs []pipelineFunc
)

func (p pipelineFuncs) Run(op *ImageOptions) string {
	var pathStr string
	for _, step := range p {
		pathStr = path.Join(pathStr, step(op))
	}
	return pathStr
}
