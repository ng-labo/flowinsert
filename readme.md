Using sflowtool csv formating output, this tool inserts records into voltdb table with timestamp. Code by golang and voltdb-client-go makes simple.

First, prepares github.com/VoltDB/voltdb-client-go/voltdbclient

script SFLOW.sql creates a table with TTL, that is new function of VoltDB.

```
$ go build
$ sqlcmd
> file SFLOW.sql
> ^D
$ ./flowinsert sflowtool -l
```


