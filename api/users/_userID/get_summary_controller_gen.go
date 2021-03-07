// Package users ...
// generated version: devel
package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
)

// GetSummaryController ...
type GetSummaryController struct {
	*props.ControllerProps
}

// NewGetSummaryController ...
func NewGetSummaryController(cp *props.ControllerProps) *GetSummaryController {
	g := &GetSummaryController{
		ControllerProps: cp,
	}
	return g
}

// GetSummary ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Success 200 {object} GetSummaryResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /users/{userID}/summary [GET]
func (g *GetSummaryController) GetSummary(
	c echo.Context, req *GetSummaryRequest,
) (res *GetSummaryResponse, err error) {

	user, err := firestore.GetUserByScreenName(req.UserID)
	if err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusNotFound, body)
	}

	res = &GetSummaryResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result: GetSummaryResponseResult{
			Profile: user.Profile,
		},
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetSummaryController) AutoBind() bool {
	return true
}
