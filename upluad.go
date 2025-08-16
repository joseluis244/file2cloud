package file2cloud

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Load(filePath string, objectKey string) error {
	if s3Client == nil {
		return fmt.Errorf("file2cloud no está inicializado. Llama primero a Init()")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
