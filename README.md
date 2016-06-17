# test_go_http_server

Все захардкожено. Локальный сервер вешаю на порт 10888. Клиент выполняет 100 тыс запросов к серверу.

`go get github.com/oke11o/test_go_http_server/server`

`go get github.com/oke11o/test_go_http_server/client`

Сперва запустить сервер

`go run server/server.go`

Потом запусить клиент

`go run client/client.go`