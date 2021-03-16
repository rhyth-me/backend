package auth

import (
	"context"
	"errors"
	"os"
	r "reflect"

	fa "firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/rhyth-me/backend/domain/model"
)

func ClaimsToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
	structVal := r.Indirect(r.ValueOf(val))
	for name, elem := range mapVal {
		structVal.FieldByName(name).Set(r.ValueOf(elem))
	}
	return
}

// Identify - Identify the user from JWT.
func Identify(c echo.Context) *CustomContext {
	ip := c.Request().RemoteAddr
	if os.Getenv("STAGING") == "true" {
		ip = "8.8.8.8"
	}

	// access - user's browsing env
	access := model.Access{
		IPAddress: ip,
		UserAgent: c.Request().Header.Get("User-Agent"),
	}

	// Get session cookie
	cookie, err := c.Cookie(SessionName)
	if err != nil {
		return &CustomContext{Context: c, Access: access}
	}

	ctx := context.Background()
	token, err := Client.VerifySessionCookieAndCheckRevoked(ctx, cookie.Value)
	if err != nil {
		return &CustomContext{Context: c, Access: access}
	}

	claims, err := ParseClaims(token)
	if err != nil {
		return &CustomContext{Context: c, Access: access}
	}
	if claims.ScreenName == "" {
		return &CustomContext{Context: c, Access: access}
	}

	cc := &CustomContext{
		Context: c,
		User: LoginUser{
			UID:        token.UID,
			ScreenName: claims.ScreenName,
			Google: model.Google{
				ID:    claims.Firebase.Identities.GoogleID[0],
				Email: claims.Firebase.Identities.Email[0],
			},
		},
		Access: access,
	}

	return cc
}

// ParseClaims - read claims from map
func ParseClaims(token *fa.Token) (*Claims, error) {
	var claims *Claims
	err := mapstructure.Decode(token.Claims, &claims)
	if err != nil {
		return nil, errors.New("Failed to parse claims")
	}
	return claims, nil
}

// IsAuthedUser - Verify if the user is logged in.
func IsAuthedUser(c echo.Context) error {
	au := c.(*CustomContext).User

	if au.UID == "" {
		return errors.New("authentication is required")
	}
	return nil
}

// GetAuthedUser - Return the identity of auth user.
func GetAuthedUser(c echo.Context) LoginUser {
	au := c.(*CustomContext).User

	return au
}

// GetAccessEnv - Return the user's ip etc.
func GetAccessEnv(c echo.Context) model.Access {
	env := c.(*CustomContext).Access

	return env
}
