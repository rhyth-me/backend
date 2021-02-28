// Package accounts ...
// generated version: 1.8.0
package accounts

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/interfaces/props"
)

// PostStripeController ...
type PostStripeController struct {
	*props.ControllerProps
}

// NewPostStripeController ...
func NewPostStripeController(cp *props.ControllerProps) *PostStripeController {
	p := &PostStripeController{
		ControllerProps: cp,
	}
	return p
}

// PostStripe ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Param email body string false ""
// @Success 200 {object} PostStripeResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/accounts/stripe [POST]
func (p *PostStripeController) PostStripe(
	c echo.Context, req *PostStripeRequest,
) (res *PostStripeResponse, err error) {

	/*
		params := &stripe.AccountParams{
			Country:      stripe.String("JP"),
			Type:         stripe.String("custom"),
			BusinessType: stripe.String("individual"),
			Capabilities: &stripe.AccountCapabilitiesParams{
				Transfers: &stripe.AccountCapabilitiesTransfersParams{
					Requested: stripe.Bool(true),
				},
			},
		}

		result, _ := p.ControllerProps.Stripe.Account.New(params)
		res = &PostStripeResponse{
			Code:        http.StatusOK,
			Message:     "Success",
			RedirectURL: result.ID,
		}
	*/

	res = &PostStripeResponse{
		Code:        http.StatusOK,
		Message:     "Success",
		RedirectURL: "",
	}

	return res, nil
}
