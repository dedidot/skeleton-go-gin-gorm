package main

import(
	"./app"
)

func main() {	
	r := routers.SetupRouters()
	r.Run()
}