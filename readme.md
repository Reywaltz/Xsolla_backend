# Тестовое задание для Xsolla School 2021. Backend

## Задача
Реализация системы управления товарами для площадки электронной коммерции (от англ. *e-commerce*).

## Структура проекта
```bash
├── cmd
│   └── item-api
│       ├── additions            # Пакет вспомогательных функций
│       │   ├── addition.go
│       │   └── addtions_test.go
│       ├── handlers             # Пакет основных обработчиков
│       │   └── item.go
│       └── main.go
├── docker-compose.yml
├── dockerfile
├── docs
│   └── swagger.yaml             # Спецификация API по стандарту OpenAPI 3.0
├── go.mod
├── go.sum
├── internal                     
│   ├── models                   # Пакет основных сущностей
│   │   ├── item.go
│   │   └── item_test.go
│   └── repositories             # Пакет репозитория
│       ├── itemrepo.go
│       └── itemrepo_test.go
├── pkg
│   ├── log                      # Пакет логера
│   │   ├── cfg.go
│   │   └── log.go
│   └── postgres                 # Пакет базы данных
│       ├── db.go
│       └── db_test.go
├── readme.md
└── sql                          # Директория основных SQL запросов
    └── schema.sql
```

## Инструкция по запуску
Для запуска приложения в локальной среде, необходимо в директории приложения выполнить команду: 
```bash
docker-compose up
````
Для настройки сервиса доступны следующие переменные среды в файле docker-compose:
  * DEV - конфигурация режима вывода сообщений логгера (True - для удобного вывода сообщений в виде текста. False для сообщения в формате JSON)
  * CONN_DB - строка подключения к базе данных.

В директории с SQL запросами определены некоторые тестовые данные, которые вставляются в базу при первой сборке контейнера. Если необходимо инициализировать приложение с пустой базой данных, то следующие строки необходимо закоментировать и пересобрать контейнер приложения, удалив ранее собранный образ:
```SQL
INSERT INTO item (SKU, name, type, cost) VALUES ('DOT-SUB', 'Dota Plus', 'Subscription', '9.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('SON-GAM', 'Sonic rangers', 'Game', '59.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('DMC-GAM', 'DMC:Devil may cry', 'Game', '69.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('OVE-GAM', 'Overwatch', 'Game', '39.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('HOT-GAM', 'Hotline Miami', 'Game', '49.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('XBO-SUB', 'Xbox Gamepass', 'Subsсription', '5.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('PLA-SUB', 'Playstation Plus', 'Subsсription', '6.99');
```

## Основные методы
Документация по реализованному API представленна в виде swagger-файла: https://app.swaggerhub.com/apis/Reywaltz/xsolla-api/1.0

## Хостинг приложения
Приложение развёрнуто на публичном хостинге reg.ru под следующим URL: https://vagu.space/api/v1/items