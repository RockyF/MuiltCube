package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
	"database/sql"
	"github.com/coopernurse/gorp"
	"time"
	"fmt"
)

func main() {
	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// create two posts
	p1 := newPost("Go 1.1 released!", "Lorem ipsum lorem ipsum")
	p2 := newPost("Go 1.2 released!", "Lorem ipsum lorem ipsum")

	// insert rows - auto increment PKs will be set properly after the insert
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "Insert failed")

	// use convenience SelectInt
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	fmt.Println(count)

	// update a row
	p2.Title = "Go 1.2 is better than ever"
	count, err = dbmap.Update(&p2)
	checkErr(err, "Update failed")
	fmt.Println(count)

	// fetch one row - note use of "post_id" instead of "Id" since column is aliased
	//
	// Postgres users should use $1 instead of ? placeholders
	// See 'Known Issues' below
	//
	err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	checkErr(err, "SelectOne failed")

	// fetch all rows
	var posts []Post
	_, err = dbmap.Select(&posts, "select * from posts order by post_id")
	checkErr(err, "Select failed")

	// delete row by PK
	count, err = dbmap.Delete(&p1)
	checkErr(err, "Delete failed")
	fmt.Println(count)

	// delete row manually via Exec
	_, err = dbmap.Exec("delete from posts where post_id=?", p2.Id)
	checkErr(err, "Exec failed")

	// confirm count is zero
	count, err = dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	fmt.Println(count)
}

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string
	Body    string
}

func newPost(title, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mymysql", "tcp:localhost:3306*gotest/root/root")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
