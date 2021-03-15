package stripe

import (
	"errors"
	"time"

	"github.com/rhyth-me/backend/domain/model"
	"github.com/stripe/stripe-go/v72"
)

// CreateAccount - Create a new account on Stripe.
func CreateAccount(user *model.User, access model.Access) (*stripe.Account, error) {
	params := &stripe.AccountParams{
		Type:    stripe.String("custom"),
		Country: stripe.String("JP"),
		Email:   stripe.String(user.Google.Email),
		Params: stripe.Params{
			Metadata: map[string]string{
				"GoogleID":    user.Google.ID,
				"Email":       user.Google.Email,
				"FirebaseUID": user.UID,
				"IPAddress":   access.IPAddress,
			},
		},
		BusinessType: stripe.String("individual"),
		Capabilities: &stripe.AccountCapabilitiesParams{
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Settings: &stripe.AccountSettingsParams{
			Payouts: &stripe.AccountSettingsPayoutsParams{
				Schedule: &stripe.PayoutScheduleParams{
					Interval: stripe.String("manual"),
				},
			},
		},
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			Date:      stripe.Int64(time.Now().Unix()),
			IP:        stripe.String(access.IPAddress),
			UserAgent: stripe.String(access.UserAgent),
		},
	}

	a, err := Client.Account.New(params)
	if err != nil {
		return nil, errors.New("Failed to create an account")
	}

	return a, nil
}

// DeleteAccount - Delete an account on Stripe.
func DeleteAccount(accountID string) error {
	_, err := Client.Account.Del(accountID, nil)
	if err != nil {
		return errors.New("Failed to delete an account")
	}
	return nil
}

// IssueLoginLink - Create a login link to log in user's connect account.
func IssueLoginLink(accountID string) (*stripe.AccountLink, error) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String(accountID),
		RefreshURL: stripe.String("https://rhyth.me/"),
		ReturnURL:  stripe.String("https://rhyth.me/"),
		Type:       stripe.String("account_onboarding"),
	}

	al, err := Client.AccountLinks.New(params)
	if err != nil {
		return nil, errors.New("Failed to create an account")
	}

	return al, nil
}
