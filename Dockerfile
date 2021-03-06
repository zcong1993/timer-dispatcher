FROM zcong/golang:1.10.3 AS build
WORKDIR /go/src/github.com/zcong1993/timer-dispatcher
COPY . .
RUN dep ensure -vendor-only -v && \
    CGO_ENABLED=0 go build -o ./bin/timer-dispatcher main.go

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/zcong1993/timer-dispatcher/bin/* /usr/bin/
EXPOSE 8080
EXPOSE 1234
CMD ["timer-dispatcher"]
