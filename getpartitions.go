// example to call stored-procedure by "database/sql" connection.
package main

import (
    "database/sql"
)

func processRows(rows *sql.Rows, processor func(map[string]interface{})) (err error) {
    columns, err := rows.Columns()
    if err != nil {
        return err
    }
    colInst := make([]interface{}, len(columns))
    colPtrs := make([]interface{}, len(columns))
    for i := 0; i < len(columns); i++ {
        colPtrs[i] = &colInst[i]
    }
    var rec = make(map[string]interface{})
    for rows.Next() {
        err := rows.Scan(colPtrs...)
        if err != nil {
            return err
        }
        for i, col := range colInst {
            rec[columns[i]] = col
        }
        processor(rec)
    }
    rows.Close()
    return nil
}

func getPartitionKeys(db *sql.DB)(p1 map[int32]int32, p2 map[string]int32) {
    rows, _ := db.Query("@GetPartitionKeys", "INTEGER")
    intpartmap := make(map[int32]int32)
    processRows(rows, func(rec map[string]interface{}) {
        intpartmap[rec["PARTITION_KEY"].(int32)] = rec["PARTITION_ID"].(int32)
    })

    rows, _ = db.Query("@GetPartitionKeys", "STRING")
    strpartmap := make(map[string]int32)
    processRows(rows, func(rec map[string]interface{}) {
        strpartmap[string(rec["PARTITION_KEY"].(uint8))] = rec["PARTITION_ID"].(int32)
    })

    return intpartmap, strpartmap
}

