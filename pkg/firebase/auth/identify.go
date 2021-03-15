package auth

import (
	"context"
	"errors"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/domain/model"
)

// Identify - Identify the user from JWT.
func Identify(c echo.Context) *CustomContext {
	var jwt string = c.Request().Header.Get("X-Token")
	ctx := context.Background()

	token, err := Client.VerifyIDToken(ctx, jwt)
	if err != nil {
		return &CustomContext{Context: c, Access: model.Access{
			IPAddress: c.Request().RemoteAddr,
			UserAgent: c.Request().Header.Get("User-Agent"),
		}}
	}

	claims := token.Claims
	identities := claims["firebase"].(map[string]interface{})["identities"].(map[string]interface{})

	var screenName string = ""
	if val, ok := claims["screen_name"].(string); ok {
		screenName = val
	}

	cc := &CustomContext{
		Context: c,
		User: LoginUser{
			UID:        token.UID,
			ScreenName: screenName,
			Google: model.Google{
				ID:    identities["google.com"].([]interface{})[0].(string),
				Email: identities["email"].([]interface{})[0].(string),
			},
		},
		Access: model.Access{
			IPAddress: c.Request().RemoteAddr,
			UserAgent: c.Request().Header.Get("User-Agent"),
		}}

	if os.Getenv("STAGING") == "true" {
		cc.Access.IPAddress = "8.8.8.8"
	}

	return cc
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
