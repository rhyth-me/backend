// Package profile ...
// generated version: 1.8.0
package profile

import (
	"context"
	"net/http"

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
// @Success 200 {object} GetResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /accounts/profile [GET]
func (g *GetController) Get(
	c echo.Context, req *GetRequest,
) (res *GetResponse, err error) {

	uid := c.(*model.CustomContext).UID
	if uid == "" {
		body := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "You need to log in.",
		}
		return nil, wrapper.NewAPIError(http.StatusBadRequest, body)
	}

	// Fetch user by uid.
	ctx := context.Background()
	dsnap, err := g.ControllerProps.Firestore.Collection("users").Doc(uid).Get(ctx)

	// If uid does not exist, create data.
	if err != nil {
		user := model.User{
			UID: uid,
			Profile: model.SocialProfile{
				ID:              uid,
				DisplayName:     "名無さん",
				ProfileImageURL: "",
				StatusMessage:   "",
			},
		}

		_, err := g.ControllerProps.Firestore.Collection("users").Doc(uid).Set(ctx, user)
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		res = &GetResponse{
			Code:    http.StatusOK,
			Message: "Success",
			Result:  user.Profile,
		}

		return res, nil
	}

	var user model.User
	dsnap.DataTo(&user)

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  user.Profile,
	}

	return res, nil
}
