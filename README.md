This is a really small demo for echo in golang on a Friday afternoon


### How to use:

```sh
curl -X POST http://localhost:1323/persons -H 'Content-Type: application/json' -d '{"name": "Luuk", "age": 30}'
curl http://localhost:1323/persons
curl http://localhost:1323/persons/81
curl -X DELETE http://localhost:1323/persons/81
curl http://localhost:1323/persons
```
