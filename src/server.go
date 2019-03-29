package main

import (
	"io/ioutil"
	"net/http"

	"./util"

	"github.com/gin-gonic/gin"
)

var routePrefix = "/api/"

func main() {
	r := gin.Default()

	services, err := ioutil.ReadDir("./src/api/")
	util.Check(err)

	for _, service := range services {
		var response []byte
		go util.CompileJSON(service.Name(), &response)

		route := routePrefix + "my-" + service.Name()
		r.GET(route, func(c *gin.Context) {
			c.JSON(http.StatusOK, string(response))
		})
	}

	r.Run()
}
