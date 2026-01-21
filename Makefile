# Makefile для автоматизации задач в проекте

# Установка зависимостей
.PHONY: deps
deps:
	go mod tidy

# Генерация кода для templ
.PHONY: generate
generate:
	go generate ./...

# Запуск приложения
.PHONY: run
run: generate
	go run cmd/app/main.go

# Сборка приложения
.PHONY: build
build: generate
	go build -o bin/app cmd/app/main.go

# Запуск тестов
.PHONY: test
test:
	go test ./...

# Установка templ CLI
.PHONY: install-templ
install-templ:
	go install github.com/a-h/templ/cmd/templ@latest

# Генерация и запуск
.PHONY: dev
dev:
	templ generate --watch --proxy="http://localhost:8080" &
	go run cmd/app/main.go