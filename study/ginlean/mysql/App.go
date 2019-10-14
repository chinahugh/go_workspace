package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// User entity
type User struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("hello mysql")
	g := gin.Default()
	g.GET("/get", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		log.Println(name)
		log.Println(age)
		i := 20
		list := Conn(name, i)

		c.JSON(200, list)
	})

	g.Run()

}

// Conn 数据库连接
func Conn(name string, age int) []User {
	db, e := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(e)

	// stmt, e := db.Prepare("insert into user (name,age) values (?,?)")
	// checkErr(e)

	// _, e := stmt.Exec(name, age)
	// checkErr(e)

	rows, e := db.Query("select id,age,name from user")
	checkErr(e)
	var users []User
	for rows.Next() {
		var id, age int
		var name string
		user := User{}

		err := rows.Scan(&id, &age, &name)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Age = age
		user.Name = name

		users = append(users, user)
	}

	db.Close()
	return users

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
