package download

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mises-id/storagesvc/app/services/awssvc/s3svc"
)

type (
	s3DownloadInput struct {
		Bucket, Key string
	}
	s3DownloadOutput struct {
		Data []byte
	}

	s3DownloadSvc struct {
	}
)

func (svc *s3DownloadSvc) Download(ctx context.Context, in *DownloadInput) (out *downloadOutput, err error) {
	//handle in.
	out = &downloadOutput{}
	str := strings.Replace(in.Url, s3svc.S3Prefix, "", 1)
	strArr := strings.Split(str, "/")
	strArrLen := len(strArr)
	if strArrLen < 2 {
		return out, errors.New("image url is invalide")
	}
	bucket := strArr[0]
	key := strings.Join(strArr[1:], "/")
	fmt.Printf("image url %v,bucket %v,key %v", in.Url, bucket, key)
	s3in := &s3DownloadInput{
		Bucket: bucket,
		Key:    key,
	}
	s3out, err := s3GetObject(ctx, s3in)
	if err != nil {
		return out, err
	}

	out.Data = s3out.Data

	return out, err
}

func s3GetObject(ctx context.Context, in *s3DownloadInput) (out *s3DownloadOutput, err error) {

	out = &s3DownloadOutput{}
	client := s3svc.GetS3Client()
	input := &s3.GetObjectInput{
		Bucket: aws.String(in.Bucket),
		Key:    aws.String(in.Key),
	}
	res, err := client.GetObject(ctx, input)
	if err != nil {
		return out, err
	}
	outP, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return out, err
	}
	defer res.Body.Close()
	out.Data = outP
	return out, err
}
