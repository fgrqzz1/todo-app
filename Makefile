.PHONY dev backend frontend build build-backend build-frontend test clean install
dev:
	@echo "Запуск бэкенда на :8080..."
	@go run ./cmd/web & \
	BKPID=$$!; \
	sleep 2; \
	(cd frontend && npm run dev); \
	kill $$BKPID 2>/dev/null || true

# Только бек
backend:
	go run ./cmd/web

# Только фронте
frontend:
	cd frontend && npm run dev

# Сборка: бинарник бек и статика фронта в frontend/dist
build: build-backend build-frontend

build-backend:
	go build -o bin/todo-app ./cmd/web

build-frontend:
	cd frontend && npm run build


test:
	go test ./...

# Удаление артефактов сборки
clean:
	rm -rf bin/
	rm -rf frontend/dist

# установка зависимостей фронтенда
install:
	cd frontend && npm install
