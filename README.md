# Сервис для хранения заметок в памяти
### Запуск
```
git clone https://github.com/InspectorVitya/note-storage.git
cd note-storage
go mod tidy
go run cmd/main.go
```

Запуск Unit-тестов
```go test -v ./...```

# HTTP API
Для аутентификация указать заголовок с названием `login` и значением `admin`.
## Get full list notes

### Request

`GET /`
```
curl --location --request GET 'localhost:8080/' \
--header 'login: admin'
```

### Response
```

HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 00:47:20 GMT
Content-Length: 41
[{"id":0,"title":"test1","text":"test"}]
```

## Create new note

### Request

`POST /`
```

curl --location --request POST 'localhost:8080/' \
--header 'login: admin' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "test1",
"text": "test",
"expire_time": "10s"
}'
```

### Response
```
HTTP/1.1 201 Created
Date: Mon, 13 Dec 2021 00:50:26 GMT
Content-Length: 0
```
## Delete note

### Request

`DELETE /`
```
curl --location --request DELETE 'localhost:8080/0' \
--header 'login: admin'
```
### Response
```
HTTP/1.1 200 OK
Date: Mon, 13 Dec 2021 00:52:48 GMT
Content-Length: 0
## Get first note
```

## Get first note
### Request
`GET /first`
```
curl --location --request GET 'localhost:8080/first' \
--header 'login: admin'
```
### Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 00:57:26 GMT
Content-Length: 39
{"id":0,"title":"test1","text":"test"}
```
## Get last note
### Request

`GET /last`
```
curl --location --request GET 'localhost:8080/last' \
--header 'login: admin'
```
### Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 13 Dec 2021 00:57:26 GMT
Content-Length: 39
{"id":1,"title":"test1","text":"test"}
```
