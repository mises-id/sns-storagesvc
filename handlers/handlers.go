package handlers

import (
	"context"
	"time"

	"github.com/mises-id/sns-storagesvc/app/logic"
	svcG "github.com/mises-id/sns-storagesvc/app/services/image"

	pb "github.com/mises-id/sns-storagesvc/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.StoragesvcServer {
	return storagesvcService{}
}

type storagesvcService struct{}

func (s storagesvcService) FUpload(ctx context.Context, in *pb.FUploadRequest) (*pb.FUploadResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var resp pb.FUploadResponse
	svc := &logic.StorageLogic{}
	res, err := svc.FUpload(ctx, in.File, in.Bucket, in.Key)
	if err != nil {
		return nil, err
	}
	resp.Path = res.Path
	resp.AttachId = res.AttachId
	resp.Url = res.Url

	return &resp, nil
}

func (s storagesvcService) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadResponse, error) {
	var resp pb.UploadResponse

	return &resp, nil
}

func (s storagesvcService) ImageUrl(ctx context.Context, in *pb.ImageUrlRequest) (*pb.ImageUrlResponse, error) {
	var resp pb.ImageUrlResponse
	op := &svcG.ImageOptions{}
	ro := in.Options.ResizeOptions
	resizeOptions := &svcG.ResizeOptions{}
	if ro != nil {
		resizeOptions = &svcG.ResizeOptions{
			Resize:     ro.Resize,
			Height:     int(ro.Height),
			Width:      int(ro.Width),
			ResizeType: ro.ResizeType,
		}

	}
	cp := in.Options.CropOptions
	cropOptions := &svcG.CropOptions{}
	if cp != nil {
		cropOptions = &svcG.CropOptions{
			Crop:   cp.Crop,
			Width:  int(cp.Width),
			Height: int(cp.Height),
		}

	}
	op.ResizeOptions = resizeOptions
	op.CropOptions = cropOptions
	imgSvc := &svcG.ImageView{Path: in.Img, Str: in.Str, Con: ctx, ImageOptions: op}
	res, err := imgSvc.Start()
	if err != nil {
		return nil, err
	}

	resp.Url = res.Url
	resp.Path = res.Path
	return &resp, nil
}
