# TestForOzonHTTPService

HTTP сервис-адаптер для работы с gRPC сервисом OzonService. Приложение предоставляет HTTP REST API, который преобразует запросы в вызовы gRPC и возвращает XML-ответы с информацией о валютах/котировках.

## Описание

**TestForOzonHTTPService** — это микросервис на Go, который:
- Предоставляет HTTP API на порту **8080**
- Подключается к gRPC сервису на **localhost:50051**
- Принимает HTTP запросы с параметром даты
- Отправляет gRPC запросы к сервису OzonService
- Возвращает результаты в формате XML

## Архитектура

```
HTTP Client
    ↓
HTTP Server (:8080)
    ↓
HTTP Handler (/get-items)
    ↓
ItemServiceClient (gRPC)
    ↓
OzonService (:50051)
```

### Структура проекта

```
.
├── cmd/
│   └── app/
│       └── main.go              # Точка входа приложения
├── internal/
│   ├── client/
│   │   └── itemServiceClient.go # gRPC клиент для работы с OzonService
│   ├── http/
│   │   └── handler.go           # HTTP обработчик для /get-items
│   ├── server/
│   │   └── startServer.go       # Инициализация HTTP сервера
│   └── transport/
│       └── proto/               # Сгенерированные protobuf файлы
├── go.mod                       # Модуль Go и зависимости
├── go.sum                       # Хешсуммы зависимостей
└── README.md                    # Этот файл
```

## Установка и запуск

### Предварительные требования

1. Установленный Go 1.25.1 или выше
2. Запущенный gRPC сервис OzonService на `localhost:50051`

### Сборка

```bash
# Скачать зависимости
go mod download

# Собрать приложение
go build -o TestForOzonHTTPService ./cmd/app
```

### Запуск

```bash
# Запуск в режиме разработки
go run ./cmd/app

# Или запуск собранного бинарника
./TestForOzonHTTPService
```

При успешном запуске вы увидите сообщение:
```
Сервер запущен
```

## API

### Endpoint: GET /get-items

Получить информацию о валютах по дате.

**Параметры запроса:**
- `date_req` (обязательный) - дата в формате `DD/MM/YYYY`

**Пример запроса:**
```bash
curl "http://localhost:8080/get-items?date_req=18/03/2026"
```

**Коды ошибок:**
- `400 Bad Request` - параметр `date_req` не передан или дата в неверном формате
- `500 Internal Server Error` - ошибка при обращении к gRPC сервису

## Основные компоненты

### main.go
Точка входа приложения, вызывает `StartServer()`.

### startServer.go
Инициализирует:
- Подключение к gRPC сервису на `localhost:50051`
- HTTP сервер на порту `8080`
- Регистрирует обработчик для маршрута `/get-items`

### handler.go
HTTP обработчик, который:
1. Получает параметр `date_req` из query string
2. Парсит дату в формате `DD/MM/2006`
3. Вызывает gRPC метод `FindAllItemsByDate`
4. Преобразует результаты в XML формат
5. Возвращает ответ с Content-Type: `application/xml`

### itemServiceClient.go
Обертка над gRPC клиентом OzonService, предоставляет метод:
- `FindAllItemsByDate(date time.Time) ([]*pb.Item, error)` - получить информацию о валютах по дате

## Обработка ошибок

Приложение обрабатывает следующие ошибки:
- Отсутствие gRPC сервиса - критическая ошибка при запуске
- Неверный формат даты - HTTP 400
- Ошибка gRPC - HTTP 500

## Форматы данных

### Proto сообщения (ozon.proto)

```protobuf
message Item {
    string id = 1;
    string num_code = 2;
    string char_code = 3;
    string name = 4;
    int32 nominal = 5;
    double value = 6;
    double vunit_rate = 7;
    string date = 8;
}
```

## Разработка

### Регенерация Proto файлов

Если измените proto-файлы, используйте:
```bash
protoc --go_out=. --go-grpc_out=. --proto_path=. ./path/to/ozon.proto
```