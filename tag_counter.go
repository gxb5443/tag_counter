package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Tag struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func initDb(config string) *gorp.DbMap {
	db, err := sql.Open("postgres", config)
	checkErr(err, "postgres.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	table := dbmap.AddTableWithName(Tag{}, "tags").SetKeys(false, "Tag")
	table.ColMap("Tag").SetNotNull(true)
	table.ColMap("Count").SetNotNull(true)

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Could not create tables")

	return dbmap
}

func main() {
	r := gin.Default()
	config, err := ioutil.ReadFile("db.conf")
	PanicIf(err)
	dbmap := initDb(string(config))
	defer dbmap.Db.Close()
	fmt.Printf("Starting Web Server...")
	r.GET("/tag/:tag", func(c *gin.Context) {
		query_tag := c.Params.ByName("tag")
		obj, err := dbmap.Get(Tag{}, query_tag)
		checkErr(err, "Couldn't get object from db")
		if obj != nil {
			c.JSON(200, obj.(*Tag))
		} else {
			t := Tag{Tag: query_tag}
			c.JSON(200, t)
		}
	})
	r.POST("/tag", func(c *gin.Context) {
		var t Tag
		c.Bind(&t)
		obj, err := dbmap.Get(Tag{}, t.Tag)
		checkErr(err, "Couldn't get object")
		if obj != nil {
			existing_tag, _ := obj.(*Tag)
			existing_tag.Count++
			_, uerr := dbmap.Update(existing_tag)
			checkErr(uerr, "Couldn't update object")
			c.JSON(200, existing_tag)
		} else {
			t.Count = 1
			dbmap.Insert(&t)
			c.JSON(200, t)
		}
	})
	r.Run(":8080")
}
