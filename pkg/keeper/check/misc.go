package check

import (
	"errors"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// Misc get the score based on cross checks
func Misc(profile *api.Profile, payment *api.PaymentInfo, summary *api.ShopSummary) (int32, error) {
	var score int32

	if profile == nil {
		return 0, errors.New("Invalid Profile")
	}

	if payment == nil {
		return 0, errors.New("Invalid PaymentInfo")
	}

	if summary == nil {
		return 0, errors.New("Invalid ShopSummary")
	}

	if profile.Name != payment.CardholderName {
		score += 10
	}

	if profile.Ddd != payment.DddBillingAddress && summary.TotalOrders < 1 {
		score += 50
	}

	return score, nil
}
