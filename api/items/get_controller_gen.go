// Package items ...
// generated version: devel
package items

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
)

// GetController ...
type GetController struct {
	*props.ControllerProps
}

// NewGetController ...
func NewGetController(cp *props.ControllerProps) *GetController {
	g := &GetController{
		ControllerProps: cp,
	}
	return g
}

// Get ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param itemID path string true ""
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /items/{itemID} [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {

	ctx := context.Background()

	snap, err := firestore.Client.Collection(firestore.Items).Doc(req.ItemID).Get(ctx)
	if err != nil {
		body := map[string]string{
			"message": "The item was not found.",
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	var item model.Item
	snap.DataTo(&item)

	user, err := firestore.GetUserByGoogleID(item.Author.GoogleID)
	if err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	item.Author.Profile = user.Profile

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  item,
	}

	return res, nil

}

// AutoBind - use echo.Bind
func (g *GetController) AutoBind() bool {
	return true
}
