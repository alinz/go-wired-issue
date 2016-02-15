package main

import (
	"github.com/alinz/goissue"
	"github.com/alinz/goissue/config"
	"github.com/alinz/goissue/service"
)

func main() {
	issue.Conf = &config.Config{}
	service.Service()
}
