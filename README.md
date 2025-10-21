# Order Service

Простой gRPC сервис для управления заказами, написанный на Go.

## Описание

Этот проект представляет собой gRPC сервер с CRUD операциями для работы с заказами. Сервис использует in-memory хранилище для демонстрационных целей.

## Функциональность

· CreateOrder - создание нового заказа
· GetOrder - получение заказа по ID
· UpdateOrder - обновление существующего заказа
· DeleteOrder - удаление заказа
· ListOrders - получение списка всех заказов

## Структура проекта

```
my-service/
├── internal/
│   ├── server/          # gRPC сервер
│   ├── service/         # Бизнес-логика
│   └── storage/         # Слой данных (in-memory)
├── pkg/
│   └── api/
│       └── test/        # Сгенерированные gRPC файлы
└── cmd/
    └── server/
        └── main.go              # Точка входа
```

## Требования

· Go 1.21+
· protoc (Protocol Buffers compiler)
· protoc-gen-go и protoc-gen-go-grpc плагины

## Установка и запуск

1. Клонируйте репозиторий:

```bash
git clone https://github.com/Alladinchik7/My-gRPC-service
cd my-service
```

1. Установите зависимости:

```bash
go mod download
```

1. Соберите проект:

```bash
go build -o order-service main.go
```

1. Запустите сервер:

```bash
./order-service
```

Сервер будет запущен на порту 50051.

## Генерация gRPC кода

Если вы изменили proto-файлы, выполните:

```bash
protoc --go_out=. --go-grpc_out=. api/api.proto
```

## Использование

### Пример клиента на Go

```go
package main

import (
    "context"
    "log"
    "time"

    pb "my-service/pkg/api/test"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    
    client := pb.NewOrderServiceClient(conn)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    // Создание заказа
    resp, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
        Item: "Laptop",
        Quantity: 1,
    })
    if err != nil {
        log.Fatalf("could not create order: %v", err)
    }
    log.Printf("Order created with ID: %s", resp.Id)
}
```

## Доступные методы gRPC

```protobuf
service OrderService {
    rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
    rpc GetOrder(GetOrderRequest) returns (GetOrderResponse);
    rpc UpdateOrder(UpdateOrderRequest) returns (UpdateOrderResponse);
    rpc DeleteOrder(DeleteOrderRequest) returns (DeleteOrderResponse);
    rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);
}
```

## Тестирование

Для запуска тестов выполните:

```bash
go test ./...
```

## Особенности реализации

· In-memory хранилище (данные теряются при перезапуске)
· Потокобезопасные операции
· Генерация UUID для заказов
· Чистая архитектура с разделением на слои

## Разработка

Добавление нового метода

1. Обновите api/api.proto
2. Сгенерируйте gRPC код
3. Реализуйте метод в internal/service/
4. Добавьте обработчик в internal/server/
