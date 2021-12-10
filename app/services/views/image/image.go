package image

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/h2non/bimg"
	"github.com/mises-id/storagesvc/app/services/views/image/imagedata"
)

var (
	cdnHost         = "http://localhost:8088/insecure/"
	cdnPrefix       = "plain/"
	s3Prefix        = "s3://"
	newImagePrefix  = "new_image"
	savePathBase    = "/Users/cg/Documents/image/"
	localViewPrefix = "local://"
)

type (
	IImageSvc interface {
	}

	imageOutput struct {
		Path, Url string
	}
	imgData struct {
		Data []byte
	}
	//URL不处理  cdn
	ImageView struct {
		Path  string
		schme string // local s3
		*ImageOptions
		optinStr          string
		Str               string
		Con               context.Context
		sourceData        *imgData
		newData           *imgData
		Errors            []error
		newImageLocalPath string
	}
)

func (svc *ImageView) Start() (out *imageOutput, err error) {

	needDo, err := svc.checkPathAndHasDo()
	if err != nil {
		return out, err
	}
	if needDo {
		err := svc.makeNewImgLocalPath().newImage()
		if err != nil {
			return out, err
		}
		svc.Path = svc.newImageLocalPath
	}

	return svc.makeOutput(), nil
}

func (svc *ImageView) makeOutput() (out *imageOutput) {
	out = &imageOutput{}
	url := svc.ViewUrl()
	fmt.Println(url)
	out.Url = url
	out.Path = svc.Path
	return out
}

func (svc *ImageView) newImage() (err error) {

	// check local file

	localPath := svc.findImgOptionsLocalPath()
	if localPath != "" {
		fmt.Println("file is exist ", svc.newImageLocalPath)
		return nil //new image is exist
	}

	//get source image data
	return svc.initSourceImageData().handleNewImage()

}

func (svc *ImageView) makeNewImgLocalPath() *ImageView {
	if svc.newImageLocalPath != "" {
		return svc
	}
	var localPath, ops, path string
	filearr := strings.Split(svc.Path, "/")
	arrlen := len(filearr)
	if arrlen > 1 {
		path = strings.Join(filearr[:len(filearr)-1], "/")
	}
	name := filearr[arrlen-1]
	ops = svc.optinsToLocalFileStr()
	if svc.schme == "s3" {
		path = imgS3PathToLocalPath(path)
	}
	localPath = fmt.Sprintf("%s/%s/%s_%s", path, newImagePrefix, ops, name)
	fmt.Println("local ", localPath)
	fmt.Println("op ", svc.optinStr)
	svc.newImageLocalPath = localPath
	return svc
}

func imgS3PathToLocalPath(path string) (localpath string) {
	localpath = savePathBase + strings.Replace(path, s3Prefix, "", 1)
	return localpath
}

func (svc *ImageView) findImgOptionsLocalPath() (localPath string) {

	if _, err := os.Stat(svc.newImageLocalPath); err != nil {
		localPath = ""
	}
	return localPath
}

func (svc *ImageView) resizeImage() *ImageView {
	if !svc.ImageOptions.Resize {
		return svc
	}
	bop := bimg.Options{
		Width:  svc.ImageOptions.ResizeOptions.Width,
		Height: svc.ImageOptions.ResizeOptions.Height,
		Embed:  true,
	}
	buf := svc.newData.Data
	new, err := bimg.NewImage(buf).Process(bop)
	if err != nil {
		svc.AddError(err)
	}
	svc.newData.Data = new
	return svc
}

func (svc *ImageView) cropImage() *ImageView {
	if !svc.ImageOptions.Crop {
		return svc
	}
	bop := bimg.Options{
		Width:  svc.ImageOptions.CropOptions.Width,
		Height: svc.ImageOptions.CropOptions.Height,
		Crop:   true,
	}
	buf := svc.newData.Data
	new, err := bimg.NewImage(buf).Process(bop)
	if err != nil {
		svc.AddError(err)
	}
	svc.newData.Data = new
	return svc
}

func (svc *ImageView) writeNewImage() *ImageView {

	//err := bimg.Write(svc.newImageLocalPath, svc.newData.Data)
	err := writeLocalFile(svc.newImageLocalPath, svc.newData.Data)
	if err != nil {
		svc.AddError(err)
	}
	return svc
}

//resize crop img ...
func (svc *ImageView) handleNewImage() (err error) {
	if svc.sourceData == nil {
		return errors.New("image data not exist")
	}
	svc.newData = svc.sourceData
	return svc.resizeImage().cropImage().writeNewImage().SvcError()
}

func localFileIsExist(localpath string) bool {
	if _, err := os.Stat(localpath); err == nil {
		return true
	}
	return false
}

//
func (svc *ImageView) initSourceImageData() *ImageView {
	var s3local string
	var s3localIsExist bool
	imgUrl := svc.Path
	// check s3
	if svc.schme == "s3" {
		s3local = imgS3PathToLocalPath(svc.Path)
		s3localIsExist = localFileIsExist(s3local)
		if s3localIsExist {
			imgUrl = s3local
		}
	}
	fmt.Println("download url: ", imgUrl)
	imageData, err := imagedata.DownLoadImageData(svc.Con, imgUrl)
	if err != nil {
		fmt.Println("download file err:", err.Error())
		svc.AddError(err)
		return svc
	}
	//TODO save s3 local file
	if svc.schme == "s3" && !s3localIsExist {
		writeLocalFile(s3local, imageData.Data)
	}
	sd := &imgData{
		Data: imageData.Data,
	}
	svc.sourceData = sd
	return svc
}

func writeLocalFile(localfile string, data []byte) (err error) {
	fmt.Println("write source image start...")
	fmt.Println("s3local file path", localfile)
	arr := strings.Split(localfile, "/")
	filePath := strings.Join(arr[:len(arr)-1], "/")

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	dst, err := os.Create(localfile)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = dst.Write(data)
	if err != nil {
		fmt.Println("write file err:", err.Error())
		return err
	}
	return nil
}

func (svc *ImageView) ViewUrl() (url string) {

	path := svc.Path
	if svc.schme == "local" {
		path = fmt.Sprintf("%s%s", localViewPrefix, path)
	}
	url = fmt.Sprintf("%s%s%s%s", cdnHost, svc.Str, cdnPrefix, path)
	return url
}

//check path  && schme type
func (svc *ImageView) checkPathAndHasDo() (hasDo bool, err error) {
	if svc.Path == "" {
		return hasDo, errors.New("path is invalid")
	}
	if svc.ImageOptions.ResizeOptions.Resize || svc.ImageOptions.CropOptions.Crop {
		hasDo = true
	}
	schme := "local"
	if strings.HasPrefix(svc.Path, s3Prefix) {
		schme = "s3"
	}
	svc.schme = schme
	return hasDo, nil
}

// optinStr
func (svc *ImageView) optinsToLocalFileStr() (str string) {
	if svc.ImageOptions.Resize {
		if svc.ImageOptions.ResizeOptions.ResizeType == "" {
			svc.ImageOptions.ResizeType = "fit"
		}
		rt := svc.ImageOptions.ResizeType
		w := svc.ImageOptions.ResizeOptions.Width
		h := svc.ImageOptions.ResizeOptions.Height
		str = fmt.Sprintf("%sresize_%s_%d_%d", str, rt, w, h)
	}
	if svc.ImageOptions.CropOptions.Crop {
		w := svc.ImageOptions.CropOptions.Width
		h := svc.ImageOptions.CropOptions.Height
		str = fmt.Sprintf("%scrop_%d_%d", str, w, h)
	}
	fmt.Println("options ", str)
	svc.optinStr = str
	return str
}
func (svc *ImageView) SvcErrors() []error {
	if svc.Errors == nil {
		return nil
	}
	err := svc.Errors
	svc.Errors = nil // reset errors so next chain will start from zero
	return err
}
func (svc *ImageView) SvcError() error {
	if svc.Errors == nil {
		return nil
	}
	err := svc.Errors[0]
	svc.Errors = nil // reset errors so next chain will start from zero
	return err
}

func (svc *ImageView) AddError(err error) {
	svc.Errors = append(svc.Errors, err)
}
