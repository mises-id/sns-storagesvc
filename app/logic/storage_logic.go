package logic

import (
	"context"

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
		return out, nil
	}

	out.AttachId = res.ID
	//get url by path
	return out, nil
}
