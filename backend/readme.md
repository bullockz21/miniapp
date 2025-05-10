# Структура проекта

## Структура папок
/cmd<br>
	/app<br>
		/main.go<br>
/admin<br>
	/handlers<br>
	/models<br>
	/service<br>
/intertal<br>
	/infrastructure<br>
		/configs<br>
			config.go<br>
		/database<br>
			/migrations<br>
			/models<br>
			/repository<br>
		/logger<br>
			logger.go<br>
	/application<br>
		/service<br>
			service.go<br>
	/domain<br>
		/modules<br>
			/menu<br>
				model.go<br>
			/order<br>
				model.go<br>
			/user<br>
				model.go<br>
	/dto<br>
		dto_models.go<br>
		dto_mappers.go<br>
	/middleware<br>
		middleware.go<br>
	/presentation<br>
		/http<br>
			handlers.go<br>
		/views<br>
			views.go<br>

## Пояснение к структуре

/cmd/app/main.go - стартовой точкой приложения.<br>
/admin — административные компоненты:<br>
	/handlers — обработчики запросов;<br>
	/models — модели данных, связанные с админкой;<br>
	/service — бизнес-логика или сервисы, используемые в административной части.<br>
/infrastructure — Внешние зависимости, настройки, работа с базой данных, логирование.<br>
	/configs/config.go — конфигурационные файлы и конфигуратор<br>
	/database<br>
		/migrations — миграции базы данных;<br>
		/models — модели для базы данных;<br>
		/repository — реализация паттерна репозитория;<br>
	/logger/logger.go — настройка/реализация логгера.<br>
/internal — логика, которая не должна использоваться за пределами приложения.<br>
	/application<br>
		/service/service.go — бизнес логика, сервисы приложения.<br>
	/domain<br>
		/modules<br>
			/menu, /order, /user — различные доменные модули. В каждой из них определена своя модель, описывающая сущности предметной области.<br>
	/dto<br>
		dto_models.go — структуры для передачи данных (Data Transfer Objects);<br>
		dto_mappers.go — функции для преобразования моделей в DTO и наоборот.<br>
	/middleware<br>
		middleware.go — промежуточное ПО для обработки запросов.<br>
	/presentation<br>
		/http/handlers.go — HTTP-обработчики, реализующие взаимодействие с внешним миром;<br>
		/views/views.go — логика представления данных, например, шаблоны для рендеринга HTML или иной формат представления данных.<br>

