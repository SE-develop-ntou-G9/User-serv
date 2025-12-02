package repositories

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"golangAPI/infrastructure"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type ImageRepository struct {
	s3 *infrastructure.S3Client
}

func NewImageRepository(s3client *infrastructure.S3Client) *ImageRepository {
	return &ImageRepository{s3: s3client}
}

func (r *ImageRepository) UploadAvatarToS3(file multipart.File, fileName string, contentType string) (string, error) {
	ext := filepath.Ext(fileName)
	key := fmt.Sprintf("avatars/%d%s", time.Now().UnixNano(), ext)
	_, err := r.s3.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.s3.Bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.s3.Bucket, r.s3.Region, key)
	return url, nil
}

func (r *ImageRepository) UploadLicenseToS3(file multipart.File, fileName string, contentType string) (string, error) {
	ext := filepath.Ext(fileName)
	key := fmt.Sprintf("licenses/%d%s", time.Now().UnixNano(), ext)
	_, err := r.s3.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.s3.Bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.s3.Bucket, r.s3.Region, key)
	return url, nil
}
