# Структура проекта

## COMMAND
* go install github.com/pressly/goose/v3/cmd/goose@latest
* goose -dir migrations postgres "postgresql://postgres:*password*@127.0.0.1:5432/postgres?sslmode=disable" up

## Пакеты
* github.com/gin-gonic/gin
* github.com/golang-jwt/jwt/v5

## Структура папок
* /cmd
	* /app
		* /main.go
* /admin
	* /handlers
	* /models
	* /service
* /intertal
	* /infrastructure
		* /configs
			* config.go
		* /database
			* /migrations
			* /models
			* /repository
		* /logger
			* logger.go
	* /application
		* /service
			* service.go
	* /domain
		* /modules
			* /menu
				* model.go
			* /order
				* model.go
			* /customer
				* model.go
	* /dto
		* dto_models.go
		* dto_mappers.go
	* /middleware
		* middleware.go
	* /presentation
		* /http
			* handlers.go
		* /views
			* views.go

## Пояснение к структуре

* /cmd/app/main.go - стартовой точкой приложения.
* /admin — административные компоненты:
	* /handlers — обработчики запросов;
	* /models — модели данных, связанные с админкой;
	* /service — бизнес-логика или сервисы, используемые в административной части.
* /internal — логика, которая не должна использоваться за пределами приложения.
	* /infrastructure — Внешние зависимости, настройки, работа с базой данных, логирование.
		* /configs/config.go — конфигурационные файлы и конфигуратор
		* /database
			* /migrations — миграции базы данных;
			* /models — модели для базы данных;
			* /repository — реализация паттерна репозитория;
		* /logger/logger.go — настройка/реализация логгера.
	* /application
		* /service/service.go — бизнес логика, сервисы приложения.
	* /domain
		* /modules
			* /menu, /order, /customer — различные доменные модули. В каждой из них определена своя модель, описывающая сущности предметной области.
	* /dto
		* dto_models.go — структуры для передачи данных (Data Transfer Objects);
		* dto_mappers.go — функции для преобразования моделей в DTO и наоборот.
	* /middleware
		* middleware.go — промежуточное ПО для обработки запросов.
	* /presentation
		* /http/handlers.go — HTTP-обработчики, реализующие взаимодействие с внешним миром;
		* /views/views.go — логика представления данных, например, шаблоны для рендеринга HTML или иной формат представления данных.

