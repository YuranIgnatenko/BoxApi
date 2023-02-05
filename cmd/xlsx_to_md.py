import pandas as pd
from prettytable import PrettyTable
import time
 
start = time.time()
# =================================== 

namefile_doc = "../docs/api.md"
namefile_ods = "api.ods"

df = pd.read_excel(namefile_ods, sheet_name=0)

Met = list(df['Method']) 
Api = list(df['Api']) 
Key = list(df['Keys']) 
Val = list(df['Values']) 
Req = list(df['Request']) 
Resp = list(df['Response']) 
Desc = list(df['Description'])

const_head = """

> # API

***

Method | Api | Keys | Values | Request | Response | Description
 ---:|:---|:---:|---|---|:---:|:---
"""

const_futer = """
</br>

***
> * [README](../README.md) 
"""

def Row(Method, Api, Keys, Values, Request, Response, Description):
    row = f"{Method}|{Api}|{Keys}|{Values}|{Request}|{Response}|{Description}"
    return row

def Table(rmet,rapi,rkey,rval,rreq,rresp,rdesc):
    count_row = len(rmet)
    body = ""
    body += const_head
    for i in range(count_row):
        row = Row(rmet[i],rapi[i],rkey[i],rval[i],rreq[i],rresp[i],rdesc[i])
        body += row+"\n"
    body += const_futer
    return body, i

def Writer(table):
    file = open(namefile_doc, "w")
    file.write(table)

def main():
    table, count_rows  = Table(Met,Api,Key,Val,Req,Resp,Desc)
    Writer(table)
    return count_rows

count_rows = main()
end = time.time()
t = str((end-start) * 10**3) + " ms"
out = PrettyTable()

out.field_names = ["input file", "output file", "count rows", "time"]
out.add_row([namefile_ods, namefile_doc, count_rows, t])

print(out)
