package main

func main() {
	go server()
	client()
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
