package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/stripe/stripe-go"
)

func init() {
	stripe.Key = "sk_test_C8kjVdfNuPTnGdKqj70IDMo7"
}

func createCharge(c buffalo.Context) error {
	amountInCents, err := strconv.ParseInt(c.Param("amount_in_cents"), 10, 64)

	if err != nil {
		return err
	}

	chargeParams := &stripe.ChargeParams{
		Amount: stripe.Int64(amountInCents),
	}
}
