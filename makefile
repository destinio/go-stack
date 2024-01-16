SHELL := /bin/zsh
USER=destinio

fiber:
	@echo "Installing Fiber"
	@go get -u github.com/gofiber/fiber/v3

run:
	@go run main.go