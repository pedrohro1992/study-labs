FROM golang:1.24

WORKDIR /app

COPY *.go ./

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /webhook

CMD ["/webhook"]
