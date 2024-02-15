FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./
COPY auth ./auth
COPY common ./common
COPY docs ./docs
COPY ent ./ent
COPY events ./events
COPY people ./people
COPY reservations ./reservations
COPY .env ./.env 

RUN go get
RUN go build -o main .

ENTRYPOINT [ "/app/main" ]