package storage

import (
	"context"
	"io"
	"os"
	"path"
	"strings"

	"github.com/mises-id/sns-storagesvc/config/env"
)

var (
	localFilePath = env.Envs.LocalFilePath
)

type FileStore struct{}

func (s *FileStore) Upload(ctx context.Context, in *StorageUploadInput) (*StorageUploadOutput, error) {

	s3in := &S3PutObjectInPut{
		Bucket: in.Bucket,
		Key:    in.FilePath,
		File:   in.File,
	}
	out := &StorageUploadOutput{}
	filePath := path.Join(s3in.Bucket, s3in.Key)
	localfile := path.Join(localFilePath, filePath)
	err := localSave(ctx, localfile, s3in.File)
	if err != nil {
		return out, err
	}
	out.FilePath = filePath
	return out, nil
}

func localSave(ctx context.Context, localfile string, file File) error {
	var err error
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
	if _, err = io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}
