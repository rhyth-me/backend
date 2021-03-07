// Package users ...
// generated version: devel
package users

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
)

// GetItemsController ...
type GetItemsController struct {
	*props.ControllerProps
}

// NewGetItemsController ...
func NewGetItemsController(cp *props.ControllerProps) *GetItemsController {
	g := &GetItemsController{
		ControllerProps: cp,
	}
	return g
}

// GetItems ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Success 200 {object} GetItemsResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/items [GET]
func (g *GetItemsController) GetItems(
	c echo.Context, req *GetItemsRequest,
) (res *GetItemsResponse, err error) {

	user, err := firestore.GetUserByScreenName(req.UserID)
	if err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	ctx := context.Background()
	iter := firestore.Client.Collection(firestore.Items).
		Select("id", "snippet.thumbnailUrl", "snippet.musicTitle", "snippet.price", "statistics").
		Where("author.googleId", "==", user.Google.ID).
		Documents(ctx)

	docs, err := iter.GetAll()
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	var result []model.Item
	var doc model.Item
	for i := 0; i < len(docs); i++ {
		docs[i].DataTo(&doc)
		doc.Author.Profile = user.Profile
		result = append(result, doc)
	}

	res = &GetItemsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  result,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetItemsController) AutoBind() bool {
	return true
}
