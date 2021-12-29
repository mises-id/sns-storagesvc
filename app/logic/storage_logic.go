package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/mises-id/sns-storagesvc/app/models"
	"github.com/mises-id/sns-storagesvc/app/services/storage"
)

type (
	StorageLogic struct {
		model *models.Attachment
	}
	StorageUploadOutput struct {
		Url, Path string
		AttachId  uint64
	}
)

func (logic *StorageLogic) FUpload(ctx context.Context, localfile, bucket, key string) (*StorageUploadOutput, error) {

	//validate
	if localfile == "" || key == "" {
		return nil, errors.New("invalid params")
	}

	out := &StorageUploadOutput{}
	//to save file
	svc := storage.New()
	sin := &storage.StorageFUploadInput{
		Bucket:    bucket,
		FilePath:  key,
		LocalFile: localfile,
	}
	upload, err := svc.FUpload(ctx, sin)
	if err != nil {
		fmt.Println("upload file error: ", err.Error())
		return out, err
	}
	data := &models.Attachment{
		Filename: upload.FilePath,
	}
	out.Path = upload.FilePath
	//return out, nil
	//create attachment log

	res, err := logic.model.Create(ctx, data)
	if err != nil {
		fmt.Println("create attachment error: ", err.Error())
		return out, nil
	}

	out.AttachId = res.ID
	//get url by path
	return out, nil
}
