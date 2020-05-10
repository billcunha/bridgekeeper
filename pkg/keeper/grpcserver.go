package keeper

import (
	"context"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// GRPCServer struct
type GRPCServer struct{}

// Ask method for calculate X + Y
func (s *GRPCServer) Ask(ctx context.Context, req *api.AskRequest) (*api.AskResponse, error) {
	return &api.AskResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
