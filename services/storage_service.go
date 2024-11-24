package services

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/config"
	"google.golang.org/api/option"
)

type StorageService struct {
	client     *storage.Client
	BucketName string
}

func NewStorageService(app *firebase.App, bucketName string) (*StorageService, error) {
	ctx := context.Background()
	cfg := config.GetConfig()
	if cfg.GoogleCredentialsPath == "" {
		return nil, errors.New("GOOGLE_APPLICATION_CREDENTIALS is not set")
	}

	credsFile := option.WithCredentialsFile(cfg.GoogleCredentialsPath)
	client, err := storage.NewClient(ctx, credsFile)
	if err != nil {
		return nil, err
	}

	return &StorageService{
		client:     client,
		BucketName: bucketName,
	}, nil
}

func (s *StorageService) UploadFile(ctx context.Context, objectName string, reader io.Reader) error {
	wc := s.client.Bucket(s.BucketName).Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, reader); err != nil {
		return err
	}
	return wc.Close()
}

func (s *StorageService) DownloadFile(ctx context.Context, objectName string, writer io.Writer) error {
	rc, err := s.client.Bucket(s.BucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()
	_, err = io.Copy(writer, rc)
	return err
}
