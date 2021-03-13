// Package account ...
// generated version: devel
package account

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/firebase/storage"
)

// PatchProfileController ...
type PatchProfileController struct {
	*props.ControllerProps
}

// NewPatchProfileController ...
func NewPatchProfileController(cp *props.ControllerProps) *PatchProfileController {
	p := &PatchProfileController{
		ControllerProps: cp,
	}
	return p
}

// PatchProfile ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param profile body model.SocialProfile false ""
// @Success 200 {object} PatchProfileResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/profile [PATCH]
func (p *PatchProfileController) PatchProfile(
	c echo.Context, req *PatchProfileRequest,
) (res *PatchProfileResponse, err error) {

	if err := auth.IsAuthedUser(c); err != nil {
		body := map[string]string{
			"message": err.Error(),
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	au := auth.GetAuthedUser(c)

	user, err := firestore.GetUserByGoogleID(au.Google.ID)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	rv := reflect.ValueOf(req.Profile)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		value := rv.FieldByName(field.Name).String()
		if field.Name != "StatusMessage" && value == "" {
			continue
		}

		switch field.Name {
		case "ScreenName":
			_, err := firestore.GetUserByScreenName(value)
			if err == nil {
				body := map[string]string{
					"message": fmt.Sprintf("@%s is already taken.", value),
				}

				return nil, wrapper.NewAPIError(http.StatusConflict, body)
			}

			if err.Error() == "Internal Server Error" {
				return nil, wrapper.NewAPIError(http.StatusInternalServerError)
			}

			ctx := context.Background()
			user.Profile.ScreenName = value
			claims := map[string]interface{}{"screen_name": user.Profile.ScreenName}

			err = auth.Client.SetCustomUserClaims(ctx, user.UID, claims)
			if err != nil {
				return nil, wrapper.NewAPIError(http.StatusInternalServerError)
			}

		case "DisplayName":
			user.Profile.DisplayName = value
		case "StatusMessage":
			user.Profile.StatusMessage = value
		case "ImageHash":
			err := storage.ActivateImage(value)
			if err != nil {
				return nil, wrapper.NewAPIError(http.StatusInternalServerError)
			}

			if user.Profile.ImageHash != "" {
				_ = storage.DeleteImage(user.Profile.ImageHash)
			}

			user.Profile.ImageHash = value
		}
	}

	_, err = firestore.StoreUser(user)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PatchProfileResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  user.Profile,
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PatchProfileController) AutoBind() bool {
	return true
}
