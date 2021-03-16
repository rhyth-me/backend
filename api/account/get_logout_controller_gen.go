// Package account ...
// generated version: devel
package account

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
)

// GetLogoutController ...
type GetLogoutController struct {
	*props.ControllerProps
}

// NewGetLogoutController ...
func NewGetLogoutController(cp *props.ControllerProps) *GetLogoutController {
	g := &GetLogoutController{
		ControllerProps: cp,
	}
	return g
}

// GetLogout ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetLogoutResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/logout [GET]
func (g *GetLogoutController) GetLogout(
	c echo.Context, req *GetLogoutRequest,
) (res *GetLogoutResponse, err error) {
	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": "You are not logged in.",
		}
		return nil, wrapper.NewAPIError(http.StatusBadRequest, body)
	}

	au := auth.GetAuthedUser(c)

	ctx := context.Background()
	if err = auth.Client.RevokeRefreshTokens(ctx, au.UID); err != nil {
		body := map[string]string{
			"message": "Failed to log out.",
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError, body)
	}

	cookie := &http.Cookie{
		Name:     auth.SessionName,
		Value:    "",
		Domain:   "rhyth.me",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	res = &GetLogoutResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetLogoutController) AutoBind() bool {
	return true
}
