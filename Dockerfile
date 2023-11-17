FROM golang:1.21-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o awesomeProject ./main.go
CMD ["*./awesomeProject"]