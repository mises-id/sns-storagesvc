package ctrl

import (
	"bytes"
	"context"
	"net/url"
	"time"

	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/app/logic"
)

type (
	ImageCtrl struct {
		logic *logic.ImageLogic
	}
)

func (ctrl *ImageCtrl) Handler(c echo.Context) error {

	//context timeout cancel
	ctx, cancel := context.WithTimeout(c.Request().Context(), 20*time.Second)
	defer cancel()

	//uri
	uri := c.Request().RequestURI
	str, err := url.QueryUnescape(uri)
	if err == nil {
		uri = str
	}

	//imgdata
	imgdata, err := ctrl.logic.Handler(ctx, uri)

	if err != nil {
		return err
	}
	return c.Stream(200, imgdata.Type.Mime(), bytes.NewReader(imgdata.Data))
}
