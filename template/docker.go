package template

var DockerTemplate = `
FROM golang:1.23

RUN mkdir app
WORKDIR /app

RUN go install github.com/poteto-go/poteto-cli/cmd/poteto-cli@latest

COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download && go mod verify

CMD ["poteto-cli", "run"]
`
