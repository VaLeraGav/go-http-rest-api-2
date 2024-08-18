## Project-2

Учебный проект на языке goland

### Curl

**Создание alias**

```sh
curl -u root:password -X POST http://localhost:8082/url/ -H "Content-Type: application/json" -d '{"url":"https://habr.com/ru/articles/", "alias":"test-habr"}'
```

**Получение и редирект на страницу**

```sh
$ curl -I -X GET http://localhost:8082/test-habr
```
