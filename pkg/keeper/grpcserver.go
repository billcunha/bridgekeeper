package keeper

import (
	"context"

	"github.com/billcunha/bridgekeeper/pkg/api"
	"github.com/billcunha/bridgekeeper/pkg/keeper/check"
)

// GRPCServer struct
type GRPCServer struct{}

// Ask method
func (s *GRPCServer) Ask(ctx context.Context, req *api.AskRequest) (*api.AskResponse, error) {

	var score int32

	score, err := check.Profile(req.Profile)
	if err != nil {
		return &api.AskResponse{
			AllowPass: false,
		}, err
	}

	score, err = check.ShopSummary(req.ShopSummary)
	if err != nil {
		return &api.AskResponse{
			AllowPass: false,
		}, err
	}

	score, err = check.PaymentInfo(req.PaymentInfo)
	if err != nil {
		return &api.AskResponse{
			AllowPass: false,
		}, err
	}

	score, err = check.Misc(req.Profile, req.PaymentInfo, req.ShopSummary)
	if err != nil {
		return &api.AskResponse{
			AllowPass: false,
		}, err
	}

	return &api.AskResponse{
		AllowPass: score > 100,
	}, nil
}
