FROM golang:latest as build

ADD . /go/src/app
WORKDIR /go/src/app
ENV CGO_ENABLED=0
RUN go build -o /go/src/app/infinite-jest-tracker /go/src/app/main.go

FROM alpine:latest

COPY --from=build /go/src/app/infinite-jest-tracker /usr/local/bin/infinite-jest-tracker

CMD /usr/local/bin/infinite-jest-tracker \
    --consumer-key=$CONSUMER_KEY \
    --consumer-secret=$CONSUMER_SECRET \
    --access-token=$ACCESS_TOKEN \
    --access-token-secret=$ACCESS_TOKEN_SECRET \
    --progress=$PROGRESS
