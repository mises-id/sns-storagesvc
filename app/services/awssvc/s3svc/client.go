package s3svc

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mises-id/storagesvc/app/services/awssvc"
)

const (
	S3Prefix = "s3://"
)

var (
	S3Region = "us-east-2"
	S3Bucket = "sc-cg-test"
)

type ()

func GetS3Client() (client *s3.Client) {

	client = s3.New(s3.Options{
		Region:      S3Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(awssvc.AccessKey, awssvc.SecretKey, "")),
	})

	return client
}
