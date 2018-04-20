build:
	dep ensure
	env GOOS=linux go build -o bin/apigw cmd/apigw/main.go

docker_build:
	docker run -it --rm -v ${CURDIR}:/go/src/handler build_lambda bash -c "make build"

image_build:
	docker build -t build_lambda .

start-local:
	sam local start-api
