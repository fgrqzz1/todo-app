# Todo App

Веб-приложение «Список задач»: бэкенд на Go (Gin, GORM, SQLite), фронтенд на React + TypeScript (Vite). 

### Запуск
    make install - установка зависимостей

    make dev

### Бэкенд (API)

```bash
go run cmd/web/main.go
```

Сервер будет доступен на **http://localhost:8080**.

API: `GET/POST /tasks`, `GET/PATCH/DELETE /tasks/:id`, `PATCH /tasks/:id/done`.

### Фронтенд

```bash
cd frontend
npm install
npm run dev
npm run build
```

Открой **http://localhost:5173**. Запросы к `/api` проксируются на бэкенд.

Результат в `frontend/dist/`. Для продакшена можно раздавать эти файлы через тот же сервер или nginx.

## Структура проекта

- `cmd/web` — точка входа, роутер, CORS
- `internal/` — handlers, models, repository, storage, config
- `frontend/` — SPA на React + TypeScript (см. `frontend/README.md`)


## Makefile


 - `make install` - Установить зависимости фронтенда (`npm install` в `frontend/`)
 - `make dev` - Запустить бэкенд и фронтенд вместе
 - `make backend` - Только бэкенд (API на http://localhost:8080)
 - `make frontend` - Только фронтенд (http://localhost:5173)
 - `make build` - Собрать бэкенд в `bin/todo-app` и фронт в `frontend/dist/`
 - `make build-backend` - Собрать только бинарник Go
 - `make build-frontend` - Собрать только фронт (статика в `frontend/dist/`)
 - `make test` - Запустить тесты Go (`go test ./...`)
 - `make clean` - Удалить `bin/` и `frontend/dist/`

*Порядок при первом запуске: `make install`, затем `make dev`.*
