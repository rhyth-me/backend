// Package account ...
// generated version: devel
package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/stripe"
)

// PostConnectController ...
type PostConnectController struct {
	*props.ControllerProps
}

// NewPostConnectController ...
func NewPostConnectController(cp *props.ControllerProps) *PostConnectController {
	p := &PostConnectController{
		ControllerProps: cp,
	}
	return p
}

// PostConnect ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} PostConnectResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/connect [POST]
func (p *PostConnectController) PostConnect(
	c echo.Context, req *PostConnectRequest,
) (res *PostConnectResponse, err error) {
	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)
	access := auth.GetAccessEnv(c)

	user, err := firestore.GetUserByGoogleID(au.Google.ID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	if user.Payment.Connect.ID != "" {
		body := map[string]string{
			"message": "You already have a connect account.",
		}
		return nil, wrapper.NewAPIError(http.StatusConflict, body)
	}

	// Create new card
	result, err := stripe.CreateAccount(user, access)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	user.Payment.Connect.ID = result.ID

	_, err = firestore.StoreUser(user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PostConnectResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PostConnectController) AutoBind() bool {
	return true
}
