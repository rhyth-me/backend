// Package interfaces ...
// generated version: 1.8.0
package interfaces

import (
	"net/http"

	"github.com/ScoreMarket/backend/interfaces/props"
	"github.com/labstack/echo/v4"
)

// GetEchoMessageController ...
type GetEchoMessageController struct {
	*props.ControllerProps
}

// NewGetEchoMessageController ...
func NewGetEchoMessageController(cp *props.ControllerProps) *GetEchoMessageController {
	g := &GetEchoMessageController{
		ControllerProps: cp,
	}
	return g
}

// GetEchoMessage ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param message path string true ""
// @Success 200 {object} GetEchoMessageResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /{message} [GET]
func (g *GetEchoMessageController) GetEchoMessage(
	c echo.Context, req *GetEchoMessageRequest,
) (res *GetEchoMessageResponse, err error) {
	res = &GetEchoMessageResponse{
		Status:  http.StatusOK,
		Message: req.Message,
	}

	return res, nil
}
