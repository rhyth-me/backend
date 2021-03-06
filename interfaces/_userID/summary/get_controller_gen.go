// Package summary ...
// generated version: 1.8.0
package summary

import (
	"context"
	"net/http"
	"os"

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
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/summary [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {
	// Fetch social profile by uid.

	ctx := context.Background()
	iter := g.ControllerProps.Firestore.Collection(os.Getenv("USERS_COLLECTION")).
		Select("profile").
		Where("profile.screenName", "==", req.UserID).
		Documents(ctx)

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

	var result model.User
	docs[0].DataTo(&result)

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  result,
	}

	return res, nil
}
