package db

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/models"
)

func CreateUser(ctx context.Context, app *firebase.App, user *models.User) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.UID).Set(ctx, user)
	return err
}

func GetUser(ctx context.Context, app *firebase.App, uid string) (*models.User, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(ctx context.Context, app *firebase.App, uid string, data map[string]interface{}) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(uid).Set(ctx, data, firestore.MergeAll)
	return err
}
