# BoxApi

> Easy micro-service for receive requests
- POST (returns json)
- GET (returns html)
    
***

> List supported API requests
+ /home
  + main list urls
+ /calc
  + get result example
+ /time
  + get time now
+ /wordplus
  + string-analizator 


***

> Requirements
``` 
go get -u github.com/gin-gonic/gin
```

> Install
``` 
git clone https://github.com/YuranIgnatenko/BoxApi
```

> Launch
```
cd BoxApi

go run main.go  
(default config)

go run main.go -config=some_file.json 
(custom config)

go run main.go >> gin.log 
(save logs requests and other)

go run main.go -config=some_file.json >> gin.log
(custom config and save logs gin)
```

***

<br>

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

> Example used API (GET) (browser-click link)

+ http://localhost:8080/home
+ http://localhost:8080/calc
+ http://localhost:8080/time
+ http://localhost:8080/wordplus


***

> Example config.json file
``` js
{
    "Host":"localhost",
    "Port":"8080",
    "PathJournal":"logs/service.log"  
}
```