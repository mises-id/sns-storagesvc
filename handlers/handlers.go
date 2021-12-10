package handlers

import (
	"context"
	"fmt"

	"github.com/mises-id/storagesvc/app/services/storage"
	svcG "github.com/mises-id/storagesvc/app/services/views/image"
	"github.com/mises-id/storagesvc/config/env"

	pb "github.com/mises-id/storagesvc/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.StoragesvcServer {
	return storagesvcService{}
}

type storagesvcService struct{}

func (s storagesvcService) FUpload(ctx context.Context, in *pb.FUploadRequest) (*pb.FUploadResponse, error) {
	var resp pb.FUploadResponse
	fmt.Println(in)
	id := env.Envs
	fmt.Println(id)
	svc := &storage.StorageSvc{}
	sin := &storage.StorageFUploadInput{
		Bucket:    in.Bucket,
		FilePath:  in.Key,
		LocalFile: in.File,
	}
	res, err := svc.FUpload(ctx, sin)
	if err != nil {
		return nil, err
	}
	if res.FilePath != "" {

	}
	resp.Path = res.FilePath

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
