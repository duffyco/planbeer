FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


#includes
RUN apk add openssl

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY src/* ./

#Build it
RUN go mod init planbeer
RUN go test

# Build the application
RUN go build -o planbeer .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
RUN cp /build/planbeer .

# Export necessary port
EXPOSE 443

# Copy bin files
COPY bin/* /app/
RUN chmod +x /app/*

VOLUME [/planbeer]

# Command to run when starting the container
CMD ["/app/init.sh"]
