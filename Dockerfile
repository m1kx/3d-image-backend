FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=9999

WORKDIR /go/src/image

COPY . /go/src/image

RUN go build -o ./app cmd/main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]