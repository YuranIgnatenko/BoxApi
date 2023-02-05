
> # API

***

 Method|Path         | Keys     |Values |body request|Response| Description
 ---:  |   :---      | :---:    |---    |---  |:---:     | :---
 GET   |/api/home    | no params|       |     |HTML| home-page
 GET   |/api/time    | no params|       |     |HTML| time-page
 GET   |/api/calc    | no params|       |     |HTML| calc-page
 GET   |/api/wordplus| no params|       |     |HTML| wordplus-page
 POST  |/api/time    | no params|       |     |JSON| time now
 POST  |/api/calc    | x,y,func |       |`{"text":"Hello , Bob! How you Are ?","params":["cw","cwl","cwu", "cwt", "cs", "verb" ,"lens"]}`  |JSON| calculate ex
 POST  |/api/wordplus| text,params |     |JSON| analizator strings


</br>

***
> * [README](../README.md) 