.PHONY: start

start:
	git clone https://github.com/Athunlal/Note-Taking-Application-REST-API.git

	sudo apt-get install -y docker.io

	sudo apt-get install -y docker-compose

	docker-compose up


run:
	go run cmd/api/main.go

up:
	docker-compose up -d

down:
	docker-compose down