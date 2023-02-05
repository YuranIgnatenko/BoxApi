> # BoxApi

Easy micro-service for receive requests

    
***



> Main requirements
``` 
go get -u github.com/gin-gonic/gin
```

> Other requirements
```
pip3 install pandas
pip3 install odfpy
pip3 install prettytable
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

> Example use browser
1. Run service
2. Click link
   ***
+ http://localhost:8080/home
+ http://localhost:8080/calc
+ http://localhost:8080/time
+ http://localhost:8080/wordplus

> Examples
* [Example curl](./docs/curl.md)
* [Example postman](./docs/postman.md)
* [Example config](./docs/config.md)

> Documentation 
* [API](./docs/api.md)
