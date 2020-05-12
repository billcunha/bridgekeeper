package check

import (
	"reflect"
	"testing"

	"github.com/billcunha/bridgekeeper/pkg/api"
)

// TestPaymentInfoCheck ...
func TestPaymentInfoCheck(t *testing.T) {
	payment := &api.PaymentInfo{
		CardholderName: "",
	}

	expected := int32(20)

	score, err := PaymentInfo(payment)
	if err != nil {
		t.Error("Fail on call check.PaymentInfo")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("PaymentInfoTest SUCCESS")
	} else {
		t.Errorf("PaymentInfoTest FAIL: Expected %v, received: %v", expected, score)
	}

	expected = int32(10)
	payment.CardholderName = "John"

	score, err = PaymentInfo(payment)
	if err != nil {
		t.Error("Fail on call check.PaymentInfo")
	}

	if reflect.DeepEqual(expected, score) {
		t.Log("PaymentInfoTest SUCCESS")
	} else {
		t.Errorf("PaymentInfoTest FAIL: Expected %v, received: %v", expected, score)
	}
}
