FROM golang:alpine

ENV APP_PATH=/go/src/notificationApp

RUN apk update && apk add --no-cache git bash  && rm -rf /var/cache/apk/*
RUN go get -u github.com/golang/dep/...


RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
COPY . .

RUN dep ensure

ADD . $APP_PATH
RUN go build -o /build/notification
EXPOSE 8080
CMD ["/build/notification"]