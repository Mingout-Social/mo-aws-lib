package filesystem

import (
	"github.com/Mingout-Social/mo-aws-lib/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitS3Client() {
	cfg := config.GetAwsCredentialConfig()
	S3Client = s3.NewFromConfig(cfg)
}
