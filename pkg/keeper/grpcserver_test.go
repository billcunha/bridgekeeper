package keeper

import (
	"context"
	"reflect"
	"testing"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// TestGrpcServer ...
func TestGrpcServer(t *testing.T) {
	server := GRPCServer{}

	request := &api.AskRequest{
		Profile: &api.Profile{
			Name:       "Jessica Tomasi",
			Ddd:        47,
			AccountAge: 10,
			Status:     api.Profile_VERIFIED,
		},
		PaymentInfo: &api.PaymentInfo{
			CardholderName:    "Jessica Tomasi",
			DddBillingAddress: 47,
		},
		ShopSummary: &api.ShopSummary{
			TotalOrders: 2,
		},
	}

	expected := &api.AskResponse{
		AllowPass: true,
		Score:     0,
	}

	response, err := server.Ask(context.Background(), request)
	if err != nil {
		t.Error("Fail on call server.Ask")
	}

	if reflect.DeepEqual(response, expected) {
		t.Log("ProfileTest SUCCESS")
	} else {
		t.Errorf("ProfileTest FAIL: Expected [%v], received: [%v]", expected, response)
	}

	request = &api.AskRequest{}

	response, err = server.Ask(context.Background(), request)
	if err != nil {
		t.Log("ProfileTest with nil values: SUCCESS")
	} else {
		t.Error("ProfileTest with nil values: FAIL")
	}
}
