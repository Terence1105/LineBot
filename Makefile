.PHONY: docker build check run help initlinebot

docker:
	docker-compose up --build -d mongodb
