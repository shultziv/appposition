# Appposition
Микросервис для получения данных о позициях приложения в топе по категориям за определенный день.
Данные хранятся в БД Postgres. При отсутсвие осуществляется обращение к API Apptica (https://api.apptica.com).
В качестве Endpoint'a используется grpc + protobuf.

## Запуск
Для запуска сервиса требуется задать следующие переменные окружения:
- GRPC_PORT - порт для grpc Endpoint'а
- POSTGRES_URL - URL к БД Postgres (пример: *postgresql://postgres:postgres@localhost:5432/app_position*)
- APPTICA_API_KEY - ключ для API Apptica (параметр "B4NKGg")

## TODO
1. 100% покрытие тестами
2. Добавить ограничение по кол-ву запросов (с использованием Redis)
3. Добавить логирование запросов на endpoint
4. Завернуть в Docker
5. Реализовать интеграционное тестирование (на основе [testcontainers](https://github.com/testcontainers/testcontainers-go))
