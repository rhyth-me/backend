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
)

// GetConnectController ...
type GetConnectController struct {
	*props.ControllerProps
}

// NewGetConnectController ...
func NewGetConnectController(cp *props.ControllerProps) *GetConnectController {
	g := &GetConnectController{
		ControllerProps: cp,
	}
	return g
}

// GetConnect ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetConnectResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/connect [GET]
func (g *GetConnectController) GetConnect(
	c echo.Context, req *GetConnectRequest,
) (res *GetConnectResponse, err error) {
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

	res = &GetConnectResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  "200",
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetConnectController) AutoBind() bool {
	return true
}
