package main

import (
	"flag"

	"github.com/jessicagreben/test-k8s-csi/pkg/driver"
)

var (
	port = flag.Int("port", 8080, "port for grpc server")
)

const (
	name    = "csi.storj.io"
	version = "v0.0.1"
)

func main() {
	flag.Parse()
	config := driver.Config{
		Name:    name,
		Version: version,
	}
	csiDriver := driver.New(config)
	csiDriver.Run()
}
