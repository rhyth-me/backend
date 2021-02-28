// Package accounts ...
// generated version: 1.8.0
package accounts

import (
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/interfaces/props"
)

// PostCreateAccountController ...
type PostCreateAccountController struct {
	*props.ControllerProps
}

// NewPostCreateAccountController ...
func NewPostCreateAccountController(cp *props.ControllerProps) *PostCreateAccountController {
	p := &PostCreateAccountController{
		ControllerProps: cp,
	}
	return p
}

// PostCreateAccount ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string true ""
// @Param email body string false ""
// @Success 200 {object} PostCreateAccountResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/accounts/create_account [POST]
func (p *PostCreateAccountController) PostCreateAccount(
	c echo.Context, req *PostCreateAccountRequest,
) (res *PostCreateAccountResponse, err error) {
	// API Error Usage: github.com/rhyth-me/backend/interfaces/wrapper
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest)
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, wrapper.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}
