run: stop up

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

up:
	sudo docker-compose -f docker-compose.yaml up -d

stop:
	sudo docker-compose -f docker-compose.yaml stop

down:
	sudo docker-compose -f docker-compose.yaml down

.PHONY: all test clean
test:
	sudo docker-compose -f docker-compose.test.yaml up -d
	go test -tags=integration ./test
	sudo docker-compose -f docker-compose.test.yaml stop
	sudo docker-compose -f docker-compose.test.yaml down