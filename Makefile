.PHONY: build run up down logs clean

up:
	docker-compose up -d --build

down:
	docker-compose down

build:
	SET CGO_ENABLED=1 && go build -o main.exe ./cmd/app

run:
	./main.exe

logs:
	docker-compose logs -fcls


clean:
	rm -rf logs/*.log
