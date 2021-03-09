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

// PutCardsController ...
type PutCardsController struct {
	*props.ControllerProps
}

// NewPutCardsController ...
func NewPutCardsController(cp *props.ControllerProps) *PutCardsController {
	p := &PutCardsController{
		ControllerProps: cp,
	}
	return p
}

// PutCards ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param token body string false ""
// @Success 200 {object} PutCardsResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/cards [PUT]
func (p *PutCardsController) PutCards(
	c echo.Context, req *PutCardsRequest,
) (res *PutCardsResponse, err error) {

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

	// Create stripe customer account if the user doesn't have one.
	if user.Payment.CustomerID == "" {
		cs, err := stripe.CreateCustomer(user)
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		user.Payment.CustomerID = cs.ID
	}

	// Create new card
	result, err := stripe.CreateCard(user.Payment.CustomerID, req.Token)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusBadRequest)
	}

	card := model.Card{
		ID:       result.ID,
		Brand:    string(result.Brand),
		Country:  result.Country,
		Last4:    result.Last4,
		ExpYear:  result.ExpYear,
		ExpMonth: result.ExpMonth,
	}

	user.Payment.Cards = append(user.Payment.Cards, card)

	_, err = firestore.StoreUser(user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutCardsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  card,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PutCardsController) AutoBind() bool {
	return true
}
