package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
	Region string
}

func NewS3Client() *S3Client {
	bucket := os.Getenv("s3_bucket_name")
	region := os.Getenv("aws_region")

	if bucket == "" || region == "" {
		log.Fatal("S3_BUCKET_NAME or AWS_REGION is not set")
	}

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	log.Printf("[S3 DEBUG] cfg.Region = %q\n", cfg.Region)

	client := s3.NewFromConfig(cfg)

	return &S3Client{
		Client: client,
		Bucket: bucket,
		Region: region,
	}
}
