# dgraph_demo

useage:
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

build nodes:
```
go run main.go -c 1000 -type nodes
```

build edges:
```
go run main.go -c 1000  -type edges 
```

schema:
```
md5: string @index(hash) .
parent: uid @reverse @count .
