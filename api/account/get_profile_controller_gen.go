// Package account ...
// generated version: 2.0.0-alpha2
package account

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
)

// GetProfileController ...
type GetProfileController struct {
	*props.ControllerProps
}

// NewGetProfileController ...
func NewGetProfileController(cp *props.ControllerProps) *GetProfileController {
	g := &GetProfileController{
		ControllerProps: cp,
	}
	return g
}

// GetProfile ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetProfileResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /accounts/profile [GET]
func (g *GetProfileController) GetProfile(
	c echo.Context, req *GetProfileRequest,
) (res *GetProfileResponse, err error) {

	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)

	user, _ := firestore.GetUserByGoogleID(au.Google.ID)
	if user != nil {
		res = &GetProfileResponse{
			Code:    http.StatusOK,
			Message: "Success",
			Result:  user.Profile,
		}
		return res, nil
	}

	ctx := context.Background()

	user = &model.User{
		UID: au.UID,
		Profile: model.SocialProfile{
			ScreenName:       au.Google.ID,
			DisplayName:      "名無しさん",
			ProfileImagePath: "/default",
			StatusMessage:    "",
		},
		Google: au.Google,
	}

	_, err = firestore.Client.Collection(firestore.Users).Doc(au.Google.ID).Set(ctx, user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	// Add custom claims
	claims := map[string]interface{}{"screen_name": user.Profile.ScreenName}
	err = auth.Client.SetCustomUserClaims(ctx, user.UID, claims)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &GetProfileResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  user.Profile,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (g *GetProfileController) AutoBind() bool {
	return true
}
