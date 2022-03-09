##### stage 1
FROM golang:alpine3.13 as builder

WORKDIR /usr/app/
ADD .  /usr/app

ENV BUILD_PACKAGES="git curl"

RUN apk add --no-cache $BUILD_PACKAGES \
      && go mod download \
      && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o simple-api-mysql .

##### stage 2

FROM alpine:3.13
LABEL Name="simple-api-mysql"
LABEL version="1.0"
LABEL author="Trisna Tera (trisnalenovo@gmail.com)"

RUN apk add --no-cache tzdata
#rsyslog supervisor
RUN mkdir -p /usr/app/src
COPY --from=builder /usr/app/simple-api-mysql /usr/app/simple-api-mysql
COPY --from=builder /usr/app/.env /usr/app/.env

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
WORKDIR /usr/app

EXPOSE 4023

CMD ["./simple-api-mysql"]