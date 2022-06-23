package options

import "fmt"

type (
	pipelineFunc  func(op *ImageOptions) string
	pipelineFuncs []pipelineFunc
)

func (p pipelineFuncs) Run(op *ImageOptions) string {
	var pathStr string
	pathStr = "version=2.0"
	for _, step := range p {
		str := step(op)
		if str == "" {
			continue
		}
		if pathStr == "" {
			pathStr = str
		} else {
			pathStr = fmt.Sprintf("%s&%s", pathStr, str)
		}
	}
	return pathStr
}
