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
go run main.go -config=some_file.json
```


> Example used API (terminal)
``` 
curl localhost:8080/home
curl localhost:8080/calc

curl -POST localhost:8080/home
curl -POST localhost:8080/time
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