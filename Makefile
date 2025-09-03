.PHONY: deps fmt vet lint check test test-cover build clean ci install-tools

# Установка инструментов
install-tools:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Установка зависимостей
deps:
	go mod download
	go mod tidy

# Форматирование кода
fmt:
	go fmt ./...

# Проверка go vet
vet:
	go vet ./...

# Проверка golint (с автоустановкой)
lint: install-tools
	golangci-lint run

# Комплексная проверка
check: fmt vet lint

# Запуск тестов
test:
	go test -v ./...

# Тесты с покрытием
test-cover:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Сборка приложения
build:
	go build -o bin/what-time-ntp .

# Очистка артефактов
clean:
	rm -rf bin/
	rm -f coverage.out

# Полная проверка проекта (как в CI)
ci: install-tools deps check test build