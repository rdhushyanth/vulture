package vultr

import (
	"strings"
	"time"
)

type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// Account represents the account info
type Account struct {
	Balance           *float32    `json:"balance,string,omitempty"`
	PendingCharges    *float32    `json:"pending_charges,string,omitempty"`
	LastPaymentDate   *customTime `json:"last_payment_date,omitempty"`
	LastPaymentAmount *float32    `json:"last_payment_amount,string,omitempty"`
}
