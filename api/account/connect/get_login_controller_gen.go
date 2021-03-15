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

// GetLoginController ...
type GetLoginController struct {
	*props.ControllerProps
}

// NewGetLoginController ...
func NewGetLoginController(cp *props.ControllerProps) *GetLoginController {
	g := &GetLoginController{
		ControllerProps: cp,
	}
	return g
}

// GetLogin ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetLoginResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/connect/login [GET]
func (g *GetLoginController) GetLogin(
	c echo.Context, req *GetLoginRequest,
) (res *GetLoginResponse, err error) {
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

	if user.Payment.ConnectID == "" {
		body := map[string]string{
			"message": "You don't have a connect account.",
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	al, err := stripe.IssueLoginLink(user.Payment.ConnectID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &GetLoginResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  al.URL,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetLoginController) AutoBind() bool {
	return true
}
