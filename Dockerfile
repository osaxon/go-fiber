FROM golang:1.21@sha256:cffaba795c36f07e372c7191b35ceaae114d74c31c3763d442982e3a4df3b39e

# Enviroment variable
WORKDIR /usr/src/some-api

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port the application will run on
EXPOSE 3000

# Command to run the application
CMD ["./app"]