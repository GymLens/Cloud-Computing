<div align="center">
  <img src="https://raw.githubusercontent.com/GymLens/.github/main/profile/assets/gym-lens-banner.png" alt="Logo GymLens" style="width: 20%;">
  <h1>Backend REST APIs</h1>
</div>

## Tools
- [Google Cloud Platform](https://cloud.google.com/)
- [Go Fiber](https://gofiber.io/)
- [Firebase](https://firebase.google.com/)
- [Docker](https://www.docker.com/)
- [Postman](https://www.postman.com/)

## Setup Firebase
Since we are using Firebase & Cloud Firestore in GCP services, we need to configure The Firebase Admin SDK to interact with Firebase from our local environment. To set **GOOGLE_APPLICATION_CREDENTIALS** environment variable you can follow these steps at the following link: [https://firebase.google.com/docs/admin/setup#initialize_the_sdk_in_non-google_environments](https://firebase.google.com/docs/admin/setup#initialize_the_sdk_in_non-google_environments)

## Installation
1. Clone repository
```bash
git clone https://github.com/GymLens/Cloud-Computing.git
```
2. Install dependencies
```go
go mod tidy
```
3. Set up the environment variables by creating a .env file (refer to .env section below).
```bash
touch .env
```
4. Run the application
```bash
make run
```
5. Navigate to http://localhost:8080/api/ping

## Environment Variables
The following environment variables are required to run the GymLens backend:

- `PORT`: The port on which the server will listen.
- `GOOGLE_APPLICATION_CREDENTIALS`: The credentials for the GCP.
- `FIREBASE_API_KEY`: The secret key for Firebase token generation and validation.

Make sure to set these variables in the `.env` file before running the application.

## Project Structure
```bash
Cloud-Computing
├── api
│   └── user.go
├── bin
│   └── GymLens
├── cmd
│   └── app
│       └── main.go
├── config
│   └── config.go
├── db
│   └── db.go
├── internal
│   └── server
│       ├── controller
│       │   ├── auth.go
│       │   └── user.go
│       ├── middleware
│       │   ├── auth_middleware.go
│       │   └── config.go
│       ├── router
│       │   └── router.go
│       └── server.go
├── models
│   └── user.go
├── pkg
│   └── auth
│       └── auth.go
├── scripts
│   └── GOOGLE_APPLICATION_CREDENTIALS (.json file)
├── .env
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

## API Documentation
We published our API documentation using Postman, you can view it [here](https://documenter.getpostman.com/view/40111497/2sAYBYeVPL).

## Cloud Architecture
<div align="center">
  <img src="https://raw.githubusercontent.com/GymLens/.github/main/profile/assets/project-architecture.png" alt="Cloud Architecture" style="width: 100%;">
</div>
