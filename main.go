package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {

	flag.Parse()
	defer glog.Flush()

	server := Server{}

	server.Build()
}
