# [ BoxApi ] (учебный)

> Простой сервис для обработки запросов
- POST
- GET
    
***

> API
+ /home
  + общий каталог
+ /calc
  + посчитать простой пример
+ /time
  + получить время онлайн 


***

> Install

```
git clone https://github.com/YuranIgnatenko/BoxApi
```

> Launch
```
cd BoxApi

go run main.go
```


> Example used API (terminal)

```
curl localhost:8080/home
curl localhost:8080/calc

curl -POST localhost:8080/home
curl -POST localhost:8080/time
```

> Example used API (browser)

http://localhost:8080/home

http://localhost:8080/calc

http://localhost:8080/time


***

