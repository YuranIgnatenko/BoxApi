

> # API

***

Method | Api | Keys | Values | Request | Response | Description
 ---:|:---|:---:|---|---|:---:|:---
GET|/api/home|no|no|curl localhost:8080/api/home|HTML|returns html -page
GET|/api/calc|no|no|curl localhost:8080/api/calc|HTML|returns calc -page
GET|/api/time|no|no|curl localhost:8080/api/time|HTML|returns time -page
GET|/api/wordplus|no|no|curl localhost:8080/api/wordplus|HTML|returns wordplus-page
POST|/api/home|no|no|curl -X POST http://localhost:8080/api/home|JSON|retruns paths api
POST|/api/time|no|no|curl -X POST http://localhost:8080/api/time|JSON|returns time now
POST|/api/calc|x, y, func|x(number),y(number),func(string)|curl -X POST \  http://localhost:8080/api/calc \  -H 'content-type: application/json' \  -d '{ "x": 23, "y":2.4, "/" }'|JSON|func-function operate:+,-,*,/
POST|/api/wordplus|text, params|text(string),params(array string)|curl -X POST \  http://localhost:8080/api/wordplus \  -H 'content-type: application/json' \  -d '{"x":"Hello , Bob! How you Are ?", \      "params":["cw","cwl", \      "cwu","cwt","cs","verb","lens"]}'|JSON|cw-count word \cwl-count word lower \cwu-count word upper \cwt-count word title \cs-count only symbols \verb-verbore retranslate \lens-len text symbols

</br>

***
> * [README](../README.md) 
