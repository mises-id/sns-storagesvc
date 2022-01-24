package storage

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mises-id/sns-storagesvc/app/services/awssvc/s3svc"
)

//var file io.Reader

type (
	S3PutObjectInPut struct {
		Bucket, Key string
		File
	}
	S3GetObjectInPut struct {
		Bucket, Key string
	}
)

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type S3Storage struct {
}

func (s *S3Storage) Upload(ctx context.Context, in *StorageUploadInput) (*StorageUploadOutput, error) {

	s3in := &S3PutObjectInPut{
		Bucket: in.Bucket,
		Key:    in.FilePath,
		File:   in.File,
	}
	out := &StorageUploadOutput{}
	err := s.S3PutObject(ctx, s3in)
	if err != nil {
		return out, err
	}
	out.FilePath = s3svc.S3Prefix + s3in.Bucket + "/" + s3in.Key
	return out, nil
}

/* func (s *S3Storage) FUpload(ctx context.Context, in *StoragFUploadInput) (error){
	file,err := os.Open(in.LocalFile)
	if err != nil {
		return err
	}
	err = s.S3PutObject(ctx, bucket, filePath, file)
	if err != nil {
		return err
	}
	return nil
} */

func (s *S3Storage) S3PutObject(ctx context.Context, in *S3PutObjectInPut) error {

	if in.Bucket == "" {
		in.Bucket = s3svc.S3Bucket
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(in.Bucket),
		Key:    aws.String(in.Key),
		Body:   in.File,
	}
	_, err := putFile(ctx, s3svc.GetS3Client(), input)
	if err != nil {
		fmt.Println("s3 put object error: ", err)
		return err
	}
	return nil
}

func (s *S3Storage) S3GetObject(ctx context.Context, in *S3GetObjectInPut) error {
	client := s3svc.GetS3Client()
	input := &s3.GetObjectInput{
		Bucket: aws.String(in.Bucket),
		Key:    aws.String(in.Key),
	}
	res, err := client.GetObject(ctx, input)
	//client.HeadObject()
	if err != nil {
		return err
	}
	//p,err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return nil
}

func putFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}
