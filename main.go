// A simple example that demonstrates the use of the VoltDB database/sql/driver and sflowtool.
package main

import (
    "database/sql/driver"
    "fmt"
    "log"
    "os/exec"
    "bufio"
    "strings"
    "strconv"
    "time"
    "os"
    "github.com/VoltDB/voltdb-client-go/voltdbclient"
)

func main() {
    if len(os.Args)<3 {
        fmt.Println("too short parameters")
        os.Exit(1)
    }
    conn, err := voltdbclient.OpenConn("localhost:21212")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    resCons := responseConsumer{}

    cmd := exec.Command(os.Args[1], os.Args[2:]...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    cmd.Start()
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        line := scanner.Text()
        if ! strings.HasPrefix(line, "FLOW") {
            continue
        }
        microsec := time.Now().UnixNano() / 1000

        // see line-by-line csv output format
        // https://github.com/sflow/sflowtool

        items := strings.Split(line, ",")
        agent := items[1]
        inind, _ := strconv.ParseInt(items[2], 10, 32)
        outind, _ := strconv.ParseInt(items[3], 10, 32)
        //srcmac, _ := item[4]
        //dstmac, _ := item[5]
        //ethtype, _ := item[6]
        //srcvlan, _ := item[7]
        //dstvlan, _ := item[8]
        srcip := items[9]
        dstip := items[10]
        ipprotocol, _ := strconv.ParseInt(items[11], 10, 32)
        srcport, _ := strconv.ParseInt(items[14], 10, 32)
        dstport, _ := strconv.ParseInt(items[15], 10, 32)
        tcpflags, _ := strconv.ParseInt(items[16], 16, 32)
        pktsz, _ := strconv.ParseInt(items[17], 10, 32)
        ipsz, _ := strconv.ParseInt(items[18], 10, 32)
        rate, _ := strconv.ParseInt(items[19], 10, 32)

        conn.ExecAsync(resCons, "SFLOW.insert",
               []driver.Value{ microsec, agent, inind, outind,
                               srcip, dstip, ipprotocol, tcpflags,
                               srcport, dstport, pktsz, ipsz, rate })
    }
    cmd.Wait()

}

type responseConsumer struct{}

func (rc responseConsumer) ConsumeError(err error) {
    fmt.Println(err)
}

func (rc responseConsumer) ConsumeResult(res driver.Result) {
    /*
    ra, _ := res.RowsAffected()
    lid, _ := res.LastInsertId()
    fmt.Printf("%d, %d\n", ra, lid)
    */
}

func (rc responseConsumer) ConsumeRows(rows driver.Rows) {
    /*
    ts, err := vrows.GetIntegerByName("ts")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ts = %dn", ts)
    */
}
