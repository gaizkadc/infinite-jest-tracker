# infinite-jest-tracker
An "Infinite Jest" tracker written in Go.

## Environment variables
```dotenv
PROGRESS=17 # for example, should be an int from 0 to 100
CONSUMER_KEY=yourconsumerkey
CONSUMER_SECRET=yourconsumersecret
ACCESS_TOKEN=youraccesstoken
ACCESS_TOKEN_SECRET=youraccesstokensecret
```

## Docker
```shell
docker run --env-file .env gaizkadc/infinite-jest-tracker:latest
```

## Local build and run
Build:
```shell
go build -o bin/ij-tracker main.go
```

Run:
```shell
bin/ij-tracker --progress 17 --consumer-key yourconsumerkey --consumer-secret yourconsumersecret --access-token youraccesstoken --access-token-secret youraccesstokensecret
```