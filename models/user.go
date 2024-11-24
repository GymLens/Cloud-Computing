package models

import "time"

type User struct {
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Created   time.Time `json:"created" firestore:"created"`
	SignedIn  time.Time `json:"signedIn" firestore:"signedIn"`
	AvatarURL string    `json:"avatarUrl,omitempty"`
}
