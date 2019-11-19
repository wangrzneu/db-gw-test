package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	var addr string
	flag.StringVar(&addr, "a", "127.0.0.1", "db gw address")
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("root:Pwd_123456_Pwd@%v/test", addr))
	if err != nil {
		log.Fatalf("fail to open sql: %v", err)
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	rows, err := db.Query("select id, arg from tbl_test_db_gw where arg = 50;/*arg=50*/")
	if err != nil {
		log.Fatalf("fail to query: %v", err)
	}
	for rows.Next() {
		var id int
		var arg int
		if err := rows.Scan(&id, &arg); err != nil {
			log.Printf("row error:%v", err)
		} else {
			log.Printf("id: %v, arg: %v", id, arg)
		}
	}
	fmt.Printf("%v:%v", db, err)
}
