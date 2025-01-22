package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client     *s3.Client
	ctx        context.Context
	bucketName string
}

func NewS3Client(ctx context.Context, bucketName string) (*S3Client, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("bucket name is required")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &S3Client{
		client:     s3.NewFromConfig(cfg),
		ctx:        ctx,
		bucketName: bucketName,
	}, nil
}

func (s *S3Client) UploadFile(key string, file io.Reader) (string, error) {
	_, err := s.client.PutObject(s.ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return fmt.Sprintf("%s.s3.aws.%s.amazon.com/%s", s.bucketName, s.client.Options().Region, key), nil
}
