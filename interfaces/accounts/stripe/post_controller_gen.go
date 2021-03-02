// Package stripe ...
// generated version: 1.8.0
package stripe

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/interfaces/props"
)

// PostController ...
type PostController struct {
	*props.ControllerProps
}

// NewPostController ...
func NewPostController(cp *props.ControllerProps) *PostController {
	p := &PostController{
		ControllerProps: cp,
	}
	return p
}

// Post ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param email body string false ""
// @Success 200 {object} PostResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /accounts/stripe [POST]
func (p *PostController) Post(
	c echo.Context, req *PostRequest,
) (res *PostResponse, err error) {

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

	res = &PostResponse{
		Code:        http.StatusOK,
		Message:     "Success",
		RedirectURL: "",
	}

	return res, nil
}
