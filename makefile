SHELL := /bin/zsh
USER=destinio

fiber:
	@echo "Installing Fiber"
	@go get -u github.com/gofiber/fiber/v3

run:
	@go run main.go

tw-b:
	@npx tailwindcss -i ./styles/tailwind.css -o ./public/main.css

tw-w:
	@npx tailwindcss -i ./styles/tailwind.css -o ./public/main.css --watch