package image

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mises-id/storagesvc/app/services/views/image/imagedata"
	"github.com/mises-id/storagesvc/app/services/views/image/imageurl"
	"github.com/mises-id/storagesvc/app/services/views/image/processing"
)

type ()

func Handler(c echo.Context) error {
	uri := c.Request().RequestURI
	fmt.Println("uri: ", uri)
	// time out
	ctx, cancel := context.WithTimeout(c.Request().Context(), 20*time.Second)
	defer cancel()
	//parse uri

	imgUrl, op, err := imageurl.ParseUri(uri)
	if err != nil {
		return err
	}
	img, err := imagedata.DownLoadImageData(ctx, imgUrl)
	if err != nil {
		return err
	}
	imgdata, err := processing.ProcessImage(ctx, img, op)
	if err != nil {
		return err
	}
	return responseImageData(c.Request(), c.Response(), imgdata)

}

func responseImageData(req *http.Request, resp http.ResponseWriter, imgdata *imagedata.ImageData) error {

	resp.WriteHeader(200)
	resp.Header().Set("Content-Type", imgdata.Type.Mime())
	_, err := resp.Write(imgdata.Data)
	if err != nil {
		return err
	}
	return nil
}
