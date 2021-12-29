package storage

import (
	"context"
	"errors"
	"os"
	"path"

	"github.com/mises-id/sns-storagesvc/config/env"
)

var (
	storageService   IStorageService
	Prefix           = "upload/"
	errFileNotExist  = errors.New("file does not exist")
	errFilePathEmpty = errors.New("file path not empty")
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
	bindSvc(env.Envs.StorageProvider)

}

func bindSvc(s string) {
	switch s {
	default:
		storageService = &FileStore{}
	case "local":
		storageService = &FileStore{}
	case "s3":
		storageService = &S3Storage{}
	}
}

func New() *StorageSvc {
	return &StorageSvc{}
}

func (svc *StorageSvc) Bind(s string) *StorageSvc {
	bindSvc(s)
	return svc
}

func (s *StorageSvc) Upload(ctx context.Context, in *StorageUploadInput) (*StorageUploadOutput, error) {
	if in.FilePath == "" {
		return nil, errFilePathEmpty
	}
	in.FilePath = path.Join(Prefix, in.FilePath)
	return storageService.Upload(ctx, in)

}
func (s *StorageSvc) FUpload(ctx context.Context, in *StorageFUploadInput) (out *StorageUploadOutput, err error) {

	//valid localfile

	file, err := os.Open(in.LocalFile)
	if err != nil {
		return out, errFileNotExist
	}
	defer file.Close()
	upin := &StorageUploadInput{
		Bucket:   in.Bucket,
		FilePath: in.FilePath,
		File:     file,
	}
	return s.Upload(ctx, upin)

}
