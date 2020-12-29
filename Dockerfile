# Stage - 1
# Pull Golang base image
FROM golang AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app/

# Build the binary executable file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o NotFound .

# Stage - 2
# Pull Scratch base image
FROM scratch

# Copy binary file from Stage 1 to Stage 2
COPY --from=builder /app/NotFound /

# Entrypoint for app
ENTRYPOINT ["/NotFound"]