package imageurl

import (
	"errors"
	"strings"

	"github.com/mises-id/sns-storagesvc/app/services/image/options"
	"github.com/mises-id/sns-storagesvc/config/env"
)

var (
	//signUrl         = env.Envs.SignURL
	errUrlMsg       = "Invalid Url"
	errSignatureMsg = "Invalid Signature "
	errUrl          = errors.New(errUrlMsg)
	errSignature    = errors.New(errSignatureMsg)
)

func ParseUri(uri string) (img string, op *options.ImageOptions, err error) {

	img, opPath, err := handlerUri(uri)
	if err != nil {
		return "", nil, err
	}
	op, err = parseOpPath(img, opPath)
	if err != nil {
		return "", nil, err
	}
	return img, op, nil
}

func parseOpPath(img, opPath string) (op *options.ImageOptions, err error) {
	if env.Envs.SignURL {
		signature := opPath
		if queryStart := strings.IndexByte(opPath, '/'); queryStart >= 0 {
			signature = opPath[:queryStart]
			opPath = opPath[queryStart:]
		} else {
			opPath = ""
		}
		if err = verifySignature(signature, img, opPath); err != nil {
			return nil, errSignature
		}
	}
	op = options.ParseOpPathToOp(opPath)
	return op, nil

}

func handlerUri(uri string) (imgPath string, opPath string, err error) {
	if uri == "" || uri == "/" {
		return "", "", errUrl
	}
	uri = strings.TrimPrefix(uri, "/")
	uri = strings.TrimSuffix(uri, "/")
	arr := strings.Split(uri, "?")
	imgPath = arr[0]
	if len(arr) > 1 {
		opPath = arr[1]
	}

	return imgPath, opPath, nil
}

func MakeUrl() {

}
