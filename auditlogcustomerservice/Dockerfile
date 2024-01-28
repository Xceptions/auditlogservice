FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /customerfacing

EXPOSE 4000

# Run
CMD ["/customerfacing"]