package driver

import (
	"fmt"
	"net"

	"github.com/jessicagreben/test-k8s-csi/pkg/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Driver is the CSI driver
type Driver struct {
	log    *log.Entry
	config Config
}

// New creates a new driver
func New(config Config) *Driver {
	return &Driver{
		log:    log.WithField("name", "driver"),
		config: config,
	}
}

// Run starts the driver
func (d *Driver) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", d.config.Port))
	if err != nil {
		d.log.Fatalf("failed to listen: %v\n", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterIdentityServer(grpcServer, NewIdentityServer(d.config))
	pb.RegisterNodeServer(grpcServer, NewNodeServer(d.config))
	err = grpcServer.Serve(listener)
	if err != nil {
		d.log.Fatalf("failed to serve: %v\n", err)
	}
}

// Close cleans up resources for the driver
func (d *Driver) Close() error {
	// stop grpc server
	return nil
}
