// Package items ...
// generated version: 1.8.0
package items

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/interfaces/wrapper"
	"github.com/rhyth-me/backend/pkg/firestore"
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
// @Param userID path string true ""
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/items [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {

	author, err := firestore.GetUserByScreenName(g.ControllerProps.Firestore, req.UserID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}
	if author == nil {
		return nil, wrapper.NewAPIError(http.StatusNotFound)
	}

	iter := g.ControllerProps.Firestore.Collection(os.Getenv("ITEMS_COLLECTION")).
		Select("id", "snippet.thumbnailUrl", "snippet.musicTitle", "snippet.price", "statistics").
		Where("author", "==", author.Google.ID).
		Documents(context.Background())

	docs, err := iter.GetAll()
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	if len(docs) < 1 {
		body := map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Not found.",
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	var result []model.Item
	var doc model.Item
	for i := 0; i < len(docs); i++ {
		docs[i].DataTo(&doc)
		result = append(result, doc)
	}

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  result,
	}

	return res, nil
}
