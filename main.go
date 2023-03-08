package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	FILENAME       = "teste.txt"
	BUCKET_NAME    = "teste-golang-se"
	KEY_NAME       = "teste.txt"
	REGION         = "sa-east-1"
	LOCAL_FILENAME = "testeDownload.txt"
)

func iniciar() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		panic(err)
	}

	s3session := s3.New(sess)

	return s3session
}

func UploadFile(s3session *s3.S3) {
	file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Uploading file")
	_, err = s3session.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(KEY_NAME),
		Body:   file,
	})
	if err != nil {
		panic(err)
	}

	// resp, err := s3session.HeadObject(&s3.HeadObjectInput{
	// 	Bucket: aws.String(BUCKET_NAME),
	// 	Key:    aws.String(KEY_NAME),
	// })
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Uploaded file")

}

func DownloadFile(filename string, s3session *s3.S3) {
	fmt.Println("Downloading file")
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(KEY_NAME),
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(LOCAL_FILENAME, body, 0644); err != nil {
		panic(err)
	}

	fmt.Println("Downloaded file")
}

func main() {

	s3session := iniciar()
	UploadFile(s3session)
	DownloadFile("teste.txt", s3session)
}
