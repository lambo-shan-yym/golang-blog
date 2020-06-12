package main

import "golang-blog/src/router"

func main() {
	router := router.InitRouter()
	router.Run()
}
