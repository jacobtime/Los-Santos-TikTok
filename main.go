package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobtime/Los-Santos-TikTok/service"
	"log"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	err := r.Run()
	if err != nil {
		log.Println(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
