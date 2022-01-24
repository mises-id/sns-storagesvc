package image

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagedata"
	"github.com/mises-id/sns-storagesvc/app/services/image/imageurl"
	"github.com/mises-id/sns-storagesvc/app/services/image/processing"
	"github.com/mises-id/sns-storagesvc/lib/codes"
)

type ()

func Handler(c echo.Context) error {
	uri := c.Request().RequestURI
	str, err := url.QueryUnescape(uri)
	if err == nil {
		uri = str
	}
	// time out
	ctx, cancel := context.WithTimeout(c.Request().Context(), 20*time.Second)
	defer cancel()
	//parse uri
	imgUrl, op, err := imageurl.ParseUri(uri)
	if err != nil {
		fmt.Println("parse uri error ", err.Error())
		return codes.ErrForbidden
	}
	img, err := imagedata.DownLoadImageData(ctx, imgUrl)
	if err != nil {
		fmt.Println("download image error ", err.Error())
		return codes.ErrNotFound
	}
	imgdata, err := processing.ProcessImage(ctx, img, op)
	if err != nil {
		fmt.Println("process image error ", err.Error())
		return err
	}
	return c.Stream(200, imgdata.Type.Mime(), bytes.NewReader(imgdata.Data))

}

func responseImageData(req *http.Request, resp http.ResponseWriter, imgdata *imagedata.ImageData) error {

	resp.WriteHeader(200)
	resp.Header().Set("Content-Type", imgdata.Type.Mime())
	resp.Header().Set("Content-Length", strconv.Itoa(len(imgdata.Data)))

	_, err := resp.Write(imgdata.Data)
	if err != nil {
		return err
	}
	return nil
}
