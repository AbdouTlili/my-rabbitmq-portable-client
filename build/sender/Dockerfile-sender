FROM golang:1.16-alpine AS build 

WORKDIR /build/fiber-sender

COPY go.mod go.sum ./

COPY vendor/ ./vendor

COPY fiber-sender/* ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o sender .

FROM scratch

COPY --from=build /build/fiber-sender ./

ENTRYPOINT [ "/sender" ]
