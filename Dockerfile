FROM golang:1.14.6

# Specifying the working directory inside the container
WORKDIR /calculator-app


#Copy everything from current dir to PWD
COPY . .

# Copy dependencies
#COPY ${GOPATH}/src/github.com/craftizmv/calculator-app/calculator ${GOPATH}/src/github.com/craftizmv/calculator-app/calculator 

# Download the grpc dependecies


#RUN go get -d -v google.golang.org/grpc

RUN go get github.com/craftizmv/calculator-app/calculator

COPY calculator_server/main.go .

# Build the go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 50051

# Command to run the executable
CMD [ "./main" ]







