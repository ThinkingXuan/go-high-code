package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	pprof.Register(r)

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World!")
	})
	r.Run()
}
