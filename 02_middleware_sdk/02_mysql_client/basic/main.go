package main

import
(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql")

func main() {
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/mysql")
	if err!=nil {
		fmt.Println("err is ",err)
	}
	defer db.Close()
	rows,err:=db.Query("show engine innodb  status")
	if err!=nil {
		fmt.Println("err is ",err)
	}
	for rows.Next(){
		var typename string
		var name string
		var status string
 		err:=rows.Scan(&typename,&name,&status)
		if err!=nil {
			fmt.Println("err is ",err)
		}
		fmt.Println(status)
	}

}
