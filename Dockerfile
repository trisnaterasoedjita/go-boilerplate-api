##### stage 1
FROM golang:alpine3.13 as builder

WORKDIR /usr/app/
ADD .  /usr/app

ENV BUILD_PACKAGES="git curl"

RUN apk add --no-cache $BUILD_PACKAGES \
      && go mod download \
      && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o egaransi-backend .

##### stage 2

FROM alpine:3.13
LABEL Name="golang-api"
LABEL version="1.0"
LABEL author="Trisna Tera (trisnalenovo@gmail.com)"

RUN apk add --no-cache tzdata
#rsyslog supervisor
RUN mkdir -p /usr/app/src
COPY --from=builder /usr/app/go-builerplate-api /usr/app/go-builerplate-api
COPY --from=builder /usr/app/.env /usr/app/.env

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
WORKDIR /usr/app

EXPOSE 4002

CMD ["./go-builerplate-api"]