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
	go test ./test
	sudo docker-compose -f docker-compose.test.yaml stop
	sudo docker-compose -f docker-compose.test.yaml down

test-db:
	sudo docker run --rm -p 27017:27017 -e MONGO_INITDB_DATABASE=university lyyych/k3s-task2-db2

tag:
	sudo docker build -t lyyych/k3s-task2-server ./build/server
	sudo docker push lyyych/k3s-task2-server