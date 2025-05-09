# Структура проекта

## Структура папок
/cmd 
	/app
		/main.go
/admin
	/handlers
	/models
	/service
/infrastructure
	/configs
		config.go
	/database
		/migrations
		/models
		/repository
	/logger
		logger.go
/intertal
	/application
		/service
			service.go
	/domain
		/modules
			/menu
				model.go
			/order
				model.go
			/user
				model.go
	/dto
		dto_models.go
		dto_mappers.go
	/middleware
		middleware.go
	/presentation
		/http
			handlers.go
		/views
			views.go

## Пояснение к структуре

/cmd/app/main.go - стартовой точкой приложения.

/admin — административные компоненты:
	/handlers — обработчики запросов;
	/models — модели данных, связанные с админкой;
	/service — бизнес-логика или сервисы, используемые в административной части.

/infrastructure — Внешние зависимости, настройки, работа с базой данных, логирование.
	/configs/config.go — конфигурационные файлы и конфигуратор
	/database
		/migrations — миграции базы данных;
		/models — модели для базы данных;
		/repository — реализация паттерна репозитория;
	/logger/logger.go — настройка/реализация логгера.

/internal — логика, которая не должна использоваться за пределами приложения.
	/application
		/service/service.go — бизнес логика, сервисы приложения.
	/domain
		/modules
			/menu, /order, /user — различные доменные модули. В каждой из них определена своя модель, описывающая сущности предметной области.
	/dto
		dto_models.go — структуры для передачи данных (Data Transfer Objects);
		dto_mappers.go — функции для преобразования моделей в DTO и наоборот.
	/middleware
		middleware.go — промежуточное ПО для обработки запросов.
	/presentation 
		/http/handlers.go — HTTP-обработчики, реализующие взаимодействие с внешним миром;
		/views/views.go — логика представления данных, например, шаблоны для рендеринга HTML или иной формат представления данных.

