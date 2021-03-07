// Package accounts ...
// generated version: 2.0.0-alpha2
package accounts

import (
	"context"
	"net/http"

	fs "cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
)

// PutProfileController ...
type PutProfileController struct {
	*props.ControllerProps
}

// NewPutProfileController ...
func NewPutProfileController(cp *props.ControllerProps) *PutProfileController {
	p := &PutProfileController{
		ControllerProps: cp,
	}
	return p
}

// PutProfile ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param profile body model.SocialProfile false ""
// @Success 200 {object} PutProfileResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /accounts/profile [PUT]
func (p *PutProfileController) PutProfile(
	c echo.Context, req *PutProfileRequest,
) (res *PutProfileResponse, err error) {

	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)

	ctx := context.Background()
	_, err = firestore.Client.Collection(firestore.Users).Doc(au.Google.ID).Update(ctx, []fs.Update{
		{
			Path:  "profile",
			Value: req.Profile,
		},
	})
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutProfileResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  req.Profile,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PutProfileController) AutoBind() bool {
	return true
}
