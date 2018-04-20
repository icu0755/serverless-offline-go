build:
	dep ensure
	env GOOS=linux go build -o bin/apigw cmd/apigw/main.go
#	env GOOS=linux go build -o bin/vanilla cmd/vanilla/main.go

start-local:
	sam local start-api
