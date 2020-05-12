package check

import "github.com/billcunha/bridgekeeper/pkg/api"

// ShopSummary get the score based on ShopSummary
func ShopSummary(summary *api.ShopSummary) (int32, error) {
	var score int32

	if summary.TotalOrders < 1 {
		score += 5
	}

	return score, nil
}
