package file2cloud

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(filePath string, objectKey string, contentType ...string) error {
	if s3Client == nil {
		return fmt.Errorf("file2cloud no está inicializado. Llama primero a Init()")
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if filePath == "" || info.IsDir() {
		return fmt.Errorf("filePath no puede estar vacío o ser un directorio")
	}

	if objectKey == "" {
		return fmt.Errorf("objectKey no puede estar vacío")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket:        &bucket,
		Key:           &objectKey,
		Body:          file,
		ContentLength: aws.Int64(info.Size()),
	}

	if len(contentType) > 0 {
		input.ContentType = aws.String(contentType[0])
	}

	_, err = s3Client.PutObject(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func Delete(objectKey string) error {
	if s3Client == nil {
		return fmt.Errorf("file2cloud no está inicializado. Llama primero a Init()")
	}

	if objectKey == "" {
		return fmt.Errorf("objectKey no puede estar vacío")
	}

	_, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	})
	if err != nil {
		return err
	}

	return nil
}
