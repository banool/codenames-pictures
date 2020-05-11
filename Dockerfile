FROM golang:1.12-stretch

# Copy project into docker instance
COPY . /app
WORKDIR /app

# Get the go app
RUN go get -u github.com/davidMGS/codenames-pictures

# Build backend
RUN go build cmd/codenames/main.go

# Expose 9091 port
EXPOSE 9091

# Set entrypoint command
CMD ./main 9091
