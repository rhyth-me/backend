// Package item ...
// generated version: 1.8.0
package item

import (
	"context"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/interfaces/wrapper"
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
// @Param itemID path string true ""
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/items/{itemID} [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {

	iter := g.ControllerProps.Firestore.Collection(os.Getenv("ITEMS_COLLECTION")).
		Where("author.screenName", "==", req.UserID).
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

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  result,
	}

	return res, nil
}
