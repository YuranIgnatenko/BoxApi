
> Example used API (terminal)
``` 
  (GET)
curl localhost:8080/home
curl localhost:8080/calc
curl localhost:8080/time
curl localhost:8080/wordplus

```
<br>

```
  (POST)
curl -POST localhost:8080/home

curl -POST localhost:8080/time

curl -X POST \
  http://localhost:8080/calc \
  -H 'content-type: application/json' \
  -d '{ "x": 23, "y":2.4, "/" }'

curl -X POST \
  http://localhost:8080/wordplus \
  -H 'content-type: application/json' \
  -d '{"x":"Hello , Bob! How you Are ?", \
      "params":["cw","cwl","cwu","cwt","cs","verb","lens"]}'
```
