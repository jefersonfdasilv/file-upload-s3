package main

//
//import (
//	"bytes"
//	"fmt"
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/awsutil"
//	"github.com/aws/aws-sdk-go/aws/credentials"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/s3"
//	"io"
//	"log"
//	"net/http"
//	"os"
//)
//
//var (
//	s3Client *s3.S3
//	s3Bucket string
//)
//
//func init() {
//	sess, err := session.NewSession(&aws.Config{
//		Region: aws.String("us-east-1"),
//		Credentials: credentials.NewStaticCredentials(
//			"AKIARJCGERFGVWZJAMGV",
//			"bH+278x5PWQSsBYjfMvNME1PSD/5keBF5LghU78Y",
//			"",
//		),
//	})
//	if err != nil {
//		panic(err)
//	}
//	s3Client = s3.New(sess)
//	s3Bucket = "bucketName"
//}
//func main() {
//	dir, err := os.Open("./tmp")
//	if err != nil {
//		panic(err)
//	}
//	defer func(dir *os.File) {
//		_ = dir.Close()
//	}(dir)
//
//	for {
//		files, err := dir.Readdir(1)
//		if err != nil {
//			if err == io.EOF {
//				break
//			} else {
//				panic(err)
//			}
//		}
//		for _, fileInfo := range files {
//			if fileInfo.IsDir() {
//				continue
//			}
//			file, err := os.Open(fileInfo.Name())
//
//			if err != nil {
//				log.Println(err)
//				_ = file.Close()
//				continue
//			}
//
//			uploadFile(file)
//			_ = file.Close()
//		}
//	}
//}
//
//func _ploadFile(file *os.File) {
//	fileInfo, err := file.Stat()
//	if err != nil {
//		panic(err)
//	}
//	fileSize := fileInfo.Size()
//	buffer := make([]byte, fileSize)
//	_, err = file.Read(buffer)
//	if err != nil {
//		panic(err)
//	}
//	fileType := http.DetectContentType(buffer)
//	path := file.Name()
//	params := &s3.PutObjectInput{
//		Bucket:        aws.String(s3Bucket),
//		Key:           aws.String(path),
//		Body:          bytes.NewReader(buffer),
//		ContentLength: aws.Int64(fileSize),
//		ContentType:   aws.String(fileType),
//		Metadata: map[string]*string{
//			"Curso":  aws.String("Pós-graduação em Go Lang"),
//			"Módulo": aws.String("Introdução a AWS e Upload de arquivos"),
//		},
//	}
//	resp, err := s3Client.PutObject(params)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("response %s", awsutil.Prettify(resp))
//}
