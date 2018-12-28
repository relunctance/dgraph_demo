# dgraph_demo
### Data is a description of process relationships 
- befor you load , you should assign max uid (eg:650000000)
```
curl -s "http://you.zero.leader:6080/assign?what=uids&num=650000000"

```

## usage:
```
go run main.go --help

-c int
    how many groutine num for run (default 100)
-maxuid int
    the uid random boundary value (default 650000000)
-num int
    how many data you want to build (default 10000)
-type string
    which type should build  , the value you can select 'nodes' or 'edges'  (default "nodes")
```

## build nodes:
```
go run main.go -c 1000 -type nodes

<0x399816484>   <md5>   "15af151e0a8900a6a5910e9ab7c3d1d9"      .
<0x88727215>    <md5>   "34aa20a185bad5c4978ef146ac22d8e8"      .
<0x123765589>   <md5>   "9f9ddcdf6d549db46604f293aeeb7056"      .
<0x202065029>   <md5>   "bcf23280fc6e057e8b174550f69064aa"      .
<0x150284490>   <md5>   "2dfa255dc0794fbfc06e9896a7ab7c8b"      .
<0x452831330>   <md5>   "32a09108de57850e15ed8714307e2ee8"      .
<0x467763304>   <md5>   "eec11dcb1b1df6f7d0054520c8b8e128"      .
...

```

## build edges:
```
go run main.go -c 1000  -type edges 

<0x143545272>   <parent>        <0x75408763>    .
<0x61168611>    <parent>        <0x616076829>   .
<0x232395722>   <parent>        <0x34282357>    .
<0x269891389>   <parent>        <0x507092845>   .
<0x48371838>    <parent>        <0x366437176>   .
...
```

## schema:
```
md5: string @index(hash) .
parent: uid @reverse @count .
```
