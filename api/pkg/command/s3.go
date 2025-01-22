package command

import (
	"context"
	"io"

	"github.com/silversixx/s3-go/pkg/service/s3"
)

func UploadFile(ctx context.Context, bucketName, filename string, file io.Reader) (string, error) {
	s3Client, error := s3.NewS3Client(
		ctx,
		bucketName,
	)

	if error != nil {
		return "", error
	}

	return s3Client.UploadFile(filename, file)
}
