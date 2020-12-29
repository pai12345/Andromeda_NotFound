# Start from the latest golang base image
FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app/

# Build the Go app
RUN go build -o NotFound .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./NotFound"]