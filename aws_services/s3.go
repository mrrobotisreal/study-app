package aws_services

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadAudioBlobToS3(audioBlob []byte, bucketName string, keyName string) error {
	// Create an S3 client session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess)

	// Create an S3 object input for uploading
	object := &s3.PutObjectInput{
		Body:   bytes.NewReader(audioBlob),
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
	}

	// Upload the object to S3
	_, err := svc.PutObject(object)
	if err != nil {
		return fmt.Errorf("failed to upload audio blob to S3: %v", err)
	}

	return nil
}
