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
	"github.com/rhyth-me/backend/pkg/authority"
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

	user := authority.GetIdentifier(c)
	if user.UID == "" {
		body := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "You need to log in.",
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	// Fetch user by uid.
	ctx := context.Background()
	dsnap, err := g.ControllerProps.Firestore.Collection("users").Doc(user.Google.ID).Get(ctx)

	// If uid does not exist, create data.
	if err != nil {
		recode := model.User{
			UID: user.UID,
			Profile: model.SocialProfile{
				ScreenName:       user.Google.ID,
				DisplayName:      "名無さん",
				ProfileImagePath: "",
				StatusMessage:    "",
			},
			Google: user.Google,
		}

		_, err := g.ControllerProps.Firestore.Collection("users").Doc(user.Google.ID).Set(ctx, recode)
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		// Add custom claims
		claims := map[string]interface{}{"screen_name": recode.Profile.ScreenName}
		err = g.ControllerProps.Auth.SetCustomUserClaims(ctx, user.UID, claims)
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		res = &GetResponse{
			Code:    http.StatusOK,
			Message: "Success",
			Result:  recode.Profile,
		}

		return res, nil
	}

	var me model.User
	dsnap.DataTo(&me)

	res = &GetResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  me.Profile,
	}

	return res, nil
}
