// Package account ...
// generated version: devel
package account

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/api/apigen/wrapper"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/firebase/firestore"
	"github.com/rhyth-me/backend/pkg/random"
)

// PostLoginController ...
type PostLoginController struct {
	*props.ControllerProps
}

// NewPostLoginController ...
func NewPostLoginController(cp *props.ControllerProps) *PostLoginController {
	p := &PostLoginController{
		ControllerProps: cp,
	}
	return p
}

// PostLogin ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param idToken body string false ""
// @Success 200 {object} PostLoginResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /account/login [POST]
func (p *PostLoginController) PostLogin(
	c echo.Context, req *PostLoginRequest,
) (res *PostLoginResponse, err error) {
	if req.IDtoken == "" {
		return nil, wrapper.NewAPIError(http.StatusBadRequest)
	}

	ctx := context.Background()
	t, err := auth.Client.VerifyIDToken(ctx, req.IDtoken)
	if err != nil {
		body := map[string]string{
			"message": "invalid token",
		}
		return nil, wrapper.NewAPIError(http.StatusBadRequest, body)
	}

	// Return error if the sign-in is older than 30 seconds.
	if time.Now().Unix()-t.AuthTime > 30 {
		body := map[string]string{
			"message": "expired token",
		}
		return nil, wrapper.NewAPIError(http.StatusBadRequest, body)
	}

	// Get claims from token
	claims, err := auth.ParseClaims(t)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	// If this is the first logging in,
	googleId := claims.Firebase.Identities.GoogleID[0]
	_, err = firestore.GetUserByGoogleID(googleId)
	if err != nil {
		if err.Error() != "User Not Found" {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		user := &model.User{
			UID: t.UID,
			Profile: model.SocialProfile{
				ScreenName:  random.String(20),
				DisplayName: "名無しさん",
			},
			Google: model.Google{
				ID:    googleId,
				Email: claims.Firebase.Identities.Email[0],
			},
		}

		_, err = firestore.StoreUser(user)
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}

		// Add custom claims
		err = auth.Client.SetCustomUserClaims(ctx, user.UID, map[string]interface{}{
			"screen_name": user.Profile.ScreenName,
		})
		if err != nil {
			return nil, wrapper.NewAPIError(http.StatusInternalServerError)
		}
	}

	// Set session expiration to 14 days.
	expiresIn := time.Hour * 24 * 14
	session, err := auth.Client.SessionCookie(ctx, req.IDtoken, expiresIn)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusBadRequest)
	}

	cookie := &http.Cookie{
		Name:     auth.SessionName,
		Value:    session,
		Domain:   "rhyth.me",
		MaxAge:   int(expiresIn.Seconds()),
		Secure:   true,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	res = &PostLoginResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result: map[string]interface{}{
			"max-age": int(expiresIn.Seconds()),
		},
	}

	return res, nil
}

// AutoBind - use echo.Bind
func (p *PostLoginController) AutoBind() bool {
	return true
}
