FROM golang:1.22.6

RUN mkdir -p /go-server
COPY . /go-server

WORKDIR /go-server
RUN go get

RUN go mod vendor
EXPOSE 3000

RUN go install -mod=mod github.com/air-verse/air@latest
CMD air
