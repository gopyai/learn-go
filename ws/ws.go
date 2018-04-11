package main

func main() {
	go server()
	client()
}

func isErr(err error) {
	if err != nil {
		panic(err)
	}
}
