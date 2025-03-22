# MUSIC LIBRARY PROJECT

REST API-приложение на Go для хранения и управления музыкальной библиотекой.

Позволяет выполнять CRUD-операции с песнями, подтягивать дополнительные данные из внешнего API и осуществлять постраничный вывод текста песен.

---

## Возможности

- Добавление песен с автоматическим обогащением (дата, текст, ссылка) через внешний API
- Получение песен с фильтрацией и пагинацией
- Получение одной песни по ID
- Обновление информации о песне
- Удаление песни
- Постраничный просмотр текста песни (lyrics)

---

## Технологии

- **Go 1.21+**
- **Gin** — HTTP-фреймворк
- **GORM** — ORM для PostgreSQL
- **Swagger** — автогенерация документации
- **Logrus** — логгер
- **GoMock + Testify** — для модульных тестов

---

## Установка и запуск

```bash
git clone https://github.com/NewMoscovv/music-library.git
cd music-library
go mod tidy
go run cmd/main.go
```
**ВАЖНО**:
- перед запуском необходимо создать .env-файл в корне для конфигурации проекта. Формат:

```.env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=postgres
API_URL=http://localhost:8080
```

---

## Swagger-документация

**http://localhost:8080/swagger/index.html**

---

## Примеры запросов

1. Добавление песни
```bash
   curl -X POST http://localhost:8080/songs \
-H "Content-Type: application/json" \
-d '{
"group": "Смешарики",
"song": "От винта!"
}'
```
2. Удаление песни
```bash
curl -X DELETE http://localhost:8080/songs/1
```
3. Получить список песен с фильтрацией и пагинацией
```bash
curl -X GET http:/localhost:8080/songs?group=ALBLAK%2052
```
Можно фильтровать по любому полю:
- group
- song
- release_date
- text
- link
4. Получить песню по ID
```bash
curl -X GET http://localhost:8080/songs/1
```
5. Обновить песню по ID
```bash
curl -X PUT http://localhost:8080/songs/1 \
-H "Content-Type: application/json" \
-d '{
  "group": "Imagine Dragons",
  "song": "Thunder",
  "release_date": "2017-04-27",
  "text": "Just a young gun...",
  "link": "https://youtube.com/example"
}'
```
6. Получить текст песни построчно
```bash
curl -X GET http://localhost:8080/songs/1/lyrics?page=1&limit=3
```