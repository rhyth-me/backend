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

// GetCardsController ...
type GetCardsController struct {
	*props.ControllerProps
}

// NewGetCardsController ...
func NewGetCardsController(cp *props.ControllerProps) *GetCardsController {
	g := &GetCardsController{
		ControllerProps: cp,
	}
	return g
}

// GetCards ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetCardsResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/cards [GET]
func (g *GetCardsController) GetCards(
	c echo.Context, req *GetCardsRequest,
) (res *GetCardsResponse, err error) {
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

	res = &GetCardsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  user.Payment.Cards,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetCardsController) AutoBind() bool {
	return true
}
