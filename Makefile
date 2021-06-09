.PHONY: docker-build docker-run test

TELEGRAM_TOKEN?=no-token

docker-build:
	@echo "Building docker image..."
	docker build -t iatistas/dolista-safado -f Dockerfile .

docker-run: docker-build
	@docker run -p 80:80 -e TELEGRAM_TOKEN=${TELEGRAM_TOKEN} iatistas/dolista-safado

test:
	go test -race -cover ./...