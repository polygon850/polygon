# Polygon demo microservice

Микросервис для демонстрации возможностей Golang стека

## Сборка исполняемого файла
```shell
go build ./cmd/polygon
```

## Сборка docker image
```shell
docker build -t polygon850/polygon .
```

## Аргументы и переменные окружения
* `CL_HTTP_PRIVATE_LISTEN` - хост и порт, который будет слушать HTTP сервер (в формате host:port)

## Пример запуска docker-контейнера
```shell
docker run --rm -it -p 8080:8080 -e CL_HTTP_PRIVATE_LISTEN=:8080 --name polygon polygon850/polygon
```