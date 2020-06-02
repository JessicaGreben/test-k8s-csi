package driver

import (
	"context"

	"github.com/jessicagreben/test-k8s-csi/pkg/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.NodeServer = (*NodeServer)(nil)

// NodeServer implements the CSI Node service as described in CSI spec
// ref: https://github.com/container-storage-interface/spec/blob/master/spec.md
type NodeServer struct {
	log    *log.Entry
	config Config
}

// NewNodeServer creates NodeServers
func NewNodeServer(config Config) *NodeServer {
	return &NodeServer{
		log:    log.WithField("name", "node-server"),
		config: config,
	}
}

// NodePublishVolume requests to publish a volume on a node
// ref: https://github.com/container-storage-interface/spec/blob/master/spec.md#nodepublishvolume
func (s *NodeServer) NodePublishVolume(context.Context, *pb.NodePublishVolumeRequest) (*pb.NodePublishVolumeResponse, error) {
	return &pb.NodePublishVolumeResponse{}, status.Error(codes.Unimplemented, "not implemented")
}

// NodeUnpublishVolume requests to unpublish a volume on a node
func (s *NodeServer) NodeUnpublishVolume(context.Context, *pb.NodeUnpublishVolumeRequest) (*pb.NodeUnpublishVolumeResponse, error) {
	return &pb.NodeUnpublishVolumeResponse{}, status.Error(codes.Unimplemented, "not implemented")
}

// NodeGetCapabilities returns the supported capabilities of node service
func (s *NodeServer) NodeGetCapabilities(context.Context, *pb.NodeGetCapabilitiesRequest) (*pb.NodeGetCapabilitiesResponse, error) {
	return &pb.NodeGetCapabilitiesResponse{}, status.Error(codes.Unimplemented, "not implemented")
}
