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
	//	panicIf(sta.Err())
	//	fmt.Println(sta.String())

	cmd := client.Get("tes")
	panicIf(cmd.Err())

	b, e := cmd.Bytes()
	panicIf(e)
	fmt.Printf("Get: %v\n", b)

	//	sta = client.FlushAll()
	//	panicIf(sta.Err())
	//	fmt.Println(sta.String())

}

func panicIf(e error) {
	if e != nil {
		log.Panic(e)
	}
}
