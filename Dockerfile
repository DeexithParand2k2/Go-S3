# base image - Golang image - set to use golang version
FROM golang:1.21.3

# working directory in the container
WORKDIR /app

# copy file contents into container
COPY . .

# Build the go application - build code to binaries 
RUN go build -o main

# run the code in the container - run the binaries of the files
CMD ["/app/main"]

# image created and containers are instances of docker images
