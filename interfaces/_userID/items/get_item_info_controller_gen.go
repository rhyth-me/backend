// Package items ...
// generated version: 1.8.0
package items

import (
	"context"
	"net/http"

	"github.com/ScoreMarket/backend/domain/model"
	"github.com/ScoreMarket/backend/interfaces/props"
	"github.com/ScoreMarket/backend/interfaces/wrapper"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

// GetItemInfoController ...
type GetItemInfoController struct {
	*props.ControllerProps
}

// NewGetItemInfoController ...
func NewGetItemInfoController(cp *props.ControllerProps) *GetItemInfoController {
	g := &GetItemInfoController{
		ControllerProps: cp,
	}
	return g
}

// GetItemInfo ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Param itemID path string true ""
// @Success 200 {object} GetItemInfoResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/items/{itemID} [GET]
func (g *GetItemInfoController) GetItemInfo(
	c echo.Context, req *GetItemInfoRequest,
) (res *GetItemInfoResponse, err error) {

	iter := g.ControllerProps.Firestore.Collection("items").
		Where("author.id", "==", req.UserID).
		Where("id", "==", req.ItemID).Documents(context.Background())

	docs, err := iter.GetAll()
	if err != nil {
		glog.Errorln("Error")
	}

	if len(docs) < 1 {
		body := map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Not found.",
		}
		return nil, wrapper.NewAPIError(http.StatusBadRequest, body)
	}

	var result model.Item
	docs[0].DataTo(&result)

	res = &GetItemInfoResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  result,
	}

	return res, nil
}
