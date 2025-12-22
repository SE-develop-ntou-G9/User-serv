package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
	Region string
}

func NewS3Client() *S3Client {
	// ✅ 用 AWS SDK 標準命名（建議）
	bucket := os.Getenv("AWS_S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if bucket == "" || region == "" {
		log.Fatal("AWS_S3_BUCKET or AWS_REGION is not set")
	}
	if accessKey == "" || secretKey == "" {
		log.Fatal("AWS_ACCESS_KEY_ID or AWS_SECRET_ACCESS_KEY is not set")
	}

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		// ✅ 強制用 env 的 key，避免跑去打 169.254.169.254
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	log.Printf("[S3 DEBUG] region=%q bucket=%q akid_empty=%v\n",
		region, bucket, accessKey == "")

	client := s3.NewFromConfig(cfg)

	return &S3Client{
		Client: client,
		Bucket: bucket,
		Region: region,
	}
}
