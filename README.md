
# Golang Echo Demo

This is a really small demo for echo in golang on a Friday afternoon.

It was hacked up in an hour to show off how to use golang and the echo framework.

*update*: Hacking up DB support in part 2 of the demo

Run MongoDB in docker with:
```sh
docker run --name demo-mongo -p 27017:27017 -d mongo:5.0.3
```


### How to run
```
go run main.go
```

### How to call using curl

```sh
curl -X POST http://localhost:1323/persons -H 'Content-Type: application/json' -d '{"name": "Luuk", "age": 30}'
curl http://localhost:1323/persons
curl http://localhost:1323/persons/81
curl -X DELETE http://localhost:1323/persons/81
curl http://localhost:1323/persons
```
