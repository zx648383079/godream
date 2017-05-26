package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    uid int
    username string
    password string
}

func main() {
    db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
    checkErr(err, "连接成功")
    stmt err := db.prepare("INSERT INT userinfo (username, departname,created) values (?, ?, ?)")
    res.err := stmt.Exec("go", "ddd", "sfa")
    checkErr(err, "插入成功")
    db.Close();
    fmt.Printf("Hello word!");
}

func checkErr(err error, msg string) {
    if err != nil {
        panic(err)
    } else {
        fmt.Println(msg)
    }
}
