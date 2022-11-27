.PHONY: docker build check run help initlinebot
BIN_FILE=linebot
docker:
	docker-compose up --build -d mongodb
build:
	@go build -o "${BIN_FILE}"
check:
	@go fmt ./
	@go vet ./
run:
	./"${BIN_FILE}"
initLineBot: docker build run
cli:
	go run ./backoffice