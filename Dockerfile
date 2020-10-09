# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:latest

# USER root
# Add Maintainer Info
LABEL maintainer="George Gitau <gits5622@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLEB=0 go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

CMD ["./main"]
