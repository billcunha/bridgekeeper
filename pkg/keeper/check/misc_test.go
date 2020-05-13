package check

import (
	"reflect"
	"testing"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// TestMiscCheck ...
func TestMiscCheck(t *testing.T) {
	profile := &api.Profile{
		Name: "John Doe",
		Ddd:  47,
	}
	summary := &api.ShopSummary{
		TotalOrders: 12,
	}
	payment := &api.PaymentInfo{
		CardholderName:    "John Doe",
		DddBillingAddress: 47,
	}

	expected := int32(0)

	score, err := Misc(profile, payment, summary)
	if err != nil {
		t.Error("Fail on call check.Misc")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("MiscTest SUCCESS")
	} else {
		t.Errorf("MiscTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(50)
	profile.Ddd = 48
	summary.TotalOrders = 0

	score, err = Misc(profile, payment, summary)
	if err != nil {
		t.Error("Fail on call check.Misc")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("MiscTest SUCCESS")
	} else {
		t.Errorf("MiscTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(60)
	profile.Name = "John"

	score, err = Misc(profile, payment, summary)
	if err != nil {
		t.Error("Fail on call check.Misc")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("MiscTest SUCCESS")
	} else {
		t.Errorf("MiscTest FAIL: Expected %v, received: %v", expected, score)
	}
}
