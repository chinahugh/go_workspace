package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
log.Printf("hello %s %d \n","web",1)
	fmt.Printf("Hello %s\n", "web test")
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func hello(rp http.ResponseWriter, rq *http.Request) {
	rq.ParseForm()
	RequestURI := rq.RequestURI
	fmt.Println(RequestURI)
	var s string = rq.Method
	closer := rq.Body
	values := rq.Form
	fmt.Println(s)
	fmt.Println(closer)
	fmt.Println(values)
	i := Conn()
	fmt.Fprint(rp, "hello web  ",i)
}

func Conn() int64{
	db, e := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(e)

	stmt, e := db.Prepare("insert into user (name,age) values (?,?)")
	checkErr(e)

	res, e := stmt.Exec("研发部门", 20)
	checkErr(e)

	i, e := res.LastInsertId()

	checkErr(e)

	fmt.Println(i)
	db.Close()
	return  i
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Aa(a int) int {
	return 1
}

