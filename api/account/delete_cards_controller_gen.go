// Package account ...
// generated version: devel
package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/stripe"
)

// DeleteCardsController ...
type DeleteCardsController struct {
	*props.ControllerProps
}

// NewDeleteCardsController ...
func NewDeleteCardsController(cp *props.ControllerProps) *DeleteCardsController {
	d := &DeleteCardsController{
		ControllerProps: cp,
	}
	return d
}

// DeleteCards ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param cardID body integer false ""
// @Success 200 {object} DeleteCardsResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/cards [DELETE]
func (d *DeleteCardsController) DeleteCards(
	c echo.Context, req *DeleteCardsRequest,
) (res *DeleteCardsResponse, err error) {
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

	// Validate if the user have some cards.
	if len(user.Payment.Cards) < 1 {
		return nil, wrapper.NewAPIError(http.StatusBadRequest)
	}

	if len(user.Payment.Cards) < req.CardID {
		return nil, wrapper.NewAPIError(http.StatusBadRequest)
	}

	card := user.Payment.Cards[req.CardID-1]
	_, err = stripe.DeleteCard(user.Payment.CustomerID, card.ID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	// Remove a card
	user.Payment.Cards = remove(user.Payment.Cards, req.CardID-1)
	_, err = firestore.StoreUser(user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &DeleteCardsResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (d *DeleteCardsController) AutoBind() bool {
	return true
}

// Remove a card
func remove(slice []model.Card, s int) []model.Card {
	return append(slice[:s], slice[s+1:]...)
}
