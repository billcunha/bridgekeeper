package check

import "github.com/billcunha/bridgekeeper/pkg/api"

// Profile get the score based on Profile
func Profile(profile *api.Profile) (int32, error) {
	var score int32

	if profile.AccountAge < 2 {
		score += 10
	} else {
		if profile.AccountAge < 10 {
			score += 5
		}
	}

	if profile.Status != api.Profile_VERIFIED {
		score += 50
	}

	return score, nil
}
