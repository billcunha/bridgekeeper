package keeper

import (
	"context"
	"errors"

	"github.com/billcunha/bridgekeeper/pkg/api"
	"github.com/billcunha/bridgekeeper/pkg/keeper/check"
)

// GRPCServer struct
type GRPCServer struct{}

// Ask method
func (s *GRPCServer) Ask(ctx context.Context, req *api.AskRequest) (*api.AskResponse, error) {
	var score int32

	if req.Profile == nil {
		doResponse(0, errors.New("Invalid call"))
	}

	result, err := check.Profile(req.Profile)
	if err != nil {
		doResponse(0, err)
	}
	score += result

	if req.ShopSummary == nil {
		doResponse(0, errors.New("Invalid call"))
	}

	result, err = check.ShopSummary(req.ShopSummary)
	if err != nil {
		doResponse(0, err)
	}
	score += result

	result, err = check.PaymentInfo(req.PaymentInfo)
	if err != nil {
		doResponse(0, err)
	}
	score += result

	result, err = check.Misc(req.Profile, req.PaymentInfo, req.ShopSummary)
	if err != nil {
		doResponse(0, err)
	}
	score += result

	return doResponse(score, err)
}

func doResponse(score int32, err error) (*api.AskResponse, error) {
	allowPass := score < 100
	if err != nil {
		allowPass = false
	}
	return &api.AskResponse{
		AllowPass: allowPass,
		Score:     score,
	}, err
}
