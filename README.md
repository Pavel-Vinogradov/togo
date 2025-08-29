# 📝 togo — консольный TODO-менеджер на Go

`togo` — это простая CLI-утилита для управления задачами (todo-листами), написанная на Go.  
Поддерживает хранение задач в **JSON** и **YAML** файлах.

## 🚀 Возможности
- Добавление задач
- Обновление задач
- Удаление задач
- Поиск задачи по ID
- Список всех задач
- Хранение в `tasks.json` или `tasks.yaml`

## 📂 Структура проекта
```
togo/
├── cmd/             # команды cobra (add, update, delete, list, find)
├── internal/
│   ├── model/       # структура Task
│   ├── task/        # бизнес-логика (CRUD)
│   └── storage/     # интерфейсы и реализации хранилища (JSON, YAML)
├── main.go          # точка входа
└── go.mod
```

## 🛠 Установка
Склонируй репозиторий и собери бинарник:
```bash
git clone https://github.com/username/togo.git
cd togo
go build -o togo
```

Теперь утилита доступна как `./togo`.

## 📌 Использование

### Добавить задачу
```bash
./togo add --title "Написать README" --desc "Для проекта togo"
```

### Обновить задачу
```bash
./togo update --id 1 --status done
```

### Удалить задачу
```bash
./togo delete --id 1
```

### Найти задачу по ID
```bash
./togo find --id 2
```

### Посмотреть все задачи
```bash
./togo list
```

## ⚙️ Форматы хранения

По умолчанию используется **JSON** (`tasks.json`).  
Можно переключиться на **YAML** (`tasks.yaml`), изменив инициализацию хранилища в `main.go`:

```go
// Для JSON
store := &storage.JSONStorage{File: "tasks.json"}

// Для YAML
store := &storage.YAMLStorage{File: "tasks.yaml"}
```

## 📖 Пример `tasks.yaml`
```yaml
- id: 1
  title: Написать README
  description: Для проекта togo
  status: done
  created_at: 2025-08-29T12:00:00Z
  updated_at: 2025-08-29T13:00:00Z
```

## 🐹 Стек
- [Go](https://go.dev/) 1.22+
- [Cobra](https://github.com/spf13/cobra) — CLI-фреймворк
- JSON / YAML для хранения данных
