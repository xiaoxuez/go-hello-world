package main

//Bolt,一个内嵌的key-value数据库
import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
