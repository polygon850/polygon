# Polygon demo microservice

![polygon workflow](https://github.com/polygon850/polygon/actions/workflows/master.yml/badge.svg?branch=master)

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
* `CL_LOG_LEVEL` - уровень логирования
* `CL_LOG_JSON` - флаг устанавливающий JSON-формат логов
* `CL_HTTP_SERVICE_LISTEN` - хост и порт, который будет слушать служебный HTTP сервер (в формате host:port)
* `CL_ENABLE_PPROF` - включение отладки при помощи pprof
 

## Пример запуска docker-контейнера
```shell
docker run \
  --rm \
  -it \
  -p 8080:8080 \
  -e CL_LOG_LEVEL=info \
  -e CL_HTTP_SERVICE_LISTEN=:8080 \
  --env-file=.env.local \
  --name polygon \
  polygon850/polygon
```