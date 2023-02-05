
> # Example curl

* GET - returns HTML
* POST - returns JSON
* HOST - host server
* PORT - port server
* PATH - path api method
* YOUR_JSON - params json from doc API

***

</br>

> * GET
```
curl HOST:PORT/api/PATH
```
```
curl localhost:8080/apihome

curl localhost:8080/api/calc
```
***
<br>

> * POST

```
curl -POST HOST:PORT/api/PATH \
        -H 'content-type: application/json' \
        -d '{YOUR_JSON}
```

```

curl -X POST \
  http://localhost:8080/api/calc \
  -H 'content-type: application/json' \
  -d '{ "x": 23, "y":2.4, "/" }'

curl -X POST \
  http://localhost:8080/api/wordplus \
  -H 'content-type: application/json' \
  -d '{"x":"Hello , Bob! How you Are ?", \
      "params":["cw","cwl", \
      "cwu","cwt","cs","verb","lens"]}'
```
***

> * [README](../README.md) 