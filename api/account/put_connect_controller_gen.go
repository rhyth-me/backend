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

// PutConnectController ...
type PutConnectController struct {
	*props.ControllerProps
}

// NewPutConnectController ...
func NewPutConnectController(cp *props.ControllerProps) *PutConnectController {
	p := &PutConnectController{
		ControllerProps: cp,
	}
	return p
}

// PutConnect ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} PutConnectResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/connect [PUT]
func (p *PutConnectController) PutConnect(
	c echo.Context, req *PutConnectRequest,
) (res *PutConnectResponse, err error) {
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

	if user.Payment.ConnectID != "" {
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

	user.Payment.ConnectID = result.ID

	_, err = firestore.StoreUser(user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutConnectResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PutConnectController) AutoBind() bool {
	return true
}
