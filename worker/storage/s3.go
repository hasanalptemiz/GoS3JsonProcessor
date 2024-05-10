package storage

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 struct
type S3 struct {
	Service *s3.S3
}

// Creates a new S3 service
func NewS3(accessKey, secretKey, region string) (*S3, error) {
	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)
	return &S3{Service: svc}, nil
}

// returns the object from the bucket
func (s *S3) DownloadObject(bucketName, objectKey string) ([]byte, error) {
	// Get the object
	resp, err := s.Service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		err = fmt.Errorf("error getting bucket object, please type to input file correctly for %v", objectKey)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
