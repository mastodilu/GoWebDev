FROM golang:latest

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8081

# Command to run when starting the container
CMD ["/dist/main"]