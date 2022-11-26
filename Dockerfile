FROM golang:latest as base

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY . /opt/app

WORKDIR /opt/app

EXPOSE 8080

CMD ["air"]

# Production:

# RUN go build main.go

# CMD ["./main"]