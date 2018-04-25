package main

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
)

func main() {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(u4)
	fmt.Println(len(u4.String()))
}
