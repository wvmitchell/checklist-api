FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /myapp

EXPOSE 8080

CMD [ "/myapp" ]
