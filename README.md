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


> Example used API (terminal)
``` 
  (GET)
curl localhost:8080/home
curl localhost:8080/calc

  (POST)
curl -POST localhost:8080/home
curl -POST localhost:8080/time
curl -v -X POST \
  http://localhost:8080/calc \
  -H 'content-type: application/json' \
  -d '{ "x": 2, "y":0, "/" }'
```

> Example used API (browser-click link)

+ http://localhost:8080/home
+ http://localhost:8080/calc
+ http://localhost:8080/time


***

> Example config.json file
``` js
{
    "Host":"localhost",
    "Port":"8080",
    "PathJournal":"logs/service.log"  
}
```