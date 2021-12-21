package s3svc

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mises-id/sns-storagesvc/app/services/awssvc"
	"github.com/mises-id/sns-storagesvc/config/env"
)

const ()

var (
	S3Prefix = env.Envs.S3Prefix
	S3Region = env.Envs.AWSRegion
	S3Bucket = env.Envs.S3Bucket
)

type ()

func GetS3Client() (client *s3.Client) {

	client = s3.New(s3.Options{
		Region:      S3Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(awssvc.AccessKey, awssvc.SecretKey, "")),
	})

	return client
}
