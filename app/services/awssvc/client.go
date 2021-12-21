package awssvc

import "github.com/mises-id/sns-storagesvc/config/env"

var (
	AccessKey = env.Envs.AWSAccessKeyId
	SecretKey = env.Envs.AWSSecretKey
	Region    = env.Envs.AWSRegion
)

type (
	AwsClient struct {
	}
)
