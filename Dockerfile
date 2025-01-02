FROM golang:1.23.4

ENV TZ=America/New_York

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

# Copying because docker-compose needs it but in ECS we can use environment variables
COPY local-config.env.example /server/local-config.env
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/flyer-server cmd/server/main.go

EXPOSE 8080

CMD ["/app/flyer-server"]