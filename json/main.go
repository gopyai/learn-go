package main

func main() {

}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
