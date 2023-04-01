package main

import "challenge-3/routers"

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}
