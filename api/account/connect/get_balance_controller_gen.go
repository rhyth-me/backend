// Package connect ...
// generated version: devel
package connect

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/stripe"
)

// GetBalanceController ...
type GetBalanceController struct {
	*props.ControllerProps
}

// NewGetBalanceController ...
func NewGetBalanceController(cp *props.ControllerProps) *GetBalanceController {
	g := &GetBalanceController{
		ControllerProps: cp,
	}
	return g
}

// GetBalance ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetBalanceResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/connect/balance [GET]
func (g *GetBalanceController) GetBalance(
	c echo.Context, req *GetBalanceRequest,
) (res *GetBalanceResponse, err error) {
	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)

	user, err := firestore.GetUserByGoogleID(au.Google.ID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	if user.Payment.Connect.ID == "" {
		body := map[string]string{
			"message": "You don't have a connect account.",
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	bal, err := stripe.GetBalance(user.Payment.Connect.ID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &GetBalanceResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result: map[string]int{
			"available": int(bal.Available[0].Value),
			"pending":   int(bal.Pending[0].Value),
		},
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetBalanceController) AutoBind() bool {
	return true
}
