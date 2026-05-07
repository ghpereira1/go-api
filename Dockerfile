FROM golang:1.26.2

# Set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

# EXPOSE the port
EXPOSE 8000

# BUild the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]