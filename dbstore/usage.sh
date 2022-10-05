curl -X POST http://localhost:8080/db/mykey -H "Content-Type:application/octet-stream" -d "foo=bar"

curl -X GET http://localhost:8080/db/mykey

curl -X POST http://localhost:8080/db/json -H "Content-Type:application/octet-stream" -d $'"json":{"name":"AnuchitO"}'

curl -X GET "http://localhost:8080/db/json?format=json"