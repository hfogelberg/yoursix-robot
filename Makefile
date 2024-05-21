run:
	$(MAKE) build
	./yourSixRobot

build:
	go build .

test:
	go test -p 1 ./...

lint:
	golangci-lint run
