package check

import (
	"reflect"
	"testing"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// TestShopSummaryCheck ...
func TestShopSummaryCheck(t *testing.T) {
	summary := &api.ShopSummary{
		TotalOrders: 0,
	}

	expected := int32(5)

	score, err := ShopSummary(summary)
	if err != nil {
		t.Error("Fail on call check.ShopSummary")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("ShopSummaryTest SUCCESS")
	} else {
		t.Errorf("ShopSummaryTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(0)
	summary.TotalOrders = 3

	score, err = ShopSummary(summary)
	if err != nil {
		t.Error("Fail on call check.ShopSummary")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("ShopSummaryTest SUCCESS")
	} else {
		t.Errorf("ShopSummaryTest FAIL: Expected %v, received: %v", expected, score)
	}
}
