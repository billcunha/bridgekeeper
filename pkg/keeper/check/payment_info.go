package check

import "github.com/billcunha/bridgekeeper/pkg/api"

// PaymentInfo get the score based on PaymentInfo
func PaymentInfo(payment *api.PaymentInfo) (int32, error) {
	var score int32

	if payment.CardholderName == "" {
		score += 20
	} else {
		if len(payment.CardholderName) <= 5 {
			score += 10
		}
	}

	return score, nil
}
