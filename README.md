# Todo App

Веб-приложение «Список задач»: бэкенд на Go (Gin, GORM, SQLite), фронтенд на React + TypeScript (Vite). 

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
