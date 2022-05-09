package main

func main() {
	httpProtocol, err := InitHttpProtocol("live")

	if err != nil {
		panic(err)
	}

	httpProtocol.Listen()
}
