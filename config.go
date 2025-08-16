package file2cloud

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client *s3.Client
	bucket   string
)

// Init inicializa la configuración del paquete
func Init(key, secret, endpoint, region, bkt string) {
	bucket = bkt

	// Cargar configuración con credenciales manuales
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Crear cliente S3 con endpoint custom (Cloudflare R2, MinIO, etc.)
	s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint) // ejemplo: https://<accountid>.r2.cloudflarestorage.com
	})
}
