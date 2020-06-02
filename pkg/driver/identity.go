package driver

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jessicagreben/test-k8s-csi/pkg/pb"
	log "github.com/sirupsen/logrus"
)

var _ pb.IdentityServer = (*IdentityServer)(nil)

// IdentityServer implements the CSI Identity service as described in CSI spec
// ref: https://github.com/container-storage-interface/spec/blob/master/spec.md
type IdentityServer struct {
	log    *log.Entry
	config Config
}

// NewIdentityServer creates a new identity server
func NewIdentityServer(config Config) *IdentityServer {
	return &IdentityServer{
		log:    log.WithField("name", "identity-server"),
		config: config,
	}
}

// GetPluginInfo returns info about the CSI driver, the name and version
func (s *IdentityServer) GetPluginInfo(context.Context, *pb.GetPluginInfoRequest) (*pb.GetPluginInfoResponse, error) {
	return &pb.GetPluginInfoResponse{
		Name:          s.config.Name,
		VendorVersion: s.config.Version,
	}, nil
}

// GetPluginCapabilities returns a list of capabilities for this CSI driver
func (s *IdentityServer) GetPluginCapabilities(context.Context, *pb.GetPluginCapabilitiesRequest) (*pb.GetPluginCapabilitiesResponse, error) {
	controllerSvc := &pb.PluginCapability_Service_{
		Service: &pb.PluginCapability_Service{
			Type: pb.PluginCapability_Service_CONTROLLER_SERVICE,
		},
	}

	controllerSvcCapability := pb.PluginCapability{
		Type: controllerSvc,
	}

	res := pb.GetPluginCapabilitiesResponse{
		Capabilities: []*pb.PluginCapability{
			&controllerSvcCapability,
		},
	}
	return &res, nil
}

// Probe returns whether or not the CSI driver is ready
func (s *IdentityServer) Probe(context.Context, *pb.ProbeRequest) (*pb.ProbeResponse, error) {
	return &pb.ProbeResponse{
		Ready: &wrappers.BoolValue{Value: true},
	}, nil
}
