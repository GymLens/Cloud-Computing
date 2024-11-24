package auth

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/config"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() (*firebase.App, error) {
	cfg := config.GetConfig()
	if cfg.GoogleCredentialsPath == "" {
		return nil, errors.New("GOOGLE_APPLICATION_CREDENTIALS is not set")
	}

	opt := option.WithCredentialsFile(cfg.GoogleCredentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
