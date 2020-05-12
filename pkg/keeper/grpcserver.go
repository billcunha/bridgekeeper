package keeper

import (
	"context"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// GRPCServer struct
type GRPCServer struct{}

// Ask method
func (s *GRPCServer) Ask(ctx context.Context, req *api.AskRequest) (*api.AskResponse, error) {
	return &api.AskResponse{
		AllowPass: true,
	}, nil
}
