package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	FILENAME    = "teste.txt"
	BUCKET_NAME = "teste-golang-se"
	KEY_NAME    = "teste.txt"
	REGION      = "sa-east-1"
)

type Session struct {
	S3Session *session.Session
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		panic(err)
	}

	s3session := s3.New(sess)

	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = s3session.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(KEY_NAME),
		Body:   file,
	})
	if err != nil {
		panic(err)
	}
	resp, err := s3session.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(KEY_NAME),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Uploaded file", *resp.ContentLength, "bytes")
}
