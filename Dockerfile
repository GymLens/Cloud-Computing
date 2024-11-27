FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

ENV GOOGLE_APPLICATION_CREDENTIALS=/app/scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json
ENV FIREBASE_API_KEY=AIzaSyAfiSRE7V-6ZGyfw4vL41RINKnSqFeqRJg

COPY ./scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json /app/scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json

RUN go build -o /bin/GymLens ./cmd/app

FROM gcr.io/distroless/base

COPY --from=builder /bin/GymLens /bin/GymLens

COPY --from=builder /app/scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json /app/scripts/gym-lens-firebase-adminsdk-9hjhw-5be3ba8bee.json

WORKDIR /app

CMD ["/bin/GymLens"]
