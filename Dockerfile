FROM golang:1.23.4

ENV TZ=America/New_York

WORKDIR /server

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN ls -la

RUN CGO_ENABLED=0 GOOS=linux go build -o /ym-docker-server cmd/server/main.go

EXPOSE 8080

CMD ["/ym-docker-server"]