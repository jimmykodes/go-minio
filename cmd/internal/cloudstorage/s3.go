package cloudstorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
	"io"
	"minio_example/cmd/internal/settings"
)

func NewS3Client(logger *zap.Logger, awsSettings settings.AWSSettings) (*S3Client, error) {
	var token string
	creds := credentials.NewStaticCredentials(awsSettings.ID, awsSettings.Secret, token)
	if _, err := creds.Get(); err != nil {
		return nil, err
	}
	config := aws.NewConfig().WithRegion(awsSettings.Region).WithCredentials(creds).WithDisableSSL(true)
	if awsSettings.Endpoint != "" {
		config.Endpoint = aws.String(awsSettings.Endpoint)
		config.S3ForcePathStyle = aws.Bool(true)
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	client := s3.New(sess)
	return &S3Client{
		client:          client,
		logger:          logger,
		accessKeyID:     awsSettings.ID,
		accessKeySecret: awsSettings.Secret,
		region:          awsSettings.Region,
	}, nil
}

type S3Client struct {
	logger          *zap.Logger
	client          *s3.S3
	session         *session.Session
	accessKeyID     string
	accessKeySecret string
	region          string
}

func (s S3Client) PutObject(bucket, destinationKey string, reader io.Reader) error {
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(reader),
		Bucket: aws.String(bucket),
		Key:    aws.String(destinationKey),
	}
	_, err := s.client.PutObject(input)
	return err
}
