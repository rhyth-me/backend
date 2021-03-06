// Package summary ...
// generated version: 1.8.0
package summary

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
// @Router /_userID/summary [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {
	// Fetch social profile by uid.

	author, err := firestore.GetUserByScreenName(g.ControllerProps.Firestore, req.UserID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}
	if author == nil {
		return nil, wrapper.NewAPIError(http.StatusNotFound)
	}

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result: GetResponseResult{
			Profile: author.Profile,
		},
	}

	return res, nil
}
