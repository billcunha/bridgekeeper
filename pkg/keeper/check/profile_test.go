package check

import (
	"reflect"
	"testing"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// TestProfileCheck ...
func TestProfileCheck(t *testing.T) {
	profile := &api.Profile{
		AccountAge: 18,
		Name:       "João da Silva",
		Ddd:        47,
		Status:     api.Profile_VERIFIED,
	}

	expected := int32(0)

	score, err := Profile(profile)
	if err != nil {
		t.Error("Fail on call check.Profile")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("ProfileTest SUCCESS")
	} else {
		t.Errorf("ProfileTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(10)
	profile.AccountAge = 1

	score, err = Profile(profile)
	if err != nil {
		t.Error("Fail on call check.Profile")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("ProfileTest SUCCESS")
	} else {
		t.Errorf("ProfileTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(50)
	profile.AccountAge = 380
	profile.Status = api.Profile_WAITING

	score, err = Profile(profile)
	if err != nil {
		t.Error("Fail on call check.Profile")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("ProfileTest SUCCESS")
	} else {
		t.Errorf("ProfileTest FAIL: Expected %v, received: %v", expected, score)
	}
}
