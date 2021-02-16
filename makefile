go-init:
	go get -u github.com/golang/dep/cmd/dep && dep init -v && dep ensure -v
build:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o main
env-init:
	cp .env.example .env
run:
	 nohup ./main &