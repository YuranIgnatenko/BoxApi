> # BoxApi

Easy micro-service for receive requests

    
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

> Example used API (GET) (browser-click link)

+ http://localhost:8080/home
+ http://localhost:8080/calc
+ http://localhost:8080/time
+ http://localhost:8080/wordplus


<!-- > List supported API requests -->
<!-- + /home -->
  <!-- + main list urls -->
<!-- + /calc -->
  <!-- + get result example -->
<!-- + /time -->
  <!-- + get time now -->
<!-- + /wordplus -->
  <!-- + string-analizator  -->



***
> Documentation & Examples
* [Example curl](./docs/curl.md)
* [Example postman](./docs/postman.md)
* [Example config](./docs/config.md)
* [API](./docs/api.md)