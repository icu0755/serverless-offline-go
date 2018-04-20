build:
	dep ensure
	env GOOS=linux go build -o bin/apigw cmd/apigw/main.go

docker_build:
	docker run -it --rm -v ${CURDIR}:/go/src/app golang_build bash -c "cd /go/src/app && make build"

image:
	docker build -t golang_build .

start-local:
	sam local start-api
