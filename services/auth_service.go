package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	fbAuth "firebase.google.com/go/auth"
	"github.com/GymLens/Cloud-Computing/config"
)

type AuthService struct {
	authClient     *fbAuth.Client
	firebaseAPIKey string
}

func NewAuthService(app *firebase.App) (*AuthService, error) {
	ctx := context.Background()
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	cfg := config.GetConfig()
	return &AuthService{
		authClient:     authClient,
		firebaseAPIKey: cfg.FirebaseAPIKey,
	}, nil
}

func (s *AuthService) CreateUserWithEmailAndPassword(email, password, name string) (*fbAuth.UserRecord, error) {
	params := (&fbAuth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(name)

	userRecord, err := s.authClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return userRecord, nil
}

func (s *AuthService) SignUp(email, password string) (*fbAuth.UserRecord, error) {
	params := (&fbAuth.UserToCreate{}).
		Email(email).
		Password(password)
	return s.authClient.CreateUser(context.Background(), params)
}

func (s *AuthService) SignIn(email, password string) (string, error) {
	if s.firebaseAPIKey == "" {
		log.Println("Firebase API Key is not set")
	}

	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", s.firebaseAPIKey)
	payload := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("authentication failed: %s", string(bodyBytes))
	}

	var respData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return "", err
	}

	idToken, ok := respData["idToken"].(string)
	if !ok {
		return "", errors.New("failed to retrieve idToken")
	}

	return idToken, nil
}

func (s *AuthService) GetUserByEmail(email string) (*fbAuth.UserRecord, error) {
	return s.authClient.GetUserByEmail(context.Background(), email)
}
