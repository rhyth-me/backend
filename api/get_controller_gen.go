// Package api ...
// generated version: 2.0.0-alpha2
package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
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
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router / [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {
	lu := auth.GetAuthedUser(c)
	if lu.ScreenName == "" {
		lu.ScreenName = "名無し"
	}

	res = &GetResponse{
		Message: fmt.Sprintf("こんにちは、 %s さん。", lu.ScreenName),
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetController) AutoBind() bool {
	return true
}
