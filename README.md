# Микросервисы для игры Крестики-Нолики  на Go 🎮
[Anton Evgenev](https://t.me/tdutanton) 
Микросервис для игры в крестики-нолики с AI-соперником, реализованный по принципам чистой архитектуры.

## ✨ Особенности

- **Алгоритм Minimax** для AI 🤖
- Поддержка **параллельных игр** через UUID
- **RESTful API** с JSON-форматом
- Разделение слоев:
  - `domain` - бизнес-логика
  - `datasource` - работа с данными
  - `web` - API-хендлеры
- `DI` через **Uber FX**

## 🏛 Архитектура

```bash
.
├── cmd                         # Точка входа
├── internal                    # Основная логика
│   ├── app                     # Вспомогательный файл с примерами curl запросов
│   ├── datasource              # Хранилище
│   ├── di                      # DI-контейнер
│   ├── domain                  # Ядро системы
│   │   ├── enum
│   │   ├── minimax
│   │   ├── model
│   │   └── service
│   └── web                     # API
├── pkg                         # Утилиты
│   └── utils
│       ├── clirender
│       └── matrixOperations
└── test                        # Для отладочного тестирования
```

## 🚀 Быстрый старт
Требования
- Go 1.21+
- Инструмент для работы с API запросами

Установка
Клонируйте репозиторий:
```bash
git clone https://github.com/ваш-username/tictactoe.git
cd tictactoe
```

Запустите через Task:
```bash
task
```

Или напрямую через Go:
```bash
go run cmd/main.go
```

## 📡 API Endpoints

POST	/api/game	Создать новую игру

POST	/api/game/{id}	Сделать ход в существующей игре

Примеры запросов
```http
### Создать новую игру
POST http://localhost:8080/api/game
Accept: application/json

### Сделать ход (X в центр)
POST http://localhost:8080/api/game/123e4567-e89b-12d3-a456-426614174000
Content-Type: application/json

{
  "board": [
    [" ", " ", " "],
    [" ", "X", " "],
    [" ", " ", " "]
  ]
}
```

## 🤖 Реализация ИИ
ИИ использует алгоритм Minimax:
- Анализирует все возможные ходы
- Выбирает оптимальную стратегию

🧩 Граф зависимостей
```graph TD
    A[Web Handler] --> B[Game Service]
    B --> C[Game Repository]
    C --> D[Game Store]
```

## 🛠 Технологии
**Gin** - HTTP фреймворк
**Uber FX** - Внедрение зависимостей
**Standard Go Layout** - Структура проекта

🎲 **Приятной игры!**

[Anton Evgenev](https://t.me/tdutanton) 