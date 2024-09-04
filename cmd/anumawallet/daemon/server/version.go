package server

import (
	"context"
	"github.com/AnumaNetwork/anumad/cmd/anumawallet/daemon/pb"
	"github.com/AnumaNetwork/anumad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
