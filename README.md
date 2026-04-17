# EffectiveMobileTest
Тестовое задание Junior Golang Developer Effective Mobile

## Технологический стек
- Язык: Golang
- База данных: PostgreSQL
  - PGX (драйвер)
  - Goose (миграции)
- Документация: Swagger OpenApi (файл api/openapi.yaml)
- Контейнеризация: Docker (docker compose)
- API: net/http
- Конфигурации: godotenv (работа с .env файлом)
- Логирование: slog + tint
- Тестирование: testing + testify

## API Эндпоинты
- POST `/subscriptions` - Создание записей о подписках, принимает JSON с периодом и автоматически разбивает на записи по месяцам
- GET `/subscriptions` - **Получение списка всех записей** о подписках с сортировкой по ID и пагинацией
- GET `/subscriptions/{id}` - **Получение записи** по ID
- PUT `/subscriptions/{id}` - **Изменение записи** по ID (*кроме ID и периода)
- DELETE `/subscriptions/{id}` - **Удаление записи** (*изменение метки об удалении)
- GET `/subscriptions/sum` - **Получение суммы трат** на подписку с указанием названия сервиса, ID пользователя и периода (*название сервиса опционально)

## Запуск
```bash
git clone https://github.com/KEKACIK/EffectiveMobileTest.git
cp .env.example .env
docker compose up --build -d
```
*Примечание: после копирования файла .env, измените настройки на актуальные

## Тестирование (+покрытость в %)
```bash
go test -cover ./...
```
