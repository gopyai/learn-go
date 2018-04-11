// main
package main

import (
	"fmt"
	"log"

	"gopkg.in/redis.v3"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//	var sta *redis.StatusCmd

	//	sta = client.Set("tes", []byte{1, 2, 3}, time.Second*10)
	//	ifErr(sta.Err())
	//	fmt.Println(sta.String())

	cmd := client.Get("tes")
	ifErr(cmd.Err())

	b, e := cmd.Bytes()
	ifErr(e)
	fmt.Printf("Get: %v\n", b)

	//	sta = client.FlushAll()
	//	ifErr(sta.Err())
	//	fmt.Println(sta.String())

}

func ifErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
