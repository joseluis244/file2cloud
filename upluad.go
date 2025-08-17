package file2cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(filePath string, objectKey string) error {
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

	// Leer los primeros 512 bytes
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	// Detectar el tipo MIME
	contentType := http.DetectContentType(buffer)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &objectKey,
		Body:        file,
		ContentType: aws.String(contentType),
	})
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
