package storage

import (
	"context"
	"fmt"
	"os"
)

var (
	storageService IStorageService
	Prefix         = "upload/"
)

type (
	IStorageService interface {
		Upload(ctx context.Context, in *StorageUploadInput) (*StorageUploadOutput, error)
		//FUpload(ctx context.Context, bucket, filePath, localFile string) (StorageOutput, error)
	}
	File interface {
		Read(p []byte) (n int, err error)
	}

	StorageUploadOutput struct {
		Url      string
		FilePath string
	}

	StorageUploadInput struct {
		Bucket, FilePath string
		File
	}

	StorageFUploadInput struct {
		Bucket, FilePath, LocalFile string
	}

	StorageSvc struct {
	}
)

func init() {
	storageService = &S3Storage{}
}

func (svc *StorageSvc) Bind(s string) *StorageSvc {
	switch s {
	default:
		storageService = &S3Storage{}
	case "local":
		storageService = &S3Storage{}
	}
	return svc
}

func (s *StorageSvc) Upload(ctx context.Context, in *StorageUploadInput) (*StorageUploadOutput, error) {

	in.FilePath = Prefix + in.FilePath
	return storageService.Upload(ctx, in)

}
func (s *StorageSvc) FUpload(ctx context.Context, in *StorageFUploadInput) (out *StorageUploadOutput, err error) {

	//valid localfile

	file, err := os.Open(in.LocalFile)
	fmt.Println(in.LocalFile)
	if err != nil {
		return out, err
	}
	defer file.Close()
	upin := &StorageUploadInput{
		Bucket:   in.Bucket,
		FilePath: in.FilePath,
		File:     file,
	}
	return s.Upload(ctx, upin)

}
