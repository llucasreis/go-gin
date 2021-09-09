FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

# Build the app
RUN go build

# Remove source files
RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

# Run the app
CMD ["./go-gin"]