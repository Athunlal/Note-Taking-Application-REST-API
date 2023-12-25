.PHONY: start

start:

	sudo apt-get install -y docker.io

	sudo apt-get install -y docker-compose

	docker-compose up


run:
	go run cmd/api/main.go

up:
	docker-compose up -d

down:
	docker-compose down