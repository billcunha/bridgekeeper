package check

import "github.com/billcunha/bridgekeeper/pkg/api"

// Misc get the score based on cross checks
func Misc(profile *api.Profile, payment *api.PaymentInfo, summary *api.ShopSummary) (int32, error) {
	var score int32

	if profile.Name != payment.CardholderName {
		score += 10
	}

	if profile.Ddd != payment.DddBillingAddress && summary.TotalOrders < 1 {
		score += 50
	}

	return score, nil
}
