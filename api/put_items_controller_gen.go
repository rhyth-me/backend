// Package api ...
// generated version: devel
package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/random"
)

// PutItemsController ...
type PutItemsController struct {
	*props.ControllerProps
}

// NewPutItemsController ...
func NewPutItemsController(cp *props.ControllerProps) *PutItemsController {
	p := &PutItemsController{
		ControllerProps: cp,
	}
	return p
}

// PutItems ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param details body model.ItemSnippet false ""
// @Success 200 {object} PutItemsResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /items [PUT]
func (p *PutItemsController) PutItems(
	c echo.Context, req *PutItemsRequest,
) (res *PutItemsResponse, err error) {

	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)

	// Generate itemID
	id := random.String(8)
	ctx := context.Background()

	// Avoid duplication
	if _, err = firestore.Client.Collection(firestore.Items).Doc(id).Get(ctx); err == nil {
		return nil, wrapper.NewAPIError(http.StatusConflict)
	}

	item := model.Item{
		ID:      id,
		Snippet: req.Details,
		Author: model.ItemAuthor{
			GoogleID: au.Google.ID,
		},
	}

	_, err = firestore.Client.Collection(firestore.Items).Doc(id).Set(ctx, item)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutItemsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  item,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PutItemsController) AutoBind() bool {
	return true
}
