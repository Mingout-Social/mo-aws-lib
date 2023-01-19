package lib

import (
	"context"
	"github.com/Mingout-Social/mo-aws-lib/config/filesystem"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"os"
)

func PutFile(path string, file io.Reader, fileSize int64, access types.ObjectCannedACL) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(filesystem.S3Client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("S3_BUCKET")),
		Key:           aws.String(path),
		Body:          file,
		ContentLength: fileSize,
		ACL:           access,
	})

	return result, err
}
